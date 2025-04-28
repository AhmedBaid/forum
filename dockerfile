# to Download the latest stable version of golang-- alpine is a Lightweight (khfifa) version of linux with a small size  5 MB
FROM golang:alpine

WORKDIR /app

# Install the Bash shell in our image
RUN apk add  bash

COPY . .

RUN go build -o forum ./cmd

LABEL maintainer="ranniz | abaid | mennaas"
LABEL version="1.0"
LABEL description="forum"

EXPOSE 8080

CMD ["./forum"]