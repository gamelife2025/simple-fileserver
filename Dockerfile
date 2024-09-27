FROM golang AS builder
WORKDIR /root
COPY go.mod go.sum ./
RUN go mod download
COPY . . 

ENV CGO_ENABLED=0

RUN go build ./cmd/simple-fileserver

FROM scratch
COPY --from=builder /root/simple-fileserver /usr/local/bin/simple-fileserver
ENTRYPOINT [ "simple-fileserver" ]