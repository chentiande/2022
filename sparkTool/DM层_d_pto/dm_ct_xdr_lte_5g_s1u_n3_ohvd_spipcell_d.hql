create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_spipcell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_spipcell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`nci_eci` bigint,
`server_ip` varchar(255),
`app_id` int,
`app_name` varchar(128),
`sub_app_id` int,
`sub_app_name` varchar(128),
`spip_total_record_abnormal_top_rate` decimal(12,6),
`spip_total_duration` bigint,
`spip_total_data` bigint,
`spip_total_record_num` bigint,
`spip_abnormal_record_rate` decimal(12,6),
`spip_abnormal_flag` int,
`spip_msisdn_num` bigint,
`spip_abnormal_msisdn_num` bigint,
`spip_ul_rtt_delay_abnormal_rate` decimal(12,6),
`spip_dl_rtt_delay_abnormal_rate` decimal(12,6),
`spip_webpage_delay_abnormal_rate` decimal(12,6),
`spip_video_download_abnormal_rate` decimal(12,6),
`spip_game_chkhand_abnormal_rate` decimal(12,6),
`spip_dns_resp_abnormal_rate` decimal(12,6),
`spipcell_total_duration` bigint,
`spipcell_total_data` bigint,
`spipcell_total_record_num` bigint,
`spipcell_abnormal_record_rate` decimal(12,6),
`spipcell_ul_rtt_delay_abnormal_rate` decimal(12,6),
`spipcell_dl_rtt_delay_abnormal_rate` decimal(12,6),
`spipcell_webpage_delay_abnormal_rate` decimal(12,6),
`spipcell_video_download_abnormal_rate` decimal(12,6),
`spipcell_game_chkhand_abnormal_rate` decimal(12,6),
`spipcell_dns_resp_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_spipcell_d/";