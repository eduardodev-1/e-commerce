# Use uma imagem base com Go instalado, baseada no Alpine Linux
FROM golang:1.22-alpine

# Instale as dependências necessárias
RUN apk add --no-cache git

# Crie e defina o diretório de trabalho no contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixe as dependências do Go
RUN go mod download

# Copie apenas as pastas necessárias para o contêiner
COPY cmd/ cmd/
COPY internal/ internal/
# COPY env/ env/

# Construa o executável
RUN go build -o main ./cmd/api

# Defina a porta em que a aplicação será executada
EXPOSE 2121

# Comando para rodar a aplicação
CMD ["./main"]
