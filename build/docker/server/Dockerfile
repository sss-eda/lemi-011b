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
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app ./cmd/server/...


# Development stage: the running container.
FROM scratch AS development

# Import the compiled executable from the first stage.
COPY --from=builder /app /
COPY --from=builder /src/web /web

ENV TIMESCALE_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"

# Expose port 8080 to our application
EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/app", "--port=8080"]


# Production stage: the running container.
FROM scratch AS production

# Import the compiled executable from the first stage.
COPY --from=builder /app /
COPY --from=builder /src/web /web
# Import the root ca-certificates (required for Let's Encrypt)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV TIMESCALE_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"

# Expose both 443 and 80 to our application
EXPOSE 443
EXPOSE 80

# Mount the certificate cache directory as a volume, so it remains even after
# we deploy a new version
# VOLUME ["/cert-cache"]

# Run the compiled binary.
ENTRYPOINT ["/app", "--tls", "--tls_dir=/certs", "--tls_host=api.sansa.dev"]