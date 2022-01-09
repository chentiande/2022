create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_fm_alarm_his;
create EXTERNAL table IF NOT EXISTS dm_ct_fm_alarm_his(
`org_alarm_id` varchar(64),
`alarm_id` varchar(64),
`event_time` timestamp,
`clear_time` timestamp,
`severity` varchar(32),
`alarm_type` varchar(32),
`clear_flag` varchar(32),
`object_id` bigint,
`object_name` varchar(255),
`object_type` varchar(64),
`object_ip` varchar(64),
`locate_object_id` bigint,
`locate_object_name` varchar(255),
`locate_object_type` varchar(64),
`locate_info` varchar(255),
`vendor_name` varchar(64),
`alarm_title` varchar(255),
`probable_cause` varchar(255),
`maintain_propose` varchar(255),
`domain` varchar(32),
`total_count` int,
`enodeb_id` varchar(32),
`eci` varchar(32),
`resource_status` varchar(32),
`province_id` varchar(32),
`province_name` varchar(128),
`region_id` varchar(32),
`region_name` varchar(128)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_fm_alarm_his/";