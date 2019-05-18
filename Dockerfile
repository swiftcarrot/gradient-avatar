FROM golang:alpine as builder
COPY . $GOPATH/src/github.com/swiftcarrot/avatar
WORKDIR $GOPATH/src/github.com/swiftcarrot/avatar
ENV GO111MODULE=on
RUN apk update && apk add git
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/avatar cmd/avatar/main.go

FROM scratch
COPY --from=builder /go/bin/avatar /go/bin/avatar
ENTRYPOINT ["/go/bin/avatar"]
