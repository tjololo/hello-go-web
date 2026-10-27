FROM golang:1.27.3@sha256:921f61120011c774a7b37fd80bd3c3d259636090d3cd002b6a344322a0328082
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
