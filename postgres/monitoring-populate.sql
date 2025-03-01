-- #############################################################################
--
--    
-- #############################################################################

-- -------
-- PROD
-- -------
insert into config_subnet(name,description,env,cidr_block) values('SubnetSlash16','SubnetProd /16','PROD','10.3.0.0/16') ;

-- DB
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDS','SubnetRDS /24','PROD','10.3.40.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDSZoneA','SubnetRDSZoneA /25','PROD','10.3.40.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDSZoneB','SubnetRDSZoneB /25','PROD','10.3.40.128/25') ;

-- Backends
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackend'     ,'SubnetBackend /24','PROD','10.3.30.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendZoneA','SubnetBackendZoneA /25','PROD','10.3.30.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendZoneB','SubnetBackendZoneB /25','PROD','10.3.30.128/25') ;

insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPriv'     ,'SubnetBackendPriv /24','PROD','10.3.50.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPrivZoneA','SubnetBackendPrivZoneA /25','PROD','10.3.50.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPrivZoneB','SubnetBackendPrivZoneB /25','PROD','10.3.50.128/25') ;

-- Frontends
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontend'     ,'SubnetFrontend /24','PROD','10.3.10.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontendZoneA','SubnetFrontendZoneA /25','PROD','10.3.10.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontendZoneB','SubnetFrontendZoneB /25','PROD','10.3.10.128/25') ;

-- CodeBuildPrivate
insert into config_subnet(name,description,env,cidr_block) values('SubnetCodeBuildPriv','SubnetCodeBuildPriv /24','PROD','10.3.20.0/24') ;

-- GoPhish / Traefik
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefik'     ,'SubnetPhishingTraefik /24','PROD','10.3.100.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefikZoneA','SubnetPhishingTraefikZoneA /25','PROD','10.3.100.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefikZoneB','SubnetPhishingTraefikZoneB /25','PROD','10.3.100.128/25') ;

-- GoPhish / GoPublic
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublic'     ,'SubnetPhishingGoPhishPublic /24','PROD','10.3.110.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublicZoneA','SubnetPhishingGoPhishPublicZoneA /25','PROD','10.3.110.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublicZoneB','SubnetPhishingGoPhishPublicZoneB /25','PROD','10.3.110.128/25') ;

-- GoPhish / GoPrivate
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivate'     ,'SubnetPhishingGoPhishPrivate /24','PROD','10.3.120.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivateZoneA','SubnetPhishingGoPhishPrivateZoneA /25','PROD','10.3.120.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivateZoneB','SubnetPhishingGoPhishPrivateZoneB /25','PROD','10.3.120.128/25') ;

-- -------
-- STAGING
-- -------
insert into config_subnet(name,description,env,cidr_block) values('SubnetSlash16','SubnetStaging /16','STAGING','10.4.0.0/16') ;

-- DB
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDS','SubnetRDS /24','STAGING','10.4.40.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDSZoneA','SubnetRDSZoneA /25','STAGING','10.4.40.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetRDSZoneB','SubnetRDSZoneB /25','STAGING','10.4.40.128/25') ;

-- Backends
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackend'     ,'SubnetBackend /24','STAGING','10.4.30.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendZoneA','SubnetBackendZoneA /25','STAGING','10.4.30.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendZoneB','SubnetBackendZoneB /25','STAGING','10.4.30.128/25') ;

insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPriv'     ,'SubnetBackendPriv /24','STAGING','10.4.50.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPrivZoneA','SubnetBackendPrivZoneA /25','STAGING','10.4.50.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetBackendPrivZoneB','SubnetBackendPrivZoneB /25','STAGING','10.4.50.128/25') ;

-- Frontends
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontend'     ,'SubnetFrontend /24','STAGING','10.4.10.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontendZoneA','SubnetFrontendZoneA /25','STAGING','10.4.10.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetFrontendZoneB','SubnetFrontendZoneB /25','STAGING','10.4.10.128/25') ;

-- CodeBuildPrivate
insert into config_subnet(name,description,env,cidr_block) values('SubnetCodeBuildPriv','SubnetCodeBuildPriv /24','STAGING','10.4.20.0/24') ;

-- GoPhish / Traefik
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefik'     ,'SubnetPhishingTraefik /24','STAGING','10.4.100.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefikZoneA','SubnetPhishingTraefikZoneA /25','STAGING','10.4.100.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingTraefikZoneB','SubnetPhishingTraefikZoneB /25','STAGING','10.4.100.128/25') ;

-- GoPhish / GoPublic
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublic'     ,'SubnetPhishingGoPhishPublic /24','STAGING','10.4.110.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublicZoneA','SubnetPhishingGoPhishPublicZoneA /25','STAGING','10.4.110.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPublicZoneB','SubnetPhishingGoPhishPublicZoneB /25','STAGING','10.4.110.128/25') ;

-- GoPhish / GoPrivate
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivate'     ,'SubnetPhishingGoPhishPrivate /24','STAGING','10.4.120.0/24') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivateZoneA','SubnetPhishingGoPhishPrivateZoneA /25','STAGING','10.4.120.0/25') ;
insert into config_subnet(name,description,env,cidr_block) values('SubnetPhishingGoPhishPrivateZoneB','SubnetPhishingGoPhishPrivateZoneB /25','STAGING','10.4.120.128/25') ;
