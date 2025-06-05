📅 Agendamento API

API REST em Go para cadastro e consulta de agendamentos por empresa, com horário disponível em faixas de 8h às 17h.

🚀 Como executar

1. Clone o repositório

git clone https://github.com/seu-usuario/agendamento-api.git
cd agendamento-api

2. Configure o banco de dados

Crie um banco MySQL com as seguintes variáveis (padrão .env):
DB_USER=seu-user
DB_PASSWORD=sua-senha
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=agendamento_db
Você pode definir essas variáveis no terminal ou usar um .env.

3. Crie a tabela necessária
Execute o SQL abaixo no seu banco:

CREATE TABLE agendas (
  id INT AUTO_INCREMENT PRIMARY KEY,
  empresa_cnpj VARCHAR(14) NOT NULL,
  horario DATETIME NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

4. Instale as dependências

go mod tidy

5. Execute a API
go run main.go
A API estará disponível em: http://localhost:8080
 - Endpoints principais - 
Método	Rota	Descrição
POST	/agendas	Cria um novo agendamento
GET	/agendas	Lista todos os agendamentos
GET	/agendas/disponibilidade	Mostra horários disponíveis

🛠 Tecnologias
Go (Golang)
MySQL
Gorilla Mux
Docker

