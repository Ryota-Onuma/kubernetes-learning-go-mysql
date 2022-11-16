# make init
.PHONY: init
init:
	make db-init
	make app-init

# make app-init
.PHONY: app-init
app-init:
	docker image build -t golang-echo-sample:latest app/src/
	kubectl delete -f app --ignore-not-found=true
	kubectl apply -f app/namespace.yaml
	kubectl apply -f app/configMap.yaml,app/secret.yaml,app/deployment.yaml,app/service.yaml

# make db-init
.PHONY: db-init
db-init:
	kubectl apply -f mysql/namespace.yaml
	kubectl apply -f mysql/configMap.yaml,mysql/deployment.yaml,mysql/service.yaml


# make up
.PHONY: up
up:
	kubectl port-forward svc/app 8080:8080 -n app


# make reset
.PHONY: reset
reset:
	kubectl delete -f app --ignore-not-found=true
	kubectl delete -f mysql --ignore-not-found=true
