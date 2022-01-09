create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_s1mme_mme_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_s1mme_mme_cell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`eci` bigint,
`mme_ip` varchar(200),
`mme_name` varchar(255)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_s1mme_mme_cell_d/";