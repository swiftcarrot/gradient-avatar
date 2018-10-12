FROM golang:alpine as builder
COPY . $GOPATH/src/avatar
WORKDIR $GOPATH/src/avatar
ENV GO111MODULE=on
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/avatar

FROM scratch
COPY --from=builder /go/bin/avatar /go/bin/avatar
ENTRYPOINT ["/go/bin/avatar"]
