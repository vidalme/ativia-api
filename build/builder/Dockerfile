# API exposing a simple CRUD
FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /ativia-api ./cmd/ativia-api

EXPOSE 1301

CMD ["/ativia-api"]
