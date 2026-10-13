FROM golang:1.27.2@sha256:007a316059246276cb214df93950be77837545dc78007a73e0291a4963d5d3bf
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]
