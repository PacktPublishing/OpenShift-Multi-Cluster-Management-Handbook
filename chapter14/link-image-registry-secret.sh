#!/bin/bash

echo "================================================"
echo ""
echo "Script to link the image registry secret which contains the credentials with pipeline serviceaccount"
echo "Created by: Giovanni Fontana"
echo ""
echo "================================================"

echo "** Input the full path to where the image registry secret file is located (e.g.: /tmp/demo-ocp-secret.yml): "
read secret_path
echo ""

secret_name=`oc apply -f $secret_path --namespace=chap14-review-cicd --dry-run=client -o jsonpath='{.metadata.name}'`
oc apply -f $secret_path --namespace=chap14-review-cicd
oc patch serviceaccount pipeline -p '{"secrets": [{"name": "'$secret_name'"}]}' --namespace=chap14-review-cicd 

oc secrets link default $secret_name --for=pull -n chap14-review-cicd

echo ""
echo "** Secret created and linked to pipeline SA"
