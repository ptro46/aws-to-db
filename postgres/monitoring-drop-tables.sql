-- #############################################################################
--
--    
-- #############################################################################

DROP TABLE public.config_subnet CASCADE;

DROP TABLE public.vpc CASCADE ;
DROP TABLE public.subnet CASCADE ;
DROP TABLE public.bucket CASCADE ;
DROP TABLE public.ami CASCADE ;

--Stack
DROP TABLE public.stack CASCADE ;
DROP TABLE public.stack_parameter CASCADE ;
DROP TABLE public.stack_output CASCADE ;
DROP TABLE public.stack_tag CASCADE ;

--Volume
DROP TABLE public.volume CASCADE ;
DROP TABLE public.volume_tag CASCADE ;
DROP TABLE public.volume_attachment CASCADE ;

--SecurityGroup
DROP TABLE public.security_group CASCADE ;
DROP TABLE public.security_group_tag CASCADE ;
DROP TABLE public.security_group_ip_permissions_egress CASCADE ;
DROP TABLE public.security_group_ip_permissions CASCADE ;
DROP TABLE public.security_group_ip_permissions_egress_ip_ranges CASCADE ;
DROP TABLE public.security_group_ip_permissions_egress_prefix_list_ids CASCADE ;
DROP TABLE public.security_group_ip_permissions_egress_user_id_group_pairs CASCADE ;
DROP TABLE public.security_group_ip_permissions_ip_ranges CASCADE ;
DROP TABLE public.security_group_ip_permissions_prefix_list_ids CASCADE ;
DROP TABLE public.security_group_ip_permissions_user_id_group_pairs CASCADE ;

--DBInstance
DROP TABLE  public.db_instance CASCADE ;
DROP TABLE  public.db_instance_associated_roles CASCADE ;
DROP TABLE  public.db_instance_processor_features CASCADE ;
DROP TABLE  public.db_instance_domain_memberships CASCADE ;
DROP TABLE  public.db_instance_status_infos CASCADE ;
DROP TABLE  public.db_instance_option_group_memberships CASCADE ;
DROP TABLE  public.db_instance_db_parameter_groups CASCADE ;
DROP TABLE  public.db_instance_vpc_security_groups CASCADE ;
DROP TABLE  public.db_instance_db_security_groups  CASCADE ;
DROP TABLE  public.db_instance_db_subnet_group  CASCADE ;
DROP TABLE  public.db_instance_db_subnet_group_subnet CASCADE ;

--Instance
DROP TABLE  public.reservation CASCADE ;

DROP TABLE  public.reservation_group  CASCADE ;
DROP TABLE  public.instance  CASCADE ;

DROP TABLE  public.instance_product_code  CASCADE ;
DROP TABLE  public.instance_block_device_mapping  CASCADE ;
DROP TABLE  public.instance_elastic_gpu_association  CASCADE ;
DROP TABLE  public.instance_elastic_inference_accelerator_association CASCADE ;
DROP TABLE  public.instance_network_interface  CASCADE ;
DROP TABLE  public.instance_network_interface_group CASCADE ;
DROP TABLE  public.instance_network_interface_private_ip_address CASCADE ;
DROP TABLE  public.instance_security_group  CASCADE ;
DROP TABLE  public.instance_tag  CASCADE ;
DROP TABLE  public.instance_license  CASCADE ;

--LoadBalancer
DROP TABLE  public.load_balancer CASCADE ;
DROP TABLE  public.load_balancer_availability_zone CASCADE ;
DROP TABLE  public.load_balancer_security_group CASCADE ;
DROP TABLE  public.load_balancer_availability_zone_load_balancer_address CASCADE ;

DROP TABLE public.load_balancer_listener CASCADE ;
DROP TABLE public.load_balancer_listener_default_actions CASCADE ;
DROP TABLE public.load_balancer_listener_certificats CASCADE ;

--AutoScalingGroups
DROP TABLE public.auto_scaling_group CASCADE ;
DROP TABLE public.auto_scaling_group_launch_template_override CASCADE ;
DROP TABLE public.auto_scaling_group_availability_zone CASCADE ;
DROP TABLE public.auto_scaling_group_load_balancer_name CASCADE ;
DROP TABLE public.auto_scaling_group_target_group_arn CASCADE ;
DROP TABLE public.auto_scaling_group_instance CASCADE ;
DROP TABLE public.auto_scaling_group_suspended_process CASCADE ;
DROP TABLE public.auto_scaling_group_enabled_metric CASCADE ;
DROP TABLE public.auto_scaling_group_tag CASCADE ;
DROP TABLE public.auto_scaling_group_termination_policy CASCADE ;
DROP TABLE public.auto_scaling_activities CASCADE ;

--LoadBalancerTargetGroups
DROP TABLE public.load_balancer_target_group CASCADE ;
DROP TABLE public.load_balancer_target_group_load_balancer_arn CASCADE ;
DROP TABLE public.load_balancer_target_group_health CASCADE ;

---Snapshot
DROP TABLE public.snapshot CASCADE ;



