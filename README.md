# cpf-cnpj

[![CI](https://github.com/lcmuniz/cpf-cnpj/actions/workflows/ci.yml/badge.svg)](https://github.com/lcmuniz/cpf-cnpj/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Two small command-line tools, written in Go, that generate random **valid** Brazilian
taxpayer IDs for testing: **CPF** (individuals) and **CNPJ** (companies). The number is
printed and copied to the clipboard, ready to paste into a form.

> 🇧🇷 [Resumo em português](#resumo-em-português) no final.

```console
$ cpf
52998224725

$ cnpj --mask
11.222.333/0001-81

$ cnpj --alpha --mask
12.ABC.345/01DE-35
```

The numbers are random. They pass the official check-digit validation but are not
meant to belong to any real person or company. Use them only for testing.

## Install

With Go 1.24 or later:

```bash
go install github.com/lcmuniz/cpf-cnpj/cmd/cpf@latest
go install github.com/lcmuniz/cpf-cnpj/cmd/cnpj@latest
```

Or from a clone:

```bash
make install
```

## Usage

| Command               | Output example       |
| --------------------- | -------------------- |
| `cpf`                 | `52998224725`        |
| `cpf --mask`          | `529.982.247-25`     |
| `cnpj`                | `11222333000181`     |
| `cnpj --mask`         | `11.222.333/0001-81` |
| `cnpj --alpha`        | `12ABC34501DE35`     |
| `cnpj --alpha --mask` | `12.ABC.345/01DE-35` |

Flags:

- `--mask`: format the number with dots, slash and dash.
- `--no-copy`: only print the number, do not touch the clipboard.
- `--alpha` (`cnpj` only): generate an [alphanumeric CNPJ](#alphanumeric-cnpj).

If the clipboard is not available (for example on a server without a display), the
number is still printed and a warning goes to stderr.

On Linux (X11/Wayland) the system clipboard only keeps the data while the program that
wrote it is running, unless a clipboard manager is installed. On macOS and Windows the
number stays on the clipboard after the tool exits.

## How it works

Both documents use a mod-11 check-digit algorithm:

- **CPF**: 9 random digits, then 2 check digits computed with weights 10..2 and 11..2.
- **CNPJ**: 12 random digits, then 2 check digits computed with the weights
  `5 4 3 2 9 8 7 6 5 4 3 2` and `6 5 4 3 2 9 8 7 6 5 4 3 2`.

### Alphanumeric CNPJ

From July 2026 the Receita Federal issues new CNPJs whose first 12 characters may
contain uppercase letters (`AA.AAA.AAA/AAAA-DV`). Existing numeric CNPJs stay valid, and
the two check digits are still numbers. The check digits use the same weights as
above, with each character worth its ASCII code minus 48: `0`–`9` are 0–9 and `A`–`Z` are
17–42. For example, `12.ABC.345/01DE-35` (the example published by the Receita) is valid.

`cnpj --alpha` always includes at least one letter. `IsValidCNPJ` accepts both formats.
The Receita also blocks some letter combinations (to avoid offensive words) but does not
publish that list, so a generated number may, in rare cases, be one that will never be
issued. It still passes check-digit validation.

Source: [Perguntas e respostas sobre o CNPJ alfanumérico](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/perguntas-e-respostas/cnpj/cnpj-alfanumerico.pdf)
(Receita Federal; IN RFB 2.229/2024).

### Repeated digits

Sequences made of a single repeated digit (such as `111.111.111-11`) satisfy the math
but are rejected by real validators, so the generator never produces them.

The `internal/generator` package also has `IsValidCPF` and `IsValidCNPJ`, which the tests
use to check every generated number.

### Project layout

```
cmd/cpf, cmd/cnpj     entry points (one binary each)
internal/cli          shared flag parsing, output and clipboard handling
internal/generator    generation, formatting and validation
```

## Development

```bash
make test      # go test -race ./...
make vet       # gofmt check + go vet
make build     # binaries in bin/
make windows   # bin/cpf.exe and bin/cnpj.exe
make clean
```

CI runs formatting, `go vet`, tests and the build on Linux, macOS and Windows for every
push and pull request.

## License

[MIT](LICENSE)

---

## Resumo em português

Ferramentas de linha de comando em Go que geram **CPF** e **CNPJ** válidos para testes.
O número aparece na tela e é copiado para a área de transferência.

```bash
go install github.com/lcmuniz/cpf-cnpj/cmd/cpf@latest
go install github.com/lcmuniz/cpf-cnpj/cmd/cnpj@latest

cpf              # 52998224725
cpf --mask       # 529.982.247-25
cnpj --mask      # 11.222.333/0001-81
cnpj --alpha     # CNPJ alfanumérico (formato novo, a partir de julho/2026)
cnpj --no-copy   # só imprime, não copia
```

Os números são aleatórios e passam na validação dos dígitos verificadores, mas não
devem ser usados fora de ambientes de teste.
