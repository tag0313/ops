#!/bin/bash

echo "📃 生成文档……"
go get github.com/swaggo/swag/cmd/swag
make swag
echo "🏞 编译并生成 Docker 镜像……"
make docker_build
echo "🗑 清除多余镜像……"
make docker_clean
echo "🚀 启动服务！"
make up
echo "⚠️ 检查容器是否正确启动？按回车键将上传镜像到云端，退出按 ctrl + c："

while [ true ] ; do

    read -t 3 -n 1

    if [ $? = 0 ] ; then
        break
    fi

done

echo "停止服务……"
make down
echo "🐌 上传中……"
make docker_push
make clean
echo "✅ 上传成功！ 服务器拉取新镜像中……"
ssh root@test.opsnft.net './pull_image_by_hand.sh'
echo "🤞 如无意外，你就可以到 http://test.opsnft.net:9000/ 部署新镜像啦……"