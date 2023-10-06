# Use a imagem oficial do Golang como base
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie o código da primeira aplicação Go para o diretório de trabalho
COPY . .

# Compile a primeira aplicação Go
RUN go get -d -v ./...
RUN go install -v ./...

# Comando para iniciar a primeira aplicação Go
CMD ["server"]