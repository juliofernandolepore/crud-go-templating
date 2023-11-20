FROM golang:1.21-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o crudbin

EXPOSE 3000

CMD ["./crudbin"]