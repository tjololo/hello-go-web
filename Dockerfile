FROM golang:1.27.1@sha256:03fd17ca31f26e3c8d5b83ae0606044133c59e0788784dd6f371da59cdb3b46f
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
