.PHONY: help

run-linux: ## Run on linux
	export DOCKER_INTERNAL_IP="$(shell ip route | grep docker0 | awk '{print $$9}')"; \
    	sudo -E ADMIN_USER=admin ADMIN_PASSWORD=admin docker-compose -f docker-compose-promgrafana.yml up -d

run-mac: ## Run on mac
	ADMIN_USER=admin ADMIN_PASSWORD=admin docker-compose -f docker-compose-promgrafana.yml up -d