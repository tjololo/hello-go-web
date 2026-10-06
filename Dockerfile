FROM golang:1.27.1@sha256:23fe8075c2e428136326703a2c63203f0c57595d8400eedbe06a29cab53055e8
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
