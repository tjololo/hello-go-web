FROM golang:1.25@sha256:698183780de28062f4ef46f82a79ec0ae69d2d22f7b160cf69f71ea8d98bf25d
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]