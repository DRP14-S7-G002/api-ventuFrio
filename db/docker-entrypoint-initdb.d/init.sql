DROP DATABASE IF EXISTS apiVentuFrio;
CREATE DATABASE apiVentuFrio CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE apiVentuFrio;

CREATE TABLE cliente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    telefone VARCHAR(15),
    cpf VARCHAR(11) UNIQUE,
    rua VARCHAR(150),
    numero VARCHAR(45),
    bairro VARCHAR(100),
    cep VARCHAR(8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    log_delete VARCHAR(300)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE agendamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    data_visita DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    cliente_id INT,
    FOREIGN KEY (cliente_id) REFERENCES cliente(id) ON DELETE CASCADE
    );

CREATE TABLE orcamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descricao_inicial VARCHAR(300),
    descricao_item VARCHAR(300),
    status VARCHAR(45),
    prazo_entrega DATE,
    valor DECIMAL(10, 2),
    cliente_id INT,
    Agendamento_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cliente_id) REFERENCES cliente(id),
    FOREIGN KEY (Agendamento_id) REFERENCES agendamento(id) ON DELETE CASCADE
);

CREATE TABLE ordem_servico (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descricao_servico VARCHAR(300),
    status VARCHAR(45),
    responsavel VARCHAR(45),
    Orcamento_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (Orcamento_id) REFERENCES orcamento(id) ON DELETE CASCADE
);

CREATE TABLE material (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(300),
    quantidade VARCHAR(45),
    valor DECIMAL(10, 2),
    ordem_de_servico_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (ordem_de_servico_id) REFERENCES ordem_servico(id) ON DELETE CASCADE
);

SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci;


-- Inclusão de dados de exemplo na tabela cliente

INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Ana Silva', '1199999100', '1234567890', 'Rua Augusta', '10', 'Jardins', '01234100');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Bruno Souza', '1199999101', '1234567891', 'Avenida Paulista', '20', 'Moema', '01234101');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Carlos Oliveira', '1199999102', '1234567892', 'Rua da Consolação', '30', 'Pinheiros', '01234102');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Daniela Santos', '1199999103', '1234567893', 'Rua Oscar Freire', '40', 'Vila Mariana', '01234103');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Eduardo Pereira', '1199999104', '1234567894', 'Rua 25 de Março', '50', 'Liberdade', '01234104');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Fernanda Costa', '1199999105', '1234567895', 'Rua Frei Caneca', '60', 'Bela Vista', '01234105');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Gabriel Rodrigues', '1199999106', '1234567896', 'Rua Bela Cintra', '70', 'Santana', '01234106');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Helena Almeida', '1199999107', '1234567897', 'Rua Haddock Lobo', '80', 'Consolação', '01234107');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Igor Nascimento', '1199999108', '1234567898', 'Avenida Ipiranga', '90', 'Ipiranga', '01234108');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Juliana Lima', '1199999109', '1234567899', 'Rua Vergueiro', '100', 'Aclimação', '01234109');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Kleber Araújo', '1199999110', '9834567890', 'Rua Domingos de Morais', '110', 'Tatuapé', '01234110');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Larissa Fernandes', '1199999111', '8794567891', 'Rua Teodoro Sampaio', '120', 'Brooklin', '01234111');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Marcos Carvalho', '1199999112', '6544567892', 'Rua Pamplona', '130', 'Perdizes', '01234112');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Natália Gomes', '1199999113', '4567567893', 'Rua Itacolomi', '140', 'Lapa', '01234113');
INSERT INTO cliente (nome, telefone, cpf, rua, numero, bairro, cep) 
VALUES ('Otávio Martins', '1199999114', '9634567894', 'Avenida Angélica', '150', 'Cambuci', '01234114');

-- inclusão de dados de exemplo na tabela agendamento

INSERT INTO agendamento (data_visita, cliente_id) VALUES
('2025-05-02', 1), ('2025-05-03', 2), ('2025-05-04', 3),
('2025-05-05', 4), ('2025-05-06', 5), ('2025-05-07', 6),
('2025-05-08', 7), ('2025-05-09', 8), ('2025-05-10', 9),
('2025-05-11', 10), ('2025-05-12', 11), ('2025-05-13', 12),
('2025-05-14', 13), ('2025-05-15', 14), ('2025-05-16', 15);

-- inclusão de dados de exemplo na tabela orcamento

INSERT INTO orcamento (descricao_inicial, descricao_item, status, prazo_entrega, valor, cliente_id, Agendamento_id) 
VALUES ('Motor com ruído incomum', 'Geladeira Electrolux', 'Pendente', '2025-06-02', 150.00, 1, 1),
('Placa eletrônica queimada', 'Fogão Brastemp', 'Cancelado', '2025-06-03', 300.00, 2, 2),
('Vazamento na parte inferior', 'Máquina de Lavar Consul', 'Finalizado', '2025-06-04', 450.00, 3, 3),
('Sistema de resfriamento não funciona', 'Micro-ondas LG', 'Pendente', '2025-06-05', 600.00, 4, 4),
('Aquecimento irregular', 'Lava e Seca Samsung', 'Pendente', '2025-06-06', 750.00, 5, 5),
('Não liga ao pressionar o botão', 'Freezer Vertical Electrolux', 'Pendente', '2025-06-07', 900.00, 6, 6),
('Barulho excessivo na centrifugação', 'Forno Elétrico Fischer', 'Cancelado', '2025-06-08', 1050.00, 7, 7),
('Luz do painel não acende', 'Cooktop 5 bocas Brastemp', 'Cancelado', '2025-06-09', 1200.00, 8, 8),
('Porta desalinhada', 'Geladeira Brastemp Frost Free', 'Em andamento', '2025-06-10', 1350.00, 9, 9),
('Falha no termostato', 'Lava-louças Electrolux', 'Cancelado', '2025-06-11', 1500.00, 10, 10),
('Erro no display digital', 'Forno de Embutir Consul', 'Finalizado', '2025-06-12', 1650.00, 11, 11),
('Cheiro de queimado ao ligar', 'Coifa de Parede Tramontina', 'Cancelado', '2025-06-13', 1800.00, 12, 12),
('Problema no compressor', 'Micro-ondas Panasonic', 'Em andamento', '2025-06-14', 1950.00, 13, 13),
('Fiação interna solta', 'Adega Climatizada Philco', 'Pendente', '2025-06-15', 2100.00, 14, 14),
('Sensor de temperatura inativo', 'Máquina de Lavar Brastemp', 'Pendente', '2025-06-16', 2250.00, 15, 15);

-- inclusão de dados de exemplo na tabela ordem_servico

INSERT INTO ordem_servico (descricao_servico, status, responsavel, Orcamento_id) 
VALUES ('Troca do motor da Geladeira', 'Finalizado', 'Wagner Ventura Gomes Machado', 1),
('Substituição da placa principal', 'Em andamento', 'Wagner Ventura Gomes Machado', 2),
('Reparo em vazamento de água', 'Finalizado', 'Wagner Ventura Gomes Machado', 3),
('Manutenção no sistema de resfriamento', 'Finalizado', 'Wagner Ventura Gomes Machado', 4),
('Ajuste no sistema de aquecimento', 'Cancelado', 'Wagner Ventura Gomes Machado', 5),
('Reparo no botão de acionamento', 'Finalizado', 'Wagner Ventura Gomes Machado', 6),
('Troca de rolamentos e balanceamento', 'Em andamento', 'Wagner Ventura Gomes Machado', 7),
('Verificação do painel de controle', 'Em andamento', 'Wagner Ventura Gomes Machado', 8),
('Alinhamento e substituição da porta', 'Cancelado', 'Wagner Ventura Gomes Machado', 9),
('Troca do termostato', 'Cancelado', 'Wagner Ventura Gomes Machado', 10),
('Substituição do display', 'Em andamento', 'Wagner Ventura Gomes Machado', 11),
('Avaliação e correção de curto interno', 'Pendente', 'Wagner Ventura Gomes Machado', 12),
('Reparo no compressor', 'Pendente', 'Wagner Ventura Gomes Machado', 13),
('Revisão da parte elétrica', 'Cancelado', 'Wagner Ventura Gomes Machado', 14),
('Substituição do sensor de temperatura', 'Cancelado', 'Wagner Ventura Gomes Machado', 15);

-- inclusão de dados de exemplo na tabela material

INSERT INTO material (nome, quantidade, valor, ordem_de_servico_id) 
VALUES ('Motor universal 1/4HP', '2', 376.77, 1),
('Placa eletrônica modelo X100', '3', 351.50, 2),
('Mangueira de drenagem reforçada', '3', 614.16, 3),
('Ventoinha de resfriamento', '3', 221.67, 4),
('Resistência cerâmica 220V', '3', 610.24, 5),
('Botão de acionamento frontal', '2', 295.08, 6),
('Kit rolamento com flange', '1', 243.57, 7),
('Painel de LED 7 segmentos', '2', 384.64, 8),
('Dobradiça de porta reforçada', '2', 478.42, 9),
('Termostato digital', '1', 196.11, 10),
('Display LCD 16x2', '2', 314.12, 11),
('Fusível térmico de segurança', '1', 161.33, 12),
('Compressor 1/3HP', '2', 285.80, 13),
('Cabos elétricos revestidos', '2', 388.04, 14),
('Sensor térmico NTC', '1', 100.59, 15);
