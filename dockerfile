FROM golang:1.21-alpine

WORKDIR /app/

RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/bin/boitata-compiler

CMD ["boitata-compiler", "/input.btt"]
