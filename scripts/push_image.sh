#!/bin/bash

docker images | grep -E $1 | awk '{print $3}' | uniq | xargs -I {} docker rmi --force {}
docker build . -f ./Dockerfile -t $1:latest
docker tag $1 registry.cn-hongkong.aliyuncs.com/zeroone/opsnft:$1
docker push registry.cn-hongkong.aliyuncs.com/zeroone/opsnft:$1
#ssh root@opsnft-test-1 "./pull_image.sh $1"