FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=8080

EXPOSE 8080

RUN CGO_ENABLED=0 go build -o /lets_go_chat

CMD ["/lets_go_chat"]