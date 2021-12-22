oc apply -k chapter07/config/cicd
oc apply -f chapter07/config/cicd/pipelinerun/build-deploy-dev-pipelinerun.yaml -n cicd
