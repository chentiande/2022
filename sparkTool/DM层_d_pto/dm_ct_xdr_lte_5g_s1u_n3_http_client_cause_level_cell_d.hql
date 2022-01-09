create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_http_client_cause_level_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_http_client_cause_level_cell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`tac` int,
`nci_eci` bigint,
`nc_ec_id` int,
`cell_name` varchar(255),
`gnb_enb_id` int,
`gnb_enb_name` varchar(128),
`longitude` decimal(12,6),
`latitude` decimal(12,6),
`client_id` bigint,
`client_name` varchar(255),
`cause_level` varchar(50),
`n3_http_cnt` bigint) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_http_client_cause_level_cell_d/";