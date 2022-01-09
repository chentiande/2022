create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_ipranusr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_ipranusr_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`msisdn` varchar(32),
`region_id` int,
`region_name` varchar(128),
`iprana_eqp_name` varchar(255),
`iprana_eqp_tot_record_ab_top_rate` decimal(12,6),
`iprana_total_duration` bigint,
`iprana_total_data` bigint,
`iprana_total_record_num` bigint,
`iprana_abnormal_record_rate` decimal(12,6),
`iprana_abnormal_flag` int,
`iprana_msisdn_num` bigint,
`iprana_abnormal_msisdn_num` bigint,
`iprana_ul_rtt_delay_abnormal_rate` decimal(12,6),
`iprana_dl_rtt_delay_abnormal_rate` decimal(12,6),
`iprana_webpage_delay_abnormal_rate` decimal(12,6),
`iprana_video_download_abnormal_rate` decimal(12,6),
`iprana_game_chkhand_abnormal_rate` decimal(12,6),
`iprana_dns_resp_abnormal_rate` decimal(12,6),
`ipranausr_total_duration` bigint,
`ipranausr_total_data` bigint,
`ipranausr_total_record_num` bigint,
`ipranausr_abnormal_record_rate` decimal(12,6),
`ipranausr_ul_rtt_delay_abnormal_rate` decimal(12,6),
`ipranausr_dl_rtt_delay_abnormal_rate` decimal(12,6),
`ipranausr_webpage_delay_abnormal_rate` decimal(12,6),
`ipranausr_video_download_abnormal_rate` decimal(12,6),
`ipranausr_game_chkhand_abnormal_rate` decimal(12,6),
`ipranausr_dns_resp_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_ipranusr_d/";