FROM golang

WORKDIR /dockerhueta

COPY . .

RUN go build -o main main.go

EXPOSE 8000

CMD ["./main"]