import {Construct} from 'constructs';
import * as cdk from 'aws-cdk-lib';
import {CognitoConstruct} from "./cognito";
import {TokenCustomizerFunction} from "./token-customizer";
import {UserTable} from "./user-table";

interface LambdaStackProps extends cdk.StackProps {
}

export class MainStack extends cdk.Stack {
    constructor(scope: Construct, id: string, props: LambdaStackProps) {
        super(scope, id, props);

        const userTable = new UserTable(this, 'UserTable');
        const customizer = new TokenCustomizerFunction(this, 'TokenCustomizer', {
            table: userTable.table
        })
        const cognito = new CognitoConstruct(this, 'Cognito', {
            tokenCustomizer: customizer.func
        });
    }
}
