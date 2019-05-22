FROM golang:1.12-alpine as builder
RUN apk update \
  && apk add --no-cache ca-certificates \
  && apk add git

ENV PROJECT github.com/tetsuya/microservice-learning/src/userservice
WORKDIR /go/src/$PROJECT
COPY . .
RUN GO111MODULE=on go install

FROM alpine:3.9 as release
RUN apk update \
  && apk add --no-cache ca-certificates
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe
COPY --from=builder /go/bin/userservice /userservice
ENV APP_PORT=8080
EXPOSE 8080
ENTRYPOINT ["/userservice"]
