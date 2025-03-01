-- #############################################################################
--
--    
-- #############################################################################

ALTER TABLE public.vpc ADD CONSTRAINT pk_vpc PRIMARY KEY(id);
ALTER TABLE public.subnet ADD CONSTRAINT pk_subnet PRIMARY KEY(id);
ALTER TABLE public.bucket ADD CONSTRAINT pk_bucket PRIMARY KEY(id);
ALTER TABLE public.ami ADD CONSTRAINT pk_ami PRIMARY KEY(id);


--Stack
ALTER TABLE public.stack ADD CONSTRAINT pk_stack PRIMARY KEY(id);
ALTER TABLE public.stack_parameter ADD CONSTRAINT pk_stack_parameter PRIMARY KEY(id);
ALTER TABLE public.stack_output ADD CONSTRAINT pk_stack_output PRIMARY KEY(id);
ALTER TABLE public.stack_tag ADD CONSTRAINT pk_stack_tag PRIMARY KEY(id);


--Volume
ALTER TABLE public.volume ADD CONSTRAINT pk_volume PRIMARY KEY(id);
ALTER TABLE public.volume_tag ADD CONSTRAINT pk_volume_tag PRIMARY KEY(id);
ALTER TABLE public.volume_attachment ADD CONSTRAINT pk_volume_attachment PRIMARY KEY(id);


--SecurityGroup
ALTER TABLE public.security_group ADD CONSTRAINT pk_security_group PRIMARY KEY(id);
ALTER TABLE public.security_group_tag ADD CONSTRAINT pk_security_group_tag PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_egress ADD CONSTRAINT pk_security_group_ip_permissions_egress PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions ADD CONSTRAINT pk_security_group_ip_permissions PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_egress_ip_ranges ADD CONSTRAINT pk_security_group_ip_permissions_egress_ip_ranges PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_egress_prefix_list_ids ADD CONSTRAINT pk_security_group_ip_permissions_egress_prefix_list_ids PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_egress_user_id_group_pairs ADD CONSTRAINT pk_security_group_ip_permissions_egress_user_id_group_pairs PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_ip_ranges ADD CONSTRAINT pk_security_group_ip_permissions_ip_ranges PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_prefix_list_ids ADD CONSTRAINT pk_security_group_ip_permissions_prefix_list_ids PRIMARY KEY(id);
ALTER TABLE public.security_group_ip_permissions_user_id_group_pairs ADD CONSTRAINT pk_security_group_ip_permissions_user_id_group_pairs PRIMARY KEY(id);


--DBInstance
ALTER TABLE  public.db_instance ADD CONSTRAINT pk_db_instance PRIMARY KEY(id);
ALTER TABLE  public.db_instance_associated_roles ADD CONSTRAINT pk_db_instance_associated_roles PRIMARY KEY(id);
ALTER TABLE  public.db_instance_processor_features ADD CONSTRAINT pk_db_instance_processor_features PRIMARY KEY(id);
ALTER TABLE  public.db_instance_domain_memberships ADD CONSTRAINT pk_db_instance_domain_memberships PRIMARY KEY(id);
ALTER TABLE  public.db_instance_status_infos ADD CONSTRAINT pk_db_instance_status_infos PRIMARY KEY(id);
ALTER TABLE  public.db_instance_option_group_memberships ADD CONSTRAINT pk_db_instance_option_group_memberships PRIMARY KEY(id);
ALTER TABLE  public.db_instance_db_parameter_groups ADD CONSTRAINT pk_db_instance_db_parameter_groups PRIMARY KEY(id);
ALTER TABLE  public.db_instance_vpc_security_groups ADD CONSTRAINT pk_db_instance_vpc_security_groups PRIMARY KEY(id);
ALTER TABLE  public.db_instance_db_security_groups  ADD CONSTRAINT pk_db_instance_db_security_groups PRIMARY KEY(id);
ALTER TABLE  public.db_instance_db_subnet_group  ADD CONSTRAINT pk_db_instance_db_subnet_group PRIMARY KEY(id);

ALTER TABLE  public.db_instance_db_subnet_group_subnet ADD CONSTRAINT pk_db_instance_db_subnet_group_subnet PRIMARY KEY(id);


--Instance
ALTER TABLE  public.reservation ADD CONSTRAINT pk_reservation PRIMARY KEY(id);

ALTER TABLE  public.reservation_group ADD CONSTRAINT  pk_reservation_group PRIMARY KEY(id);
ALTER TABLE  public.instance ADD CONSTRAINT  pk_instance PRIMARY KEY(id);

ALTER TABLE  public.instance_product_code ADD CONSTRAINT  pk_instance_product_code PRIMARY KEY(id);
ALTER TABLE  public.instance_block_device_mapping ADD CONSTRAINT  pk_instance_block_device_mapping PRIMARY KEY(id);
ALTER TABLE  public.instance_elastic_gpu_association ADD CONSTRAINT  pk_instance_elastic_gpu_association PRIMARY KEY(id);
ALTER TABLE  public.instance_elastic_inference_accelerator_association ADD CONSTRAINT  pk_instance_elastic_inference_accelerator_association PRIMARY KEY(id);
ALTER TABLE  public.instance_network_interface ADD CONSTRAINT  pk_instance_network_interface PRIMARY KEY(id);
ALTER TABLE  public.instance_network_interface_group ADD CONSTRAINT pk_instance_network_interface_group PRIMARY KEY(id);
ALTER TABLE  public.instance_network_interface_private_ip_address ADD CONSTRAINT pk_instance_network_interface_private_ip_address PRIMARY KEY(id);
ALTER TABLE  public.instance_security_group ADD CONSTRAINT  pk_instance_security_group PRIMARY KEY(id);
ALTER TABLE  public.instance_tag ADD CONSTRAINT  pk_instance_tag PRIMARY KEY(id);
ALTER TABLE  public.instance_license ADD CONSTRAINT  pk_instance_license PRIMARY KEY(id);

--LoadBalancer
ALTER TABLE  public.load_balancer ADD CONSTRAINT pk_load_balancer PRIMARY KEY(id);
ALTER TABLE  public.load_balancer_availability_zone ADD CONSTRAINT pk_load_balancer_availability_zone PRIMARY KEY(id);
ALTER TABLE  public.load_balancer_security_group ADD CONSTRAINT pk_load_balancer_security_group PRIMARY KEY(id);
ALTER TABLE  public.load_balancer_availability_zone_load_balancer_address ADD CONSTRAINT pk_load_balancer_availability_zone_load_balancer_address PRIMARY KEY(id);

ALTER TABLE public.load_balancer_listener ADD CONSTRAINT pk_load_balancer_listener PRIMARY KEY(id);
ALTER TABLE public.load_balancer_listener_default_actions ADD CONSTRAINT pk_load_balancer_listener_default_actions PRIMARY KEY(id);
ALTER TABLE public.load_balancer_listener_certificats ADD CONSTRAINT pk_load_balancer_listener_certificats PRIMARY KEY(id);

--AutoScalingGroups
ALTER TABLE public.auto_scaling_group ADD CONSTRAINT pk_auto_scaling_group PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_launch_template_override ADD CONSTRAINT pk_auto_scaling_group_launch_template_override PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_availability_zone ADD CONSTRAINT pk_auto_scaling_group_availability_zone PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_load_balancer_name ADD CONSTRAINT pk_auto_scaling_group_load_balancer_name PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_target_group_arn ADD CONSTRAINT pk_auto_scaling_group_target_group_arn PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_instance ADD CONSTRAINT pk_auto_scaling_group_instance PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_suspended_process ADD CONSTRAINT pk_auto_scaling_group_suspended_process PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_enabled_metric ADD CONSTRAINT pk_auto_scaling_group_enabled_metric PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_tag ADD CONSTRAINT pk_auto_scaling_group_tag PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_group_termination_policy ADD CONSTRAINT pk_auto_scaling_group_termination_policy PRIMARY KEY(id);
ALTER TABLE public.auto_scaling_activities ADD CONSTRAINT pk_auto_scaling_activities PRIMARY KEY(id);

--LoadBalancerTargetGroups
ALTER TABLE ONLY public.load_balancer_target_group ADD CONSTRAINT pk_load_balancer_target_group PRIMARY KEY(id);
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn ADD CONSTRAINT pk_load_balancer_target_group_load_balancer_arn PRIMARY KEY(id);
ALTER TABLE ONLY public.load_balancer_target_group_health ADD CONSTRAINT pk_load_balancer_target_group_health PRIMARY KEY (id) ;

--Snapshot
ALTER TABLE ONLY public.snapshot ADD CONSTRAINT pk_snapshot PRIMARY KEY(id);

