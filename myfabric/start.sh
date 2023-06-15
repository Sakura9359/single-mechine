docker-compose down -v       # 关闭节点

#docker rm -f $(docker ps -aq)

# docker volume prune

./create-artifacts.sh      # 生成证书文件

docker-compose up -d       # 开启节点

sleep 2
./createChannel.sh         # 创建通道

sleep 2
./deployChaincode.sh       # 部署链码