FROM golang:1.24 AS go
WORKDIR /src/todo
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o todo-app

FROM ubuntu:latest
WORKDIR /app/todo
COPY --from=go /src/todo/.env /src/todo/todo-app /src/todo/go.mod /src/todo/go.sum ./
EXPOSE 7540
ENV POSTGRES_USER=postgres POSTGRES_PASSWORD=123 POSTGRES_PORT=5432 POSTGRES_DB=testdb
CMD ["./todo-app"]