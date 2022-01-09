package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type SparkTables struct {
	tablename []string
	sp        []SparkTable
}

type SparkTable struct {
	SourceisDwd      bool //多维度汇总
	SourceisDwdCount int  //多维度数量
	SourcePmCount    int
	SourcePmTable    string
	SourceWhere      string
	SourcePmTables   []string
	DestTableZh      string
	SourceCmTable    string
	DestTable        string
	DestColumnName   []string
	DestColumnType   []string
	DestColumnExpr   []string
	DestColumnDim    []string
	DestColumnKJ     []string
	DestColumnSJ     []string
}

func main() {
	var filename, sheetname string
	var version bool
	flag.StringVar(&filename, "filename", "aaaa.xlsx", "输入的excel文件名称")
	flag.StringVar(&sheetname, "sheetname", "DM层_d_pto", "excel文件的sheet名称")
	flag.BoolVar(&version, "v", false, "查看版本")

	flag.Parse()

	if version {
		msg :=
			` ----------------------------     
| author:chentd@tongtech.com |
 ----------------------------
     2022-01-09 v0.1 实现根据模板自动生成执行脚本
	1、生成function标签
	2、生成临时表，根据多张表聚合
	3、生成多种维度汇总
	4、生成性能表和工参表进行关联汇总
		 `
		fmt.Println(msg)
		os.Exit(0)

	}

	st := makeSparkTable(filename, sheetname)

	for i := 1; i < len(st.sp); i++ {
		makeHql(st.sp[i], sheetname)
		makeSh(st.sp[i], sheetname)
		makeXml(st.sp[i], sheetname)
	}

	fmt.Println("生成成功")
}
func isfuncstr(st []string) string {

	function := ""
	m3 := map[string]string{
		"weightedAverage": "weightedAverage",
	}

	for _, v := range m3 {
		for _, y := range st {
			if strings.Index(y, v) > 0 {
				function += "<function>" + v + "</function>\n"
				goto Loop
			}

		}
	Loop:
	}
	return function
}

func makeXml(st SparkTable, filedir string) {
	function := isfuncstr(st.DestColumnExpr)
	resulttable := ""
	temptable := ""
	dim := ""
	selectcol := ""
	groupby := ""
	tmptable := ""
	selectall := ""

	if st.SourcePmCount > 1 || st.SourceisDwd {
		temptable = " <tableName>tmp_" + st.DestTable[0:20] + "</tableName>\n<sql>\n<![CDATA[select \n"

		allcol := getScounter(st, "")

		for _, v := range strings.Split(st.SourcePmTable, "#") {

			vcol := getScounter(st, v)

			for i := 0; i < len(allcol); i++ {
				if instrings(allcol[i], vcol) {
					temptable += allcol[i] + ",\n"
				} else {
					temptable += "0 as " + allcol[i] + ",\n"
				}

			}
			for i := 0; i < len(st.DestColumnExpr); i++ {
				if IsNum(st.DestColumnExpr[i]) {
					temptable += st.DestColumnExpr[i] + " as " + st.DestColumnName[i] + ",\n"
				}
			}
			temptable = temptable[:len(temptable)-2] + "\n from ${source_db}." + v + "\nwhere p_date>='${p_date_start}' and p_date<'${p_date_end}'"

			aaa := strings.Split(st.SourceWhere, "/")

			for vv := range aaa {

				bb := strings.Split(aaa[vv], "#")

				if len(bb) > 1 && bb[0] == v {

					temptable += " and " + bb[1]
				}
			}

			temptable += "\n union all\n select\n"
		}

		temptable = temptable[:len(temptable)-18] + `	]]>
		</sql>`
	}

	for z := 0; z < st.SourceisDwdCount; z++ {

		selectcol = ""
		dim = ""
		groupby = ""

		for i := 0; i < len(st.DestColumnExpr); i++ {

			if st.DestTable[len(st.DestTable)-2:] == st.SourcePmTable[len(st.SourcePmTable)-2:] {
				switch strings.ToLower(st.DestColumnKJ[i]) {
				case "max":
					selectcol += "max(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "min":
					selectcol += "min(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "sum":
					if IsNum(st.DestColumnExpr[i]) {
						selectcol += "sum(" + st.DestColumnName[i] + ") as " + st.DestColumnName[i] + ",\n"

					} else {
						selectcol += "sum(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
					}
				case "avg":
					selectcol += "avg(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "dim":
					aa := strings.Split(st.DestColumnExpr[i], "/")
					if aa[z] == "置空" {
						aa[z] = "''"
					}
					if aa[z] == "空" {
						aa[z] = "NULL"
					}
					dim += aa[z] + " as " + st.DestColumnName[i] + ",\n"
					groupby += aa[z] + ","
				case "time":
					dim += "substr('${p_start_time}',1,8) as " + st.DestColumnName[i] + ",\n"

				default:
					if st.DestColumnName[i] == "starttime" {
						dim += "substr('${p_start_time}',1,8) as " + st.DestColumnName[i] + ",\n"
					}
					if st.DestColumnExpr[i] == "置空" {
						selectcol += "''" + " as " + st.DestColumnName[i] + ",\n"
					} else if st.DestColumnExpr[i] == "无" {
						selectcol += "0" + " as " + st.DestColumnName[i] + ",\n"
					} else if st.DestColumnName[i] == "starttime" {
						dim += "substr('${p_start_time}',1,8) as " + st.DestColumnName[i] + ",\n"
					} else {
						selectcol += st.DestColumnExpr[i] + " as " + st.DestColumnName[i] + ",\n"

					}
				}

			} else {
				switch strings.ToLower(st.DestColumnSJ[i]) {
				case "max":
					selectcol += "max(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "min":
					selectcol += "min(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "sum":
					if IsNum(st.DestColumnExpr[i]) {
						selectcol += "sum(" + st.DestColumnName[i] + ") as " + st.DestColumnName[i] + ",\n"

					} else {
						selectcol += "sum(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
					}
				case "avg":
					selectcol += "avg(" + st.DestColumnExpr[i] + ") as " + st.DestColumnName[i] + ",\n"
				case "dim":
					aa := strings.Split(st.DestColumnExpr[i], "/")
					if aa[z] == "置空" {
						aa[z] = "''"
					}
					if aa[z] == "空" {
						aa[z] = "NULL"
					}
					dim += aa[z] + " as " + st.DestColumnName[i] + ",\n"
					groupby += aa[z] + ","
				case "time":
					dim += "substr('${p_start_time}',1,8) as " + st.DestColumnName[i] + ",\n"
				default:
					if st.DestColumnExpr[i] == "置空" {
						selectcol += "''" + " as " + st.DestColumnName[i] + ",\n"
					} else if st.DestColumnExpr[i] == "无" {
						selectcol += "0" + " as " + st.DestColumnName[i] + ",\n"
					} else if st.DestColumnName[i] == "starttime" {
						dim += "substr('${p_start_time}',1,8) as " + st.DestColumnName[i] + ",\n"
					} else {
						selectcol += st.DestColumnExpr[i] + " as " + st.DestColumnName[i] + ",\n"

					}
				}

			}
		}

		if len(groupby) > 2 {
			aa := RemoveDuplicatesAndEmpty(strings.Split(groupby, ","))
			groupby = strings.Join(aa, ",")
			if len(groupby) > 0 {
				groupby = "\n group by " + groupby
			}

		}

		if st.SourcePmCount > 1 || st.SourceisDwd {
			tmptable = "tmp_" + st.DestTable[0:20]
		} else {
			tmptable = "${source_db}." + st.SourcePmTable + " where p_date>='${p_date_start}' and p_date<'${p_date_end}' "
			aaa := strings.Split(st.SourceWhere, "/")

			for vv := range aaa {

				bb := strings.Split(aaa[vv], "#")

				if len(bb) > 1 && bb[0] == st.SourcePmTable {

					tmptable += " and " + bb[1]
				}
			}

		}

		selectall += "\nselect " + dim + selectcol[:len(selectcol)-2] + "\nfrom " + tmptable + groupby + "\nunion all"
	}
	if st.SourceCmTable != "" {

		selectall = "\n select " + dim + selectcol[:len(selectcol)-2] + "\nfrom \n(select * from " + tmptable + " \n) t1 \n " + st.SourceWhere[strings.Index(st.SourceWhere, "inner join"):] + "\n" + groupby
	}

	if selectall[len(selectall)-9:] == "union all" {
		selectall = selectall[:len(selectall)-9]
	}

	resulttable = `        <resultTable>
	<sql>
		<![CDATA[
			insert overwrite table ${result_db}.` + st.DestTable + ` partition(p_date='${p_date}')
		` +
		selectall + `
						  ]]>
            </sql>
        </resultTable>`

	strxml := `<AggrSqlConfig topicId="` + st.DestTable + `" period="1` + st.DestTable[len(st.DestTable)-1:] + `"
               desc="` + st.DestTableZh + `">
      <functions>
	` + function + `  </functions>

    <tempTables>
	` + temptable + `
    </tempTables>

    <resultTables>
	` + resulttable + `
	</resultTables>
	</AggrSqlConfig>`

	fw(filedir, st.DestTable+".xml", strxml)

}

func fw(filedir, filename string, content string) {
	os.Mkdir(filedir, 660)
	f, _ := os.Create(filedir + "/" + filename) //创建文件

	defer f.Close()
	f.WriteString(content)
	f.Sync()

}

func makeSh(st SparkTable, filedir string) {
	sh := `#!/bin/bash
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
	topic_id="` + st.DestTable + `"
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
	`
	fw(filedir, st.DestTable+".sh", sh)
}

func makeHql(st SparkTable, filedir string) {
	column := ""
	for i := 0; i < len(st.DestColumnName); i++ {
		column += "`" + st.DestColumnName[i] + "` " + st.DestColumnType[i] + ",\n"
	}

	column = column[:len(column)-2]

	hsql := `create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists ` + st.DestTable + `;
create EXTERNAL table IF NOT EXISTS ` + st.DestTable + `(
` + column + `) 
PARTITIONED BY ( 
  ` + "`p_date`" + ` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/` + st.DestTable + `/";`
	fw(filedir, st.DestTable+".hql", hsql)

}

func makeSparkTable(xlsname string, sheetname string) SparkTables {
	f, err := excelize.OpenFile(xlsname)

	if err != nil {
		log.Fatalf("无法打开文件 %s\n", xlsname)

	}

	recode := f.GetRows(sheetname)

	if len(recode) == 0 {
		log.Fatalln("sheetname 错误")
	}

	sts := SparkTables{}
	st := SparkTable{}
	var i int
	for z, row := range recode[1:] {

		if len(sts.tablename) == 0 || sts.tablename[len(sts.tablename)-1] != row[4] {
			i = 0
			sts.sp = append(sts.sp, st)
			st = SparkTable{}
		}

		if i == 0 {
			sts.tablename = append(sts.tablename, row[4])
			st.SourceWhere = row[15]
			st.SourceisDwdCount = 1

			tmpsourcetable := strings.Split(strings.ReplaceAll(row[16], "\\", "/"), "/")
			pm := ""
			cm := ""
			cmcount := 0

			if len(tmpsourcetable) == 1 {
				pm = tmpsourcetable[0]
			} else {
				for _, v := range tmpsourcetable {
					if strings.ToLower(v[0:3]) == "dim" {
						cm += v + "#"
						cmcount += 1
					} else {
						pm += v + "#"
					}
				}
				pm = pm[:len(pm)-1]

			}

			st.SourcePmCount = len(tmpsourcetable) - cmcount
			st.SourcePmTable = pm[:len(pm)-1]
			if len(cm) > 0 {
				st.SourceCmTable = cm[:len(cm)-1]
			}

			st.DestTable = row[4]
			st.DestTableZh = row[3]

		}
		i++
		if strings.ToLower(row[11]) == "dim" && len(strings.Split(row[13], "/")) > 1 {
			st.SourceisDwd = true
			st.SourceisDwdCount = len(strings.Split(row[13], "/"))
		}
		st.SourcePmTables = append(st.SourcePmTables, row[16])
		st.DestColumnName = append(st.DestColumnName, row[6])
		st.DestColumnType = append(st.DestColumnType, row[7])
		st.DestColumnExpr = append(st.DestColumnExpr, strings.TrimRight(row[13], " "))
		st.DestColumnKJ = append(st.DestColumnKJ, row[11])
		st.DestColumnSJ = append(st.DestColumnSJ, row[12])

		if z == len(recode)-2 {
			sts.sp = append(sts.sp, st)
		}
	}

	return sts
}
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	sort.Strings(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		if IsNum(a[i]) {
			continue
		}
		if a[i] == "''" || a[i] == "NULL" || a[i] == "array" || a[i] == "max" || a[i] == "min" || a[i] == "avg" || a[i] == "else" || a[i] == "nvl" || a[i] == "case" || a[i] == "when" || a[i] == "then" || a[i] == "sum" || a[i] == "and" || a[i] == "end" {
			continue
		}
		m3 := map[string]string{
			"weightedAverage": "weightedAverage",
		}
		for _, v := range m3 {

			if a[i] == v {

				goto loop
			}
		}

		ret = append(ret, a[i])
	loop:
	}
	for i := 0; i < len(ret); i++ {
		if i > 0 && len(ret[i-1]) > len(ret[i]) {
			ret[i-1], ret[i] = ret[i], ret[i-1]
		}
	}

	return
}

func getScounter(st SparkTable, table string) []string {
	str := ""
	if table == "" {
		str = strings.Join(st.DestColumnExpr, "#")

	} else {
		for i := 0; i < len(st.DestColumnExpr); i++ {
			if strings.Index(st.SourcePmTables[i], table) >= 0 {
				str += st.DestColumnExpr[i] + "#"
			}
		}
	}

	def1 := str
	def2 := strings.ReplaceAll(def1, "(", "#")
	def2 = strings.ReplaceAll(def2, "置空", "")
	def2 = strings.ReplaceAll(def2, "starttime(例原话单为2020-11-29 00:00:00~2020-11-30 00:00:00,开始时间为2020-11-29 00:00:00)", "")
	def2 = strings.ReplaceAll(def2, "空", "")
	def2 = strings.ReplaceAll(def2, "无", "")
	def2 = strings.ReplaceAll(def2, ")", "#")
	def2 = strings.ReplaceAll(def2, "/", "#")
	def2 = strings.ReplaceAll(def2, "*", "#")
	def2 = strings.ReplaceAll(def2, "+", "#")
	def2 = strings.ReplaceAll(def2, "-", "#")
	def2 = strings.ReplaceAll(def2, ">", "#")
	def2 = strings.ReplaceAll(def2, "<", "#")
	def2 = strings.ReplaceAll(def2, "=", "#")
	def2 = strings.ReplaceAll(def2, ",", "#")
	def2 = strings.ReplaceAll(def2, " ", "#")
	def2 = strings.ReplaceAll(def2, "'", "#")

	def3 := strings.Split(def2, "#")
	def3 = RemoveDuplicatesAndEmpty(def3)
	return def3
}

func instrings(a string, b []string) bool {

	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}
