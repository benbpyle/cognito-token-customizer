# Cognito Token Customer Sample

Simple Repos that supports the blog article on Binaryheap.com. 

### CDK Infra

Run `cdk deploy npx ts-node bin/app.ts` from the root directory
* Creates a Cognito/User Pool combo
* Setups up an AppClient for use with Password login like a web app
* Creates a DynamoDB Table which holds the additional user information
* Creates a Lambda function that pulls from that table and customizes the ID Token
* Attaches the Lambda to the PreTokenGeneration Trigger in the User Pool
  