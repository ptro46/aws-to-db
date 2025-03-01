#!/bin/bash

aws --profile oppens --region eu-west-3 ec2 describe-vpcs --filters Name=tag:Name,Values=OppensPRODVPC,OppensSTAGINGVPC

