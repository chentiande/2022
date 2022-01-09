create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_http_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_http_cell_d(
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
`tcp_build_success_rat` decimal(12,6),
`http_service_success_rat` decimal(12,6),
`http_avg_respdelay` decimal(24,4),
`tcp_connect_avg_delay` decimal(24,4),
`tcp_renum_rat` decimal(12,6),
`tcp_discordnum_rat` decimal(12,6),
`tcp_estab_excellent_delay_rate` decimal(12,6),
`service_good_rat` decimal(12,6),
`rtt_ul_server_bat` bigint,
`rtt_dl_termeral_bat` bigint) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_http_cell_d/";