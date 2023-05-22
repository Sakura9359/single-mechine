'use strict';

const {WorkloadModuleBase} = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
    }

    async submitTransaction() {
        const requester = 'user2'
        const role = 'querier'
        const pk = '-----BEGIN PUBLIC KEY-----\\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEr6k2ouQMK94cQnRWq5YIYlTKgFT0mXI5\\nOvIev15MtL5Nnc4bPxFx/g77rnMmacd+1SmuIz9Ujaup8CtrG2HpeT/saVjHwqpR\\nct76ghbYDiRjPaA0FYQDDp40Iu8J30ZG\\n-----END PUBLIC KEY-----'
        const dataID = `data_${this.workerIndex}_${this.txIndex.toString()}`;
        const dataOwner = 'user1'
        const level = '15'
        const requestID = `request_${this.workerIndex}_${this.txIndex.toString()}`;
        this.txIndex++;
         // console.log(`Worker ${this.workerIndex}: Creating asset ${assetID}`);

        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'AddRequest',
            invokerIdentity: 'User1',
            contractArguments: [requester, role, pk, dataID, dataOwner, level, requestID],
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