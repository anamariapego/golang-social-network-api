# Golang Social Network API


Este é um projeto de uma **rede social** desenvolvida em **Golang**, que simula funcionalidades básicas de uma rede social, como criação de publicações e interação entre usuários (seguir, parar de seguir, curtir e descurtir publicações, alterar senha). A aplicação inclui uma **API RESTful**, com autenticação e segurança, e utiliza um banco de dados **PostgreSQL** para armazenar os dados e uma **documentação Swagger** das rotas.

## Contexto

Este projeto foi desenvolvido como parte de um **curso de Golang** na plataforma [Udemy](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/?couponCode=BFCPSALE24). Durante o curso, aprendi diversos conceitos e técnicas de desenvolvimento back-end utilizando Golang, e esse projeto implementa todos os conceitos abordados ao longo do aprendizado, além de incluir algumas funcionalidades extras por minha conta.

Apesar de o curso abranger também o desenvolvimento de uma aplicação web, optei por focar no desenvolvimento back-end para aprofundar meus conhecimentos na construção de APIs robusta e eficiente.

## Funcionalidades

**Usuários**

* **CRUD de Usuários**: Criar, buscar, atualizar e deletar usuários.
* **Seguir e Parar de Seguir**: Funcionalidade para seguir e deixar de seguir usuários.
* **Buscar Seguidores e Seguidos**: Consultar os usuários que um usuário está seguindo e os que estão seguindo ele.
* **Alteração de Senha**: Permite ao usuário alterar sua senha com validação e segurança.

**Publicações**

* **CRUD de Publicações**: Criar, buscar, atualizar e deletar publicações.
* **Publicações dos Seguidores**: Consultar as publicações dos usuários que um usuário está seguindo.
* **Curtir e Descurtir Publicações**: Interagir com publicações através de likes e dislikes.

## Tecnologias e Pacotes GoLang Utilizadas

* **Golang**: Linguagem principal para o desenvolvimento da API.
* **PostgreSQL**: Banco de dados relacional para armazenar os dados de usuários e publicações, acessado com o driver `pq`.
* **JWT (JSON Web Tokens)**: Para autenticação e segurança na API, implementado com o pacote `jwt-go`.
* **Hashing de Senhas**: Utilizado para segurança ao armazenar senhas no banco, implementado com `bcrypt` do pacote `x/crypto`.
* **Gorilla Mux**: Framework de roteamento e middleware, usado para criar e organizar as rotas da API.
* **Godotenv**: Para carregar variáveis de ambiente de um arquivo `.env`, facilitando o gerenciamento de configurações.
* **Checkmail**: Para validação de endereços de e-mail durante o registro de usuários.
* **Swagger**:
  - **github.com/swaggo/swag/cmd/swag**: Gera documentação Swagger a partir de comentários no código.
  - **github.com/swaggo/http-swagger**: Serve a documentação Swagger no navegador.

## Estrutura do Projeto

A estrutura do projeto é organizada para manter a separação de responsabilidades, seguindo boas práticas de desenvolvimento.

### Arquivos Principais

- **`main.go`**: Arquivo principal que inicia a aplicação, configurando as rotas e inicializando o servidor.
- **`go.mod`**: Define os pacotes externos e dependências utilizadas no projeto.
- **`go.sum`**: Contém os hashes das versões dos pacotes especificados em `go.mod`.
- **`api.exe`**: Arquivo binário gerado após a compilação do projeto (aplicável apenas em ambientes Windows).
- **`.env`**: Arquivo para armazenar variáveis de ambiente, como credenciais do banco de dados e chaves secretas.
- **`sql/`**: Pasta que contém scripts **SQL** para criação e inserção de dados nas tabelas necessárias (exemplo: usuários, publicações, seguidores).
- **`docs/`**: Pasta que contém a documentação **Swagger** da API.

### Estrutura de Diretórios (dentro de `src`)

- **`src/auth/`**: 
  - Gerenciamento e validação de **tokens JWT** para autenticação de usuários.
  - Exemplo: geração e validação de tokens de acesso.

- **`src/config/`**: 
  - Configurações globais e carregamento de variáveis de ambiente do arquivo `.env`.
  - Exemplo: configuração de porta do servidor, host do banco de dados.

- **`src/controller/`**: 
  - Contém funções que implementam a lógica de cada rota da API.
  - Exemplo: criar usuário, buscar publicações, seguir usuários.

- **`src/database/`**: 
  - Configuração da conexão com o banco de dados.
  - Exemplo: inicialização e pooling de conexões para PostgreSQL.

- **`src/middlewares/`**: 
  - Implementa funções intermediárias que são executadas antes de processar as rotas.
  - Exemplos:
    - Escrever logs no terminal.
    - Verificar se o token de autenticação é válido.

- **`src/models/`**: 
  - Define as estruturas de dados que representam as entidades principais do projeto.
  - Exemplos: estruturas para **Usuário**, **Publicação**, **Seguidor**.

- **`src/repositories/`**: 
  - Responsável por interagir diretamente com o banco de dados.
  - Exemplo: executar queries SQL para criação, leitura, atualização e exclusão de dados.

- **`src/routes/`**: 
  - Configura e organiza as rotas da API RESTful.
  - Exemplo: endpoints para usuários e publicações.

- **`src/security/`**: 
  - Funções relacionadas à segurança, como:
    - Criação e validação de hashes de senhas.
    - Validação de senhas no momento do login.

**Autenticação**

* `POST /login`: Usuário faz login e obtém um token de autenticação.

**Usuários**

Rotas para gerenciamento de usuários, incluindo criação, listagem, atualização e operações de seguimento.

* `POST /users`: Criação de um novo usuário.
* `GET /users`: Listar todos os usuários.
* `GET /users/{id}`: Buscar um usuário por ID.
* `PUT /users/{id}`: Atualizar as informações de um usuário.
* `DELETE /users/{id}`: Deletar um usuário.
* `POST /users/{id}/follower`: Seguir um usuário.
* `POST /users/{id}/stop-follower`: Parar de seguir um usuário.
* `GET /users/{id}/followers`: Buscar os seguidores de um usuário.
* `GET /users/{id}/following`: Buscar os usuários que um usuário está seguindo.
* `POST /users/{id}/update-password`: Atualizar a senha do usuário.

**Publicações**

Rotas para criar, listar, buscar e interagir com publicações.

* `POST /publications`: Criar uma nova publicação.
* `GET /publications`: Listar todas as publicações.
* `GET /publications/{id}`: Buscar uma publicação por ID.
* `PUT /publications/{id}`: Atualizar uma publicação.
* `DELETE /publications/{id}`: Deletar uma publicação.
* `GET /users/{id}/publications`: Buscar todas as publicações de um usuário.
* `POST /publications/{id}/like`: Curtir uma publicação.
* `POST /publications/{id}/dislike`: Descurtir uma publicação.


## Como Rodar o Projeto

### Pré-requisitos

* Golang (>= 1.22)
* PostgreSQL


### Passos para rodar

1. Clone o repositório:

    ```bash
    git clone https://github.com/anamariapego/golang-social-network-api.git
    cd golang-social-network-api
    ```

2. Configure o banco de dados PostgreSQL:

* Crie um banco de dados chamado social_network (ou outro nome de sua preferência).
* Configure as variáveis de ambiente no arquivo `.env` com as credenciais do banco:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=seu_usuario
    DB_PASSWORD=sua_senha
    DB_NAME=social_network
    ```

3. Instale as dependências do projeto:

    ```bash
    go mod tidy
    ```

4. Execute o projeto:

    ```bash
    go run main.go
    ```

    A aplicação estará disponível em: http://localhost:2468.

## Considerações Finais
Esse projeto foi uma excelente oportunidade de aplicar meus conhecimentos de Golang em uma aplicação real e prática. Durante o desenvolvimento, explorei conceitos de segurança, autenticação, manipulação de banco de dados, construção de APIs RESTful e documentação, além de aprender a utilizar diversas ferramentas e pacotes para melhorar a arquitetura e a performance da aplicação.

