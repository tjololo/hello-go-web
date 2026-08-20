FROM golang:1.27.0@sha256:65b6f280bf050ec5af12716857e8ea8439d694dbba8f31ceeb7630670071f2bb
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
