create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_http_geo_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_http_geo_d(
`starttime` timestamp,
`geo_id` int,
`ne_type` int,
`province_id` int,
`province_name` varchar(128),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`tcp_build_success_rat` decimal(12,6),
`http_service_success_rat` decimal(12,6),
`http_avg_respdelay` decimal(24,4),
`tcp_connect_avg_delay` decimal(24,4),
`tcp_renum_rat` decimal(12,6),
`tcp_discordnum_rat` decimal(12,6),
`tcp_estab_excellent_delay_rate` decimal(12,6),
`service_good_rat` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_http_geo_d/";