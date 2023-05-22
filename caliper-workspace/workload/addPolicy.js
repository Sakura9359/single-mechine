'use strict';

const {WorkloadModuleBase} = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
    }

    async submitTransaction() {
        const policyID = `policy_${this.workerIndex}_${this.txIndex.toString()}`;
        const requestID = `request_${this.workerIndex}_${this.txIndex.toString()}`;
        const responseID = `response_${this.workerIndex}_${this.txIndex.toString()}`;
        this.txIndex++;
         // console.log(`Worker ${this.workerIndex}: Creating asset ${assetID}`);

        const policyStr = '{\"subjectA\":{\"userID\":\"user1\",\"role\":\"owner\",\"PK\":\"-----BEGIN PUBLIC KEY-----\\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEj2sdlxgXUoFBcNTXyA3mQNsmze4dEhUp\\neQ24KqVEeBieJ1wf+rra0/bp/wO6Wph0HXxnxcG/Lqfgk5V5v+2Z5+MBOZm94aBI\\nd3s4z8dMgNEMkWGZqr0DZKY3Y8vmP9Nv\\n-----END PUBLIC KEY-----\"},\"dataA\":{\"dataID\":\"豫E-MJ893\",\"owner\":\"uder1\",\"key\":\"豫E-MJ893\"},\"actionA\":{\"level\":15},\"environmentalA\":{\"createTime\":\"1670078438061014000\",\"endTime\":\"4071049445000000000\",\"address\":\"0.0.0.0\"}}'
        const signature = '3066023100c5651c5c2df0e32e0d1e5cff62847577466ac54bdae0d5818e0a0730470e531ddad154a78136cb4f3d2470bb0b1c73eb02310086dc906775b361f04207a9d31c6316468083ac0402f60643e8efe522bacfa89571558fb23859bd29c323e4010374a093'

        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'AddPolicy',
            invokerIdentity: 'User1',
            contractArguments: [policyID, requestID, responseID, policyStr,signature],
            readOnly: false
        };

        await this.sutAdapter.sendRequests(myArgs);
    }

    /*    async cleanupWorkloadModule() {
            for (let i=0; i<this.roundArguments.assets; i++) {
                const assetID = `${this.workerIndex}_${i}`;
                console.log(`Worker ${this.workerIndex}: Deleting asset ${assetID}`);
                const request = {
                    contractId: this.roundArguments.contractId,
                    contractFunction: 'deleteRecord',
                    invokerIdentity: 'User1',
                    contractArguments: [assetID],
                    readOnly: false
                };

                await this.sutAdapter.sendRequests(request);
            }
        }*/

}

function randomString(length) {
    const str = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    let result = '';
    for (let i = length; i > 0; --i)
        result += str[Math.floor(Math.random() * str.length)];
    return result;
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;