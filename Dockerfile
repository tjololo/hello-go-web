FROM golang:1.26.4@sha256:478231bfd9677835606c249208483a3c43b31e941c1040c48747b111c7ab871c
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
