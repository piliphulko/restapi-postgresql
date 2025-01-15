REGISTRY = localhost:5000
OS_VERSION = 0.1.0
PfCO_VERSION = 0.1.0
VP_VESION = 0.1.0

helm-push-outboard-services:

	helm package ./outboard-services

	helm push outboard-services-$(OS_VERSION).tgz oci://$(REGISTRY)/helm-repo

helm-push-prometheus-for-cores-api:

	helm package ./prometheus-for-cores-api

	helm push prometheus-for-cores-api-$(PfCO_VERSION).tgz oci://$(REGISTRY)/helm-repo

helm-push-vector-pipelines:

	helm package ./vector-pipelines
 
	helm push vector-pipelines-$(VP_VESION).tgz oci://$(REGISTRY)/helm-repo