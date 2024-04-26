FROM golang:1.21.6-alpine AS BuildStage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o residential cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY ./.env .
COPY ./assets ./assets
COPY --from=BuildStage /app/residential residential

EXPOSE 80 8000 8080

#USER nonroot:nonroot
ENTRYPOINT ["./residential"]