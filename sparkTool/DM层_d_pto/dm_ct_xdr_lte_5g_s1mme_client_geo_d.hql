create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1mme_client_geo_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1mme_client_geo_d(
`starttime` timestamp,
`geo_id` int,
`ne_type` int,
`province_id` int,
`province_name` varchar(128),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`client_id` bigint,
`client_name` varchar(255),
`eps_attach_succ_rate` decimal(12,6),
`ue_pdnconn_succ_rate` decimal(12,6),
`tau_succ_rate` decimal(12,6),
`s1_ho_suc_rate` decimal(12,6),
`x2_ho_suc_rate` decimal(12,6),
`initial_cont_avg_dealy` decimal(24,4)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1mme_client_geo_d/";