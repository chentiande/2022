create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_pm_ipran_iprana_d;
create EXTERNAL table IF NOT EXISTS dm_ct_pm_ipran_iprana_d(
`starttime` timestamp,
`a_eqp_name` varchar(255),
`cpu_use_rat` decimal(12,6),
`memory_use_rat` decimal(12,6),
`peak_bandwidth_rat` decimal(12,6),
`data_packloss_rat` decimal(12,6),
`service_delay` decimal(24,4)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_pm_ipran_iprana_d/";