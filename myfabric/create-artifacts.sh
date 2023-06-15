#export PATH=${PWD}/bin:${PWD}:$PATH

chmod -R 0777 ./crypto-config
# Delete existing artifacts
rm -rf ./crypto-config
rm -rf ./channel-artifacts/*

#Generate Crypto artifactes for organizations
cryptogen generate --config=./crypto-config.yaml --output=./crypto-config/
./ccp-generate.sh         # 生成ccp文件，与caliper链接

# System channel
SYS_CHANNEL="system-channel"

# channel name defaults to "mychannel"
CHANNEL_NAME="mychannel"

echo $CHANNEL_NAME

# Generate System Genesis block
configtxgen -profile OrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID $SYS_CHANNEL

# Generate channel configuration block
configtxgen -profile BasicChannel -outputCreateChannelTx ./channel-artifacts/${CHANNEL_NAME}.tx -channelID $CHANNEL_NAME

echo "#######    Generating anchor peer update for Org1MSP  ##########"
configtxgen -profile BasicChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

echo "#######    Generating anchor peer update for Org2MSP  ##########"
configtxgen -profile BasicChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

echo "#######    Generating anchor peer update for Org3MSP  ##########"
configtxgen -profile BasicChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org3MSP