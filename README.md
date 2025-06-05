📅 Agendamento API

API REST em Go para cadastro e consulta de agendamentos por empresa, com horário disponível em faixas de 8h às 17h.

🚀 Como executar

1. Clone o repositório

$ git clone https://github.com/flavio-ph/tc

2. Configure o banco de dados

Crie um banco MySQL com as seguintes variáveis (padrão .env):
DB_USER=root
DB_PASSWORD=root
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
Endpoints principais  
POST	http://localhost:8080/agendas	Cria um novo agendamento | GET	http://localhost:8080/agendas	Lista todos os agendamentos | GET	http://localhost:8080/agendas/disponibilidade	Mostra horários disponíveis

🛠 Tecnologias
Go (Golang), MySQL, Gorilla Mux e Docker.

