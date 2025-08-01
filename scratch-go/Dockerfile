FROM golang AS base

RUN apt-get update && \
    apt-get install -y --no-install-recommends musl-dev musl-tools

WORKDIR /usr/src/app

COPY src/hello-web.go ./

ARG CGO_ENABLED=1
ARG CC=musl-gcc
RUN go build --ldflags '-linkmode=external -extldflags=-static' -v -o /usr/local/bin/app ./hello-web.go
RUN strip /usr/local/bin/app

FROM scratch

COPY --from=base /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
