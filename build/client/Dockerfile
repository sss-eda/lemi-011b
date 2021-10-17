FROM golang:alpine AS builder

ENV LEMI011B_CLIENT_REST_URL="http://localhost:8080"
ENV LEMI011B_CLIENT_SERIAL_NAME="COM5"
ENV LEMI011B_CLIENT_SERIAL_BAUD=19200

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod vendor

COPY cmd/client/* ./cmd/
RUN echo $(ls -1 /tmp/dir)
# ENTRYPOINT ["ls"]
# COPY pkg/* .
# RUN go build -o main cmd/client/main.go

# FROM scratch

# COPY --from=builder /build/main /
# # COPY ./database/data.json /database/data.json

# # Command to run the executable
# ENTRYPOINT ["/main"]