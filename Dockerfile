FROM golang:1.26.4@sha256:d184d9be4c13614e28498d632eeaaac704d662f18ad357e1df74a44424236cea
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
