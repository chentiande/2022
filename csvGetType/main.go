package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type tableinfo struct {
	filename     []string
	tablename    []string
	tablezhname  []string
	colsoucename []string
	colname      []string
	colzhname    []string
	coltype      []string
	collen       []int
	colmax       []int
	colmin       []int
	coldata      []string
	isstand      []string
}

func main() {
	var dirname string
	var wcomma string
	var isgbk int
	var isheader int
	var filename string
	var pk string
	var version bool
	var ylcount int

	flag.StringVar(&dirname, "dir", "csv", "csv所在的目录")
	flag.StringVar(&wcomma, "wcomma", "^", "csv字段分隔符")
	flag.StringVar(&filename, "mf", "Model.xlsx", "标准模型表")
	flag.StringVar(&pk, "pk", "int_id", "建表主键指定")
	flag.IntVar(&isgbk, "isgbk", 0, "输入csv文件格式，如果是gbk则配置为1")
	flag.IntVar(&isheader, "isheader", 1, "输入csv文件第一行是否为表头，1为是")
	flag.IntVar(&ylcount, "ylcount", 1000, "识别样例数据的类型抽样比例，设置为1全部，1000为千分之一")

	flag.BoolVar(&version, "v", false, "查看版本")

	flag.Parse()

	if version {
		msg :=
			` ----------------------------     
| author:chentd@tongtech.com |
 ----------------------------
     2022-01-19 v0.1 根据样例数据和标准模型生成模型
	1、根据csv样例数据定义数据类型，int bigint  varchar timestamp
	2、根据模型中的样例文件名和表名进行比对分析
	3、根据标准模型进行分析
	4、生成单个模型的比对结果，生成单个模型的建表脚本，生成总的建表脚本

	2022-01-21 v0.2 
	1、修正识别长度超过2000的基于识别+50
	2、增加类型识别样例条数设置，默认识别1000行

	2022-02-11 v0.3 
	1、根据抽样进行类型识别
		 `
		fmt.Println(msg)
		os.Exit(0)

	}

	files1, _ := ioutil.ReadDir(dirname)

	sqlall := ""
	allerr := ""
	allcomment := ""
	for _, v := range files1 {

		if strings.ToLower(v.Name()[len(v.Name())-4:]) == ".csv" {
			xx := getcsvinfo(ylcount, filename, dirname, v.Name(), wcomma, isgbk, isheader)

			unloadxls(v.Name()[:len(v.Name())-4]+".xlsx", xx)
			a, b, c := mksql(pk,v.Name()[:len(v.Name())-4], xx)
			sqlall += a
			allerr += b
			allcomment += c
			fmt.Println(v.Name(), "完成转化")
		}

	}

	fw("sql", "all.sql", sqlall)
	fw("sql", "allcomment.sql", allcomment)
	fw("sql", "allerr.csv", "\xEF\xBB\xBF文件名\t表名\t字段名\t类型\t备注\n"+allerr)

}
func getlen(length int) int {
	switch {
	case length <= 16:
		return 16
	case length <= 32:
		return 32
	case length <= 64:
		return 64
	case length <= 128:
		return 128
	case length <= 256:
		return 255
	case length <= 1000:
		return 1000
	default:
		return length + 50
	}
}
func fw(filedir, filename string, content string) {
	os.Mkdir(filedir, 660)
	f, _ := os.Create(filedir + "/" + filename) //创建文件

	defer f.Close()
	f.WriteString(content)
	f.Sync()

}
func mksql(pk string,filename string, tableinfo2 tableinfo) (allsql, allerr, allcomment string) {

	if len(tableinfo2.coltype) == 0 {
		return "", "", ""
	}

	os.Mkdir("sql", 0660)
	commont := "comment on table cadb." + tableinfo2.tablename[0] + " is '" + tableinfo2.tablezhname[0] + "';\n"
	sql := "drop table if exists cadb." + tableinfo2.tablename[0] + ";\n"

	sql += "create table cadb." + tableinfo2.tablename[0] + " (\n"
	for i, _ := range tableinfo2.colsoucename {

		if tableinfo2.isstand[i] == "缺失规范字段" || tableinfo2.isstand[i] == "规范字段，类型不匹配" {
			allerr += tableinfo2.filename[i] + "\t" + tableinfo2.tablename[i] + "\t" + tableinfo2.colname[i] + "\t" + tableinfo2.coltype[i] + "\t" + tableinfo2.isstand[i] + "\n"
		}
		commont += "comment on column cadb." + tableinfo2.tablename[0] + "." + tableinfo2.colname[i] + " is '" + tableinfo2.colzhname[i] + "';\n"

		sql += tableinfo2.colname[i] + " "

		if tableinfo2.coltype[i] == "varchar" || tableinfo2.coltype[i] == "" {
			sql += "varchar(" + strconv.Itoa(getlen(tableinfo2.collen[i])) + "),\n"
		} else if tableinfo2.coltype[i] == "int" && tableinfo2.colmax[i] > 2147483647 {
			sql += "bigint,\n"

		} else {
			sql += tableinfo2.coltype[i] + ",\n"

		}

	}
	sql = sql + "PRIMARY KEY ("+pk+"));\n"
	fw("sql", filename+".sql", sql+commont)
	return sql, allerr, commont

}

func unloadxls(filename string, tableinfo2 tableinfo) {
	os.Mkdir("xlsx", 0660)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()

	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "表名"
	cell = row.AddCell()
	cell.Value = "厂家字段"
	cell = row.AddCell()
	cell.Value = "定义字段"
	cell = row.AddCell()
	cell.Value = "字段类型"
	cell = row.AddCell()
	cell.Value = "字段长度"
	cell = row.AddCell()
	cell.Value = "最大值"
	cell = row.AddCell()
	cell.Value = "最小值"
	cell = row.AddCell()
	cell.Value = "样例数据"
	cell = row.AddCell()
	cell.Value = "备注"
	for i, _ := range tableinfo2.colsoucename {
		if len(tableinfo2.collen) > i {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = tableinfo2.tablename[i]
			cell = row.AddCell()
			cell.Value = tableinfo2.colsoucename[i]
			cell = row.AddCell()
			cell.Value = tableinfo2.colname[i]
			cell = row.AddCell()
			if tableinfo2.coltype[i] == "varchar" || tableinfo2.coltype[i] == "" {
				cell.Value = "varchar(" + strconv.Itoa(getlen(tableinfo2.collen[i])) + ")"
				cell = row.AddCell()
			} else if tableinfo2.coltype[i] == "int" && tableinfo2.colmax[i] > 2147483647 {
				cell.Value = "bigint"
				cell = row.AddCell()
			} else {
				cell.Value = tableinfo2.coltype[i]
				cell = row.AddCell()
			}
			if tableinfo2.coltype[i] == "int" && tableinfo2.colmax[i] <= 2147483647 {
				cell.Value = "32"
				cell = row.AddCell()
			} else if tableinfo2.coltype[i] == "int" && tableinfo2.colmax[i] > 2147483647 {
				cell.Value = "64"
				cell = row.AddCell()
			} else if tableinfo2.coltype[i] == "timestamp" {
				cell.Value = "64"
				cell = row.AddCell()
			} else {
				cell.Value = strconv.Itoa(getlen(tableinfo2.collen[i]))
				cell = row.AddCell()
			}
			cell.Value = strconv.Itoa(tableinfo2.colmax[i])
			cell = row.AddCell()
			cell.Value = strconv.Itoa(tableinfo2.colmin[i])
			cell = row.AddCell()

			cell.Value = tableinfo2.coldata[i]
			cell = row.AddCell()
			cell.Value = tableinfo2.isstand[i]

		}

	}
	err = file.Save("xlsx/" + filename)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func getzhname(recode [][]string, filename string, colname string) string {

	if len(recode) == 0 {
		fmt.Println("sheetname 错误")
	} else {
		for i := range recode {

			if recode[i][7] == filename && recode[i][2] == colname {

				return recode[i][1]
			}
		}

	}
	return "未定义"
}

func getcsvinfo(ylcount int, mf string, dirname string, filename string, fgf string, isgbk int, isheader int) tableinfo {
	tablename := ""
	tablezhname := ""
	var recodesn [][]string
	if mf != "" {

		f, err := excelize.OpenFile(mf)

		if err != nil {
			fmt.Println("无法打开文件 ", mf)

		}

		recode := f.GetRows("文件表对应")

		recodesn = f.GetRows("综资接口")

		if len(recode) == 0 {
			fmt.Println("sheetname 错误")
		} else {
			for i := range recode {
				if recode[i][0] == filename {
					tablename = recode[i][1]
					tablezhname = recode[i][2]
				}
			}

		}

	}

	ti := tableinfo{}
	fp, err := os.Open(dirname + "/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	nr := csv.NewReader(fp)

	wcomma := []rune(fgf)
	nr.Comma = wcomma[0]
	nr.LazyQuotes = true

	recode, err := nr.Read()
	if err != nil {
		fmt.Println("csv读取错误,err=%s\n", err)
	}

	for mm, v := range recode {
		ti.filename = append(ti.filename, filename)
		ti.tablename = append(ti.tablename, tablename)
		ti.tablezhname = append(ti.tablezhname, tablename)
		ti.isstand = append(ti.isstand, "省内个性化字段")

		if isheader == 1 {
			ti.colsoucename = append(ti.colsoucename, v)
			ti.colname = append(ti.colname, strings.TrimSpace(strings.ToLower(v)))
			ti.colzhname = append(ti.colzhname, getzhname(recodesn, filename, strings.ToLower(v)))
		} else {
			ti.colsoucename = append(ti.colsoucename, "column"+strconv.Itoa(mm))
			ti.colname = append(ti.colname, "column"+strconv.Itoa(mm))
			ti.colzhname = append(ti.colzhname, "")
		}

	}

	var i int
	for i = 0; ; i++ {

		rows, err := nr.Read()

		if err == io.EOF {

			goto xxx
		}

		if err != nil {

			fmt.Println(rows, err)
			break
		}
		for k, v := range rows {

			if i == 0 {
				ti.filename = append(ti.filename, filename)
				ti.coltype = append(ti.coltype, gettype(v))
				ti.collen = append(ti.collen, len(v))

				if isgbk == 1 {
					tmp_v, _ := GbkToUtf8([]byte(v))
					ti.coldata = append(ti.coldata, string(tmp_v))
				} else {
					ti.coldata = append(ti.coldata, v)
				}

				ti.colmax = append(ti.colmax, 0)
				ti.colmin = append(ti.colmin, 0)
			} else {
				if (i%ylcount == 0 || i < 1000) && ti.coltype[k] != "varchar" {

					tmpz := gettype(v)
					if ti.coltype[k] != tmpz && tmpz == "varchar" {

						ti.coldata[k] = v
						ti.coltype[k] = tmpz
					}

				}

				if ti.collen[k] < len(v) {
					//fmt.Println(i,k,ti.colname[k],ti.collen[k] ,len(v))
					ti.collen[k] = len(v)
				}
				if ti.coldata[k] == "" {

					if isgbk == 1 {
						tmp_v, _ := GbkToUtf8([]byte(v))
						ti.coldata[k] = string(tmp_v)
					} else {
						ti.coldata[k] = v
					}

				}

				if ti.coltype[k] == "int" || ti.coltype[k] == "float" {
					vv, _ := strconv.Atoi(v)
					if ti.colmax[k] < vv {
						ti.colmax[k] = vv
					}
					if ti.colmin[k] > vv {
						ti.colmin[k] = vv
					}

				}

			}

		}

	}

xxx:
	fmt.Println("分析数据量为", i+1)
	if len(ti.coltype) == 0 {
		for _ = range ti.tablename {
			ti.coltype = append(ti.coltype, "varchar")
			ti.collen = append(ti.collen, 0)
			ti.colmax = append(ti.colmax, 0)
			ti.colmin = append(ti.colmin, 0)
			ti.coldata = append(ti.coldata, "")

		}
	}
	//fmt.Printf("%v", ti)
	if mf != "" {
		f, err := excelize.OpenFile(mf)

		if err != nil {
			fmt.Println("无法打开文件 ", mf)

		}

		recode := f.GetRows("标准模型")

		if len(recode) == 0 {
			fmt.Println("sheetname 错误")
		} else {
			for i := range recode {
				ishave := false

				if recode[i][0] == tablename {
					tablezhname = recode[i][1]
					for i2 := range ti.colname {
						ti.tablezhname[i2] = tablezhname

						if ti.colname[i2] == strings.TrimSpace(strings.ToLower(recode[i][2])) {
							ishave = true
							ti.colzhname[i2] = recode[i][3]

							if (len(recode[i][4]) > 6 && recode[i][4][:7] == "varchar") || ti.coltype[i2] == "" || ti.coltype[i2] == recode[i][4] || (len(recode[i][4]) > 6 && ti.coltype[i2] == recode[i][4][:7]) {
								ti.isstand[i2] = "规范字段，类型匹配"

								if (len(recode[i][4]) > 6 && recode[i][4][:7] == "varchar") || ti.coltype[i2] == "" {
									ti.coltype[i2] = "varchar"
								}
								strlen, _ := strconv.Atoi(recode[i][5])
								if ti.collen[i2] == 0 {
									ti.collen[i2] = strlen
								}
							} else {
								ti.isstand[i2] = "规范字段，类型不匹配"
								ti.coltype[i2] = recode[i][4]
							}
						}

					}
					if !ishave {
						ti.filename = append(ti.filename, filename)
						ti.tablename = append(ti.tablename, tablename)
						ti.tablezhname = append(ti.tablezhname, tablezhname)
						ti.colsoucename = append(ti.colsoucename, "")
						ti.colname = append(ti.colname, strings.TrimSpace(strings.ToLower(recode[i][2])))
						ti.colzhname = append(ti.colzhname, recode[i][3])
						ti.coltype = append(ti.coltype, recode[i][4])
						strlen, _ := strconv.Atoi(recode[i][5])
						ti.collen = append(ti.collen, strlen)
						ti.colmax = append(ti.colmax, 0)
						ti.colmin = append(ti.colmin, 0)
						ti.coldata = append(ti.coldata, "")
						ti.isstand = append(ti.isstand, "缺失规范字段")

					}

				}
			}

		}

	}

	return ti
}

func gettype(colvalue string) string {
	if colvalue == "" {
		return ""
	}

	reg := regexp.MustCompile(`^\d+\.\d+$`) // +表示匹配前一个字符的一次或者多次
	if reg == nil {
		fmt.Println("MustCompile err")
		return "err"
	}
	// 提取关键信息
	aa := reg.MatchString(colvalue)
	if aa {
		return "decimal(12,6)"
	}

	reg = regexp.MustCompile(`^\d+$`) // +表示匹配前一个字符的一次或者多次
	if reg == nil {
		fmt.Println("MustCompile err")
		return "err"
	}
	// 提取关键信息
	aa = reg.MatchString(colvalue)
	if aa {
		return "int"
	}
	reg = regexp.MustCompile("^((\\d{2}(([02468][048])|([13579][26]))[\\-\\/\\s]?((((0?[13578])|(1[02]))[\\-\\/\\s]?((0?[1-9])|([1-2][0-9])|(3[01])))|(((0?[469])|(11))[\\-\\/\\s]?((0?[1-9])|([1-2][0-9])|(30)))|(0?2[\\-\\/\\s]?((0?[1-9])|([1-2][0-9])))))|(\\d{2}(([02468][1235679])|([13579][01345789]))[\\-\\/\\s]?((((0?[13578])|(1[02]))[\\-\\/\\s]?((0?[1-9])|([1-2][0-9])|(3[01])))|(((0?[469])|(11))[\\-\\/\\s]?((0?[1-9])|([1-2][0-9])|(30)))|(0?2[\\-\\/\\s]?((0?[1-9])|(1[0-9])|(2[0-8]))))))(\\s(((0?[0-9])|([0][0])|([1][0-9])|([2][0-3]))\\:([0-5]?[0-9])((\\s)|(\\:([0-5]?[0-9])))))?$")
	if reg == nil {
		fmt.Println("MustCompile err")
		return "err"
	}
	// 提取关键信息
	aa = reg.MatchString(colvalue)
	if aa {
		return "timestamp"
	}
	return "varchar"

}
