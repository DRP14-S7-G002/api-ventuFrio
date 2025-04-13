CREATE DATABASE IF NOT EXISTS apiVentuFrio;
USE apiVentuFrio;

CREATE TABLE IF NOT EXISTS Enderecos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    logradouro VARCHAR(150) NOT NULL,
    bairro VARCHAR(100) NOT NULL,
    numero VARCHAR(10),
    cidade VARCHAR(100) NOT NULL,
    estado VARCHAR(2) NOT NULL,
    cep CHAR(9) NOT NULL
);

CREATE TABLE IF NOT EXISTS Clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    sobrenome VARCHAR(255) NOT NULL,
    cpf CHAR(14) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    endereco_id INT,
    FOREIGN KEY (endereco_id) REFERENCES Enderecos(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS Orcamentos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    valor_total DECIMAL(10, 2) NOT NULL,
    prazo_entrega DATE NOT NULL,
    admin_user_name VARCHAR(50) NOT NULL,
    status_orcamento VARCHAR(50) NOT NULL,
    pedido_id_orcamento INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Agendamentos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_id INT,
    endereco_id INT,
    data DATE NOT NULL,
    hora TIME NOT NULL,
    orcamento_id INT,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (cliente_id) REFERENCES Clientes(id) ON DELETE CASCADE,
    FOREIGN KEY (endereco_id) REFERENCES Enderecos(id) ON DELETE CASCADE,
    FOREIGN KEY (orcamento_id) REFERENCES Orcamentos(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS Admin (
    id VARCHAR(16) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(32) NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO Enderecos (logradouro, bairro, numero, cidade, estado, cep) VALUES
('Rua das Flores', 'Centro', '123', 'São Paulo', 'SP', '01001-000'),
('Av. Brasil', 'Jardins', '456', 'Rio de Janeiro', 'RJ', '20001-000');

INSERT INTO Clientes (nome, sobrenome, cpf, telefone, email, endereco_id) VALUES
('João', 'Silva', '123.456.789-00', '11999999999', 'joao@email.com', 1),
('Maria', 'Oliveira', '987.654.321-00', '21988888888', 'maria@email.com', 2);

