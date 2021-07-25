#FROM golang:1.16.3 as build
#
#WORKDIR /home/opsnft
#
#COPY ./contract /home/opsnft

#build at China
#RUN go env -w GOPROXY=https://goproxy.cn,direct && \
#    make test && \
#    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o contract

#RUN go env -w GOPROXY=https://goproxy.cn,direct && \
#   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o contract

FROM alpine

COPY ./glibc-2.33-r0.apk /tmp/
RUN touch repo.list && apk add --repositories-file=repo.list --allow-untrusted --no-network --no-cache /tmp/glibc-2.33-r0.apk &&\
#   apk add --no-cache libc6-compat && \
   apk update && \
   apk add ca-certificates && \
   apk add -U tzdata && \
   mkdir -p /home/opsnft && \
   cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /home/opsnft

#COPY --from=build /home/opsnft/contract /home/opsnft/contract

COPY  ./contract /home/opsnft/contract
COPY ./conf/conf_online.yaml /home/opsnft/conf.yaml

ENTRYPOINT [ "./contract" ]