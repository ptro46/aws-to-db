-- #############################################################################
--
--    
-- #############################################################################

--Ptro Logs--
CREATE TABLE public.logentry (
	id 					bigserial,
	folder 				text,
	file 				text,
	backend 			text,
	hsize 				text,
	size 				bigint,
	nblines 			bigint
) ;
--Network Parameters---

CREATE TABLE public.config_subnet (
	id 					bigserial,
	name 				text,
	description 		text,
	env 				text,
	cidr_block 			text
) ;

--Vpc-------------------
CREATE TABLE public.vpc (
	id 					bigserial,
	vpc_id 				text,
	vpc_name 			text,
	cidr_association_id	text,
	cidr_block 			text,
	cidr_block_state	text,
	vpc_state			text,
	owner_id 			text
) ;
--Subnet-------------------
CREATE TABLE public.subnet (
	id 							bigserial,
	ref_vpc_id					bigint,
	subnet_id 					text,
	map_public_at_launch		boolean,
	availability_zone_id		text,
	name 						text,
	cloudformation_stack_name	text,
	available_ip_address_count	int,
	state 						text,
	availability_zone 			text,
	owner_id					text,
	cidr_block 					text 
) ;

--Bucket-------------------
CREATE TABLE public.bucket (
	id 					bigserial,
    creation_date 		text,
    name 				text
);
--Ami-------------------
CREATE TABLE public.ami (
	id bigserial,
	virtualization_type text,
	description text,
	hypervisor text,
	ena_support boolean,
	image_id text,
	state text,
	device_name text,
	snap_shot_id text,
	ref_snapshot_id bigint,
	delete_on_termination boolean,
	volume_type text,
	volume_size int,
	encrypted boolean,
	architecture text,
	image_location text,
	root_device_type text,
	owner_id text,
	root_device_name text,
	creation_date text,
	public boolean,
	image_type text,
	name text
);
--Stack-------------------
CREATE TABLE public.stack (
    id 									bigserial,
    stack_id 							text,
    stack_name 							text,
    change_set_id 						text,
    description 						text,
    creation_time 						text,
    deletion_time 						text,
    last_update_time 					text,
    stack_status 						text,
    stack_status_raison 				text,
    disable_rollback 					boolean,
    timeout_in_minutes 					bigint,
    role_arn 							text,
    enable_termination_protection 		boolean,
    parent_id							text,
    root_id 							text,
    next_token 							text
) ;
--StackParam-------------------
CREATE TABLE public.stack_parameter (
    id 						bigserial,
    stack_id 				bigint,
    parameter_key			text,
    parameter_value			text,
    use_previous_value		boolean,
    resolved_value			text
) ;
--StackOut-------------------
CREATE TABLE public.stack_output (
    id 						bigserial,
    stack_id 				bigint,
    output_key				text,
    output_value			text,
    description 			text,
    export_name 			text
) ;
--StackTag-------------------
CREATE TABLE public.stack_tag (
    id 						bigserial,
    stack_id 				bigint,
    key						text,
    value					text
) ;
--Volume-------------------
CREATE TABLE public.volume (
	id 						bigserial,
	availability_zone 		text,
	encrypted 				boolean,
	volume_type 			text,
	volume_id 				text,
	state 					text,
	iops 					int,
	snap_shot_id 			text,
	create_time 			text,
	size 					int
) ;
--VolumeAtt-------------------
CREATE TABLE public.volume_attachment (
	id bigserial,
	volume_id bigint,
	attach_time text,
	instance_id text,
	ref_instance_id bigint,
	state text,
	delete_on_termination boolean,
	device text
) ;
--VolumeTag-------------------
CREATE TABLE public.volume_tag (
	id 			bigserial,
	volume_id 	bigint,
	key 		text,
	value 		text
) ;

--SecurityGroup-------------------
CREATE TABLE public.security_group (
	id 			bigserial,
	description text,
	group_name text,
	vpc_id text,
	ref_vpc_id bigint,
	owner_id text,
	group_id text
) ;

--SecurityGroupTag-------------------
CREATE TABLE public.security_group_tag (
	id 			bigserial,
	security_group_id bigint,
	key text,
	value text
) ;

--SecurityGroupIpPermissionsEgress-------------------
CREATE TABLE public.security_group_ip_permissions_egress (
	id 			bigserial,
	security_group_id bigint,
	from_port int,
	ip_protocol text,
	to_port int
) ;

--SecurityGroupIpPermissions-------------------
CREATE TABLE public.security_group_ip_permissions (
	id 			bigserial,
	security_group_id bigint,
	from_port int,
	ip_protocol text,
	to_port int
) ;

--SecurityGroupIpPermissionsEgressIpRanges-------------------
CREATE TABLE public.security_group_ip_permissions_egress_ip_ranges (
	id 			bigserial,
	security_group_ip_permissions_egress_id bigint,
	cidr_ip text,
	description text
) ;

--SecurityGroupIpPermissionsIpRanges-------------------
CREATE TABLE public.security_group_ip_permissions_ip_ranges (
	id 			bigserial,
	security_group_ip_permissions_id bigint,
	cidr_ip text,
	description text
) ;

--SecurityGroupIpPermissionsEgressListIds-------------------
CREATE TABLE public.security_group_ip_permissions_egress_prefix_list_ids (
	id 			bigserial,
	security_group_ip_permissions_egress_id bigint,
	description text,
	prefix_list_id text
) ;

--SecurityGroupIpPermissionsPrefixListIds-------------------
CREATE TABLE public.security_group_ip_permissions_prefix_list_ids (
	id 			bigserial,
	security_group_ip_permissions_id bigint,
	description text,
	prefix_list_id text
) ;

--SecurityGroupIpPermissionsEgressUserIdGroupPairs-------------------
CREATE TABLE public.security_group_ip_permissions_egress_user_id_group_pairs (
	id 			bigserial,
	security_group_ip_permissions_egress_id bigint,
	description text,
	group_id text,
	group_name text,
	peering_status text,
	user_id text,
	vpc_id text,
	ref_vpc_id bigint,
	vpc_peering_connection_id text
) ;

--SecurityGroupIpPermissionsUserIdGroupPairs-------------------
CREATE TABLE public.security_group_ip_permissions_user_id_group_pairs (
	id 			bigserial,
	security_group_ip_permissions_id bigint,
	description text,
	group_id text,
	group_name text,
	peering_status text,
	user_id text,
	vpc_id text,
	ref_vpc_id bigint,
	vpc_peering_connection_id text
) ;

--DBInstances-------------------
CREATE TABLE public.db_instance (
	id bigserial,
	db_instance_identifier text,
	db_instance_class text,
	engine text,
	db_instance_status text,
	master_username text,
	db_name text,
	endpoint_address text,
	endpoint_port int,
	endpoint_hosted_zone_id text,
	allocated_storage int,
	instance_create_time text,
	preferred_backup_window text,
	backup_retention_period int,
	availability_zone text,
	preferred_maintenance_window text,
	pending_modified_values_db_instance_class text,
	pending_modified_values_allocated_storage int,
	pending_modified_values_master_user_password text,
	pending_modified_values_port int,
	pending_modified_values_backup_retention_period int,
	pending_modified_values_multi_az boolean,
	pending_modified_values_engine_version text,
	pending_modified_values_license_model text,
	pending_modified_values_iops int,
	pending_modified_values_db_instance_identifier text,
	pending_modified_values_storage_type text,
	pending_modified_values_ca_certificate_identifier text,
	pending_modified_values_db_subnet_group_name text,
	latest_restorable_time text,
	multi_az boolean,
	engine_version text,
	auto_minor_version_upgrade boolean,
	read_replica_source_db_instance_identifier text,
	license_model text,
	iops int,
	character_set_name text,
	secondary_availability_zone text,
	publicly_accessible boolean,
	storage_type text,
	tde_credential_arn text,
	db_instance_port int,
	db_cluster_identifier text,
	storage_encrypted boolean,
	kms_key_id text,
	db_i_resource_id text,
	ca_certificate_identifier text,
	copy_tags_to_snapshot boolean,
	monitoring_interval int,
	enhanced_monitoring_resource_arn text,
	monitoring_role_arn text,
	promotion_tier int,
	db_instance_arn text,
	timezone text,
	iam_database_authentication_enabled boolean,
	performance_insights_enabled boolean,
	performance_insights_kms_key_id text,
	performance_insights_retention_period int,
	deletion_protection text,
	listener_endpoint_address text,
	listener_endpoint_port int,
	listener_endpoint_hosted_zone_id text,
	max_allocated_storage int
) ;

--DBInstancesAssociatedRoles-------------------
CREATE TABLE public.db_instance_associated_roles (
	id bigserial,
	db_instance_id bigint,
	role_arn text,
	feature_name text,
	status text
) ;

--DBInstancesProcessorFeatures-------------------
CREATE TABLE public.db_instance_processor_features (
	id bigserial,
	db_instance_id bigint,
	name text,
	value text
) ;

--DBInstancesDomainMemberships-------------------
CREATE TABLE public.db_instance_domain_memberships (
	id bigserial,
	db_instance_id bigint,
	domain text,
	status text,
	fqdn text,
	iam_role_name text
) ;

--DBInstancesStatusInfos-------------------
CREATE TABLE public.db_instance_status_infos (
	id bigserial,
	db_instance_id bigint,
	status_type text,
	normal boolean,
	status text,
	message text
) ;

--DBInstancesOptionGroupMemberships-------------------
CREATE TABLE public.db_instance_option_group_memberships (
	id bigserial,
	db_instance_id bigint,
	option_group_name text,
	status text
) ;

--DBInstancesParameterGroups-------------------
CREATE TABLE public.db_instance_db_parameter_groups (
	id bigserial,
	db_instance_id bigint,
	db_parameter_group_name text,
	parameter_apply_status text

) ;
--DBInstancesVpcSecurityGroups-------------------
CREATE TABLE public.db_instance_vpc_security_groups (
	id bigserial,
	db_instance_id bigint,
	vpc_security_group_id text,
	ref_vpc_security_group_id bigint,
	status text
) ;
-- ALTER TABLE public.db_instance_vpc_security_groups add column ref_vpc_security_group_id bigint ;

--DBInstancesDBSecurityGroups-------------------
CREATE TABLE public.db_instance_db_security_groups (
	id bigserial,
	db_instance_id bigint,
	db_security_group_name text,
	ref_security_group_id bigint,
	status text
) ;

--DBInstancesDBSubnetGroup-------------------
CREATE TABLE public.db_instance_db_subnet_group (
	id bigserial,
	db_instance_id bigint,
	db_subnet_group_name text,
	db_subnet_group_description text,
	vpc_id text,
	ref_vpc_id bigint,
	subnet_group_status text,
	db_subnet_group_arn text
) ;

--DBInstancesDBSubnetGroupSubnet-------------------
CREATE TABLE public.db_instance_db_subnet_group_subnet (
	id bigserial,
	db_instance_db_subnet_group_id bigint,
	subnet_identifier text,
	ref_subnet_id bigint,
	name text,
	subnet_status text
) ;

--Reservations-------------------
CREATE TABLE public.reservation (
	id bigserial,
	owner_id text,
	requester_id text,
	reservation_id text
) ;

--ReservationsGroup-------------------
CREATE TABLE public.reservation_group (
	id bigserial,
	reservation_id bigint,
	group_name text,
	group_id text
) ;

--ReservationsInstance-------------------
CREATE TABLE public.instance (
	id bigserial,
	reservation_id bigint,
	ami_launch_index int,
	image_id text,
	ref_image_id bigint,
	instance_id text,
	instance_type text,
	kernel_id text,
	key_name text,
	launch_time text,
	monitoring_state text,
	placement_availability_zone text,
	placement_affinity text,
	placement_group_name text,
	placement_partition_number text,
	placement_host_id text,
	placement_tenancy text,
	placement_spread_domain text,
	platform text,
	private_dns_name text,
	private_ip_address text,
	public_dns_name text,
	public_ip_address text,
	ramdisk_id text,
	state_code int,
	state_name text,
	state_transition_reason text,
	subnet_id text,
	ref_subnet_id bigint,
	vpc_id text,
	ref_vpc_id bigint,
	architecture text,
	client_token text,
	ebs_optimized boolean,
	ena_support boolean,
	hypervisor text,
	iam_instance_profile_id text ,
	iam_instance_profile_arn text,
	instance_lifecycle text,
	root_device_name text,
	root_device_type text,
	source_dest_check boolean,
	spot_instance_request_id text,
	sriov_net_support text,
	state_reason_code text,
	state_reason_message text,
	virtualization_type text,
	cpu_option_core_count int,
	cpu_option_threads_per_core int,
	capacity_reservation_id text,
	capacity_reservation_preference text,
	capacity_reservation_target_id text,
	hibernation_options_configured boolean
) ;
--InstanceProductCode-------------------
CREATE TABLE public.instance_product_code (
	id bigserial,
	instance_id bigint,
	product_code_id text,
	product_code_type text
) ;

--InstanceBlockDeviceMapping-------------------
CREATE TABLE public.instance_block_device_mapping (
	id bigserial,
	instance_id bigint,
	device_name text,
	ebs_attach_time text,
	ebs_delete_on_termination boolean,
	ebs_status text,
	ebs_volume_id text,
	ref_volume_id bigint	
) ;

--InstanceElasticGPUAssociation-------------------
CREATE TABLE public.instance_elastic_gpu_association (
	id bigserial,
	instance_id bigint,
	elastic_gpu_id text,
	elastic_gpu_association_id text,
	elastic_gpu_association_state text,
	elastic_gpu_association_time text
) ;

--InstanceElasticInferenceAcceleratorAssociation-------------------
CREATE TABLE public.instance_elastic_inference_accelerator_association (
	id bigserial,
	instance_id bigint,
	elasticInference_accelerator_arn text,
	elasticInference_accelerator_association_id text,
	elasticInference_accelerator_association_state text,
	elasticInference_accelerator_association_time text
) ;

--InstanceNetorkInterface-------------------
CREATE TABLE public.instance_network_interface (
	id bigserial,
	instance_id bigint,
	association_ip_owner_id text,
	association_public_dns_name text,
	association_public_ip text,
	attachment_attach_time text,
	attachment_attachment_id text,
	attachment_delete_on_termination boolean,
	attachment_device_index int,
	attachment_status text,
	description text,
	mac_address text,
	network_interface_id text,
	owner_id text,
	private_dns_name text,
	private_ip_address text,
	source_dest_check boolean,
	status text,
	subnet_id text,
	ref_subnet_id bigint,
	vpc_id text,
	ref_vpc_id bigint,
	interface_type text
) ;

--InstanceNetworkInterfaceGroup-------------------
CREATE TABLE public.instance_network_interface_group (
	id bigserial,
	instance_network_interface_id bigint,
	group_name text,
	group_id text,
	ref_security_group_id bigint
) ;

--InstanceNetworkInterfacePrivateIpAddress-------------------
CREATE TABLE public.instance_network_interface_private_ip_address (
	id bigserial,
	instance_network_interface_id bigint,
	association_ip_owner_id text,
	association_public_dns_name text,
	association_public_ip text,
	is_primary boolean,
	private_dns_name text,
	private_ip_address text
) ;

--InstanceSecurityGroup-------------------
CREATE TABLE public.instance_security_group (
	id bigserial,
	instance_id bigint,
	group_name text,
	group_id text,
	ref_security_group_id bigint
) ;

--InstanceTag-------------------
CREATE TABLE public.instance_tag (
	id bigserial,
	instance_id bigint,
	key text,
	value text
) ;

--InstanceLicense-------------------
CREATE TABLE public.instance_license (
	id bigserial,
	instance_id bigint,
	license_configuration_arn text
) ;

--LoadBalancers-------------------
CREATE TABLE public.load_balancer (
	id bigserial,
	load_balancer_arn text,
	dns_name text,
	canonical_hosted_zone_id text,
	created_time text,
	load_balancer_name text,
	scheme text,
	vpc_id text,
	ref_vpc_id bigint,
	state_code text,
	state_reason text,
	type text,
	ip_address_type text
) ;

--LoadBalancersAvailablityZone-------------------
CREATE TABLE public.load_balancer_availability_zone (
	id bigserial,
	load_balancer_id bigint,
	zone_name text,
	subnet_id text,
	ref_subnet_id bigint
) ;

CREATE TABLE public.load_balancer_availability_zone_load_balancer_address (
	id bigserial,
	load_balancer_availability_zone_id bigint,
	ip_address text,
	allocation_id text
) ;
--LoadBalancersSecurityGroup-------------------
CREATE TABLE public.load_balancer_security_group (
	id bigserial,
	load_balancer_id bigint,
	security_group_id text,
	ref_security_group_id bigint
) ;

--command-load-balancers-listeners.sh
CREATE TABLE public.load_balancer_listener (
	id bigserial,
	load_balancer_id bigint,
	load_balancer_arn text,
	listener_arn text,
	ssl_policies text,
	port bigint,
	protocol text
) ;


CREATE TABLE public.load_balancer_listener_default_actions (
	id bigserial,
	load_balancer_listener_id bigint,
	target_group_arn text,
	target_group_id bigint,
	type text
) ;

CREATE TABLE public.load_balancer_listener_certificats (
	id bigserial,
	load_balancer_listener_id bigint,
	certificat_arn text
) ;

--command-load-balancers-target-health.sh
-- CREATE TABLE public.load_balancer_target_health (
-- 	id bigserial,
-- 	load_balancer_id bigint
-- ) ;
-- 
-- CREATE TABLE public.load_balancer_target_health_descr (
-- 	id bigserial,
-- 	load_balancer_target_health_id bigint,
-- 	health_check_port text,
-- 	target_id  text,
-- 	target_instance_id bigint,
-- 	target_port bigint,
-- 	target_health_state text
-- ) ;

--AutoScalingGroups-------------------
CREATE TABLE public.auto_scaling_group (
	id bigserial,
	auto_scaling_group_name text,
	auto_scaling_group_arn text,
	launch_configuration_name text,
	launch_template_id text,
	launch_template_name text,
	launch_template_version text,
	mx_inst_pol_launch_template_specification_id text,
	mx_inst_pol_launch_template_specification_name text,
	mx_inst_pol_launch_template_specification_version text,
	mx_inst_pol_inst_dist_on_demand_allocation_strategy text,
	mx_inst_pol_inst_dist_on_demand_base_capacity int,
	mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity int,
	mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy text,
	mx_inst_pol_inst_dist_on_demand_spot_instance_pools int,
	mx_inst_pol_inst_dist_on_demand_spot_max_price text,
	minSize int,
	maxSize int,
	desired_capacity int,
	default_cooldown int,
	health_check_type text,
	health_check_grace_period int,
	created_time text,
	placement_group text,
	vpc_zone_identifier text,
	ref_subnet_id bigint,
	status text,
	new_instances_protected_from_scale_in boolean,
	service_linked_role_arn text
) ;

--AutoScalingGroupsLauchTemplateOverride-------------------
CREATE TABLE public.auto_scaling_group_launch_template_override (
	id bigserial,
	auto_scaling_group_id bigint,
	instance_type text
) ;

--AutoScalingGroupsAvailabilityZone-------------------
CREATE TABLE public.auto_scaling_group_availability_zone (
	id bigserial,
	auto_scaling_group_id bigint,
	availability_zone text
) ;

--AutoScalingGroupsLoadBalancerName-------------------
CREATE TABLE public.auto_scaling_group_load_balancer_name (
	id bigserial,
	auto_scaling_group_id bigint,
	load_balancer_name text
) ;

--AutoScalingGroupsTargetGroupARN-------------------
CREATE TABLE public.auto_scaling_group_target_group_arn(
	id bigserial,
	auto_scaling_group_id bigint,
	arn_name text
) ;

--AutoScalingInstance-------------------
CREATE TABLE public.auto_scaling_group_instance (
	id bigserial,
	auto_scaling_group_id bigint,
	instance_id text,
	ref_instance_id bigint,
	availability_zone text,
	lifecycle_state text,
	health_status text,
	launch_configuration_name text,
	launch_template_id text,
	launch_template_name text,
	launch_template_version text,
	protected_from_scale_in boolean
) ;

--AutoScalingGroupsSuspendedProcesses-------------------
CREATE TABLE public.auto_scaling_group_suspended_process (
	id bigserial,
	auto_scaling_group_id bigint,
	process_name text,
	suspension_reason text
) ;

--AutoScalingGroupsEnabledMetric-------------------
CREATE TABLE public.auto_scaling_group_enabled_metric (
	id bigserial,
	auto_scaling_group_id bigint,
	metric text,
	granularity text
) ;

--AutoScalingGroupsTag-------------------
CREATE TABLE public.auto_scaling_group_tag (
	id bigserial,
	auto_scaling_group_id bigint,
	resource_id text,
	resource_type text,
	key text,
	value text,
	propagate_at_launch boolean
) ;

--AutoScalingGroupsTerminationPolicy-------------------
CREATE TABLE public.auto_scaling_group_termination_policy (
	id bigserial,
	auto_scaling_group_id bigint,
	termination_policy text
) ;

--AutoScalingActivities-------------------
CREATE TABLE public.auto_scaling_activities (
	id bigserial,
	activity_id text,
	description text,
	auto_scaling_group_name text,
	ref_auto_scaling_group_id bigint,
	details text,
	start_time text,
	progress int,
	end_time text,
	cause text,
	status_code text
) ;


--LoadBalancerTargetGroup-------------------
CREATE TABLE public.load_balancer_target_group (
	id bigserial,
	target_group_arn text,
	target_group_name text,
	protocol text,
	port int,
	vpc_id text,
	ref_vpc_id bigint,
	health_check_protocol text,
	health_check_port text,
	health_check_enabled boolean,
	health_check_interval_seconds int,
	health_check_timeout_seconds int,
	healthy_threshold_count int,
	unhealthy_threshold_count int,
	health_check_path text,
	http_code text,
	target_type text

) ;

--LoadBalancerTargetGroupLoadBalancerArn-------------------
CREATE TABLE public.load_balancer_target_group_load_balancer_arn (
	id bigserial,
	load_balancer_target_group_id bigint,
	load_balancer_arn text,
	ref_load_balancer_id bigint
) ;

--LoadBalancerTargetGroupHealth-------------------
CREATE TABLE public.load_balancer_target_group_health (
	id bigserial,
	load_balancer_target_group_id bigint,
	health_check_port text,
	target_id text,
	ref_instance_id int,
	port int,
	target_health_state text,
	target_health_reason text,
	target_health_description text
) ;

--Snapshot-------------------
CREATE TABLE public.snapshot (
	id bigserial,
	snapshot_id text,
	owner_id text,
	volume_id text,
	ref_volume_id int,
	encrypted boolean,
	description text,
	state text,
	volume_size int,
	start_time text,
	progress text
) ;







