# ============================================================
# Makefile para gerenciar ambiente Docker (app Go + PostgreSQL)
# ============================================================

# Caminhos e variáveis
DOCKER_COMPOSE := sudo docker-compose -f /home/pedro/estudo/facul/volei/socialVoleiAPI/deployments/docker-compose.yml

# =============================
# 🧠 Ajuda / documentação
# =============================
help: ## Mostra esta ajuda
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "----------------------"
	@grep -E '^[a-zA-Z_-]+:.*?##' Makefile | awk 'BEGIN {FS = "##"}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""

# =============================
# 🐳 Comandos Docker
# =============================

build: ## Constrói as imagens Docker
	@$(DOCKER_COMPOSE) build

up: ## Sobe os containers (app + banco)
	@$(DOCKER_COMPOSE) up --build

up-d: ## Sobe em modo background (detached)
	@$(DOCKER_COMPOSE) up --build -d
	@echo "🚀 Containers em execução:"
	@$(DOCKER_COMPOSE) ps

down: ## Para os containers
	@$(DOCKER_COMPOSE) down

logs: ## Mostra os logs do app e banco
	@$(DOCKER_COMPOSE) logs -f

restart: ## Reinicia o ambiente
	@$(DOCKER_COMPOSE) down
	@$(DOCKER_COMPOSE) up --build

clean: ## Remove containers, volumes e imagens órfãs
	@$(DOCKER_COMPOSE) down -v --rmi all --remove-orphans

ps: ## Lista os containers em execução
	@$(DOCKER_COMPOSE) ps
	
db-up: ## Sobe apenas o container do PostgreSQL
	@$(DOCKER_COMPOSE) up -d db
	@echo "🟢 Banco PostgreSQL em execução"
	@$(DOCKER_COMPOSE) ps

run:
	@go run cmd/server/main.go 

# =============================
# 🧰 Auxiliares
# =============================

db-shell: ## Abre um shell dentro do container PostgreSQL
	@docker exec -it postgres_db psql -U app_user -d app_db

app-shell: ## Abre um shell dentro do container da aplicação
	@docker exec -it go_app /bin/bash
