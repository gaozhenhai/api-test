package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

//结构体
type Job struct {
	db    *sql.DB
	ch    chan int
	total int
	n     int
}

func CreateTestData(cmd *cobra.Command, args []string) {
	dsn, _ := cmd.Flags().GetString("dsn")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		glog.Errorln(err)
		return
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
	db.SetMaxOpenConns(100)                  //设置最大连接数

	total, _ := cmd.Flags().GetInt("total")
	gonum, _ := cmd.Flags().GetInt("connection")

	jobChan := make(chan Job, 20)
	go worker(jobChan, total)

	start := time.Now()
	ch := make(chan int, gonum)
	for n := 0; n < gonum; n++ {
		job := Job{
			db:    db,
			ch:    ch,
			total: total,
			n:     n,
		}
		jobChan <- job
	}

	ii := 0
	for {
		<-ch
		ii++
		if ii >= gonum {
			break
		}
	}
	glog.Infof("running time: %v", time.Since(start))
}

func worker(jobChan <-chan Job, total int) {
	for job := range jobChan {
		go sqlExec(job, total)
	}
}

func sqlExec(job Job, total int) {
	buf := make([]byte, 0, job.total)
	buf = append(buf, " insert into t_file_info(file_name,file_path) values "...)
	for i := 0; i < job.total; i++ {
		if i == job.total-1 {
			buf = append(buf, fmt.Sprintf("('m%08d','/tmp/m%08d'),", i, i)...)
		} else {
			buf = append(buf, fmt.Sprintf("('m%08d','/tmp/m%08d'),", i, i)...)
		}
	}

	ss := string(buf)

	ss = strings.TrimSuffix(ss, ",")

	glog.Infof("start install %d data with %d", total, job.n)
	if _, err := job.db.Exec(ss + ";"); err != nil {
		glog.Errorln(err)
	}
	glog.Infof("install %d data done with %d", total, job.n)
	job.ch <- 1
}
