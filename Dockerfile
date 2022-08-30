FROM golang1.18 AS builder
COPY . /bin
RUN make /app
CMD go run main.go
