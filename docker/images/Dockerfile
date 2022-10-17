FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /root
COPY goHello.Dockerfile /root/
RUN go build .
