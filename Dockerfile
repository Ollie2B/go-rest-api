FROM golang:1.20

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . /src
RUN swag init

EXPOSE 8080

CMD go run main.go