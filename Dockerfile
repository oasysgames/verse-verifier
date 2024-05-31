# Build stage
FROM  golang:1.22.3-bookworm as builder

RUN apt update && apt install -y git ca-certificates

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

ADD . ./
ENV CGO_ENABLED=1
RUN go build -a -o oasvlfy -tags netgo -installsuffix netgo --ldflags='-s -w -extldflags "-static"' -buildvcs=false

# Final stage
FROM gcr.io/distroless/base-debian12
COPY --from=builder /build/oasvlfy /usr/local/bin/oasvlfy
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

ENTRYPOINT ["/usr/local/bin/oasvlfy"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""

LABEL commit="$COMMIT" version="$VERSION"
