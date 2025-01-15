include env_secrets.env
export

REGISTRY = localhost:5000
IMAGE_SERVICE_AA = service-acct-auth:V0.0.2
IMAGE_VAULT_SERVER =vault-server:V0.1.0
PORT_MINIO = 9000
HELM_SAA_VERSIOM = 0.1.0 # chart version | service-acct-auth
HELM_OS_VERSIOM = 0.1.0 # chart version | outboard-services
HELM_PfCA_VERSIOM = 0.1.0 # chart version | prometheus-for-cores-api
HELM_VP_VERSIOM = 0.1.0 # chart version | vector-pipelines

docker-build-service-acct-auth:
	docker build -t $(IMAGE_SERVICE_AA) -f cmd/service-acct-auth/Dockerfile .

docker-image-put-registry-service-acct-auth:
	docker tag $(IMAGE_SERVICE_AA) $(REGISTRY)/$(IMAGE_SERVICE_AA)
	docker push $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-service-acct-auth:
	docker run -d --name test-$(IMAGE_NAME) -p 50051:50051 $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-registry:
	docker run -d -p 5000:5000 --name registry -v registry-data:/var/lib/registry registry:2

docker-run-minio:
	docker run -d --name minio -p $(PORT_MINIO):$(PORT_MINIO) -p 9001:9001 -e "MINIO_ROOT_USER=admin" -e "MINIO_ROOT_PASSWORD=secretpassword" minio/minio server /data --console-address ":9001"

docker-run-prometheus:
	docker run -d --name prometheus -p 9090:9090 -v "$(PWD)/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yaml" -v prometheus-data:/prometheus prom/prometheus   --config.file=/etc/prometheus/prometheus.yaml --storage.tsdb.retention.time=2h --storage.tsdb.min-block-duration=2h --storage.tsdb.max-block-duration=2h

docker-run-thanos-sidecar:
	docker run -d --name thanos-sidecar -p 10902:10902 -v prometheus-data:/prometheus -v "$(PWD)/prometheus/thanos/thanos.yaml:/etc/thanos/thanos.yaml"  thanosio/thanos:main-2024-09-02-dfeaf6e sidecar --tsdb.path=/prometheus --objstore.config-file=/etc/thanos/thanos.yaml --prometheus.url=http://host.docker.internal:9090

# vault server
docker-build-vault-server:
	docker build -t $(IMAGE_VAULT_SERVER) -f vault/Dockerfile .

docker-image-put-registry-vault-server:
	docker tag $(IMAGE_VAULT_SERVER) $(REGISTRY)/$(IMAGE_VAULT_SERVER)
	docker push $(REGISTRY)/$(IMAGE_VAULT_SERVER)

docker-start-vault-server:
#	docker run -d --cap-add IPC_LOCK -p 8200:8200 -p 8201:8201 --name vault-server -e VAULT_POSTGRES_URL=$(VAULT_POSTGRES_URL) -v $(PWD)\vault\certificate.crt:/usr/local/bin/certificate.crt:ro -v $(PWD)\vault\private.key:/usr/local/bin/private.key:ro -v $(PWD)\vault\config.hcl:/vault/config/config.hcl:ro vault:1.13.3 server -config=/vault/config/config.hcl
	docker run -d --cap-add IPC_LOCK -p 8200:8200 -p 8201:8201 --name vault-server -e VAULT_POSTGRES_URL=$(VAULT_POSTGRES_URL) -v $(PWD)\config\certificate.crt:/usr/local/bin/certificate.crt:ro -v $(PWD)\config\private.key:/usr/local/bin/private.key:ro -v $(PWD)\vault\config.hcl:/vault/config/config.hcl:ro $(REGISTRY)/$(IMAGE_VAULT_SERVER)

# launching the greptimedb
docker-start-greptimedb:
	docker run -p 127.0.0.1:4000-4003:4000-4003 -v "$(PWD)/greptimedb:/tmp/greptimedb" --name greptime --rm greptime/greptimedb:latest standalone start --http-addr 0.0.0.0:4000 --rpc-addr 0.0.0.0:4001 --mysql-addr 0.0.0.0:4002 --postgres-addr 0.0.0.0:4003

# CHARTS ARE AS FOLLOWS:

# service-acct-auth
helm-install-service-acct-auth:
	helm install release-saa oci://$(REGISTRY)/helm-repo/service-acct-auth --version $(HELM_SAA_VERSIOM)

helm-uninstall-service-acct-auth:
	helm uninstall release-saa

# outboard-services
helm-install-outboard-services:
	helm install release-test oci://$(REGISTRY)/helm-repo/outboard-services --version $(HELM_OS_VERSIOM) -n outboard-services

helm-uninstall-outboard-services:
	helm uninstall release-test -n outboard-services

# prometheus-for-cores-api
helm-install-prometheus-for-cores-api:
	helm install release-test oci://$(REGISTRY)/helm-repo/prometheus-for-cores-api --version $(HELM_PfCA_VERSIOM) -n prometheus-for-cores-api

helm-uninstall-prometheus-for-cores-api:
	helm uninstall release-test -n prometheus-for-cores-api

# vector-pipelines
helm-install-vector-pipelines:
	helm install release-test oci://$(REGISTRY)/helm-repo/vector-pipelines --version $(HELM_VP_VERSIOM) -n vector-pipelines

helm-uninstall-vector-pipelines:
	helm uninstall release-test -n vector-pipelines
