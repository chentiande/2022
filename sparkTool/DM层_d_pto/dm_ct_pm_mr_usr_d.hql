create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_pm_mr_usr_d;
create EXTERNAL table IF NOT EXISTS dm_ct_pm_mr_usr_d(
`starttime` timestamp,
`msisdn` varchar(32),
`rsrp_avg` decimal(24,4),
`rsrp_tolal_samples_cnt` int,
`rsrp_110_samples_cnt` int) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_pm_mr_usr_d/";