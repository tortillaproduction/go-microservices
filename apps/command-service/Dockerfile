FROM golang:1.26-alpine
RUN apk update && apk add git curl alpine-sdk
WORKDIR /go/src
COPY . /go/src/

# TODO: sqlboiler CLI is only needed when regenerating models from the DB schema
# RUN go install github.com/aarondl/sqlboiler/v4@latest && \
#     go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \

RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

RUN go mod download

CMD ["go", "run", "./apps/command-service/main.go"]