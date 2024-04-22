# builder
FROM golang:1.18.10-bullseye as builder

RUN apt update && apt install -y git ca-certificates

ADD . /build
WORKDIR /build

ENV CGO_ENABLED=1
RUN go build -a -o oasvlfy -tags netgo -installsuffix netgo --ldflags='-s -w -extldflags "-static"' -buildvcs=false

# runner
FROM debian:11.9-slim
COPY --from=builder /build/oasvlfy /usr/local/bin/oasvlfy
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

ENTRYPOINT ["/usr/local/bin/oasvlfy"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""

LABEL commit="$COMMIT" version="$VERSION"
