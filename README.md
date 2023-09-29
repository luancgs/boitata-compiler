# boitata-compiler

## Build

Para usar o compilador, rode o comando `go build`. Isso irá gerar um arquivo binário no repositório chamado `boitata-compiler`.

## Usando exemplos

O diretório [input](./input/) possui alguns exemplos de código em _btt_ que podem ser utilizados com o compilador. Para isso basta usar o comando:

```sh
# Executando o código do arquivo ola_mundo.btt
$ ./boitata-compiler ./input/ola_mundo.btt
```

Este comando também funciona com scripts _btt_ além dos listados neste repositório.

## Usando Docker

Para fazer uso do docker ao executar scripts _btt_, você precisa primeiro montar a imagem presente neste repositório:

```sh
# Montando a imagem do compilador
$ docker build --tag boitata .
```

Na sequencia, basta iniciar um container com esta imagem, passando o script _.btt_ nos volumes da seguinte forma:

```sh
# Executando meu_script.btt dentro do container
$ docker run -i --rm -v ./meu_script.btt:/input.btt boitata
```

## To Do

- [ ] Melhorar checagem de tokens
- [ ] Adicionar substituição por regex para funções nativas
- [ ] Criar dockerfile para rodar o compilador direto de um container
