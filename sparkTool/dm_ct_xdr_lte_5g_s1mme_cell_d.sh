#!/bin/bash

script_home=$(dirname $0)
source ${script_home}/../config.sh

###for test
deploy_mode=client
###for product
#deploy_mode=cluster

##input time
#part_time=202010081200
part_time=$1
echo "part_time->"$part_time
topic_id="dm_ct_xdr_lte_5g_s1mme_cell_d"
app_name=$topic_id"-"$part_time
jar_file=${app_home}/"bigdata-telcom-hadoop-app-0.0.1-jar-with-dependencies.jar"


spark-submit --name $app_name --class com.tong.bigdata.spark.sql.SparkSqlApp \
--conf spark.network.timeout=10000001 \
--conf spark.executor.heartbeatInterval=10000000 \
--conf spark.executor.memoryOverhead=1536 \
--conf spark.dynamicAllocation.executorIdleTimeout=600 \
--conf spark.sql.adaptive.enabled=true \
--conf spark.sql.adaptive.shuffle.targetPostShuffleInputSize=134217728b \
--conf spark.sql.adaptive.join.enabled=true \
--master yarn --deploy-mode $deploy_mode \
--driver-cores 1 \
--driver-memory 1g \
--num-executors 4 \
--executor-memory 4g \
--executor-cores 4 \
--queue $queue_name \
${jar_file} \
--mq_server ${mq_server} --mq_topic ${mq_topic} \
--id $topic_id --part_time $part_time \
--source_db ${hive_db_name} --result_db ${hive_db_name}

