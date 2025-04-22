# 📦 Projeto Integrador - API VentuFrio

> Software de gerenciamento técnico de manutenção de eletrodomésticos de linha branca.

---


## 📚 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Equipe](#equipe)
- [Organização](#organização)
- [Como Executar](#como-executar)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Documentação Swagger](#documentação-swagger)
- [Endpoints da API](#endpoints-da-api)

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
$DB_HOST=ventufrio-mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=ventufrio

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

api-ventuFrio/
│
├── cmd/                  # Arquivo main.go e rotas
├── internal/
│   ├── db/               # Conexão com o banco de dados
│   ├── models/           # Definições das entidades
│   ├── handlers/         # Handlers das rotas (controllers)
│   ├── repositories/     # Acesso ao banco de dados (camada de dados)
│   └── services/         # Lógica de negócio
├── docs/                 # Arquivos Swagger
├── Dockerfile            # Dockerfile da aplicação
├── docker-compose.yml    # Compose com o banco e a API
└── README.md             # Documentação do projeto


## 📑 Documentação Swagger
