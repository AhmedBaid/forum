FROM golang:alpine

# Install needed packages
RUN apk add --no-cache gcc musl-dev sqlite-libs

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1

RUN go mod tidy
RUN go build -o main ./cmd/main.go

LABEL maintainer="ranniz | abaid | mennaas"
LABEL version="1.0"
LABEL description="Forum"

EXPOSE 8080

CMD ["./main"]
