create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_sgwusr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_sgwusr_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`msisdn` varchar(32),
`sgw_upf_ip` varchar(200),
`sgw_upf_name` varchar(255),
`sgw_total_record_abnormal_top_rate` decimal(12,6),
`sgw_total_duration` bigint,
`sgw_total_data` bigint,
`sgw_total_record_num` bigint,
`sgw_abnormal_record_rate` decimal(12,6),
`sgw_abnormal_flag` int,
`sgw_msisdn_num` bigint,
`sgw_abnormal_msisdn_num` bigint,
`sgw_ul_rtt_delay_abnormal_rate` decimal(12,6),
`sgw_dl_rtt_delay_abnormal_rate` decimal(12,6),
`sgw_webpage_delay_abnormal_rate` decimal(12,6),
`sgw_video_download_abnormal_rate` decimal(12,6),
`sgw_game_chkhand_abnormal_rate` decimal(12,6),
`sgw_dns_resp_abnormal_rate` decimal(12,6),
`sgwusr_total_duration` bigint,
`sgwusr_total_data` bigint,
`sgwusr_total_record_num` bigint,
`sgwusr_abnormal_record_rate` decimal(12,6),
`sgwusr_ul_rtt_delay_abnormal_rate` decimal(12,6),
`sgwusr_dl_rtt_delay_abnormal_rate` decimal(12,6),
`sgwusr_webpage_delay_abnormal_rate` decimal(12,6),
`sgwusr_video_download_abnormal_rate` decimal(12,6),
`sgwusr_game_chkhand_abnormal_rate` decimal(12,6),
`sgwusr_dns_resp_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_sgwusr_d/";