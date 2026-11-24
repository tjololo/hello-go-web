FROM golang:1.27.4@sha256:87f49b77a7b6a99ad7dc656be010ae453322459dd2ed902ac177649a302bc101
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
