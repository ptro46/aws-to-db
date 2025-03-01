-- #############################################################################
--
--    
-- #############################################################################

--Subnet
ALTER TABLE ONLY public.subnet ADD CONSTRAINT fk_subnet_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);

--Ami
ALTER TABLE ONLY public.ami ADD CONSTRAINT fk_ami_snapshot FOREIGN KEY (ref_snapshot_id) REFERENCES public.snapshot(id);


--Stack
ALTER TABLE ONLY public.stack_parameter ADD CONSTRAINT fk_stack_parameter_stack FOREIGN KEY (stack_id) REFERENCES public.stack(id);
ALTER TABLE ONLY public.stack_output ADD CONSTRAINT fk_stack_output_stack FOREIGN KEY (stack_id) REFERENCES public.stack(id);
ALTER TABLE ONLY public.stack_tag ADD CONSTRAINT fk_stack_tag_stack FOREIGN KEY (stack_id) REFERENCES public.stack(id);

--Volume
ALTER TABLE ONLY public.volume_tag ADD CONSTRAINT fk_volume_tag_volume FOREIGN KEY (volume_id) REFERENCES public.volume(id);
ALTER TABLE ONLY public.volume_attachment ADD CONSTRAINT fk_volume_attachment_volume FOREIGN KEY (volume_id) REFERENCES public.volume(id);
ALTER TABLE ONLY public.volume_attachment ADD CONSTRAINT fk_volume_attachment_instance FOREIGN KEY (ref_instance_id) REFERENCES public.instance(id);


--SecurityGroup
ALTER TABLE ONLY public.security_group ADD CONSTRAINT fk_security_group_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);

ALTER TABLE ONLY public.security_group_tag ADD CONSTRAINT fk_security_group_tag_security_group FOREIGN KEY (security_group_id) REFERENCES public.security_group(id);
ALTER TABLE ONLY public.security_group_ip_permissions_egress ADD CONSTRAINT fk_security_group_ip_permissions_egress_security_group FOREIGN KEY (security_group_id) REFERENCES public.security_group(id);
ALTER TABLE ONLY public.security_group_ip_permissions ADD CONSTRAINT fk_security_group_ip_permissions_security_group FOREIGN KEY (security_group_id) REFERENCES public.security_group(id);

ALTER TABLE ONLY public.security_group_ip_permissions_egress_ip_ranges ADD CONSTRAINT fk_sec_grp_ip_perm_egress_ip_rng_sec_grp_ip_perm_egress FOREIGN KEY (security_group_ip_permissions_egress_id) REFERENCES public.security_group_ip_permissions_egress(id);
ALTER TABLE ONLY public.security_group_ip_permissions_egress_prefix_list_ids ADD CONSTRAINT fk_sec_grp_ip_perm_egress_prefix_lst_ids_sec_grp_ip_perm_egress FOREIGN KEY (security_group_ip_permissions_egress_id) REFERENCES public.security_group_ip_permissions_egress(id);
ALTER TABLE ONLY public.security_group_ip_permissions_egress_user_id_group_pairs ADD CONSTRAINT fk_sec_grp_ip_perm_egr_user_id_grp_pairs_sec_grp_ip_perm_egr FOREIGN KEY (security_group_ip_permissions_egress_id) REFERENCES public.security_group_ip_permissions_egress(id);

ALTER TABLE ONLY public.security_group_ip_permissions_egress_user_id_group_pairs ADD CONSTRAINT fk_sec_grp_ip_perm_egr_user_id_grp_pairs_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);


ALTER TABLE ONLY public.security_group_ip_permissions_ip_ranges ADD CONSTRAINT fk_sec_grp_ip_perm_ip_rng_sec_grp_ip_perm FOREIGN KEY (security_group_ip_permissions_id) REFERENCES public.security_group_ip_permissions(id);
ALTER TABLE ONLY public.security_group_ip_permissions_prefix_list_ids ADD CONSTRAINT fk_sec_grp_ip_perm_prefix_lst_ids_sec_grp_ip_perm FOREIGN KEY (security_group_ip_permissions_id) REFERENCES public.security_group_ip_permissions(id);
ALTER TABLE ONLY public.security_group_ip_permissions_user_id_group_pairs ADD CONSTRAINT fk_sec_grp_ip_perm_user_id_grp_pairs_sec_grp_ip_perm FOREIGN KEY (security_group_ip_permissions_id) REFERENCES public.security_group_ip_permissions(id);
ALTER TABLE ONLY public.security_group_ip_permissions_user_id_group_pairs ADD CONSTRAINT fk_sec_grp_ip_perm_user_id_grp_pairs_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);


--DBInstance
ALTER TABLE ONLY public.db_instance_associated_roles ADD CONSTRAINT fk_db_instance_associated_role_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_processor_features ADD CONSTRAINT fk_db_instance_processor_features_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_domain_memberships ADD CONSTRAINT fk_db_instance_domain_memberships_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_status_infos ADD CONSTRAINT fk_db_instance_status_infos_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_option_group_memberships ADD CONSTRAINT fk_db_instance_option_group_memberships_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_db_parameter_groups ADD CONSTRAINT fk_db_instance_db_parameter_groups_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_vpc_security_groups ADD CONSTRAINT fk_db_instance_vpc_security_groups_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_vpc_security_groups ADD CONSTRAINT fk_db_instance_vpc_security_groups_sg FOREIGN KEY (ref_vpc_security_group_id) REFERENCES public.security_group(id);
ALTER TABLE ONLY public.db_instance_db_security_groups  ADD CONSTRAINT fk_db_instance_db_security_groups_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);
ALTER TABLE ONLY public.db_instance_db_security_groups  ADD CONSTRAINT fk_db_instance_db_security_groups_security_group FOREIGN KEY (ref_security_group_id) REFERENCES public.security_group(id);

ALTER TABLE ONLY public.db_instance_db_subnet_group  ADD CONSTRAINT fk_db_instance_db_subnet_group_db_instance FOREIGN KEY (db_instance_id) REFERENCES public.db_instance(id);

ALTER TABLE ONLY public.db_instance_db_subnet_group_subnet ADD CONSTRAINT fk_db_i_db_subnet_grp_subnet_db_i_db_subnet_grp FOREIGN KEY (db_instance_db_subnet_group_id) REFERENCES public.db_instance_db_subnet_group(id);
ALTER TABLE ONLY public.db_instance_db_subnet_group_subnet ADD CONSTRAINT fk_db_i_db_subnet_grp_subnet_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id);
ALTER TABLE ONLY public.db_instance_db_subnet_group ADD CONSTRAINT fk_db_i_db_subnet_grp_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);

--Instance
ALTER TABLE ONLY public.reservation_group ADD CONSTRAINT fk_reservation_group_reservation FOREIGN KEY (reservation_id) REFERENCES public.reservation(id);
ALTER TABLE ONLY public.instance ADD CONSTRAINT fk_instance_reservation FOREIGN KEY (reservation_id) REFERENCES public.reservation(id);
ALTER TABLE ONLY public.instance ADD CONSTRAINT fk_instance_ami FOREIGN KEY (ref_image_id) REFERENCES public.ami(id) ;
ALTER TABLE ONLY public.instance ADD CONSTRAINT fk_instance_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id) ;
ALTER TABLE ONLY public.instance ADD CONSTRAINT fk_instance_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id) ;
ALTER TABLE ONLY public.instance_product_code ADD CONSTRAINT fk_instance_product_code_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_block_device_mapping ADD CONSTRAINT fk_instance_block_device_mapping_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_block_device_mapping ADD CONSTRAINT fk_instance_block_device_mapping_volume FOREIGN KEY (ref_volume_id) REFERENCES public.volume(id);

ALTER TABLE ONLY public.instance_elastic_gpu_association ADD CONSTRAINT fk_instance_elastic_gpu_association_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_elastic_inference_accelerator_association ADD CONSTRAINT fk_instance_elastic_inference_accelerator_association_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_network_interface ADD CONSTRAINT fk_instance_network_interface_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_network_interface ADD CONSTRAINT fk_instance_network_interface_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id);
ALTER TABLE ONLY public.instance_network_interface ADD CONSTRAINT fk_instance_network_interface_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);


ALTER TABLE ONLY public.instance_network_interface_group ADD CONSTRAINT fk_instance_network_interface_group_instance_network_interface FOREIGN KEY (instance_network_interface_id) REFERENCES public.instance_network_interface(id);
ALTER TABLE ONLY public.instance_network_interface_group ADD CONSTRAINT fk_instance_network_interface_group_security_group FOREIGN KEY (ref_security_group_id) REFERENCES public.security_group(id);

ALTER TABLE ONLY public.instance_network_interface_private_ip_address ADD CONSTRAINT fk_inst_net_if_priv_ip_address_inst_net_if FOREIGN KEY (instance_network_interface_id) REFERENCES public.instance_network_interface(id);
ALTER TABLE ONLY public.instance_security_group ADD CONSTRAINT fk_instance_security_group_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_security_group ADD CONSTRAINT fk_instance_security_group_security_group FOREIGN KEY (ref_security_group_id) REFERENCES public.security_group(id);

ALTER TABLE ONLY public.instance_tag ADD CONSTRAINT fk_instance_tag_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_license ADD CONSTRAINT fk_instance_license_instance FOREIGN KEY (instance_id) REFERENCES public.instance(id);

--LoadBalancers
ALTER TABLE ONLY public.load_balancer ADD CONSTRAINT fk_load_balancer_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id) ;
ALTER TABLE ONLY public.load_balancer_availability_zone ADD CONSTRAINT fk_load_balancer_availability_zone_load_balancer FOREIGN KEY (load_balancer_id) REFERENCES public.load_balancer(id) ;
ALTER TABLE ONLY public.load_balancer_availability_zone ADD CONSTRAINT fk_load_balancer_availability_zone_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id) ;

ALTER TABLE ONLY public.load_balancer_security_group ADD CONSTRAINT fk_load_balancer_security_group_load_balancer FOREIGN KEY (load_balancer_id) REFERENCES public.load_balancer(id) ;
ALTER TABLE ONLY public.load_balancer_security_group ADD CONSTRAINT fk_load_balancer_security_group_security_group FOREIGN KEY (ref_security_group_id) REFERENCES public.security_group(id) ;

ALTER TABLE ONLY public.load_balancer_availability_zone_load_balancer_address ADD CONSTRAINT fk_lb_availability_zone_lb_address_lb_availability_zone FOREIGN KEY (load_balancer_availability_zone_id) REFERENCES public.load_balancer_availability_zone(id);

ALTER TABLE ONLY public.load_balancer_listener ADD CONSTRAINT fk_lbl_load_balancer FOREIGN KEY (load_balancer_id) REFERENCES public.load_balancer(id) ;
ALTER TABLE ONLY public.load_balancer_listener_default_actions ADD CONSTRAINT fk_lbl_default_actions_lbl FOREIGN KEY (load_balancer_listener_id) REFERENCES public.load_balancer_listener(id) ;
ALTER TABLE ONLY public.load_balancer_listener_default_actions ADD CONSTRAINT fk_lbl_default_actions_target_group FOREIGN KEY (target_group_id) REFERENCES public.load_balancer_target_group(id) ;
ALTER TABLE ONLY public.load_balancer_listener_certificats ADD CONSTRAINT fk_lbl_certificats_lbl FOREIGN KEY (load_balancer_listener_id) REFERENCES public.load_balancer_listener(id) ;

--AutoScalingGroups
ALTER TABLE ONLY public.auto_scaling_group ADD CONSTRAINT fk_auto_scaling_group_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id) ;
ALTER TABLE ONLY public.auto_scaling_group_launch_template_override ADD CONSTRAINT fk_asg_launch_template_override_asg FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_availability_zone ADD CONSTRAINT fk_auto_scaling_group_availability_zone_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_load_balancer_name ADD CONSTRAINT fk_auto_scaling_group_load_balancer_name_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_target_group_arn ADD CONSTRAINT fk_auto_scaling_group_target_group_arn_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_instance ADD CONSTRAINT fk_auto_scaling_group_instance_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_instance ADD CONSTRAINT fk_auto_scaling_group_instance_instance FOREIGN KEY (ref_instance_id) REFERENCES public.instance(id) ;

ALTER TABLE ONLY public.auto_scaling_group_suspended_process ADD CONSTRAINT fk_auto_scaling_group_suspended_process_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_enabled_metric ADD CONSTRAINT fk_auto_scaling_group_enabled_metric_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_tag ADD CONSTRAINT fk_auto_scaling_group_tag_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_group_termination_policy ADD CONSTRAINT fk_auto_scaling_group_termination_policy_auto_scaling_group FOREIGN KEY (auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;
ALTER TABLE ONLY public.auto_scaling_activities ADD CONSTRAINT fk_auto_scaling_activities_auto_scaling_group FOREIGN KEY (ref_auto_scaling_group_id) REFERENCES public.auto_scaling_group(id) ;

--LoadBalancerTargetGroup
ALTER TABLE ONLY public.load_balancer_target_group ADD CONSTRAINT fk_load_balancer_target_group_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id) ;
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn ADD CONSTRAINT fk_lb_tgt_grp_lb_arn_load_balancer FOREIGN KEY (ref_load_balancer_id) REFERENCES public.load_balancer(id) ;
ALTER TABLE ONLY public.load_balancer_target_group_load_balancer_arn ADD CONSTRAINT fk_lb_tgt_grp_lb_arn_load_balancer_target_group FOREIGN KEY (load_balancer_target_group_id) REFERENCES public.load_balancer_target_group(id) ;
ALTER TABLE ONLY public.load_balancer_target_group_health ADD CONSTRAINT fk_load_balancer_target_group_health_load_balancer_target_group FOREIGN KEY (load_balancer_target_group_id) REFERENCES public.load_balancer_target_group(id) ;
ALTER TABLE ONLY public.load_balancer_target_group_health ADD CONSTRAINT fk_load_balancer_target_group_health_instance FOREIGN KEY (ref_instance_id) REFERENCES public.instance(id) ;

--Snapshot
ALTER TABLE ONLY public.snapshot ADD CONSTRAINT fk_snapshot_volume FOREIGN KEY (ref_volume_id) REFERENCES public.volume(id) ;



