FROM golang:1.25-alpine AS builder
WORKDIR /use/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-service ./cmd/server/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /usr/src/app/product-service .

EXPOSE 50051

CMD [ "./product-service" ]