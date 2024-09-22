# Imagem base
FROM golang:bullseye as build

# Diretório de trabalho dentro do contêiner
WORKDIR /usr/src/app

# Copie o diretório atual para o diretório de trabalho
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Execute o comando de compilação
RUN git config --global --add safe.directory /usr/src/app
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/app
RUN > /usr/local/bin/time

## Production Stage
FROM gcr.io/distroless/static-debian11 AS runtime

WORKDIR /
# USER nonroot:nonroot
# COPY --from=build --chown=nonroot:nonroot /usr/local/bin/app /app

COPY --from=build /usr/local/bin/app /app


# Comando de execução padrão
EXPOSE 2000
ENTRYPOINT ["/app"]