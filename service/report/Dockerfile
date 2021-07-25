FROM golang:1.16.3 as build

WORKDIR /home/opsnft

COPY . /home/opsnft

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o report

FROM alpine

RUN apk update && \
    apk add ca-certificates && \
    apk add libc6-compat && \
    mkdir -p /home/opsnft && \
    apk add -U tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /home/opsnft

COPY --from=build /home/opsnft/report /home/opsnft/report

COPY ./conf/conf_online.yaml /home/opsnft/conf.yaml

ENTRYPOINT [ "./report" ]