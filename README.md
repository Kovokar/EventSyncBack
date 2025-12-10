# Projeto API Golang com Gin Framework

Uma API REST desenvolvida em Golang utilizando o framework Gin, seguindo princípios de Clean Architecture.

## 🚀 Tecnologias

- **Golang** - Linguagem de programação
- **Gin** - Framework web HTTP
- **GORM** - ORM para banco de dados
- **JWT** - Autenticação baseada em tokens
- **PostgreSQL** - Banco de dados relacional
<!-- - **Docker** - Containerização -->

## 📁 Estrutura do Projeto

```
projeto-gin/
├── cmd/server/          # Ponto de entrada da aplicação
├── internal/            # Código interno da aplicação
│   ├── controllers/     # Handlers das rotas HTTP
│   ├── models/         # Modelos de dados
│   ├── routes/         # Definição de rotas
│   ├── services/       # Lógica de negócio
│   ├── repositories/   # Camada de dados
│   ├── middleware/     # Middlewares customizados
│   └── config/         # Configurações
├── tests/              # Testes unitários e integração
└── configs/            # Arquivos de configuração
```

## ⚙️ Pré-requisitos

- Go 1.21+
- PostgreSQL 14+
<!-- - Docker (opcional) -->

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/Kovokar/socialVoleiAPI
cd socialVoleiAPI
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

## 🏃‍♂️ Executando a aplicação

### Modo desenvolvimento
```bash
go run cmd/server/main.go
```

<!-- ### Com Docker
```bash
docker-compose up -d
``` -->

<!-- ### Usando Makefile
```bash
make run
``` -->

A aplicação estará disponível em `http://localhost:8080`

## 📚 Endpoints da API

### Autenticação
- `POST /auth/login` - Login do usuário
- `POST /auth/register` - Registro de usuário

### Usuários
- `GET /users` - Lista usuários
- `GET /users/:id` - Busca usuário por ID
- `PUT /users/:id` - Atualiza usuário
- `DELETE /users/:id` - Remove usuário

### Produtos
- `GET /products` - Lista produtos
- `POST /products` - Cria produto
- `GET /products/:id` - Busca produto por ID
- `PUT /products/:id` - Atualiza produto
- `DELETE /products/:id` - Remove produto

<!-- ## 🧪 Testes

Execute todos os testes:
```bash
make test
```

Testes com coverage:
```bash
make test-coverage
```

## 📝 Comandos Make

```bash
make run          # Executa a aplicação
make build        # Compila a aplicação
make test         # Executa testes
make migrate      # Executa migrações
make seed         # Popula banco com dados de teste
make docker-build # Constrói imagem Docker
make clean        # Limpa arquivos temporários
``` -->

## 🌱 Variáveis de Ambiente

```env
# Servidor
PORT=8080
ENV=development

# Banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gin_db

# JWT
JWT_SECRET=sua-chave-secreta
JWT_EXPIRES_IN=24h

```

<!-- ## 🐳 Docker

Para executar com Docker Compose:

```bash
docker-compose up -d
```

Isso irá subir:
- Aplicação Golang na porta 8080
- PostgreSQL na porta 5432
- Redis na porta 6379

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

Seu Nome - [@seuusuario](https://github.com/seuusuario)

---

⭐ Se este projeto te ajudou, considere dar uma estrela! -->