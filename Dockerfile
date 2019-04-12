# build base stage
FROM golang AS build_base
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download

# build binary stage
FROM build_base AS builder
COPY . .
RUN go test -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server

# image stage
FROM scratch
COPY --from=builder /workspace/server /go/bin/server
CMD ["/go/bin/server"]
