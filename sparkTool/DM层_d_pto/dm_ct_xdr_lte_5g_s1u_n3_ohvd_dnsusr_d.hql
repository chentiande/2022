create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_dnsusr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_dnsusr_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`msisdn` varchar(32),
`query_domain_name` varchar(255),
`dns_total_record_abnormal_top_rate` decimal(12,6),
`dns_total_duration` bigint,
`dns_total_data` bigint,
`dns_total_record_num` bigint,
`dns_abnormal_record_rate` decimal(12,6),
`dns_abnormal_flag` int,
`dns_msisdn_num` bigint,
`dns_abnormal_msisdn_num` bigint,
`dns_resp_fail_abnormal_rate` decimal(12,6),
`dns_reply_delay_abnormal_rate` decimal(12,6),
`dnsusr_total_duration` bigint,
`dnsusr_total_data` bigint,
`dnsusr_total_record_num` bigint,
`dnsusr_abnormal_record_rate` decimal(12,6),
`dnsusr_resp_fail_abnormal_rate` decimal(12,6),
`dnsusr_reply_delay_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_dnsusr_d/";