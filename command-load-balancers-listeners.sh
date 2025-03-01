#!/bin/bash

arn=$1

aws --profile oppens --region eu-west-3 elbv2 describe-listeners --load-balancer-arn "${arn}"



