FROM golang:1.19-alpine

WORKDIR /dockerclone
RUN export TMPDIR=~/tmp/
COPY ./ ./
RUN export GOPATH=./
RUN go build container.go

CMD ["./container", "child", "/bin/bash"]