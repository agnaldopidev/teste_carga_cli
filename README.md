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
│ └── usecase # Regras de negócio (LoadTester)
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
git clone https://github.com/agnaldopidev/loadtester.git
cd loadtester
 

# Rodar localmente
go run ./cmd/cli \
  --url=http://google.com \
  --requests=100 \
  --concurrency=10
``` 
## 2️⃣  Executar com Docker
```bash
# Construir imagem
docker build -t loadtester .

# Executar teste de carga
docker run --rm loadtester \
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
=== RUN   TestLoadTester_All200
--- PASS: TestLoadTester_All200 (0.00s)
=== RUN   TestLoadTester_MixedResponses
--- PASS: TestLoadTester_MixedResponses (0.00s)
=== RUN   TestLoadTester_ConcurrentExecution
--- PASS: TestLoadTester_ConcurrentExecution (0.00s)
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
