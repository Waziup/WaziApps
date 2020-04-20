FROM golang:alpine AS development

COPY . /go/src/
WORKDIR /go/src/
ENV GOPATH=/go/

RUN apk add --no-cache \
    git \
    curl \
    gcc \
    musl-dev \
    linux-headers \
    zip \
    && mkdir /build/ \
    && zip /build/index.zip docker-compose.yml package.json \
    && cp -r ui /build/ui \
    && go build -o /build/start -i . 

#----------------------------#

FROM alpine:latest AS production

WORKDIR /root/
COPY --from=development /build .

#Just to keep the source code inside the container
COPY --from=development /go/src /go/src

RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    curl \
    && mv index.zip /

ENTRYPOINT ["./start"]
