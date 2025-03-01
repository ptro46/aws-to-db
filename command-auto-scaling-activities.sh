#!/bin/bash

# aws autoscaling describe-auto-scaling-instances > asg-instances.json
# 		aws autoscaling describe-load-balancers > asg-load-balancers.json
# 		aws autoscaling describe-load-balancer-target-groups > asg-load-balancer-target-groups.json
# aws autoscaling describe-scaling-activities > asg-scaling-activities.json
# aws autoscaling describe-auto-scaling-groups > asg.jsonscaling-groups.json


aws --profile oppens --region eu-west-3 autoscaling describe-scaling-activities
