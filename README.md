# boitata-compiler

## Build

Para usar o compilador, rode o comando `go build`. Isso irá gerar um arquivo binário no repositório chamao `boitata-compiler`.

## Usando exemplos

O diretório [input](./input/) possui alguns exemplos de código em _btt_ que podem ser utilizados com o compilador. Para isso basta usar o comando:

```sh
# Executando o código do arquivo ola_mundo.btt
$ ./boitata-compiler ./input/ola_mundo.btt
```

Este comando também funciona com scripts _btt_ além dos listados neste repositório.

## To Do

- [ ] Melhorar checagem de tokens
- [ ] Adicionar substituição por regex para funções nativas
- [ ] Criar dockerfile para rodar o compilador direto de um container
