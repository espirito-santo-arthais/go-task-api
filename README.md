# Go Task API

Uma API RESTful simples desenvolvida em Go para gerenciamento de tarefas.

## Tecnologias
- Go (Golang)
- Gorilla Mux
- JSON API

## Endpoints

- `GET /tasks` — Lista tarefas
- `POST /tasks` — Cria nova tarefa
- `DELETE /tasks/:id` — Remove tarefa

## Como rodar

```bash
go mod tidy
go run main.go
