create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_pm_lte_kpi_geo_d;
create EXTERNAL table IF NOT EXISTS dm_ct_pm_lte_kpi_geo_d(
`starttime` timestamp,
`geo_id` int,
`ne_type` int,
`province_id` int,
`province_name` varchar(128),
`region_id` int,
`region_name` varchar(128),
`city_id` int,
`city_name` varchar(128),
`puschprbtotmeandl_rate` decimal(12,6),
`puschprbtotmeanul_rate` decimal(12,6),
`erab_drop_rate` decimal(12,6),
`rsrp_avg` decimal(24,4)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_pm_lte_kpi_geo_d/";