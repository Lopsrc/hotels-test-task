FROM golang:latest AS gobuild

WORKDIR /build

COPY ./ ./

RUN go mod download && \
    CGO_ENABLED=0 go build main.go

CMD [ "./main" ]