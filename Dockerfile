FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
 
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /build/bin/app /app/app

COPY ./config/docker.yaml /app/docker.yaml  

COPY ./api /app/
COPY ./public ./app

EXPOSE 8089
CMD ["/app/app"]
