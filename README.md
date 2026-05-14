# CPF/CNPJ Generator

Utilitários de linha de comando para gerar CPF e CNPJ válidos.

## Instalação

```bash
make install
```

Ou instale individualmente:

```bash
go install ./cmd/cpf
go install ./cmd/cnpj
```

## Uso

### Gerar CPF

```bash
cpf
```

Gera um CPF válido sem formatação, exibe na tela e copia para a área de transferência.

Exemplo de saída:
```
12345678909
```

Para gerar com formatação:
```bash
cpf --mask
```

Exemplo de saída:
```
123.456.789-09
```

### Gerar CNPJ

```bash
cnpj
```

Gera um CNPJ válido sem formatação, exibe na tela e copia para a área de transferência.

Exemplo de saída:
```
12345678000195
```

Para gerar com formatação:
```bash
cnpj --mask
```

Exemplo de saída:
```
12.345.678/0001-95
```

## Desenvolvimento

### Build

```bash
make build
```

Os binários serão criados em `bin/cpf` e `bin/cnpj`.

### Build para Windows

```bash
make windows
```

Os binários para Windows serão criados em `bin/cpf.exe` e `bin/cnpj.exe`.

### Limpar

```bash
make clean
```

## Como funciona

Os geradores implementam os algoritmos de validação de CPF e CNPJ:

- **CPF**: Gera 9 dígitos aleatórios e calcula os 2 dígitos verificadores
- **CNPJ**: Gera 12 dígitos aleatórios e calcula os 2 dígitos verificadores

Os números gerados são válidos segundo as regras da Receita Federal.
