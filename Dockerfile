FROM golang:latest AS BUILD

WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fakadog .

FROM scratch
WORKDIR /app
COPY --from=BUILD /app/fakadog .
EXPOSE 8126
CMD ["./fakadog"]
