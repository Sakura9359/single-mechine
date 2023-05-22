'use strict';

const {WorkloadModuleBase} = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
    }

    async submitTransaction() {
        const owner = 'user1'
        const endtime = '2099-11-11 15:04:05'
        const responseID = `response_${this.workerIndex}_${this.txIndex.toString()}`;
        const requestID = `request_${this.workerIndex}_${this.txIndex.toString()}`;
        const policyID = `policy_${this.workerIndex}_${this.txIndex.toString()}`;
        this.txIndex++;
         // console.log(`Worker ${this.workerIndex}: Creating asset ${assetID}`);

        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'AddResponse',
            invokerIdentity: 'User1',
            contractArguments: [owner, endtime, requestID, responseID, policyID],
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