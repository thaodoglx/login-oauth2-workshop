FROM golang:alpine AS builder

LABEL maintainers="Anggit M Ginanjar <anggit@isi.co.id>"

WORKDIR $GOPATH/src/
COPY . .

RUN ls -alh

RUN go mod download
RUN GOOS=linux go build -o /go/bin/login-provider main.go
COPY views /go/bin/views
COPY .env /go/bin/.env
RUN ls -alh /go/bin

FROM alpine
COPY --from=builder /go/bin/ /srv/
WORKDIR /srv/
RUN ls -alh
CMD ./login-provider
