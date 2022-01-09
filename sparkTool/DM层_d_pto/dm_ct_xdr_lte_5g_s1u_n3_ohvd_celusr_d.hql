create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_celusr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_celusr_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`msisdn` varchar(32),
`region_id` int,
`region_name` varchar(128),
`nci_eci` bigint,
`cell_name` varchar(255),
`gnb_enb_id` int,
`gnb_enb_name` varchar(128),
`celkqi_total_record_abnormal_top_rate` decimal(12,6),
`celkqi_total_duration` bigint,
`celkqi_total_data` bigint,
`celkqi_total_record_num` bigint,
`celkqi_abnormal_record_rate` decimal(12,6),
`celkqi_abnormal_flag` int,
`celkqi_msisdn_num` bigint,
`celkqi_abnormal_msisdn_num` bigint,
`celkqi_ul_rtt_delay_abnormal_rate` decimal(12,6),
`celkqi_dl_rtt_delay_abnormal_rate` decimal(12,6),
`celkqi_webpage_delay_abnormal_rate` decimal(12,6),
`celkqi_video_download_abnormal_rate` decimal(12,6),
`celkqi_game_chkhand_abnormal_rate` decimal(12,6),
`celkqi_dns_resp_abnormal_rate` decimal(12,6),
`celkqiusr_total_duration` bigint,
`celkqiusr_total_data` bigint,
`celkqiusr_total_record_num` bigint,
`celkqiusr_abnormal_record_rate` decimal(12,6),
`celkqiusr_ul_rtt_delay_abnormal_rate` decimal(12,6),
`celkqiusr_dl_rtt_delay_abnormal_rate` decimal(12,6),
`celkqiusr_webpage_delay_abnormal_rate` decimal(12,6),
`celkqiusr_video_download_abnormal_rate` decimal(12,6),
`celkqiusr_game_chkhand_abnormal_rate` decimal(12,6),
`celkqiusr_dns_resp_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_celusr_d/";