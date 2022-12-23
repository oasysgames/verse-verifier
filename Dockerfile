# builder
FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /build
WORKDIR /build

ENV CGO_ENABLED=1
RUN go build -a -o oasvlfy -tags netgo -installsuffix netgo --ldflags='-s -w -extldflags "-static"' -buildvcs=false

# runner
FROM alpine:3.17

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/oasvlfy /usr/local/bin/oasvlfy

ENTRYPOINT ["/usr/local/bin/oasvlfy"]
