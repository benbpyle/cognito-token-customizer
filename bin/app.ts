import {MainStack} from "../lib/main-stack";
import * as cdk from 'aws-cdk-lib';

const app = new cdk.App();
new MainStack(app,
    `CognitoTokenCustomer`,
    {});
