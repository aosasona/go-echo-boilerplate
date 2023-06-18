FROM golang:1.20.1 as base

WORKDIR /go/src/app
COPY go.* .
RUN go mod download

COPY . .
RUN make build TARGET_DIR=/go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=base /go/bin/app /app

CMD ["/app"]