docker-compose down -v

#docker rm -f $(docker ps -aq)

# docker volume prune

./create-artifacts.sh

docker-compose up -d

sleep 2
./createChannel.sh

sleep 2
./deployChaincode.sh