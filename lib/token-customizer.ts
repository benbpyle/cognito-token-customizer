import {Construct} from "constructs";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";
import * as path from "path";
import {Duration, Tags} from "aws-cdk-lib";
import {ITable, Table} from "aws-cdk-lib/aws-dynamodb";

interface TokenCustomizerProps {
    table: ITable
}

export class TokenCustomizerFunction extends Construct {
    private readonly _func: GoFunction;

    get func(): GoFunction {
        return this._func;
    }

    constructor(scope: Construct, id: string, props: TokenCustomizerProps) {
        super(scope, id);
        this._func = new GoFunction(this, `TokenCustomizerFunction`, {
            entry: path.join(__dirname, `../src/token-customizer`),
            functionName: `token-customizer`,
            timeout: Duration.seconds(10),
            environment: {
                "LOG_LEVEL": "debug",
                "TABLE_NAME": props.table.tableName,
            },
        });

        // add permissions and event sources
        props.table.grantReadWriteData(this._func)
    }

}
