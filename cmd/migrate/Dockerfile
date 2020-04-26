FROM golang:1.14.2 AS builder

COPY ./ /opt/watdo
WORKDIR /opt/watdo
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o migrate cmd/migrate/main.go

FROM scratch
COPY --from=builder /opt/watdo/migrate /opt/migrate/migrate
COPY --from=builder /opt/watdo/migration /opt/migrate/migration
WORKDIR /opt/migrate
ENTRYPOINT ["./migrate"]
CMD ["up"]
