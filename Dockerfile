FROM golang:1.25@sha256:cc737435e2742bd6da3b7d575623968683609a3d2e0695f9d85bee84071c08e6
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]