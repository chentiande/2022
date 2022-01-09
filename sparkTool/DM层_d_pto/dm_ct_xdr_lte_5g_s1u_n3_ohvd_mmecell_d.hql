create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_mmecell_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_mmecell_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`nci_eci` bigint,
`mme_ip` varchar(200),
`mme_name` varchar(255),
`mme_total_record_abnormal_top_rate` decimal(12,6),
`mme_total_record_num` bigint,
`mme_abnormal_record_rate` decimal(12,6),
`mme_abnormal_flag` int,
`mme_msisdn_num` bigint,
`mme_abnormal_msisdn_num` bigint,
`mme_service_registration_abnormal_rate` decimal(12,6),
`mme_service_reques_abnormal_rate` decimal(12,6),
`mme_pdnconn_abnormal_rate` decimal(12,6),
`mme_s1_ho_out_abnormal_rate` decimal(12,6),
`mmecell_total_record_num` bigint,
`mmecell_abnormal_record_rate` decimal(12,6),
`mmecell_service_registration_abnormal_rate` decimal(12,6),
`mmecell_service_reques_abnormal_rate` decimal(12,6),
`mmecell_pdnconn_abnormal_rate` decimal(12,6),
`mmecell_s1_ho_out_abnormal_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_mmecell_d/";