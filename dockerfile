FROM golang:1.22.2

WORKDIR /app

COPY  go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o main

EXPOSE 5000

CMD [ "./main" ]