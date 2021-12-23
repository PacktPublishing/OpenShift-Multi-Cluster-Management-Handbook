#!/bin/bash
clear
echo "================================================"
echo ""
echo "Change the references to your forked Git repo"
echo "Created by: Giovanni Fontana"
echo ""
echo "================================================"

echo "Input your github username: "
read repo_user
echo ""
echo "Changing ArgoCD and Pipeline manifests to your forked repo (https://github.com/$repo_user/Openshift-Multi-Cluster-management/)"

sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-dev.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-qa.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/argocd/argocd-app-prod.yaml
sed -i "s/PacktPublishing/$repo_user/" ./config/cicd/pipelinerun/build-v1.yaml

echo ""
echo "Manifest files changed. Pushing changes to GitHub (please inform your GitHub user and password if asked)."

git add --quiet ./config/argocd/* ./config/cicd/pipelinerun/build-v1.yaml
git commit --quiet -m 'Changed ArgoCD and Pipeline manifests to forked repo'
git push --quiet -u origin main

echo ""
echo "References changed your forked repository"