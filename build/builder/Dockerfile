# API exposing a simple CRUD
FROM golang:1.23 AS build-stage

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ../../go.mod ./
COPY ../../go.sum ./
RUN go mod download && go mod verify

COPY ../../. .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /ativia-api ./cmd/ativia-api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /ativia-api /ativia-api

EXPOSE 2113

USER nonroot:nonroot

ENTRYPOINT ["/ativia-api"]