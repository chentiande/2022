create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_termcell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_termcell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`nci_eci` bigint,
`term_tac` int,
`term_vendor_name` varchar(64),
`term_model_name` varchar(100),
`term_total_record_abnormal_top_rate` decimal(12,6),
`term_total_duration` bigint,
`term_total_data` bigint,
`term_total_record_num` bigint,
`term_abnormal_record_rate` decimal(12,6),
`term_abnormal_flag` int,
`term_msisdn_num` bigint,
`term_abnormal_msisdn_num` bigint,
`term_ul_rtt_delay_abnormal_rate` decimal(12,6),
`term_dl_rtt_delay_abnormal_rate` decimal(12,6),
`term_webpage_delay_abnormal_rate` decimal(12,6),
`term_video_download_abnormal_rate` decimal(12,6),
`term_game_chkhand_abnormal_rate` decimal(12,6),
`term_dns_resp_abnormal_rate` decimal(12,6),
`termcell_total_duration` bigint,
`termcell_total_data` bigint,
`termcell_total_record_num` bigint,
`termcell_abnormal_record_rate` decimal(12,6),
`termcell_ul_rtt_delay_abnormal_rate` decimal(12,6),
`termcell_dl_rtt_delay_abnormal_rate` decimal(12,6),
`termcell_webpage_delay_abnormal_rate` decimal(12,6),
`termcell_video_download_abnormal_rate` decimal(12,6),
`termcell_game_chkhand_abnormal_rate` decimal(12,6),
`termcell_dns_resp_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_termcell_d/";