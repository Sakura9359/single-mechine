# 声明变量
#export PATH=${PWD}/bin:${PWD}:$PATH
export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_ORG1_CA=${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export PEER0_ORG2_CA=${PWD}/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export PEER0_ORG3_CA=${PWD}/crypto-config/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export FABRIC_CFG_PATH=${PWD}/config/

#export PRIVATE_DATA_CONFIG=${PWD}/private-data/collections_config.json

export CHANNEL_NAME=mychannel

setGlobalsForOrderer() {
    export CORE_PEER_LOCALMSPID="OrdererMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp
}

setGlobalsForPeer0Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

setGlobalsForPeer1Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:8051

}

setGlobalsForPeer0Org2(){
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051

}

setGlobalsForPeer1Org2(){
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:10051

}

setGlobalsForPeer0Org3(){
    export CORE_PEER_LOCALMSPID="Org3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
    export CORE_PEER_ADDRESS=localhost:11051

}

setGlobalsForPeer1Org3(){
    export CORE_PEER_LOCALMSPID="Org3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
    export CORE_PEER_ADDRESS=localhost:12051

}

# 打包链码
packageChaincode() {
    rm -rf ${CC_NAME}.tar.gz
    setGlobalsForPeer0Org1
    peer lifecycle chaincode package ${CC_NAME}.tar.gz \
        --path ${CC_SRC_PATH} \
        --label ${CC_NAME}_${VERSION}
    echo "===================== Chaincode is packaged on peer0.org1 ===================== "
}
# packageChaincode

# 安装链码
installChaincode() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org1 ===================== "

    setGlobalsForPeer1Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org1 ===================== "

    setGlobalsForPeer0Org2
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org2 ===================== "

    setGlobalsForPeer1Org2
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org2 ===================== "

    setGlobalsForPeer0Org3
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org3 ===================== "

    setGlobalsForPeer1Org3
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org3 ===================== "
}

# installChaincode

# 查询已安装的链码
queryInstalled() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode queryinstalled >&log.txt
    cat log.txt
    PACKAGE_ID=$(sed -n "/${CC_NAME}_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    echo PackageID is ${PACKAGE_ID}
    echo "===================== Query installed successful on peer0.org1 on channel ===================== "
}

# queryInstalled

# 验证链码
approveForMyOrg1() {
    setGlobalsForPeer0Org1
    # set -x
    peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls true \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --package-id ${PACKAGE_ID} --init-required \
        --sequence ${VERSION}
    # set +x

    echo "===================== chaincode approved from org 1 ===================== "

}

# approveForMyOrg1

#checkCommitReadyness() {
#    setGlobalsForPeer0Org1
#    peer lifecycle chaincode checkcommitreadiness \
#        --collections-config $PRIVATE_DATA_CONFIG \
#        --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
#        --sequence ${VERSION} --output json --init-required
#    echo "===================== checking commit readyness from org 1 ===================== "
#}

# checkCommitReadyness

approveForMyOrg2() {
    setGlobalsForPeer0Org2

    peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} \
        --version ${VERSION} --package-id ${PACKAGE_ID} --init-required \
        --sequence ${VERSION}

    echo "===================== chaincode approved from org 2 ===================== "
}

# approveForMyOrg2

checkCommitReadyness() {

    setGlobalsForPeer0Org1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from org 1 ===================== "
}

# checkCommitReadyness

approveForMyOrg3() {
    setGlobalsForPeer0Org3

    peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} \
        --version ${VERSION} --package-id ${PACKAGE_ID} --init-required \
        --sequence ${VERSION}

    echo "===================== chaincode approved from org 3 ===================== "
}

# approveForMyOrg3

checkCommitReadyness() {

    setGlobalsForPeer0Org1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from org 1 ===================== "
}

# checkCommitReadyness

# 提交链码
commitChaincodeDefination() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        --channelID $CHANNEL_NAME --name ${CC_NAME} --init-required \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:11051 --tlsRootCertFiles $PEER0_ORG3_CA \
        --version ${VERSION} --sequence ${VERSION}
}

# commitChaincodeDefination

queryCommitted() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}

}

# queryCommitted

# 链码初始化
chaincodeInvokeInit() {
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:11051 --tlsRootCertFiles $PEER0_ORG3_CA \
        --isInit -c '{"Args":["-----BEGIN PRIVATE KEY-----\nMIGkAgEBBDCjyI/wNTU0HEA4/m+t8a+2/XLnv3rkZ5c3VYOHRX+5Lv5155GKmC6J\neqm1s0EcVnegBwYFK4EEACKhZANiAASPax2XGBdSgUFw1NfIDeZA2ybN7h0SFSl5\nDbgqpUR4GJ4nXB/6utrT9un/A7pamHQdfGfFwb8up+CTlXm/7Znn4wE5mb3hoEh3\nezjPx0yA0QyRYZmqvQNkpjdjy+Y/028=\n-----END PRIVATE KEY-----",
        "-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEj2sdlxgXUoFBcNTXyA3mQNsmze4dEhUp\neQ24KqVEeBieJ1wf+rra0/bp/wO6Wph0HXxnxcG/Lqfgk5V5v+2Z5+MBOZm94aBI\nd3s4z8dMgNEMkWGZqr0DZKY3Y8vmP9Nv\n-----END PUBLIC KEY-----",
        "-----BEGIN PRIVATE KEY-----\nMIGkAgEBBDDKE0ro5fodcecsT3PkZfBFInxVlUTgDnZoXtRO3kxQlCKQn/gftZc/\nDWgMr2UBLz2gBwYFK4EEACKhZANiAASvqTai5Awr3hxCdFarlghiVMqAVPSZcjk6\n8h6/Xky0vk2dzhs/EXH+DvuucyZpx37VKa4jP1SNq6nwK2sbYel5P+xpWMfCqlFy\n3vqCFtgOJGM9oDQVhAMOnjQi7wnfRkY=\n-----END PRIVATE KEY-----",
        "-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEr6k2ouQMK94cQnRWq5YIYlTKgFT0mXI5\nOvIev15MtL5Nnc4bPxFx/g77rnMmacd+1SmuIz9Ujaup8CtrG2HpeT/saVjHwqpR\nct76ghbYDiRjPaA0FYQDDp40Iu8J30ZG\n-----END PUBLIC KEY-----"]}'

}

# chaincodeInvokeInit

#chaincodeInvoke() {
#    setGlobalsForPeer0Org1
#    ## Init ledger
#    peer chaincode invoke -o localhost:7050 \
#        --ordererTLSHostnameOverride orderer.example.com \
#        --tls $CORE_PEER_TLS_ENABLED \
#        --cafile $ORDERER_CA \
#        -C $CHANNEL_NAME -n ${CC_NAME} \
#        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
#        --peerAddresses localhost:10051 --tlsRootCertFiles $PEER0_ORG2_CA \
#        -c '{"function": "initLedger","Args":[]}'
#
#}

# chaincodeInvoke

#chaincodeQuery() {
#    setGlobalsForPeer0Org2
#    peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"function": "queryCar","Args":["CAR0"]}'
#}

# chaincodeQuery

lifecycle1(){
    CHANNEL_NAME="mychannel"
    VERSION="1"
    CC_SRC_PATH="./chaincode/go/ABAC"
    CC_NAME="ABAC"
    packageChaincode
    installChaincode
    queryInstalled
    approveForMyOrg1
    checkCommitReadyness
    approveForMyOrg2
    checkCommitReadyness
    approveForMyOrg3
    checkCommitReadyness
    commitChaincodeDefination
    queryCommitted
    chaincodeInvokeInit
}

#lifecycle2(){
#    CHANNEL_NAME="mychannel"
#    VERSION="1"
#    CC_SRC_PATH="./chaincode/go/point"
#    CC_NAME="point"
#    packageChaincode
#    installChaincode
#    queryInstalled
#    approveForMyOrg1
#    checkCommitReadyness
#    approveForMyOrg2
#    checkCommitReadyness
#    commitChaincodeDefination
#    queryCommitted
#
#}
#
#lifecycle3(){
#    CHANNEL_NAME="mychannel"
#    VERSION="1"
#    CC_SRC_PATH="./chaincode/go/case"
#    CC_NAME="case"
#    packageChaincode
#    installChaincode
#    queryInstalled
#    approveForMyOrg1
#    checkCommitReadyness
#    approveForMyOrg2
#    checkCommitReadyness
#    commitChaincodeDefination
#    queryCommitted
#
#}
#
#lifecycle4(){
#    CHANNEL_NAME="mychannel"
#    VERSION="1"
#    CC_SRC_PATH="./chaincode/go/log"
#    CC_NAME="log"
#    packageChaincode
#    installChaincode
#    queryInstalled
#    approveForMyOrg1
#    checkCommitReadyness
#    approveForMyOrg2
#    checkCommitReadyness
#    commitChaincodeDefination
#    queryCommitted
#
#}

lifecycle1

#sleep 8
#lifecycle2
#sleep 8
#lifecycle3
#sleep 8
#lifecycle4
#sleep 8
#chaincodeInvoke
#sleep 5
#chaincodeQuery
