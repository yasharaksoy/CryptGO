FROM gcr.io/cloud-builders/go:latest as builder

RUN sed -i -e 's/v[[:digit:]]\.[[:digit:]]/edge/g' /etc/apk/repositories && \
    apk add --update --update-cache pkgconfig librdkafka-dev

WORKDIR /go/src/CryptGO/
COPY . .

ENV GOPATH=/go
RUN GO111MODULE=on go install -v ./gateway
RUN GO111MODULE=on go install -v ./match
RUN GO111MODULE=on go install -v ./memorystore
RUN GO111MODULE=on go install -v ./sql
RUN GO111MODULE=on go install -v ./wallet
RUN GO111MODULE=on go install -v ./wallet/ethereum
RUN GO111MODULE=on go install -v ./ws

FROM alpine:3.8

WORKDIR /CryptGO/
COPY --from=builder /go/bin/wallet .

ENTRYPOINT ["/CryptGO/wallet"]
