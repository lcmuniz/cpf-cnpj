.PHONY: build install clean test vet windows

build:
	@echo "Building cpf..."
	@go build -o bin/cpf ./cmd/cpf
	@echo "Building cnpj..."
	@go build -o bin/cnpj ./cmd/cnpj
	@echo "Build complete!"

install:
	@echo "Installing cpf..."
	@go install ./cmd/cpf
	@echo "Installing cnpj..."
	@go install ./cmd/cnpj
	@echo "Install complete!"

windows:
	@echo "Building Windows binaries..."
	@GOOS=windows GOARCH=amd64 go build -o bin/cpf.exe ./cmd/cpf
	@GOOS=windows GOARCH=amd64 go build -o bin/cnpj.exe ./cmd/cnpj
	@echo "Windows build complete! (cpf.exe, cnpj.exe)"

clean:
	@rm -rf bin/
	@echo "Clean complete!"

test:
	@go test -race ./...

vet:
	@gofmt -l . | (! grep .) || (echo "gofmt needed on the files above" && exit 1)
	@go vet ./...
