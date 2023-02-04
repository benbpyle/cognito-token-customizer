import {Construct} from "constructs";
import {AttributeType, BillingMode, ITable, Table} from "aws-cdk-lib/aws-dynamodb";
import {RemovalPolicy} from "aws-cdk-lib";

export class UserTable extends Construct {
    private readonly _table: ITable;

    get table(): ITable {
        return this._table;
    }

    constructor(scope: Construct, id: string) {
        super(scope, id);
        this._table = new Table(this, id, {
            billingMode: BillingMode.PAY_PER_REQUEST,
            removalPolicy: RemovalPolicy.DESTROY,
            partitionKey: {name: 'PK', type: AttributeType.STRING},
            sortKey: {name: 'SK', type: AttributeType.STRING},
            pointInTimeRecovery: true,
            tableName: 'SampleUsers',
        });


    }
}