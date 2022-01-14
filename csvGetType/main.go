package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tableinfo struct {
	filename     []string
	colsoucename []string
	colname      []string
	coltype      []string
	collen       []int
	colmax       []int
	colmin       []int
	coldata      []string
}

func main() {
	var dirname string
	var wcomma string
	var isgbk int
	var isheader int
	flag.StringVar(&dirname, "dir", "csv", "csv所在的目录")
	flag.StringVar(&wcomma, "wcomma", "^", "csv字段分隔符")
	flag.IntVar(&isgbk, "isgbk", 1, "输入csv文件格式，如果是gbk则配置为1")
	flag.IntVar(&isheader, "isheader", 1, "输入csv文件第一行是否为表头，1为是")
	flag.Parse()
	fmt.Println(dirname)
	files1, _ := ioutil.ReadDir(dirname)
	for _, v := range files1 {

		if strings.ToLower(v.Name()[len(v.Name())-4:]) == ".csv" {
			unloadxls(v.Name()[:len(v.Name())-4]+".xlsx", getcsvinfo(dirname+"/"+v.Name(), wcomma, isgbk, isheader))
			fmt.Println(v.Name(), "转化成功")
		}

	}

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
		return 2000
	}
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
	for i, _ := range tableinfo2.colsoucename {
		if len(tableinfo2.collen) > i {
			row = sheet.AddRow()

			cell = row.AddCell()
			cell.Value = tableinfo2.colsoucename[i]
			cell = row.AddCell()
			cell.Value = tableinfo2.colname[i]
			cell = row.AddCell()
			if tableinfo2.coltype[i] == "varchar" {
				cell.Value = tableinfo2.coltype[i] + "(" + strconv.Itoa(getlen(tableinfo2.collen[i])) + ")"
				cell = row.AddCell()
			} else if tableinfo2.coltype[i] == "int" && getlen(tableinfo2.collen[i]) == 64 {
				cell.Value = "bigint"
				cell = row.AddCell()
			} else {
				cell.Value = tableinfo2.coltype[i]
				cell = row.AddCell()
			}
			if tableinfo2.coltype[i] == "int" && getlen(tableinfo2.collen[i]) <= 32 {
				cell.Value = "32"
				cell = row.AddCell()
			} else if tableinfo2.coltype[i] == "int" && getlen(tableinfo2.collen[i]) == 64 {
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

func getcsvinfo(filename string, fgf string, isgbk int, isheader int) tableinfo {
	ti := tableinfo{}
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	nr := csv.NewReader(fp)

	wcomma := []rune(fgf)
	nr.Comma = wcomma[0]

	recode, err := nr.Read()
	if err != nil {
		fmt.Println("csv读取错误,err=%s\n", err)
	}

	for mm, v := range recode {
		if isheader == 1 {
			ti.colsoucename = append(ti.colsoucename, v)
			ti.colname = append(ti.colname, strings.ToLower(v))
		} else {
			ti.colsoucename = append(ti.colsoucename, "column"+strconv.Itoa(mm))
			ti.colname = append(ti.colname, "column"+strconv.Itoa(mm))
		}

	}
	for i := 0; ; i++ {
		rows, err := nr.Read()
		if err != nil {
			goto xxx
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

				if ti.collen[k] < len(v) {
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
	return ti
}

func gettype(colvalue string) string {
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
