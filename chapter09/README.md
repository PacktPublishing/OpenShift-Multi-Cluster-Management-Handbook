# Chapter 9 - OpenShift Pipelines - Tekton

List of commands to copy and paste:

**Install the tkn CLI**

```
tar -xvzf tkn-linux-amd64-0.17.2.tar.gz
sudo cp tkn /usr/local/bin
tkn version
```

**Using tkn**

oc login -u <user> https://<ocp-api-url>:6443
tkn clustertasks ls

**Creating a new (custom) Task**

oc new-project pipelines-sample
oc get serviceaccount pipeline
oc apply -f  https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Tasks/apply-manifests.yaml
oc apply -f  https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Tasks/update-image-version.yaml 
oc apply -f  https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Tasks/check-route-health.yaml 
tkn tasks ls 

**TaskRun**

oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/PipelineRun/pvc.yaml
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Tasks/git-clone-taskrun.yaml
tkn taskrun logs git-clone -f
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Tasks/apply-manifests-taskrun.yaml
tkn taskrun logs run-apply-manifests -f

**Pipeline**

oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Pipeline/build-deploy.yaml
tkn pipelines ls

**PipelineRun**

oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/PipelineRun/clouds-api-build-deploy.yaml
tkn pipelinerun logs build-deploy-api-pipelinerun -f

**Using a Trigger with GitHub Webhook**

oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Trigger/clouds-api-tb.yaml
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Trigger/clouds-api-tb.yaml
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Trigger/clouds-api-tt.yaml
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Trigger/clouds-api-trigger.yaml
oc apply -f https://raw.githubusercontent.com/PacktPublishing/Openshift-Multi-Cluster-management/main/chapter06/Trigger/clouds-api-el.yaml
oc expose svc el-clouds-api-el
echo "$(oc  get route el-clouds-api-el --template='http://{{.spec.host}}')"
git commit -m "empty-commit" --allow-empty && git push origin main
