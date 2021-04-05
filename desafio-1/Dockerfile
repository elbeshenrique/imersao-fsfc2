FROM golang:1.16 as build

WORKDIR /go/src

COPY . .

ENV PATH="/go/bin:${PATH}"

EXPOSE 8080

CMD ["go", "run", "app/main.go"]