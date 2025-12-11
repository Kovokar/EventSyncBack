# Comandos CURL - Métodos PUT, PATCH e DELETE

Base URL: `http://localhost:8081/api/v1/user`

---

## 🔄 PATCH - Atualização Parcial

Atualiza apenas os campos enviados no body.

### Atualizar apenas o nome
```bash
curl -X PATCH http://localhost:8081/api/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Novo Nome"}'
```

### Atualizar email e telefone
```bash
curl -X PATCH http://localhost:8081/api/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "novoemail@example.com",
    "phone": "9876543210"
  }'
```

### Atualizar senha
```bash
curl -X PATCH http://localhost:8081/api/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{"password": "novaSenha123"}'
```

### Atualizar múltiplos campos
```bash
curl -X PATCH http://localhost:8081/api/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao.silva@example.com",
    "phone": "11987654321",
    "gender": "male",
    "visible_in_public_list": true
  }'
```

---

## 🔄 PUT - Atualização Completa

Atualiza todos os campos. **Todos os campos obrigatórios devem ser enviados.**

### Atualização completa
```bash
curl -X PUT http://localhost:8081/api/v1/user/1 \
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

### Atualização completa (formato compacto)
```bash
curl -X PUT http://localhost:8081/api/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Maria Santos","email":"maria@example.com","password":"senha123","birthdate":"1992-08-20T00:00:00Z","phone":"11912345678","gender":"female","photo":"","visible_in_public_list":false}'
```

---

## 🗑️ DELETE - Deletar Usuário

Realiza soft delete (marca como deletado, não remove do banco).

### Deletar usuário
```bash
curl -X DELETE http://localhost:8081/api/v1/user/1
```

### Deletar com verbose (mostra detalhes)
```bash
curl -X DELETE http://localhost:8081/api/v1/user/1 -v
```

---

## 📋 Respostas Esperadas

### PATCH/PUT - Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao.silva@example.com",
  "phone": "11987654321",
  "gender": "male",
  "photo": "",
  "visible_in_public_list": true,
  "birthdate": "1990-01-15T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-02T00:00:00Z"
}
```

### PATCH/PUT - Erro 404 (Usuário não encontrado)
```json
{
  "error": "usuário não encontrado"
}
```

### DELETE - Sucesso (204 No Content)
```
(Resposta vazia)
```

### DELETE - Erro 404 (Usuário não encontrado)
```json
{
  "error": "usuário não encontrado"
}
```

---

## ⚠️ Observações

- **PATCH**: Apenas campos enviados são atualizados
- **PUT**: Todos os campos obrigatórios devem ser enviados
- **DELETE**: Soft delete (não remove fisicamente do banco)
- **Formato de data**: `YYYY-MM-DDTHH:MM:SSZ` (ex: `1990-01-15T00:00:00Z`)
- **Gender**: `"male"`, `"female"` ou `"other"`
