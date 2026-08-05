FROM golang:1.26.5@sha256:b004b9c35c68c8bcf5420e5164423073a1de5f0c7d4b9121784780ceb7f9961f
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
