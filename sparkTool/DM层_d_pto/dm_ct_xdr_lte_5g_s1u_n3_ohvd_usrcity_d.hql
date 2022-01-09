create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_usrcity_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_usrcity_d(
`starttime` timestamp,
`imsi_supi` varchar(32),
`msisdn` varchar(32),
`province_id` int,
`province_name` varchar(128),
`area_id` int,
`area_name` varchar(255),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`duration` bigint,
`n3_cnt` bigint,
`total_data` bigint,
`f_packets_response_good_rat` decimal(12,6),
`http_dns_res_good_rat` decimal(12,6),
`page_open_good_rat` decimal(12,6),
`page_screen_good_rat` decimal(12,6),
`im_good_rat` decimal(12,6),
`game_good_rat` decimal(12,6),
`vide_speed_good_rat` decimal(12,6),
`vide_halt_good_rat` decimal(12,6),
`bit_code_rate_good_rat` decimal(12,6),
`http_service_good_rat` decimal(12,6),
`vide_service_good_rat` decimal(12,6),
`rtt_ul_server_good_rat` decimal(12,6),
`rtt_dl_termeral_good_rat` decimal(12,6),
`rtt_service_good_rat` decimal(12,6),
`ul_upload_good_rat` decimal(12,6),
`dl_download_good_rat` decimal(12,6),
`tran_service_good_rate` decimal(12,6),
`service_perceive_good_rat` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_usrcity_d/";