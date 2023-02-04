import {Duration, RemovalPolicy} from "aws-cdk-lib";
import {Construct} from "constructs";
import {AttributeType, BillingMode, ITable, Table} from "aws-cdk-lib/aws-dynamodb";
import {AccountRecovery, IUserPool, UserPool} from "aws-cdk-lib/aws-cognito";
import {IFunction} from "aws-cdk-lib/aws-lambda";

interface CognitoProps {
    tokenCustomizer: IFunction
}

export class CognitoConstruct extends Construct {
    constructor(scope: Construct, id: string, props: CognitoProps) {
        super(scope, id);

const userPool = new UserPool(this, 'SampleUserPool', {
    lambdaTriggers: {
      preTokenGeneration: props.tokenCustomizer
    },
    userPoolName: 'SamplePool',
    signInAliases: {
        email: true,
        username: true,
        preferredUsername: true
    },
    autoVerify: {
        email: false,
    },
    standardAttributes: {
        email: {
            required: true,
            mutable: true,
        }
    },
    passwordPolicy: {
        minLength: 12,
        requireLowercase: true,
        requireDigits: true,
        requireUppercase: true,
        requireSymbols: true,
    },
    accountRecovery: AccountRecovery.EMAIL_ONLY,
    removalPolicy: RemovalPolicy.DESTROY,
});

        userPool.addClient("web-client",
            {
                userPoolClientName: "web-client",
                authFlows: {
                    adminUserPassword:true,
                    custom:true,
                    userPassword:true,
                    userSrp:false
                },
                idTokenValidity: Duration.minutes(60),
                refreshTokenValidity: Duration.days(30),
                accessTokenValidity: Duration.minutes(60),
            })
    }
}