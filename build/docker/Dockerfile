FROM golang:latest AS builder

WORKDIR $GOPATH/src/github.com/sss-eda/lemi011b

COPY go.mod .
COPY go.sum .
COPY cmd cmd/
COPY pkg pkg/
COPY web web/

RUN go mod tidy
RUN go install ./cmd/client/...
RUN go install ./cmd/server/...

RUN env CGO_ENABLED=0 go build -o /go/bin/client ./cmd/client/...
RUN env CGO_ENABLED=0 go build -o /go/bin/server ./cmd/server/...


FROM scratch AS client

ENV LEMI011B_CLIENT_SENSOR_ID=1
ENV LEMI011B_CLIENT_REST_URL="http://localhost:8081"
ENV LEMI011B_CLIENT_SERIAL_PORT="/dev/ttyUSB0"
ENV LEMI011B_CLIENT_SERIAL_BAUD=19200

COPY --from=builder /go/bin/client /go/bin/import

ENTRYPOINT ["/go/bin/import"]


FROM scratch AS server

ENV LEMI011B_SERVER_TIMESCALEDB_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"
ENV LEMI011B_SERVER_REST_PORT=8080

COPY --from=builder /go/bin/server /go/bin/import

ENTRYPOINT ["/go/bin/import"]