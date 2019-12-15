# BUILDER
FROM alpine:3.7 as builder

RUN apk upgrade \
  && apk add --no-cache go musl-dev git

ENV GOPATH=/go
RUN go get github.com/gorilla/mux
COPY . /go/src/github.com/alb/petstore
RUN go install -v github.com/alb/petstore

# RUNTIME
FROM alpine:3.7
COPY --from=builder /go/bin/petstore /
EXPOSE 8080
ENTRYPOINT ["/petstore", "--port", "8080", "--host", "0.0.0.0"]
