DROP DATABASE IF EXISTS apiVentuFrio;
CREATE DATABASE apiVentuFrio;
USE apiVentuFrio;

CREATE TABLE cliente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    sobrenome VARCHAR(150) NOT NULL,
    telefone VARCHAR(15),
    cpf VARCHAR(11) UNIQUE,
    rua VARCHAR(150),
    numero VARCHAR(45),
    bairro VARCHAR(100),
    cep VARCHAR(8),
    criate_at DATE DEFAULT CURRENT_DATE,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    log_delete VARCHAR(300)
);

CREATE TABLE agendamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    data_visita VARCHAR(45),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE orcamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descricao_inicial VARCHAR(300),
    descricao_item VARCHAR(300),
    status VARCHAR(45),
    prazo_entrega DATE,
    cliente_id INT,
    Agendamento_id INT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletado_log VARCHAR(300),
    FOREIGN KEY (cliente_id) REFERENCES cliente(id),
    FOREIGN KEY (Agendamento_id) REFERENCES agendamento(id)
);

CREATE TABLE ordem_servico (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descricao_servico VARCHAR(300),
    status VARCHAR(45),
    responsavel VARCHAR(45),
    Orcamento_id INT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (Orcamento_id) REFERENCES orcamento(id)
);

CREATE TABLE material (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(300),
    quantidade VARCHAR(45),
    ordem_servico_id INT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (ordem_servico_id) REFERENCES ordem_servico(id)
);

