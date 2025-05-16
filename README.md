# Challenge Stress Test

## 🧪 Objetivo

Este projeto fornece uma aplicação de linha de comando (CLI) escrita em Go para executar **testes de carga em serviços web**. O objetivo é simular múltiplas requisições concorrentes a uma URL, coletar estatísticas e apresentar um relatório claro e útil.

---

## 🚀 Como Usar

### Via Docker

```bash
docker build -t stress-test .
docker run stress-test --url=http://example.com --requests=1000 --concurrency=10
```

### Via Go (local)

```bash
go run ./cmd/stress-cli --url=http://example.com --requests=1000 --concurrency=10
```

---

## ⚙️ Parâmetros da Linha de Comando

| Flag            | Descrição                               | Obrigatório |
|------------------|-------------------------------------------|-------------|
| `--url`         | URL do serviço a ser testado              | ✅           |
| `--requests`    | Total de requisições a serem realizadas   | ✅           |
| `--concurrency` | Número de requisições simultâneas         | ✅           |

---

## 📊 Relatório Gerado

Ao final da execução, será exibido no console um relatório contendo:

- ⏱️ Tempo total de execução
- ✅ Total de requisições realizadas
- 📗 Quantidade de respostas com status HTTP 200
- 📉 Distribuição dos demais códigos HTTP (404, 500, etc.)

---

## 🧱 Estrutura do Projeto

```
challenge-stress-test/
├── cmd/stress-cli/              # Ponto de entrada principal
├── internal/
│   ├── domain/                  # Contratos e interfaces
│   ├── application/             # Lógica de negócio (caso de uso)
│   └── infra/                   # Implementações técnicas
│       ├── httpclient/          # Cliente HTTP
│       ├── reporter/            # Relatório de resultados
│       └── runner/              # Execução concorrente
├── tests/                       # Testes unitários
├── Dockerfile                   # Docker
├── Makefile                     # Build, test e execução
├── go.mod / go.sum              # Gerenciamento de dependências
└── README.md                    # Documentação
```

---

## 🧪 Testes

Execute os testes com:

```bash
make test
```

---

## 🛠️ Boas Práticas

O projeto segue:

- Clean Code
- Clean Architecture (inspirada na Hexagonal)
- Design Patterns como Injeção de Dependência e Strategy
- Testes unitários para lógica de negócio
- Docker para facilitar execução
- Makefile para automação de build, testes e execução

---

## 👨‍💻 Autor

[Well Alencar](https://github.com/wellalencarweb)

---