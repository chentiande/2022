create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_alarm_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_alarm_d(
`alarm_id` varchar(128),
`alarm_nms` varchar(64),
`alarm_messageid` varchar(128),
`province_id` int,
`province_name` varchar(128),
`area_id` int,
`area_name` varchar(255),
`region_id` int,
`region_name` varchar(128),
`alarm_unit` varchar(64),
`alarm_node_name` varchar(255),
`alarm_type` varchar(16),
`alarm_severity` int,
`alarm_time` bigint,
`alarm_check_flag` int,
`alarm_text` varchar(1000),
`alarm_status` int,
`alarm_exist_day_cnt` int,
`sheet_status` int) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_alarm_d/";