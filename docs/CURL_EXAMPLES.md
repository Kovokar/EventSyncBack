# Exemplos de CURL para API de Usuários

Base URL: `http://localhost:8081/api/user`

## PATCH - Atualização Parcial (UpdateUser)

Atualiza apenas os campos fornecidos no body da requisição.

### Exemplo 1: Atualizar apenas o nome
```bash
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Nome"
  }'
```

### Exemplo 2: Atualizar email e telefone
```bash
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "novoemail@example.com",
    "phone": "9876543210"
  }'
```

### Exemplo 3: Atualizar senha
```bash
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "password": "novaSenha123"
  }'
```

### Exemplo 4: Atualizar múltiplos campos
```bash
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao.silva@example.com",
    "phone": "11987654321",
    "gender": "male",
    "visible_in_public_list": true
  }'
```

### Exemplo 5: Atualizar data de nascimento
```bash
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "birthdate": "1995-05-15T00:00:00Z"
  }'
```

## PUT - Atualização Completa (PutUser)

Atualiza todos os campos do usuário. Todos os campos obrigatórios devem ser fornecidos.

### Exemplo 1: Atualização completa básica
```bash
curl -X PUT http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao.silva@example.com",
    "password": "novaSenha123",
    "birthdate": "1990-01-15T00:00:00Z",
    "phone": "11987654321",
    "gender": "male",
    "photo": "https://example.com/photo.jpg",
    "visible_in_public_list": true
  }'
```

### Exemplo 2: Atualização completa com todos os campos
```bash
curl -X PUT http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Maria Santos",
    "email": "maria.santos@example.com",
    "password": "senhaSegura456",
    "birthdate": "1992-08-20T00:00:00Z",
    "phone": "11912345678",
    "gender": "female",
    "photo": "https://example.com/maria.jpg",
    "visible_in_public_list": false
  }'
```

### Exemplo 3: Atualização com gender "other"
```bash
curl -X PUT http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alex Costa",
    "email": "alex.costa@example.com",
    "password": "senha123",
    "birthdate": "1988-12-10T00:00:00Z",
    "phone": "11999999999",
    "gender": "other",
    "photo": "",
    "visible_in_public_list": true
  }'
```

## DELETE - Deletar Usuário (DeleteUser)

Realiza soft delete do usuário (marca como deletado, mas não remove do banco).

### Exemplo 1: Deletar usuário por ID
```bash
curl -X DELETE http://localhost:8081/api/user/1
```

### Exemplo 2: Deletar usuário com verbose
```bash
curl -X DELETE http://localhost:8081/api/user/1 -v
```

### Exemplo 3: Deletar usuário e mostrar headers
```bash
curl -X DELETE http://localhost:8081/api/user/1 -i
```

## Respostas Esperadas

### PATCH - Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Novo Nome",
  "email": "usuario@example.com",
  "phone": "1234567890",
  "gender": "male",
  "photo": "",
  "visible_in_public_list": true,
  "birthdate": "1990-01-01T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-02T00:00:00Z"
}
```

### PATCH - Erro 404 (Usuário não encontrado)
```json
{
  "error": "usuário não encontrado"
}
```

### PATCH - Erro 400 (Validação)
```json
{
  "error": "Key: 'UpdateUserRequest.Name' Error:Field validation for 'Name' failed on the 'min' tag"
}
```

### PUT - Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao.silva@example.com",
  "phone": "11987654321",
  "gender": "male",
  "photo": "https://example.com/photo.jpg",
  "visible_in_public_list": true,
  "birthdate": "1990-01-15T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-02T00:00:00Z"
}
```

### PUT - Erro 400 (Campos obrigatórios faltando)
```json
{
  "error": "Key: 'CreateUserRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

### DELETE - Sucesso (204 No Content)
```
(Resposta vazia, status 204)
```

### DELETE - Erro 404 (Usuário não encontrado)
```json
{
  "error": "usuário não encontrado"
}
```

## Notas Importantes

1. **PATCH**: Apenas os campos fornecidos serão atualizados. Campos não fornecidos permanecem inalterados.

2. **PUT**: Todos os campos obrigatórios devem ser fornecidos, mesmo que não estejam sendo alterados.

3. **DELETE**: Realiza soft delete. O usuário não será removido fisicamente do banco de dados, apenas marcado como deletado.

4. **Formato de Data**: Use o formato ISO 8601: `YYYY-MM-DDTHH:MM:SSZ` ou `YYYY-MM-DDTHH:MM:SS+00:00`

5. **Gender**: Valores aceitos: `"male"`, `"female"`, `"other"`

6. **Validações**:
   - `name`: mínimo 3 caracteres, máximo 100
   - `email`: deve ser um email válido
   - `phone`: mínimo 9 caracteres, máximo 20
   - `password`: será hasheado automaticamente

## Exemplos com Autenticação (se necessário)

Se a API requer autenticação via JWT:

```bash
# PATCH com token
curl -X PATCH http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_JWT_AQUI" \
  -d '{
    "name": "Novo Nome"
  }'

# PUT com token
curl -X PUT http://localhost:8081/api/user/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_JWT_AQUI" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "senha123",
    "birthdate": "1990-01-15T00:00:00Z",
    "phone": "11987654321",
    "gender": "male"
  }'

# DELETE com token
curl -X DELETE http://localhost:8081/api/user/1 \
  -H "Authorization: Bearer SEU_TOKEN_JWT_AQUI"
```
