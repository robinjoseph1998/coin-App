FROM golang:1.22.2 AS builder

WORKDIR /app

COPY  go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o main

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql postgresql-client bash

ENV PGDATA=/var/lib/postgresql/data

EXPOSE 5000

EXPOSE 5432

COPY --from=builder /app/main /app/main

RUN initdb --auth=trust -D /var/lib/postgresql/data

CMD ["sh", "-c", "pg_ctl -D /var/lib/postgresql/data start && ./main"]