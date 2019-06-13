# Creating Custom AWS Config Rules

This code is intended as a companion to this article: https://medium.com/spaceapetech/custom-aws-config-rules-20d7995561a8.

Together with the article, it provides an example of a custom AWS Config Rule written in Golang and deployed using SAM.

The facile policy check it implements is:

_All Production Database EC2 Instances must have EBS Optimization enabled._

## Pre-Requisites

   * Go 1.x
   * The [aws-sam-cli](https://github.com/awslabs/aws-sam-cli)

## Build and Deploy

The code can be deployed with `S3_BUCKET=<an-s3-bucket-you-own> make deploy` although it is recommended to read the
associated article to understand what you are deploying.

