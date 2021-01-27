BUILDPATH=$(CURDIR)
PROYECT = tlv-to-map

build: 
	@echo "ESTAMOS CREANDO EL BINARIO DE ESTA INCREIBLE FUNCIONALIDAD .... !!!"
	@echo "--------------------------------------------------"
	@echo "--------------------------------------------------"
	@echo "--------------------------------------------------"
	@echo "--------------------------------------------------"
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${PROYECT} cmd/main.go
	@echo "BINARIO GENERADO CORRECTAMENTE EN:       build/bin/${PROYECT}"

test: 
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html


.PHONY: test build