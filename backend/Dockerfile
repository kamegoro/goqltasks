FROM golang:1.19.4

RUN mkdir /backend
WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "server.go"]
