FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o forum cmd/main.go

LABEL maintainer="ranniz | abaid | mennaas"
LABEL version="1.0"
LABEL description="forum"

EXPOSE 8080

CMD ["./forum"]