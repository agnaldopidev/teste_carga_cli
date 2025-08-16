# Teste de Carga CLI

Aplicação **CLI em Go** para realizar **testes de carga** em serviços web, utilizando **arquitetura limpa**, **Clean Code** e princípios **SOLID**.

O usuário informa:
- **URL** do serviço
- **Número total de requisições**
- **Nível de concorrência**

Ao final, o sistema gera um relatório com:
- Tempo total gasto
- Total de requisições
- Quantidade de status **HTTP 200**
- Distribuição de outros códigos HTTP

---

## 🚀 Funcionalidades

- Execução de requisições HTTP concorrentes
- Controle do número total de requisições
- Coleta de métricas de status HTTP
- Geração de relatório resumido
- Testes unitários com **mock** (sem dependência de internet)
- Suporte à execução via **Docker**

---

## 📂 Estrutura do Projeto
```
├── cmd
│ └── cli # Interface CLI (entrada do programa)
├── internal
│ ├── domain # Entidades e interfaces
│ ├── infra # Implementações (HTTP, Reporter)
│ └── usecase # Regras de negócio (testeCargaCLITester)
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

---

## ⚙️ Uso

## 1️⃣ Instalar e executar localmente
```bash
# Clonar o repositório
git clone https://github.com/agnaldopidev/teste_carga_cli.git
cd testeCargaCLItester
 

# Rodar localmente
go run ./cmd/cli \
  --url=http://google.com \
  --requests=100 \
  --concurrency=10
``` 
## 2️⃣  Executar com Docker
```bash
# Construir imagem
docker build -t testeCargaCLItester .

# Executar teste de carga
docker run --rm testeCargaCLItester \
  --url=http://google.com \
  --requests=1000 \
  --concurrency=10
``` 
## 🧪 Testes
## Rodar todos os testes
```bash
go test ./... -v
``` 
## Saída esperada
```
=== RUN   TesttesteCargaCLI_All200
--- PASS: TesttesteCargaCLI_All200 (0.00s)
=== RUN   TesttesteCargaCLI_MixedResponses
--- PASS: TesttesteCargaCLI_MixedResponses (0.00s)
=== RUN   TesttesteCargaCLI_ConcurrentExecution
--- PASS: TesttesteCargaCLI_ConcurrentExecution (0.00s)
=== RUN   TestReporter_Output
--- PASS: TestReporter_Output (0.00s)
PASS
ok      example.com/loadtest/internal/usecase  0.004s
```

## 📦 Tecnologias

- Go 1.22+

- Arquitetura Limpa

- Clean Code

- Princípios SOLID

- Docker
