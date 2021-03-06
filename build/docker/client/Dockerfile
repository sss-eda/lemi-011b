# First stage: build the executable.
FROM golang:alpine AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app ./cmd/client/...


# # Development stage: the running container.
# FROM scratch AS development

# # Import the compiled executable from the first stage.
# COPY --from=builder /app /app

# ENV API_URL="api.sansa.dev"

# # Run the compiled binary.
# ENTRYPOINT ["/app"]


# Production stage: the running container.
FROM scratch AS production

#ENV SERIAL_NAME="/dev/ttyUSB0"
#ENV SERIAL_BAUD=19200
#ENV API_URL="sansa.dev/api"
#ENV INSTRUMENT_ID=1

# VOLUME /dev/ttyUSB0

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Import the root ca-certificates (required for Let's Encrypt)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the compiled binary.
ENTRYPOINT ["/app", "--serial_name=/dev/ttyUSB0", "--serial_baude=19200", "--api_url=https://sansa.dev/api", "--instrument_id=1"]
