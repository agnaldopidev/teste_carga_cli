APP_NAME := loadtester
DOCKER_IMAGE := loadtester:latest

.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  make build           -> Compila o binário localmente"
	@echo "  make run             -> Executa o CLI localmente"
	@echo "  make test            -> Executa testes unitários"
	@echo "  make docker-build    -> Constrói imagem Docker"
	@echo "  make docker-run      -> Executa container Docker"
	@echo "  make clean           -> Remove binário local"

.PHONY: build
build:
	go build -o $(APP_NAME) ./cmd/cli

.PHONY: run
run: build
	./$(APP_NAME) --url=http://google.com --requests=100 --concurrency=10

.PHONY: test
test:
	go test ./... -v

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	docker run --rm $(DOCKER_IMAGE) \
		--url=http://google.com \
		--requests=100 \
		--concurrency=10

.PHONY: clean
clean:
	rm -f $(APP_NAME)
