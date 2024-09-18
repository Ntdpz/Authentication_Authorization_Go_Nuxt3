# ใช้ Go version 1.23.1
FROM golang:1.23.1-alpine

WORKDIR /app  

COPY ./backend/go.mod ./backend/go.sum ./  

RUN go mod download  

COPY ./backend /app  

RUN go build -o main .

CMD ["/app/main"]
