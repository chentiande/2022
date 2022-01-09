create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_http_sgw_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_http_sgw_cell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`nci_eci` bigint,
`sgw_upf_ip` varchar(200),
`sgw_upf_name` varchar(255)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_http_sgw_cell_d/";