create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_pm_4gc_sgw_mme_d;
create EXTERNAL table IF NOT EXISTS dm_ct_pm_4gc_sgw_mme_d(
`starttime` timestamp,
`equip_name` varchar(64),
`node_name` varchar(64),
`equip_type` varchar(64),
`os_type` varchar(255),
`equip_attribute` varchar(64),
`equip_ip` varchar(64),
`disk_use_rat` decimal(12,6),
`cpu_use_rat` decimal(12,6),
`memory_use_rat` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_pm_4gc_sgw_mme_d/";