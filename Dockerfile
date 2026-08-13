FROM golang:1.26.5@sha256:5822931cf78fe98a97edcf73a0c54c29fa2386b99c8136468e274ae9fab8cfba
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
