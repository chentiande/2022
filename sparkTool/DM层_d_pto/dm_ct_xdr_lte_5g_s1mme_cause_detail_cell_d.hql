create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1mme_cause_detail_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1mme_cause_detail_cell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`tac` int,
`eci` bigint,
`cell_id` int,
`cell_name` varchar(255),
`enb_id` int,
`enb_name` varchar(128),
`longitude` decimal(12,6),
`latitude` decimal(12,6),
`cause_value` int,
`cause_name` VARCHAR(255),
`cause_may` varchar(50),
`s1_mme_cnt` int) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1mme_cause_detail_cell_d/";