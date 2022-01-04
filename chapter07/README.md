# Chapter 7 - OpenShift GitOps - ArgoCD

List of commands to copy and paste:

**Intalling argocd cli**

sudo curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
sudo chmod +x /usr/local/bin/argocd
argocd version

**Configuring ArgoCD for multi-cluster**

API_NEWCLUSTER=api.ocp.example.com
NEWCLUSTER_USER=admin

API_OCPARGOCLUSTER=api.ocp-argo.example.com
OCPARGOCLUSTER_USER=admin

oc login -u $NEWCLUSTER_USER https://$API_NEWCLUSTER:6443
oc login -u $OCPARGOCLUSTER_USER https://$API_OCPARGOCLUSTER:6443
oc config get-contexts
oc config set-context prd-cluster --cluster=newcluster --user=admin
oc get route openshift-gitops-server -n openshift-gitops -o jsonpath='{.spec.host}'
oc extract secret/openshift-gitops-cluster -n openshift-gitops --to=-
argocd login --insecure openshift-gitops-server-openshift-gitops.apps.example.com
argocd cluster add prd-cluster -y

**Building new image version**

GITHUB_USER=<your_user>
git clone https://github.com/$GITHUB_USER/Openshift-Multi-Cluster-management.git
cd Openshift-Multi-Cluster-management/chapter07
./change-repo-urls.sh
cd ..
git checkout -b dev
vim ./sample-go-app/clouds-api/clouds.go 
    func homePage(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "Welcome to the HomePage! Version=1.0") 
    fmt.Println("Endpoint Hit: homePage") 
    } 
git add ./sample-go-app/clouds-api/clouds.go
git commit -m 'Version 1.0 changes'
git push -u origin dev
oc apply -k ./chapter07/config/cicd
oc apply -f ./chapter07/config/cicd/pipelinerun/build-v1.yaml -n cicd
tkn pipelinerun logs build-v1-pipelinerun -f -n cicd

**Deploying in Development**

sed -i 's/changeme/v1.0/' ./chapter07/clouds-api-gitops/overlays/dev/kustomization.yaml
git add chapter07/clouds-api-gitops/overlays/dev/kustomization.yaml 
git commit -m 'updating kustomization file for v1.0' 
git push -u origin dev 
oc apply -f ./chapter07/config/argocd/argocd-project.yaml
oc apply -f ./chapter07/config/argocd/argocd-app-dev.yaml
echo "$(oc  get route openshift-gitops-server -n openshift-gitops --template='https://{{.spec.host}}')"
oc extract secret/openshift-gitops-cluster -n openshift-gitops --to=-
curl $(oc get route clouds-api -n clouds-api-dev --template='http://{{.spec.host}}')

**Promoting to QA**

git checkout -b qa 
cp -r ./chapter07/clouds-api-gitops/overlays/dev/ ./chapter07/clouds-api-gitops/overlays/qa/ 
sed -i 's/dev/qa/' ./chapter07/clouds-api-gitops/overlays/qa/namespace.yaml ./chapter07/clouds-api-gitops/overlays/qa/kustomization.yaml 
git add ./chapter07/clouds-api-gitops/overlays/qa 
git commit -m 'Promoting v1.0 to QA' 
git push -u origin qa 
oc apply -f ./chapter07/config/argocd/argocd-app-qa.yaml 
curl $(oc get route clouds-api -n clouds-api-qa --template='http://{{.spec.host}}')

**Promoting to Production**

git checkout -b pre-prod 
cp -r chapter07/clouds-api-gitops/overlays/dev/ chapter07/clouds-api-gitops/overlays/prod/ 
sed -i 's/dev/prod/' ./chapter07/clouds-api-gitops/overlays/prod/namespace.yaml ./chapter07/clouds-api-gitops/overlays/prod/kustomization.yaml 
git add ./chapter07/clouds-api-gitops/overlays/prod 
git commit -m 'Promoting v1.0 to Prod' 
git push -u origin pre-prod 
git checkout main 
oc apply -f ./chapter07/config/argocd/argocd-app-prod.yaml 
curl $(oc get route clouds-api -n clouds-api-prod --template='http://{{.spec.host}}')
