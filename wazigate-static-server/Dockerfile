FROM golang:alpine AS development

COPY . /go/src/
WORKDIR /go/src/
ENV GOPATH=/go/

RUN apk add --no-cache \
    git \
#    curl \
#    gcc \
#    musl-dev \
#    linux-headers \
    && mkdir /build/ \
    && go build -o /build/start -i . 

#----------------------------#

FROM alpine:latest AS production

WORKDIR /root/
COPY --from=development /build .

RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    curl \
    zip 

ENTRYPOINT ["./start"]
