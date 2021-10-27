FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/sss-eda/lemi011b

COPY go.mod .
COPY go.sum .
COPY cmd cmd/
COPY pkg pkg/
COPY web web/

ENTRYPOINT ["sh"]

RUN go mod tidy
RUN go install ./cmd/client/...
ARG opts
RUN env ${opts} go build -o /go/bin/import ./cmd/client/...


FROM scratch AS lemi011b-client

ENV LEMI011B_CLIENT_SENSOR_ID=1
ENV LEMI011B_CLIENT_REST_URL="http://localhost:8081"
ENV LEMI011B_CLIENT_SERIAL_PORT="/dev/ttyUSB0"
ENV LEMI011B_CLIENT_SERIAL_BAUD=19200

COPY --from=builder /go/bin/import /go/bin/import

ENTRYPOINT ["/go/bin/import"]