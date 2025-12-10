# ==========================
# 1. Etapa de build (compilação)
# ==========================
FROM golang:1.25 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependência primeiro (para aproveitar cache de build)
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila o binário da aplicação
RUN go build -o server ./cmd/server

# ==========================
# 2. Etapa final (imagem leve)
# ==========================
FROM debian:bookworm-slim

# Define o diretório de trabalho na imagem final
WORKDIR /app

# Copia apenas o binário compilado da etapa anterior
COPY --from=builder /app/server .

# Copia o arquivo de configuração (ex: app.yaml)
COPY configs/app.yaml ./configs/app.yaml

# Expõe a porta que o seu servidor Go usa (ajuste se for diferente)
EXPOSE 8080

# Comando padrão que roda quando o container inicia
CMD ["./server"]
