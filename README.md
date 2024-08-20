go run main.go http




// Para atualizar o mock após modificar o arquivo product.go
mockgen -destination=application/mocks/application.go -source=application/product.go application
