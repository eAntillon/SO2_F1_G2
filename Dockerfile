FROM golang:1.17-alpine
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 5000
CMD ["go", "run", "main.go"]