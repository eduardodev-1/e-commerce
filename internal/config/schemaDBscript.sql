
-- Criar tabelas principais
CREATE SCHEMA IF NOT EXISTS adm;

CREATE TABLE adm.usuario (
                             id serial PRIMARY KEY,
                             username varchar NOT NULL,
                             "password" varchar NOT NULL,
                             enabled bool DEFAULT true,
                             account_expired bool DEFAULT FALSE NOT null,
                             account_locked bool default FALSE NOT null,
                             password_expired bool default FALSE NOT null,
                             projeto_padrao_id int NOT null,
                             tenant_id int NOT null
);

CREATE TABLE adm.projeto (
                             id serial PRIMARY KEY,
                             nome varchar NOT NULL,
                             longitude numeric(19, 6) NOT NULL,
                             latitude numeric(19, 6) NOT NULL,
                             zoom_inicial int NOT NULL,
                             tenant_id int NOT NULL
);

CREATE TABLE adm.grupo (
                           id serial PRIMARY KEY,
                           nome VARCHAR(255),
                           tenant_id int NOT NULL
);

CREATE TABLE adm.atribuicao (
                                id serial PRIMARY KEY,
                                authority varchar NOT NULL,
                                tenant_id int NOT NULL
);

CREATE TABLE adm.usuario_grupo (
                                   usuario_id int NOT NULL,
                                   grupo_id int NOT NULL,
                                   tenant_id int NOT NULL,
                                   CONSTRAINT fk_usuario FOREIGN KEY (usuario_id) REFERENCES adm.usuario(id),
                                   CONSTRAINT fk_grupo FOREIGN KEY (grupo_id) REFERENCES adm.grupo(id),
                                   CONSTRAINT uc_usuario_grupo UNIQUE (usuario_id, grupo_id, tenant_id)
);

CREATE TABLE adm.grupo_atribuicao (
                                      grupo_id int NOT NULL,
                                      atribuicao_id int NOT NULL,
                                      tenant_id int NOT NULL,
                                      CONSTRAINT fk_grupo FOREIGN KEY (grupo_id) REFERENCES adm.grupo(id),
                                      CONSTRAINT fk_atribuicao FOREIGN KEY (atribuicao_id) REFERENCES adm.atribuicao(id),
                                      CONSTRAINT uc_grupo_atribuic UNIQUE (grupo_id, atribuicao_id, tenant_id)
);

CREATE TABLE IF NOT EXISTS tb_wms (
                                      id SERIAL PRIMARY KEY,
                                      url VARCHAR(255),
                                      tipo VARCHAR(255),
                                      nome VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tb_categoria (
                                            id SERIAL PRIMARY KEY,
                                            nome VARCHAR(255),
                                            ordem INT,
                                            categoria_pai_id INT REFERENCES tb_categoria(id)
);

CREATE TABLE IF NOT EXISTS tb_estilo (
                                         id SERIAL PRIMARY KEY,
                                         nome VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tb_layer_capability (
                                                   id SERIAL PRIMARY KEY,
                                                   nome VARCHAR(255),
                                                   wms_id INT REFERENCES tb_wms(id)
);

CREATE TABLE IF NOT EXISTS tb_layer (
                                        id SERIAL PRIMARY KEY,
                                        nome VARCHAR(255),
                                        tipo VARCHAR(255),
                                        extent VARCHAR(255),
                                        layer_wms VARCHAR(255),
                                        wms_id INT REFERENCES tb_wms(id),
                                        tiled BOOLEAN,
                                        opacidade INT,
                                        ordem INT,
                                        visible BOOLEAN,
                                        layer_capability_id INT REFERENCES tb_layer_capability(id),
                                        estilo_id INT REFERENCES tb_estilo(id),
                                        tabela VARCHAR(255),
                                        categoria_id INT REFERENCES tb_categoria(id)
);

CREATE TABLE IF NOT EXISTS tb_consulta (
                                           id SERIAL PRIMARY KEY,
                                           nome VARCHAR(255),
                                           principal BOOLEAN,
                                           query VARCHAR(255),
                                           campos JSON,
                                           layer_id INT REFERENCES tb_layer(id)
);
CREATE TABLE IF NOT EXISTS tb_consulta_atribuicao (
                                                      consulta_id INT NOT NULL,
                                                      atribuicao_id INT NOT NULL,
                                                      PRIMARY KEY (consulta_id, atribuicao_id),
                                                      FOREIGN KEY (consulta_id) REFERENCES tb_consulta(id),
                                                      FOREIGN KEY (atribuicao_id) REFERENCES adm.atribuicao(id)
);

CREATE TABLE IF NOT EXISTS tb_layer_base (
                                             id SERIAL PRIMARY KEY,
                                             tipo VARCHAR(255),
                                             chave VARCHAR(255),
                                             nome VARCHAR(255),
                                             ativo BOOLEAN
);

CREATE TABLE IF NOT EXISTS tb_formulario (
                                             id SERIAL PRIMARY KEY,
                                             titulo VARCHAR(255),
                                             tabela VARCHAR(255),
                                             tipo_geometria VARCHAR(255),
                                             campos VARCHAR(255),
                                             layer_id INT REFERENCES tb_layer(id),
                                             controller_url VARCHAR(255)
);

-- Criar tabela de associação many-to-many entre tb_layer e adm.atribuicao
CREATE TABLE IF NOT EXISTS tb_layer_atribuicao (
                                                   layer_id INT NOT NULL,
                                                   atribuicao_id INT NOT NULL,
                                                   PRIMARY KEY (layer_id, atribuicao_id),
                                                   FOREIGN KEY (layer_id) REFERENCES tb_layer(id),
                                                   FOREIGN KEY (atribuicao_id) REFERENCES adm.atribuicao(id)
);
CREATE TABLE tb_memorial (
                             id SERIAL PRIMARY KEY,
                             arquivo BYTEA,
                             nome_do_arquivo VARCHAR,
                             comunidade VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS tb_certidao_legitimacao (
                                                       id SERIAL PRIMARY KEY,
                                                       identificacao_imovel VARCHAR NOT NULL,
                                                       proprietario VARCHAR NOT NULL,
                                                       matricula VARCHAR,
                                                       area_contida_na_letigimacao VARCHAR,
                                                       area_remanescente VARCHAR,
                                                       area_total VARCHAR,
                                                       usuario_id INT REFERENCES adm.usuario(id),
                                                       resposta_solicitacao_legitimacao_id INT REFERENCES tb_resposta_solicitacao_legitimacao(id),
                                                       situacao_legitimacao VARCHAR,
                                                       observacao VARCHAR
);
CREATE TABLE IF NOT EXISTS tb_resposta_solicitacao_legitimacao (
                                                                   id SERIAL PRIMARY KEY,
                                                                   numero INTEGER NOT NULL,
                                                                   data_oficio DATE DEFAULT CURRENT_DATE NOT NULL,
                                                                   usuario_id INT REFERENCES adm.usuario(id),
                                                                   arquivo BYTEA,
                                                                   nome_do_arquivo VARCHAR,
                                                                   solicitacao_legitimacao_id INT REFERENCES tb_solicitacao_legitimacao(id),
                                                                   status VARCHAR,
                                                                   observacao VARCHAR
);
CREATE TABLE IF NOT EXISTS tb_solicitacao_legitimacao (
                                                          id SERIAL PRIMARY KEY,
                                                          numero INTEGER NOT NULL,
                                                          data_oficio DATE DEFAULT CURRENT_DATE NOT NULL,
                                                          comunidade VARCHAR(255) NOT NULL,
                                                          usuario_id INT REFERENCES adm.usuario(id),
                                                          arquivo BYTEA,
                                                          nome_do_arquivo VARCHAR,
                                                          cartorio VARCHAR
);
CREATE TABLE IF NOT EXISTS tb_processamento (
                                                id SERIAL PRIMARY KEY,
                                                comunidade VARCHAR(255) NOT NULL,
                                                ponto_campo VARCHAR(255) NOT NULL,
                                                responsavel_tecnico INTEGER REFERENCES adm.usuario(id),
                                                precisao_x DOUBLE PRECISION,
                                                precisao_y DOUBLE PRECISION,
                                                precisao_z DOUBLE PRECISION,
                                                altura_antena DOUBLE PRECISION,
                                                descricao VARCHAR(255),
                                                localizacao VARCHAR(255),
                                                rbmc1 VARCHAR(255),
                                                rbmc2 VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tb_anexo_processamento (
                                                      id SERIAL PRIMARY KEY,
                                                      tipo VARCHAR NOT NULL,
                                                      arquivo BYTEA NOT NULL,
                                                      nome_do_arquivo VARCHAR NOT NULL,
                                                      processamento_id INT NOT NULL,
                                                      CONSTRAINT fk_processamento
                                                          FOREIGN KEY (processamento_id)
                                                              REFERENCES tb_processamento (id)
);
CREATE TABLE cadastro.titular (
                                  id SERIAL PRIMARY KEY,
                                  nome_titulo BOOLEAN,
                                  unidade_id INTEGER NOT NULL REFERENCES cadastro.unidade(id),
                                  pessoa_id INTEGER NOT NULL REFERENCES bureau.pessoa(id),
                                  principal BOOLEAN,
                                  parentesco VARCHAR(255),
                                  parentesco_outro VARCHAR(255),
                                  representado_legalmente BOOLEAN
);