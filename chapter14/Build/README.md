
oc apply -f Pipeline/namespace.yaml
oc apply -f Pipeline/quarkus-build-pi.yaml

oc apply -f PipelineRun/quarkus-build-pr.yaml