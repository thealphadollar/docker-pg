FROM golang

WORKDIR /dockerclone

COPY ./ ./

RUN export GOPATH=./
RUN go build container.go

ENTRYPOINT ["./container", "run", "/bin/sh"]