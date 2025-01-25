FROM golang:1.23-alpine

WORKDIR /app

RUN apk update && apk add --no-cache nano curl coreutils libc-dev gcc make

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download
# RUN go mod download && go mod verify

# RUN go get github.com/githubnemo/CompileDaemon
# RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

# COPY ./entrypoint.sh /entrypoint.sh

# RUN chmod +rx ./entrypoint.sh

# EXPOSE 3000

# ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
# RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

# ENTRYPOINT ["sh", "/entrypoint.sh"]

# ENTRYPOINT ["./entrypoint.sh"]
# CMD ["sh", "-c", "go run ./cmd/api/main.go"]

# RUN ls -l /app

RUN go build -o budgetapi ./cmd/api/main.go

CMD ["./budgetapi"]