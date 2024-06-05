FROM golang:latest

WORKDIR /app

COPY server-sync/go.mod ./
RUN go mod download

COPY server-sync/*.go ./

RUN go build -o /docker-gs-ping

EXPOSE 50000

# Run
CMD ["/docker-gs-ping"]