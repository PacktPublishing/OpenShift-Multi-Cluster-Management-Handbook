#!/bin/bash

echo "================================================"
echo ""
echo "Change the references to your forked Git repo"
echo "Created by: Giovanni Fontana"
echo ""
echo "================================================"

echo "** Input your github username: "
read repo_user
echo ""
echo "** Changing ArgoCD and Pipeline manifests to your forked repo (https://github.com/$repo_user/Openshift-Multi-Cluster-management/)"

sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-dev.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-qa.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-prod.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/cicd/pipelinerun/build-v1.yaml

echo ""
echo "** Manifest files changed. Pushing changes to GitHub (please inform your GitHub user and password if asked)."

git add ./config/argocd/* ./config/cicd/pipelinerun/build-v1.yaml >>/dev/null
git commit -m 'Changed ArgoCD and Pipeline manifests to forked repo' >>/dev/null
git push -u origin main >>/dev/null

echo ""
echo "** References changed to your forked repository"