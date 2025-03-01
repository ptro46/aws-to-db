#!/bin/bash

arn=$1

aws --profile oppens --region eu-west-3 elbv2 describe-target-health --target-group-arn "${arn}"

