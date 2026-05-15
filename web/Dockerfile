FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /out/web .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /out/web /web
EXPOSE 8080
ENTRYPOINT ["/web"]
