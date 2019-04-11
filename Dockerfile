# build stage
FROM golang AS builder
WORKDIR /workspace
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go test
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server

# image stage
FROM scratch
COPY --from=builder /workspace/server /go/bin/server
CMD ["/go/bin/server"]
