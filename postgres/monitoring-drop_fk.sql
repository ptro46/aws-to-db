-- #############################################################################
--
--    
-- #############################################################################

--Subnet
ALTER TABLE ONLY public.subnet DROP CONSTRAINT fk_subnet_vpc ;

--Ami
ALTER TABLE ONLY public.ami DROP CONSTRAINT fk_ami_snapshot;

--Stack
ALTER TABLE ONLY public.stack_parameter DROP CONSTRAINT fk_stack_parameter_stack ;
ALTER TABLE ONLY public.stack_output DROP CONSTRAINT fk_stack_output_stack ;
ALTER TABLE ONLY public.stack_tag DROP CONSTRAINT fk_stack_tag_stack ;

--Volume
ALTER TABLE ONLY public.volume_tag DROP CONSTRAINT fk_volume_tag_volume ;
ALTER TABLE ONLY public.volume_attachment DROP CONSTRAINT fk_volume_attachment_volume ;
ALTER TABLE ONLY public.volume_attachment DROP CONSTRAINT fk_volume_attachment_instance ;

--SecurityGroup
ALTER TABLE ONLY public.security_group DROP CONSTRAINT fk_security_group_vpc ;

ALTER TABLE ONLY public.security_group_tag DROP CONSTRAINT fk_security_group_tag_security_group ;
ALTER TABLE ONLY public.security_group_ip_permissions_egress DROP CONSTRAINT fk_security_group_ip_permissions_egress_security_group ;
ALTER TABLE ONLY public.security_group_ip_permissions DROP CONSTRAINT fk_security_group_ip_permissions_security_group ;

ALTER TABLE ONLY public.security_group_ip_permissions_egress_ip_ranges DROP CONSTRAINT fk_sec_grp_ip_perm_egress_ip_rng_sec_grp_ip_perm_egress ;
ALTER TABLE ONLY public.security_group_ip_permissions_egress_prefix_list_ids DROP CONSTRAINT fk_sec_grp_ip_perm_egress_prefix_lst_ids_sec_grp_ip_perm_egress ;
ALTER TABLE ONLY public.security_group_ip_permissions_egress_user_id_group_pairs DROP CONSTRAINT fk_sec_grp_ip_perm_egr_user_id_grp_pairs_sec_grp_ip_perm_egr ;

ALTER TABLE ONLY public.security_group_ip_permissions_egress_user_id_group_pairs DROP CONSTRAINT fk_sec_grp_ip_perm_egr_user_id_grp_pairs_vpc ;

ALTER TABLE ONLY public.security_group_ip_permissions_ip_ranges DROP CONSTRAINT fk_sec_grp_ip_perm_ip_rng_sec_grp_ip_perm ;
ALTER TABLE ONLY public.security_group_ip_permissions_prefix_list_ids DROP CONSTRAINT fk_sec_grp_ip_perm_prefix_lst_ids_sec_grp_ip_perm ;
ALTER TABLE ONLY public.security_group_ip_permissions_user_id_group_pairs DROP CONSTRAINT fk_sec_grp_ip_perm_user_id_grp_pairs_sec_grp_ip_perm ;
ALTER TABLE ONLY public.security_group_ip_permissions_user_id_group_pairs DROP CONSTRAINT fk_sec_grp_ip_perm_user_id_grp_pairs_vpc ;

--DBInstance
ALTER TABLE ONLY public.db_instance_associated_roles DROP CONSTRAINT fk_db_instance_associated_role_db_instance ;
ALTER TABLE ONLY public.db_instance_processor_features DROP CONSTRAINT fk_db_instance_processor_features_db_instance ; 
ALTER TABLE ONLY public.db_instance_domain_memberships DROP CONSTRAINT fk_db_instance_domain_memberships_db_instance ;
ALTER TABLE ONLY public.db_instance_status_infos DROP CONSTRAINT fk_db_instance_status_infos_db_instance ;
ALTER TABLE ONLY public.db_instance_option_group_memberships DROP CONSTRAINT fk_db_instance_option_group_memberships_db_instance ;
ALTER TABLE ONLY public.db_instance_db_parameter_groups DROP CONSTRAINT fk_db_instance_db_parameter_groups_db_instance ;
ALTER TABLE ONLY public.db_instance_vpc_security_groups DROP CONSTRAINT fk_db_instance_vpc_security_groups_db_instance ;
ALTER TABLE ONLY public.db_instance_vpc_security_groups DROP CONSTRAINT fk_db_instance_vpc_security_groups_sg ;
ALTER TABLE ONLY public.db_instance_db_security_groups  DROP CONSTRAINT fk_db_instance_db_security_groups_db_instance ;
ALTER TABLE ONLY public.db_instance_db_security_groups  DROP CONSTRAINT fk_db_instance_db_security_groups_security_group ;

ALTER TABLE ONLY public.db_instance_db_subnet_group  DROP CONSTRAINT fk_db_instance_db_subnet_group_db_instance ;

ALTER TABLE ONLY public.db_instance_db_subnet_group_subnet DROP CONSTRAINT fk_db_i_db_subnet_grp_subnet_db_i_db_subnet_grp ;
ALTER TABLE ONLY public.db_instance_db_subnet_group_subnet DROP CONSTRAINT fk_db_i_db_subnet_grp_subnet_subnet ;
ALTER TABLE ONLY public.db_instance_db_subnet_group DROP CONSTRAINT fk_db_i_db_subnet_grp_vpc ;

--Instance
ALTER TABLE ONLY public.reservation_group DROP CONSTRAINT fk_reservation_group_reservation ;
ALTER TABLE ONLY public.instance DROP CONSTRAINT fk_instance_reservation ;
ALTER TABLE ONLY public.instance DROP CONSTRAINT fk_instance_ami ;
ALTER TABLE ONLY public.instance DROP CONSTRAINT fk_instance_vpc ;
ALTER TABLE ONLY public.instance DROP CONSTRAINT fk_instance_subnet ;
ALTER TABLE ONLY public.instance_product_code DROP CONSTRAINT fk_instance_product_code_instance ;
ALTER TABLE ONLY public.instance_block_device_mapping DROP CONSTRAINT fk_instance_block_device_mapping_instance ;
ALTER TABLE ONLY public.instance_block_device_mapping DROP CONSTRAINT fk_instance_block_device_mapping_volume ;

ALTER TABLE ONLY public.instance_elastic_gpu_association DROP CONSTRAINT fk_instance_elastic_gpu_association_instance ;
ALTER TABLE ONLY public.instance_elastic_inference_accelerator_association DROP CONSTRAINT fk_instance_elastic_inference_accelerator_association_instance ;
ALTER TABLE ONLY public.instance_network_interface DROP CONSTRAINT fk_instance_network_interface_instance ;
ALTER TABLE ONLY public.instance_network_interface DROP CONSTRAINT fk_instance_network_interface_subnet ;
ALTER TABLE ONLY public.instance_network_interface DROP CONSTRAINT fk_instance_network_interface_vpc ;


ALTER TABLE ONLY public.instance_network_interface_group DROP CONSTRAINT fk_instance_network_interface_group_instance_network_interface ;
ALTER TABLE ONLY public.instance_network_interface_group DROP CONSTRAINT fk_instance_network_interface_group_security_group ;


ALTER TABLE ONLY public.instance_network_interface_private_ip_address DROP CONSTRAINT fk_inst_net_if_priv_ip_address_inst_net_if ;
ALTER TABLE ONLY public.instance_security_group DROP CONSTRAINT fk_instance_security_group_instance ;
ALTER TABLE ONLY public.instance_security_group DROP CONSTRAINT fk_instance_security_group_security_group ;

ALTER TABLE ONLY public.instance_tag DROP CONSTRAINT fk_instance_tag_instance ;
ALTER TABLE ONLY public.instance_license DROP CONSTRAINT fk_instance_license_instance ;

--LoadBalancers
ALTER TABLE ONLY public.load_balancer DROP CONSTRAINT fk_load_balancer_vpc ;
ALTER TABLE ONLY public.load_balancer_availability_zone DROP CONSTRAINT fk_load_balancer_availability_zone_load_balancer ;
ALTER TABLE ONLY public.load_balancer_availability_zone DROP CONSTRAINT fk_load_balancer_availability_zone_subnet ;

ALTER TABLE ONLY public.load_balancer_security_group DROP CONSTRAINT fk_load_balancer_security_group_load_balancer ;
ALTER TABLE ONLY public.load_balancer_security_group DROP CONSTRAINT fk_load_balancer_security_group_security_group ;

ALTER TABLE ONLY public.load_balancer_availability_zone_load_balancer_address DROP CONSTRAINT fk_lb_availability_zone_lb_address_lb_availability_zone ;

ALTER TABLE ONLY public.load_balancer_listener DROP CONSTRAINT fk_lbl_load_balancer ;
ALTER TABLE ONLY public.load_balancer_listener_default_actions DROP CONSTRAINT fk_lbl_default_actions_lbl ;
ALTER TABLE ONLY public.load_balancer_listener_default_actions DROP CONSTRAINT fk_lbl_default_actions_target_group ;
ALTER TABLE ONLY public.load_balancer_listener_certificats DROP CONSTRAINT fk_lbl_certificats_lbl ;

--AutoScalingGroups
ALTER TABLE ONLY public.auto_scaling_group DROP CONSTRAINT fk_auto_scaling_group_subnet ;
ALTER TABLE ONLY public.auto_scaling_group_launch_template_override DROP CONSTRAINT fk_asg_launch_template_override_asg ;
ALTER TABLE ONLY public.auto_scaling_group_availability_zone DROP CONSTRAINT fk_auto_scaling_group_availability_zone_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_load_balancer_name DROP CONSTRAINT fk_auto_scaling_group_load_balancer_name_auto_scaling_group ; 
ALTER TABLE ONLY public.auto_scaling_group_target_group_arn DROP CONSTRAINT fk_auto_scaling_group_target_group_arn_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_instance DROP CONSTRAINT fk_auto_scaling_group_instance_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_instance DROP CONSTRAINT fk_auto_scaling_group_instance_instance ;

ALTER TABLE ONLY public.auto_scaling_group_suspended_process DROP CONSTRAINT fk_auto_scaling_group_suspended_process_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_enabled_metric DROP CONSTRAINT fk_auto_scaling_group_enabled_metric_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_tag DROP CONSTRAINT fk_auto_scaling_group_tag_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_group_termination_policy DROP CONSTRAINT fk_auto_scaling_group_termination_policy_auto_scaling_group ;
ALTER TABLE ONLY public.auto_scaling_activities DROP CONSTRAINT fk_auto_scaling_activities_auto_scaling_group ;

--LoadBalancerTargetGroup
ALTER TABLE ONLY public.load_balancer_target_group DROP CONSTRAINT fk_load_balancer_target_group_vpc ;
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn DROP CONSTRAINT fk_lb_tgt_grp_lb_arn_load_balancer ;
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn DROP CONSTRAINT fk_lb_tgt_grp_lb_arn_load_balancer_target_group ;
ALTER TABLE ONLY public.load_balancer_target_group_health DROP CONSTRAINT fk_load_balancer_target_group_health_load_balancer_target_group ;
ALTER TABLE ONLY public.load_balancer_target_group_health DROP CONSTRAINT fk_load_balancer_target_group_health_instance ;

--Snapshot
ALTER TABLE ONLY public.snapshot DROP CONSTRAINT fk_snapshot_volume  ;

