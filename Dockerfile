FROM golang:1.27.3@sha256:6928d94d1142af58816fed5a62338d605a266ccd53a0deaab6b1096f19059261
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
