package main

//从mysql数据库导出文件，可导出xlsx和csv两种格式
//作者：chentiande
//

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)


var TIME_LOCATION_CST *time.Location

//全局错误标识，如果有错误，显示执行失败
var errflag bool

//增加GBK到utf8函数转换，将数据库取出的数据转成uft8然后保存到excel
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func createtable(filename string, tablename string) string {

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("无法打开文件 %s\n", filename)
	}
	nr := csv.NewReader(fp)
	wcomma := []rune("|")
	nr.Comma = wcomma[0]

	recode, err := nr.Read()
	if err != nil {
		log.Fatalf("csv读取错误,err=%s\n", err)
	}
	colstr := strings.Join(recode, " varchar(255),\n") + " varchar(255)"
	colstr = strings.Replace(colstr, "\xEF\xBB\xBF", "", -1)

	return "create table " + tablename + "\n (" + colstr + "\n)"
}

func querytocsv(db *sql.DB) {

	f, err := os.Create("xxx.csv") //创建文件
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)         //创建一个新的写入文件流

	wcomma := []rune("|")
	w.Comma = wcomma[0]
	str := os.Args[1]
	//解决密码中有特殊字符进行转义后去掉转义斜杠
	str = strings.Replace(str, "\\", "", -1)

	strsql := "select a_dn,a_master_phy_cell_id,a_related_gnb_dn,a_related_gnb_id from pm"
	fmt.Println("sql=", "select * from pm")

	num := 0

	rows, err := db.Query(strsql)

	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	w.Write(cols) //写入表头数据
	w.Flush()

	for rows.Next() {

		err = rows.Scan(dest...)
		num = num + 1

		for i, raw := range rawResult {

			if raw == nil || string(raw) == "" {
				result[i] = ""
			} else {

			}

		}

		w.Write(result) //写入数据

		if num%100000 == 0 {
			w.Flush()
			fmt.Println("导出" + strconv.Itoa(num) + "行")

		}

	}

	w.Flush()
	fmt.Println("总数据量为" + strconv.Itoa(num) + "行")
	rows.Close()
	fmt.Println("导出数据成功")
}

func init1(tablename string ) {
	TIME_LOCATION_CST, _ = time.LoadLocation("Asia/Shanghai")
	errflag = false
	_ = os.Mkdir("log", 755)
	file := "./log/"+tablename+"_" + time.Now().Format("200601021504") + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[loadgp]")
	//日志标识
	log.SetFlags(log.LstdFlags)
	return
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
func loadcsv(commitcount int,debug bool, allnum int, wcomma1 string, isgbk int, i int, ch_run chan int, db *sql.DB, filename string, tablename string) {

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("无法打开文件 %s\n", filename)
	}
	nr := csv.NewReader(fp)

	wcomma := []rune(wcomma1)
	nr.Comma = wcomma[0]
	nr.LazyQuotes = true

	recode, err := nr.Read()
	if err != nil {
		log.Fatalf("csv读取错误,err=%s\n", err)
	}
	colstr := strings.Join(recode, ",")
	colstr = strings.Replace(colstr, "\xEF\xBB\xBF", "", -1)
	valstr := ""
	for x, _ := range recode {

		valstr = valstr + fmt.Sprintf(",$%v", x+1)

	}
	valstr = valstr[1:]

	strsql := "insert  into " + tablename + " (" + colstr + ") values (" + valstr + ")"
	log.Printf("第%d个线程:开始对表%s插入数据,导入文件名=%s\n",i+1, tablename, filename)
	fmt.Printf("第%d个线程:开始对表%s插入数据,导入文件名=%s\n",i+1, tablename, filename)
	_, err = db.Exec("select 1 from " + tablename + " where 1=2")
	if err != nil {
		log.Fatalf("打开数据库错误,err=%s\n", err)
	}
	var txbool bool = false
	var tx *sql.Tx
	filenum := 0
	j := 1
	m := 1
	for {
		rows, err := nr.Read()
        if err == io.EOF {
			goto xxx
		}

		if err != nil {
			if debug{
				log.Println("err:",err)
			}
			
			 continue
		}
		if j%allnum == i {
			filenum += 1
			var dest []interface{} = make([]interface{}, len(rows))

			for m, v := range rows {

				if v == "NullInt" {
					dest[m] = interface{}(sql.NullInt64{Int64: 0, Valid: false})
				} else if v == "NullString" {
					dest[m] = interface{}(sql.NullString{String: "", Valid: false})
				} else if v == "NULL" {
					dest[m] = interface{}(sql.NullString{String: "", Valid: false})
				} else if v == "" {
					dest[m] = interface{}(sql.NullString{String: "", Valid: false})

				} else {
					if isgbk == 1 {
						vv, _ := GbkToUtf8([]byte(v))
						dest[m] = interface{}(string(vv))
					} else {
						dest[m] = interface{}(v)
					}
				}

			}

			if !txbool {

				tx, _ = db.Begin()
			}

			_, err = tx.Exec(strsql, dest...)

			if err != nil {
				if debug{
					fmt.Printf("提交发生错误：err=%s,dest=%v,strsql=%s\n", err, dest, strsql)
					log.Printf("提交发生错误：err=%s,dest=%v,strsql=%s\n", err, dest, strsql)
				}
				tx.Rollback()
				txbool = false
				j++
				continue
			} else {
				txbool = true

			}
			if (m-1)%commitcount == 0 && txbool {
				if commitcount>=1000{
					fmt.Printf("%d插入%d条数据\n", i, m-1)
				}
				if txbool == true {

				
				tx.Commit()
				txbool = false
				}
			}
			m++
		}
		j++
	}
xxx:
	if txbool == true {
		tx.Commit()
	
	}

	log.Printf("第%d个线程:文件%s有数据%d条,插入%s数据成功,共插入%d条数据\n", i+1, filename, filenum, tablename, m-1)
	fmt.Printf("第%d个线程:文件%s有数据%d条,插入%s数据成功,共插入%d条数据\n", i+1, filename, filenum, tablename, m-1)

	err = fp.Close()
	if err != nil {
		log.Fatalf("文件关闭错误,err=%s\n", err)
	}
	ch_run <- m - 1
}

func main() {
	tt := time.Now()

	var csvfile = flag.String("f", "pm1.csv", "要导入的csv文件")
	var tablename = flag.String("t", "pm1", "要导入的数据库表名")
	var showVer = flag.Bool("v", false, "显示版本信息")
	var showsql = flag.Bool("sql", false, "生成创建表sql")
	var wcomma = flag.String("wcomma", "^", "csv字段分隔符")
	var isgbk = flag.Int("isgbk", 0, "输入csv文件格式，如果是gbk则配置为1")
	var debug = flag.Bool("debug", false, "打开调测，异常会打印")
	var del = flag.Bool("del", true, "插入数据前先清除数据")
	var num = flag.Int("num", 1, "并发配置数量")
	var commitcount = flag.Int("commit", 1000, "批次提交数据量")
	var conf=flag.String("conf","db.conf","数据库连接配置文件")

	flag.Parse()
	init1(*tablename)
	dbconf, dberr := getconf(*conf)
	if dberr != nil {
		fmt.Println("get db.conf err:", dberr)
	}
	db, err := sql.Open("postgres", "user="+dbconf.Dbuser+" password="+dbconf.Dbpwd+" dbname="+dbconf.Dbname+" host="+dbconf.Dbip+" port="+dbconf.Dbport+" sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(100)
	//querytocsv(db)
	if *showVer {
		// Printf( "build name:\t%s\nbuild ver:\t%s\nbuild time:\t%s\nCommitID:%s\n", BuildName, BuildVersion, BuildTime, CommitID )
		fmt.Printf("build name:\t%s\n", "tongtech loadgp")
		fmt.Printf("build ver:\t%s\n", "20210525")

		os.Exit(0)
	}

	if *showsql {
		// Printf( "build name:\t%s\nbuild ver:\t%s\nbuild time:\t%s\nCommitID:%s\n", BuildName, BuildVersion, BuildTime, CommitID )
		fmt.Println(createtable(*csvfile, *tablename))

		os.Exit(0)
	}

	var i int
	var intChan chan int
	intChan = make(chan int, *num)
	if *del {
		_, err = db.Exec("truncate table " + *tablename)
		log.Println(*tablename, "删除数据成功")
	}

	if err != nil {
		log.Fatalf("执行sql错误,err=%s\n", err)
	}

	//如果加了num参数，将取模后对数据进行拆分，输入文件前面为1_ 2_
	for i = 0; i < *num; i++ {

		if *num == 1 {
			go loadcsv(*commitcount,*debug, *num, *wcomma, *isgbk, i, intChan, db, *csvfile, *tablename)
		} else {
			go loadcsv(*commitcount,*debug, *num, *wcomma, *isgbk, i, intChan, db, *csvfile, *tablename)
		}

	}
	succount := 0
	for i = 0; i < *num; i++ {
		succount += <-intChan
	}

	alltime := int(time.Since(tt)) / 1000000
	log.Printf("执行成功,time=%dms,共插入数据%d条\n", alltime, succount)
	fmt.Printf("执行成功,time=%dms,共插入数据%d条\n", alltime, succount)

}
