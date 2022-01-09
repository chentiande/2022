create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_cell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_cell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`nci_eci` bigint,
`region_id` int,
`region_name` varchar(128),
`sgwcell_mme_abnormal_flag` int,
`spipcell_abnormal_flag` int,
`ipranacell_abnormal_flag` int,
`dnscell_abnormal_flag` int,
`celkqicell_abnormal_flag` int,
`termcell_abnormal_flag` int,
`cell_total_record_num` bigint,
`cell_abnormal_record_num` bigint,
`cell_abnormal_record_rate` decimal(12,6),
`cell_ul_rtt_delay_abnormal_num` bigint,
`cell_ul_rtt_delay_abnormal_rate` decimal(12,6),
`cell_dl_rtt_delay_abnormal_num` bigint,
`cell_dl_rtt_delay_abnormal_rate` decimal(12,6),
`cell_webpage_delay_abnormal_num` bigint,
`cell_webpage_delay_abnormal_rate` decimal(12,6),
`cell_video_download_abnormal_num` bigint,
`cell_video_download_abnormal_rate` decimal(12,6),
`cell_game_chkhand_abnormal_num` bigint,
`cell_game_chkhand_abnormal_rate` decimal(12,6),
`cell_reply_delay_abnormal_num` bigint,
`cell_reply_delay_abnormal_rate` decimal(12,6),
`delimitation_conclusion` varchar(32),
`treatment_suggestion` varchar(32)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_delimit_cell_d/";