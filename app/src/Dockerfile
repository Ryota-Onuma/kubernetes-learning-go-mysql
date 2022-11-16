FROM golang:1.19

WORKDIR /app
COPY ./main.go /app/main.go

RUN go mod init main \
  && go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080


CMD ["go", "run", "/app/main.go"]
