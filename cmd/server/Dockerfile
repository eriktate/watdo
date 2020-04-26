FROM golang:1.14.2 AS builder
ARG WATDO_PORT
ENV WATDO_PORT=$WATDO_PORT

COPY ./ /opt/watdo
WORKDIR /opt/watdo
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/server/main.go

FROM scratch
COPY --from=builder /opt/watdo/server /opt/watdo/server
COPY --from=builder /opt/watdo/queries /opt/watdo/queries
WORKDIR /opt/watdo
EXPOSE $WATDO_PORT
ENTRYPOINT ["./server"]
