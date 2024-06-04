# Etapa 1: Construção
FROM golang:1.22-alpine AS builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos de configuração do Go
COPY go.mod go.sum ./

# Baixa as dependências necessárias
RUN go mod download

# Copia o código-fonte da aplicação
COPY . .

# Copia os arquivos HTML para dentro do contêiner
COPY views/ ./views

# Compila a aplicação
RUN go build -o docker-volumes

# Etapa 2: Execução
FROM alpine:latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copia o binário da etapa de construção para a imagem de execução
COPY --from=builder /app/docker-volumes .

# Copia os arquivos HTML para dentro do contêiner
COPY --from=builder /app/views/ ./views

# Expõe a porta que a aplicação utiliza (ajuste conforme necessário)
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./docker-volumes"]