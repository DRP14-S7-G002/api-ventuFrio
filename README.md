# 📦 Projeto Integrador - API VentuFrio

> Software de gerenciamento técnico de manutenção de eletrodomésticos de linha branca.

---


## 📚 Índice

- [Sobre o Projeto](#-sobre-o-projeto)
- [Tecnologias Utilizadas](#-tecnologias-utilizadas)
- [Equipe](#-equipe)
- [Organização](#-organização)
- [Como Executar](#%EF%B8%8F-como-executar)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Documentação Swagger](#-documentação-swagger)
- [Endpoints da API](#-endpoints-da-api)
- [Link - Repositório Frontend](#-repositório-frontend)
- - [Link - Protótipo](#-protótipo)

---

## 🧩 Sobre o Projeto

Este projeto tem como objetivo o desenvolvimento de um sistema completo para técnicos autônomos que prestam manutenção em eletrodomésticos de linha branca. A API permite o gerenciamento de clientes, ordens de serviço, produtos, técnicos e agendamentos.

Projeto desenvolvido como parte do curso da **UNIVESP - Universidade Virtual do Estado de São Paulo**, Polo Capão Redondo.

---

## 💻 Tecnologias Utilizadas

| Ferramenta         | Descrição                                  |
|--------------------|--------------------------------------------|
| [Go](https://golang.org)         | Backend da aplicação              |
| [Gin](https://gin-gonic.com/)   | Framework web para Go             |
| [GORM](https://gorm.io/)        | ORM para Go                       |
| [MySQL](https://www.mysql.com/) | Banco de dados relacional         |
| [Docker](https://www.docker.com/)| Containerização                   |
| [Swagger](https://swagger.io/)  | Documentação da API               |

---

## 👨‍👩‍👧‍👦 Equipe

- Daniela Martins Costa
- Guilherme da Silveira Santos
- Guilherme Fontainha Machado
- Ítalo Oliveira Almeida
- João Vitor Alves Ribeiro
- José Lucas Silva Rodrigues

---


## 🗓 Organização 
Organizamos o projeto utilizando a metodologia ágil Kaban, através da ferramenta Trello. Separamos como Tarefas Backlog, A Fazer, Em Desenvolvimento, Review (Code Review), Concluído e Finished. <a href="https://trello.com/b/LwRp41bi/pi-projeto-integrador">Clique aqui para visualizar o quadro</a></p>

---

## ▶️ Como Executar

Para rodar o projeto, é necessário ter instalado na máquina:

### 1. Docker Desktop
> Download:  
https://www.docker.com/products/docker-desktop/

> Guia de instalação:  
https://docs.docker.com/desktop/install/windows-install/

### 2. Git
> Download:  
https://git-scm.com/downloads

### 3. Go (para desenvolvimento local)
> Download Go:  
https://go.dev/dl/


```bash

#Clone o repositório

git clone https://github.com/DRP14-S7-G002/api-ventuFrio.git

#Acesse a pasta do projeto no terminal
cd api-ventuFrio

#Instale as dependências do projeto
go mod tidy

# Configure o ambiente criando um arquivo .env na raiz do projeto
DB_HOST=ventufrio-mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=apiVentuFrio

#Subir o projeto com Docker
docker compose up --build

# Quando o projeto estiver sendo executado, acesse no navegador
http://localhost:8080/swagger/index.html

# Acesse a API no navegador
http://localhost:8080

#Encerrar os containers
docker compose down

```

## 📁 Estrutura do projeto

```
api-ventuFrio/
├── cmd/
│   └── main.go
│   └── routes/
│       └── routes.go
├── db/
│   └── init.go
│
├── internal/
│   ├── db/
│   │   └── database.go
│   ├── dto/
│   │   ├── agendamento_dto.go
│   │   ├── cliente_dto.go
│   │   ├── material_dto.go
│   │   ├── orcamento_dto.go
│   │   └── ordem_servico_dto.go
│   ├── handler/
│   │   ├── agendamento_handler.go
│   │   ├── cliente_handler.go
│   │   ├── material_handler.go
│   │   ├── orcamento_handler.go
│   │   └── ordem_servico_handler.go
│   ├── models/
│   │   ├── agendamento.go
│   │   ├── cliente.go
│   │   ├── material.go
│   │   ├── orcamento.go
│   │   ├── ordem_de_servico.go
│   ├── repository/
│   │   ├── agendamento_repository.go
│   │   ├── cliente_repository.go
│   │   ├── material_repository.go
│   │   ├── orcamento_repository.go
│   │   └── ordem_servico_repository.go
│   └── service/
│       ├── agendamento_service.go
│       ├── cliente_service.go
│       ├── material_service.go
│       ├── orcamento_service.go
│       └── ordem_servico_service.go
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md

```

## 📑 Documentação Swagger
![image](https://github.com/user-attachments/assets/6713f6e9-9e48-4ba0-9f54-dda3353d1161)

---


## 📚 Endpoints da API



| Entidade           | Método  | Rota                        | Descrição                          |
|--------------------|---------|------------------------------|------------------------------------|
| Cliente            | GET     | `/api/v1/clientes`           | Lista todos os clientes           |
| Cliente            | GET     | `/api/v1/clientes/{id}`      | Busca cliente pelo ID             |
| Cliente            | POST    | `/api/v1/clientes`           | Cria um novo cliente              |
| Cliente            | PUT     | `/api/v1/clientes/{id}`      | Atualiza cliente pelo ID          |
| Cliente            | DELETE  | `/api/v1/clientes/{id}`      | Deleta cliente pelo ID            |
| Agendamento        | GET     | `/api/v1/agendamentos`       | Lista todos os agendamentos       |
| Agendamento        | GET     | `/api/v1/agendamentos/{id}`  | Busca agendamento pelo ID         |
| Agendamento        | POST    | `/api/v1/agendamentos`       | Cria um novo agendamento          |
| Agendamento        | PUT     | `/api/v1/agendamentos/{id}`  | Atualiza agendamento pelo ID      |
| Agendamento        | DELETE  | `/api/v1/agendamentos/{id}`  | Deleta agendamento pelo ID        |
| Orçamento          | GET     | `/api/v1/orcamentos`         | Lista todos os orçamentos         |
| Orçamento          | GET     | `/api/v1/orcamentos/{id}`    | Busca orçamento pelo ID           |
| Orçamento          | POST    | `/api/v1/orcamentos`         | Cria um novo orçamento            |
| Orçamento          | PUT     | `/api/v1/orcamentos/{id}`    | Atualiza orçamento pelo ID        |
| Orçamento          | DELETE  | `/api/v1/orcamentos/{id}`    | Deleta orçamento pelo ID          |
| Ordem de Serviço   | GET     | `/api/v1/ordens`             | Lista todas as ordens de serviço  |
| Ordem de Serviço   | GET     | `/api/v1/ordens/{id}`        | Busca ordem de serviço pelo ID    |
| Ordem de Serviço   | POST    | `/api/v1/ordens`             | Cria uma nova ordem de serviço    |
| Ordem de Serviço   | PUT     | `/api/v1/ordens/{id}`        | Atualiza ordem de serviço pelo ID |
| Ordem de Serviço   | DELETE  | `/api/v1/ordens/{id}`        | Deleta ordem de serviço pelo ID   |
| Material           | GET     | `/api/v1/materiais`          | Lista todos os materiais          |
| Material           | GET     | `/api/v1/materiais/{id}`     | Busca material pelo ID            |
| Material           | POST    | `/api/v1/materiais`          | Cria um novo material             |
| Material           | PUT     | `/api/v1/materiais/{id}`     | Atualiza material pelo ID         |
| Material           | DELETE  | `/api/v1/materiais/{id}`     | Deleta material pelo ID           |

--- 

## 🔗 Repositório Frontend

[LINK - Projeto Webapp VentuFrio](https://github.com/DRP14-S7-G002/webapp-ventuFrio)

---

## 🔗 Protótipo

[LINK - Protótipo](https://www.figma.com/design/1a5X22NvuV6VODQ6GDCKnC/Projeto-Integrador?node-id=0-1&p=f&t=HMP40m3aAZw9vZtR-0)

---

[Início](#-projeto-integrador---api-ventufrio)
