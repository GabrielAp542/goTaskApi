FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod download


RUN go get -u github.com/swaggo/swag/cmd/swag 
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/files
RUN go get -u github.com/swaggo/gin-swagger
RUN swag init

RUN go build -o app .

EXPOSE 80

CMD ["./app"]
