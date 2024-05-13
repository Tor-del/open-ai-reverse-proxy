FROM golang:latest as builder

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

FROM alpine:latest  

WORKDIR /app 
COPY --from=builder /app/main ./main
RUN apk --no-cache add ca-certificates


EXPOSE 8080

ENV PORT=8080 \
    PROXY_URL="https://api.openai.com" \
    EXT_PROXY_URL="http://your-http-proxy-address:proxy-port"

ENTRYPOINT ["./main"]
