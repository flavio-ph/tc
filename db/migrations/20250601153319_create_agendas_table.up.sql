-- +goose Up
CREATE TABLE IF NOT EXISTS agendas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    empresa_cnpj VARCHAR(14) NOT NULL,
    horario DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_empresa_cnpj (empresa_cnpj),
    UNIQUE (empresa_cnpj, horario)
);

-- +goose Down
DROP TABLE IF EXISTS agendas;
