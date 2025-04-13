<h1 align="center"> # Projeto Integrador - API VentuFrio</h1>

# Índice

- [Sobre o Projeto](#--sobre-projeto)
- [Equipe do Projeto](#--equipe-do-projeto--)
- [Tecnologias utilizadas](#--tecnologias-utilizadas-)
- [Organização](#--organização-)
- [Execução do Projeto](#--execução-o-projeto-)

##

<br>
<h2> 👨🏻‍💻 Sobre projeto</h2>
<p>Este projeto visa o desenvolvimento de um software para otimizar a organização de demandas de serviços técnicos de manutenção de eletrodomésticos de linha branca, como solução de um problema, que ajude os profissionais autônomos a organizar e administrar melhor seus empreendimentos e serviços prestados e a complementação regimental dos Cursos do Eixo de Tecnologia apresentado à Universidade Virtual do Estado de São Paulo (UNIVESP), do Polo Capão Redondo.</p>
<br>

##

<h2> 👩‍💻 Equipe do projeto </h2>

- Daniela Martins Costa
- Guilherme da Silveira  Santos
- Guilherme Fontainha Machado
- Ítalo Oliveira Almeida
- João Vitor Alves Ribeiro
- José Lucas Silva Reis
- Luiz Henrique Gomes Santos
- Rafael Peixoto de Carvalho


##


<h2> 💻 Tecnologias utilizadas: </h2>

- [Golang 1.22.1](https://go.dev/)
- [MySQL - usado via Docker](https://dev.mysql.com/downloads/mysql/)
- [GORM - ORM](https://gorm.io/)
- [Gin - Router HTTP](https://github.com/gin-gonic/gin)
- [Docker](https://www.docker.com/get-started/)
- [Swagger](https://swagger.io/)


##


<h2> 🗓 Organização </h2>
<p>Organizamos o projeto utilizando a metodologia ágil Kaban, através da ferramenta Trello. Separamos como Tarefas Backlog, A Fazer, Em Desenvolvimento, Review (Code Review), Concluído e Finished. <a href="https://trello.com/b/LwRp41bi/pi-projeto-integrador">Clique aqui para visualizar o quadro</a></p>

##

<h2> 🎲 Execução o projeto </h2>

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
$ git clone https://github.com/DRP14-S7-G002/api-ventuFrio.git

#Acesse a pasta do projeto no terminal
$ cd api-ventuFrio

#Instale as dependências do projeto
$ go mod tidy

# Configure o ambiente criando um arquivo .env na raiz do projeto
$ DB_USER=root
$ DB_PASSWORD=root
$ DB_HOST=mysql
$ DB_PORT=3306
$ DB_NAME=apiVentuFrio

#Subir o projeto com Docker
$ docker compose up --build

# Quando o projeto estiver sendo executado, acesse no navegador
$ http://localhost:8080/swagger/index.html


#Encerrar os containers
$ docker compose down

```



