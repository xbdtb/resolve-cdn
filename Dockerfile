FROM golang:1.14.4-alpine3.12 as builder
MAINTAINER xbdtb <xbdtb@163.com>
WORKDIR /app
ADD . /app
ENV GOFLAGS " -mod=vendor"
RUN go build -o app

FROM alpine:3.12
WORKDIR /app
ENV LISTEN_ADDRESS ":80"
EXPOSE 80
COPY --from=builder /app/app /app/
COPY --from=builder /app/public /app/public
COPY --from=builder /app/config/local.yml /app/config/
CMD /app/app
