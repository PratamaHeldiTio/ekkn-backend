FROM golang:1.20-alpine3.17

WORKDIR /app

COPY . .

RUN go build -o ekknutm-backend

EXPOSE 8080

CMD ./ekknutm-backend