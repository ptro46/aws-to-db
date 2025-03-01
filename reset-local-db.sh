#!/bin/bash
./aws-to-db aws-to-db-config.json purge
./aws-to-db aws-to-db-config.json amis
./aws-to-db aws-to-db-config.json vpc
./aws-to-db aws-to-db-config.json subnets
./aws-to-db aws-to-db-config.json volumes
./aws-to-db aws-to-db-config.json security_groups
./aws-to-db aws-to-db-config.json instances
./aws-to-db aws-to-db-config.json buckets
./aws-to-db aws-to-db-config.json stacks
./aws-to-db aws-to-db-config.json snapshots
./aws-to-db aws-to-db-config.json db_instances
./aws-to-db aws-to-db-config.json load_balancers
./aws-to-db aws-to-db-config.json load_balancers_target_groups
./aws-to-db aws-to-db-config.json auto_scaling_groups
./aws-to-db aws-to-db-config.json auto_scaling_activities
