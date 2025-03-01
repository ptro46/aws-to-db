#!/bin/bash

export PGPASSWORD=PledgerPass2022
export SQL="select distinct(i.id),i.instance_id,i.private_ip_address,i.public_ip_address,t.key,t.value from instance i, instance_tag t where i.id=t.instance_id and t.key='Name' order by t.value;"

psql --host=10.1.1.18 --username=pledger --command="$SQL" pledger_monitoring
