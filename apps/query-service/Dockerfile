FROM golang:1.26-alpine
RUN apk update && apk add git curl alpine-sdk
WORKDIR /go/src
COPY . /go/src

RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

RUN go mod download

CMD ["go", "run", "./apps/query-service/main.go"]