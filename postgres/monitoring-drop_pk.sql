-- #############################################################################
--
--    
-- #############################################################################

ALTER TABLE public.vpc DROP CONSTRAINT pk_vpc ;
ALTER TABLE public.subnet DROP CONSTRAINT pk_subnet ;
ALTER TABLE public.bucket DROP CONSTRAINT pk_bucket ;
ALTER TABLE public.ami DROP CONSTRAINT pk_ami ;

--Stack
ALTER TABLE public.stack DROP CONSTRAINT pk_stack ;
ALTER TABLE public.stack_parameter DROP CONSTRAINT pk_stack_parameter ;
ALTER TABLE public.stack_output DROP CONSTRAINT pk_stack_output ;
ALTER TABLE public.stack_tag DROP CONSTRAINT pk_stack_tag ;

--Volume
ALTER TABLE public.volume DROP CONSTRAINT pk_volume ;
ALTER TABLE public.volume_tag DROP CONSTRAINT pk_volume_tag ;
ALTER TABLE public.volume_attachment DROP CONSTRAINT pk_volume_attachment ;

--SecurityGroup
ALTER TABLE public.security_group DROP CONSTRAINT pk_security_group ;
ALTER TABLE public.security_group_tag DROP CONSTRAINT pk_security_group_tag ;
ALTER TABLE public.security_group_ip_permissions_egress DROP CONSTRAINT pk_security_group_ip_permissions_egress ;
ALTER TABLE public.security_group_ip_permissions DROP CONSTRAINT pk_security_group_ip_permissions ;
ALTER TABLE public.security_group_ip_permissions_egress_ip_ranges DROP CONSTRAINT pk_security_group_ip_permissions_egress_ip_ranges ;
ALTER TABLE public.security_group_ip_permissions_egress_prefix_list_ids DROP CONSTRAINT pk_security_group_ip_permissions_egress_prefix_list_ids ;
ALTER TABLE public.security_group_ip_permissions_egress_user_id_group_pairs DROP CONSTRAINT pk_security_group_ip_permissions_egress_user_id_group_pairs ;
ALTER TABLE public.security_group_ip_permissions_ip_ranges DROP CONSTRAINT pk_security_group_ip_permissions_ip_ranges ;
ALTER TABLE public.security_group_ip_permissions_prefix_list_ids DROP CONSTRAINT pk_security_group_ip_permissions_prefix_list_ids ;
ALTER TABLE public.security_group_ip_permissions_user_id_group_pairs DROP CONSTRAINT pk_security_group_ip_permissions_user_id_group_pairs ;

--DBInstance
ALTER TABLE  public.db_instance DROP CONSTRAINT pk_db_instance ;
ALTER TABLE  public.db_instance_associated_roles DROP CONSTRAINT pk_db_instance_associated_roles ;
ALTER TABLE  public.db_instance_processor_features DROP CONSTRAINT pk_db_instance_processor_features ;
ALTER TABLE  public.db_instance_domain_memberships DROP CONSTRAINT pk_db_instance_domain_memberships ;
ALTER TABLE  public.db_instance_status_infos DROP CONSTRAINT pk_db_instance_status_infos ;
ALTER TABLE  public.db_instance_option_group_memberships DROP CONSTRAINT pk_db_instance_option_group_memberships ;
ALTER TABLE  public.db_instance_db_parameter_groups DROP CONSTRAINT pk_db_instance_db_parameter_groups ;
ALTER TABLE  public.db_instance_vpc_security_groups DROP CONSTRAINT pk_db_instance_vpc_security_groups ;
ALTER TABLE  public.db_instance_db_security_groups  DROP CONSTRAINT pk_db_instance_db_security_groups;
ALTER TABLE  public.db_instance_db_subnet_group  DROP CONSTRAINT pk_db_instance_db_subnet_group ;

ALTER TABLE  public.db_instance_db_subnet_group_subnet DROP CONSTRAINT pk_db_instance_db_subnet_group_subnet ;

--Instance
ALTER TABLE  public.reservation DROP CONSTRAINT pk_reservation ;
ALTER TABLE  public.reservation_group DROP CONSTRAINT  pk_reservation_group ;
ALTER TABLE  public.instance DROP CONSTRAINT  pk_instance ;
ALTER TABLE  public.instance_product_code DROP CONSTRAINT  pk_instance_product_code ;
ALTER TABLE  public.instance_block_device_mapping DROP CONSTRAINT  pk_instance_block_device_mapping ;
ALTER TABLE  public.instance_elastic_gpu_association DROP CONSTRAINT  pk_instance_elastic_gpu_association ;
ALTER TABLE  public.instance_elastic_inference_accelerator_association DROP CONSTRAINT  pk_instance_elastic_inference_accelerator_association ;
ALTER TABLE  public.instance_network_interface DROP CONSTRAINT  pk_instance_network_interface ;
ALTER TABLE  public.instance_network_interface_group DROP CONSTRAINT pk_instance_network_interface_group ;
ALTER TABLE  public.instance_network_interface_private_ip_address DROP CONSTRAINT pk_instance_network_interface_private_ip_address ;
ALTER TABLE  public.instance_security_group DROP CONSTRAINT  pk_instance_security_group ;
ALTER TABLE  public.instance_tag DROP CONSTRAINT  pk_instance_tag ;
ALTER TABLE  public.instance_license DROP CONSTRAINT  pk_instance_license ;

--LoadBalancers
ALTER TABLE  public.load_balancer DROP CONSTRAINT pk_load_balancer ;
ALTER TABLE  public.load_balancer_availability_zone DROP CONSTRAINT pk_load_balancer_availability_zone ;
ALTER TABLE public.load_balancer_security_group DROP CONSTRAINT pk_load_balancer_security_group ;
ALTER TABLE  public.load_balancer_availability_zone_load_balancer_address DROP CONSTRAINT pk_load_balancer_availability_zone_load_balancer_address ;

ALTER TABLE public.load_balancer_listener DROP CONSTRAINT pk_load_balancer_listener ;
ALTER TABLE public.load_balancer_listener_default_actions DROP CONSTRAINT pk_load_balancer_listener_default_actions ;
ALTER TABLE public.load_balancer_listener_certificats DROP CONSTRAINT pk_load_balancer_listener_certificats ;

--AutoScalingGroups
ALTER TABLE ONLY public.auto_scaling_group DROP CONSTRAINT pk_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_launch_template_override DROP CONSTRAINT pk_auto_scaling_group_launch_template_override;
ALTER TABLE ONLY public.auto_scaling_group_availability_zone DROP CONSTRAINT pk_auto_scaling_group_availability_zone;
ALTER TABLE ONLY public.auto_scaling_group_load_balancer_name DROP CONSTRAINT pk_auto_scaling_group_load_balancer_name;
ALTER TABLE ONLY public.auto_scaling_group_target_group_arn DROP CONSTRAINT pk_auto_scaling_group_target_group_arn;
ALTER TABLE ONLY public.auto_scaling_group_instance DROP CONSTRAINT pk_auto_scaling_group_instance;
ALTER TABLE ONLY public.auto_scaling_group_suspended_process DROP CONSTRAINT pk_auto_scaling_group_suspended_process;
ALTER TABLE ONLY public.auto_scaling_group_enabled_metric DROP CONSTRAINT pk_auto_scaling_group_enabled_metric;
ALTER TABLE ONLY public.auto_scaling_group_tag DROP CONSTRAINT pk_auto_scaling_group_tag;
ALTER TABLE ONLY public.auto_scaling_group_termination_policy DROP CONSTRAINT pk_auto_scaling_group_termination_policy;
ALTER TABLE ONLY public.auto_scaling_activities DROP CONSTRAINT pk_auto_scaling_activities;

--LoadBalancerTargetGroups
ALTER TABLE ONLY public.load_balancer_target_group DROP CONSTRAINT pk_load_balancer_target_group ;
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn DROP CONSTRAINT pk_load_balancer_target_group_load_balancer_arn ;
ALTER TABLE ONLY public.load_balancer_target_group_health DROP CONSTRAINT pk_load_balancer_target_group_health ;

--Snapshot
ALTER TABLE ONLY public.snapshot DROP CONSTRAINT pk_snapshot ;

