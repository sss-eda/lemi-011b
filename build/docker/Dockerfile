# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.17
ARG GO_VERSION=1.17

# First stage: build the executable.
FROM golang:${GO_VERSION} AS builder

# Git is required for fetching the dependencies.
# RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app ./cmd/server/...

# Final stage: the running container.
FROM scratch AS final

# Import the compiled executable from the first stage.
COPY --from=builder /app /app
# Import the root ca-certificates (required for Let's Encrypt)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV TIMESCALEDB_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"
ENV PRODUCTION=False
ENV REDIRECT_TO_HTTPS=False

# Expose both 443 and 80 to our application
EXPOSE 443
EXPOSE 80

# Mount the certificate cache directory as a volume, so it remains even after
# we deploy a new version
VOLUME ["/cert-cache"]

# Run the compiled binary.
ENTRYPOINT ["/app"]

# FROM golang:latest AS builder

# WORKDIR $GOPATH/src/github.com/sss-eda/lemi011b

# COPY go.mod .
# COPY go.sum .
# COPY cmd cmd/
# COPY pkg pkg/
# COPY web web/

# RUN go mod tidy
# RUN go install ./cmd/client/...
# RUN go install ./cmd/server/...

# RUN env CGO_ENABLED=0 go build -o /go/bin/client ./cmd/client/...
# RUN env CGO_ENABLED=0 go build -o /go/bin/server ./cmd/server/...


# FROM scratch AS client

# ENV LEMI011B_CLIENT_SENSOR_ID=1
# ENV LEMI011B_CLIENT_REST_URL="http://localhost:8081"
# ENV LEMI011B_CLIENT_SERIAL_PORT="/dev/ttyUSB0"
# ENV LEMI011B_CLIENT_SERIAL_BAUD=19200

# COPY --from=builder etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /go/bin/client /go/bin/import

# ENTRYPOINT ["/go/bin/import"]


# FROM scratch AS server

# ENV LEMI011B_SERVER_TIMESCALEDB_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"

# EXPOSE 443
# VOLUME "certs"

# COPY --from=builder etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY /web .
# COPY --from=builder /go/bin/server /go/bin/import

# ENTRYPOINT ["/go/bin/import"]