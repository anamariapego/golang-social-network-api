-- Criando o banco de dados
-- Database: SocialNetwork

CREATE DATABASE "SocialNetwork"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Portuguese_Brazil.1252'
    LC_CTYPE = 'Portuguese_Brazil.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Verificando se existe a tabela e criando a tabela usuário
DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;


CREATE TABLE users(
    id              SERIAL PRIMARY KEY,
    name_user       VARCHAR(50) NOT NULL,
    nick            VARCHAR(50) NOT NULL UNIQUE,
    email           VARCHAR(50) NOT NULL UNIQUE,
    password_user   VARCHAR(100) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE followers (
    user_id     INT NOT NULL,
    follower_id INT NOT NULL,
    PRIMARY KEY (user_id, follower_id), -- chave primária composta

    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE,

    FOREIGN KEY (follower_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);

CREATE TABLE publications (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(50) NOT NULL,
    text        VARCHAR(500) NOT NULL,
    author_id   INT NOT NULL,
    likes       INT DEFAULT 0,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (author_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);
