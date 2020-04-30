FROM golang:1.14.2 AS builder
ARG WRKHUB_PORT
ENV WRKHUB_PORT=$WRKHUB_PORT

COPY ./ /opt/wrkhub
WORKDIR /opt/wrkhub
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/server/main.go

FROM scratch
COPY --from=builder /opt/wrkhub/server /opt/wrkhub/server
COPY --from=builder /opt/wrkhub/queries /opt/wrkhub/queries
WORKDIR /opt/wrkhub
EXPOSE $WRKHUB_PORT
ENTRYPOINT ["./server"]
