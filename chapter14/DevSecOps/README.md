
oc apply -f Task/s2i-java.yaml
oc apply -f Pipeline/quarkus-devsecops-v1-pi.yaml
oc apply -f Pipeline/quarkus-devsecops-v2-pi.yaml

oc apply -f PipelineRun/quarkus-devsecops-v1-pr.yaml
oc apply -f PipelineRun/quarkus-devsecops-v2-pr.yaml