create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_usr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_usr_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`msisdn` varchar(32),
`term_vendor_name` varchar(64),
`term_model_name` varchar(100),
`term_i_os` varchar(128),
`term_type` varchar(64),
`sgwusr_mme_abnormal_flag` int,
`spipusr_abnormal_flag` int,
`ipranausr_abnormal_flag` int,
`dnsusr_abnormal_flag` int,
`celkqiusr_abnormal_flag` int,
`termusr_abnormal_flag` int,
`usr_total_record_num` bigint,
`usr_abnormal_record_num` bigint,
`usr_abnormal_record_rate` decimal(12,6),
`usr_ul_rtt_delay_abnormal_num` bigint,
`usr_ul_rtt_delay_abnormal_rate` decimal(12,6),
`usr_dl_rtt_delay_abnormal_num` bigint,
`usr_dl_rtt_delay_abnormal_rate` decimal(12,6),
`usr_webpage_delay_abnormal_num` bigint,
`usr_webpage_delay_abnormal_rate` decimal(12,6),
`usr_video_download_abnormal_num` bigint,
`usr_video_download_abnormal_rate` decimal(12,6),
`usr_game_chkhand_abnormal_num` bigint,
`usr_game_chkhand_abnormal_rate` decimal(12,6),
`usr_reply_delay_abnormal_num` bigint,
`usr_reply_delay_abnormal_rate` decimal(12,6),
`term_good_rate` decimal(12,6),
`termusr_good_rate` decimal(12,6),
`termusr_beat_rate` decimal(12,6),
`term_change_flag` int,
`term_change_time` timestamp,
`delimitation_conclusion` varchar(32),
`treatment_suggestion` varchar(32)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_usr_d/";