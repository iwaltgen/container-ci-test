
# build stage
FROM golang AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server

# image stage
FROM gcr.io/distroless/base
# FROM gcr.io/distroless/base
COPY --from=builder /workspace/server /go/bin/server
CMD ["/go/bin/server"]
