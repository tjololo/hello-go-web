FROM golang:1.26.5@sha256:d52df9c279840adf958d017ebb275651ed8338b953d39817bc3633a2e6b1bbcc
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
