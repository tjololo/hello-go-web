FROM golang:1.25@sha256:c0a2bd0756d92462a0d449124b039100ce447ebf69dc6c80a6d877503b36935e
WORKDIR /go/src/github.com/tjololo/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/app/hello-go-web ./app
ENTRYPOINT ["/app"]