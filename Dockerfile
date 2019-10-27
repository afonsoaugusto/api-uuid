FROM golang:latest as builder

LABEL maintainer="Afonso Rodrigues <afonsoaugustoventura@gmail.com>"

WORKDIR /app

RUN go get go.elastic.co/apm/module/apmgorilla && \
    go get github.com/gorilla/mux && \
    go get github.com/google/uuid

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"] 