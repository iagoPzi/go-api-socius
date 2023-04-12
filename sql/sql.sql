CREATE DATABASE IF NOT EXISTS devbook;

-- USE devbook;

-- DROP TABLE IF EXISTS users;

-- CREATE TABLE

--     users(

--         id INT AUTO_INCREMENT PRIMARY KEY,

--         nome VARCHAR(50) NOT NULL,

--         nick VARCHAR(50) NOT NULL UNIQUE,

--         email VARCHAR(50) NOT NULL UNIQUE,

--         senha VARCHAR(50) NOT NULL UNIQUE,

--         createdAt TIMESTAMP DEFAULT current_timestamp()

--     ) ENGINE = INNODB;

DROP TABLE IF EXISTS users;

CREATE TABLE
    users(
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        nome VARCHAR(50) NOT NULL,
        nick VARCHAR(50) NOT NULL UNIQUE,
        email VARCHAR(50) NOT NULL UNIQUE,
        senha VARCHAR(50) NOT NULL,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

DROP TABLE IF EXISTS seguidores;

CREATE TABLE
    seguidores(
        user_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        seguidor_id INTEGER NOT NULL,
        FOREIGN KEY (seguidor_id) REFERENCES users(id) ON DELETE CASCADE,
        PRIMARY KEY (user_id, seguidor_id)
    );

DROP TABLE IF EXISTS publicacoes;

CREATE TABLE
    publicacoes(
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        titulo VARCHAR(50) NOT NULL,
        conteudo VARCHAR(30) NOT NULL,
        autor_id INTEGER NOT NULL,
        FOREIGN KEY (autor_id) REFERENCES users(id) ON DELETE CASCADE,
        curtidas INTEGER DEFAULT 0,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );