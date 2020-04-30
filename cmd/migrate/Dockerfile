FROM golang:1.14.2 AS builder

COPY ./ /opt/wrkhub
WORKDIR /opt/wrkhub
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o migrate cmd/migrate/main.go

FROM scratch
COPY --from=builder /opt/wrkhub/migrate /opt/migrate/migrate
COPY --from=builder /opt/wrkhub/migration /opt/migrate/migration
WORKDIR /opt/migrate
ENTRYPOINT ["./migrate"]
CMD ["up"]
