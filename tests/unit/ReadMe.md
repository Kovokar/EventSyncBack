# Todos os testes
go test ./tests/unit/... -v

# Apenas repositories
go test ./tests/unit/repositories/... -v

# Apenas services
go test ./tests/unit/services/... -v

# Apenas controllers
go test ./tests/unit/controllers/... -v