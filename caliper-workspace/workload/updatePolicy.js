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

        const policyStr = '{\"subjectA\":{\"userID\":\"user1\",\"role\":\"owner\",\"PK\":\"-----BEGIN PUBLIC KEY-----\\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEj2sdlxgXUoFBcNTXyA3mQNsmze4dEhUp\\neQ24KqVEeBieJ1wf+rra0/bp/wO6Wph0HXxnxcG/Lqfgk5V5v+2Z5+MBOZm94aBI\\nd3s4z8dMgNEMkWGZqr0DZKY3Y8vmP9Nv\\n-----END PUBLIC KEY-----\"},\"dataA\":{\"dataID\":\"豫E-MJ893\",\"owner\":\"uder1\",\"key\":\"豫E-MJ893\"},\"actionA\":{\"level\":15},\"environmentalA\":{\"createTime\":\"1670159991300492000\",\"endTime\":\"4098092645000000000\",\"address\":\"0.0.0.0\"}}'
        const signature = '306402301d19359ed3772bdfb9a41fa5621d173d2e9b7ccd23e45cf963258727f868d9911d08acf98cb78899909809694698c362023079ac2fd256cd47c4aaf5f7cb5df7511e7d766fd1715bbdb010a0122e335b5b5e4e8d54cfc6ffa065215004b57e173d38'

        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'UpdatePolicy',
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