-- #############################################################################
--
--    
-- #############################################################################
CREATE VIEW public.view_elb_input_cidr AS
select elb.id as load_balancer_id, elb.load_balancer_name as load_balancer_name, sgipir.cidr_ip as cidr_ip
from load_balancer elb, load_balancer_security_group elbsg, security_group_ip_permissions sgip, security_group_ip_permissions_ip_ranges sgipir
where elb.id=elbsg.load_balancer_id and elbsg.ref_security_group_id=sgip.security_group_id and sgip.id=sgipir.security_group_ip_permissions_id ;

CREATE VIEW public.view_elb_by_cidr AS 
select s.id as subnet_id, s.cidr_block as subnet_cidr_block, elb.id as load_balancer_id, elb.load_balancer_name as load_balancer_name
from subnet s, load_balancer elb, load_balancer_availability_zone lbaz 
where s.id=lbaz.ref_subnet_id and lbaz.load_balancer_id=elb.id ;

CREATE VIEW public.view_instance_by_cidr AS
select i.id as instance_id, it.key as instance_key, it.value as instance_name, s.cidr_block as cidr_block
from subnet s, instance_network_interface ini, instance i, instance_tag it
where s.id=ini.ref_subnet_id and ini.instance_id=i.id and i.id=it.instance_id and it.key='Name' ;

CREATE VIEW public.view_db_instance_input_cidr AS
select di.id as db_instance_id, di.db_name as db_instance_name, sgipir.cidr_ip as subnet_cidr_block
from db_instance di, db_instance_vpc_security_groups divsg, security_group sg, security_group_ip_permissions sgip, security_group_ip_permissions_ip_ranges sgipir
where di.id=divsg.db_instance_id and divsg.ref_vpc_security_group_id=sg.id and sg.id=sgip.security_group_id and sgip.id=sgipir.security_group_ip_permissions_id;

CREATE VIEW public.view_elb_targets_health AS
select lb.id,lb.load_balancer_arn,lb.load_balancer_name,lbl.port as load_balancer_port,lbl.protocol as load_balancer_protocol,lblda.type,lbtg.health_check_interval_seconds,lbtg.health_check_timeout_seconds,lbtg.healthy_threshold_count,lbtg.unhealthy_threshold_count,lbtg.health_check_path,lbtg.http_code,lbtg.target_type,lbtgh.target_health_state,i.instance_id,it.value as instance_name
from load_balancer lb, load_balancer_listener lbl, load_balancer_listener_default_actions lblda, load_balancer_target_group lbtg, load_balancer_target_group_health lbtgh, instance i, instance_tag it
where lb.id=lbl.load_balancer_id and lbl.id=lblda.load_balancer_listener_id and lblda.target_group_id=lbtg.id and lbtg.id=lbtgh.load_balancer_target_group_id and lbtgh.ref_instance_id=i.id and i.id=it.instance_id and it.key='Name';
-- select distinct(instance_id),instance_name from view_instance_by_cidr where cidr_block like '%10.1.30.%' ;
-- select distinct(instance_id),instance_name from view_instance_by_cidr where cidr_block like '%10.1.31.%' ;
-- select distinct(id),load_balancer_name,load_balancer_port,load_balancer_protocol,type,http_code,target_type,target_health_state,instance_id,instance_name from view_elb_targets_health ;
