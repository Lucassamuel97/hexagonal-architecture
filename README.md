# Exemplo de desenvolvimento utilizando hexagonal architecture com Golang

Este repositório contém um exemplo de implementação da `Hexagonal Architecture` desenvolvido em `Golang`.

## Configuração do Ambiente

### 1. Clonando o Repositório

Clone este repositório para a sua máquina local:
```bash
git clone https://github.com/Lucassamuel97/hexagonal-architecture
cd hexagonal-architecture
```
### 2. Executando aplicação com docker-compose e entrando no container
```bash
    docker-compose up -d --build
    # Entrando no container
    docker exec -it appproduct bash
```

### 3. Comandos após entrar no container
```bash
    #Executa os testes da aplicação
    go test ./...

    # Comando para criar um produto via CLI
    go run main.go cli -a=create -n="Produto teste" -p=60.00
    # Buscar produto via CLI
    go run main.go cli -a=get -i="eef1d2d0-4530-4aef-a8b7-4c8842a13f8d"
    # desabilitar produto via CLI
    go run main.go cli -a=disable -i="eef1d2d0-4530-4aef-a8b7-4c8842a13f8d"
    # habilitar produto via CLI
    go run main.go cli -a=enable -i="eef1d2d0-4530-4aef-a8b7-4c8842a13f8d"

    #Sobe um WebServer na em localhost:9000
    go run main.go http
```
## Referencial teórico

## Comandos uteis
// Para atualizar o mock após modificar o arquivo product.go
mockgen -destination=application/mocks/application.go -source=application/product.go application
