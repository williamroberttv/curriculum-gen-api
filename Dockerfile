FROM golang:1.20.5

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

# RUN go build -o main ./src/main.go

# CMD ["air"]

# RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy

# CMD ["air", "-c", ".air.toml"]