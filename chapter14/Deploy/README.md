
oc apply -f Pipeline/quarkus-build-and-deploy-pi.yaml
oc apply -f Rolebindings/pipeline-cicd-service-rolebinding.yaml
oc apply -f Rolebindings/pipeline-service-role.yaml

oc apply -f PipelineRun/quarkus-build-and-deploy-pr.yaml
