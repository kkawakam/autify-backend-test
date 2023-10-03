# syntax=docker/dockerfile:1
FROM golang:1.21 as builder
WORKDIR /src

# Download dependencies and verify
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code and Makefile into the image
COPY main.go ./
COPY Makefile ./
COPY internal ./internal

# Build the executable
RUN make build

# Put executable lightweight scratch image
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /src/fetch /fetch
ENTRYPOINT [ "/fetch" ]