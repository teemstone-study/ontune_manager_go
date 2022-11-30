package app

import (
	"fmt"
	"log"
	"manager/data"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBHandler struct {
	name string
	db   *sqlx.DB
	demo DemoInfo
}

func (d *DBHandler) CheckTable() {
	// Table Initialization
	var info_tables = [...]string{"tableinfo", "agentinfo", "lastrealtimeperf", "deviceid", "descid"}
	var metric_ref_tables = [...]string{"proccmd", "procuserid"}
	var metric_tables = [...]string{"realtimeperf", "realtimecpu", "realtimedisk", "realtimenet", "realtimepid", "realtimeproc"}

	// Info Table Check
	for _, tb := range info_tables {
		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM pg_tables where tablename=$1", tb).Scan(&exist_count)
		if err != nil {
			log.Fatal(err)
		}

		if exist_count == 0 {
			tx := d.db.MustBegin()
			switch tb {
			case "tableinfo":
				tx.MustExec(data.TableinfoStmt)
			case "agentinfo":
				tx.MustExec(data.AgentinfoStmt)
			case "lastrealtimeperf":
				tx.MustExec(data.LastrealtimeperfStmt)
			case "deviceid":
				tx.MustExec(data.DeviceidStmt)
			case "descid":
				tx.MustExec(data.DescidStmt)
			}
			tx.MustExec(data.InsertTableinfo, tb, time.Now().Unix())
			tx.Commit()
		}
	}

	// Metric Reference Table Check
	for _, tb := range metric_ref_tables {
		var tablename string
		switch d.name {
		case "pg-hour", "pg-day":
			tablename = tb + "_" + time.Now().Format("060102") + "00"
		case "pg-week":
			_, week := time.Now().ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			tablename = tb + "_" + time.Now().Format("0601") + weekstr
		case "pg-month":
			tablename = tb + "_" + time.Now().Format("0601")
		default:
			tablename = tb
		}

		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM public.tableinfo where _tablename=$1", tablename).Scan(&exist_count)
		if err != nil {
			log.Fatal(err)
		}

		if exist_count == 0 {
			tx := d.db.MustBegin()
			tx.MustExec(fmt.Sprintf(data.ProcStmt, tablename))
			tx.MustExec(data.InsertTableinfo, tablename, time.Now().Unix())
			tx.Commit()
		}
	}

	// Metric Table Check
	for _, tb := range metric_tables {
		var tablename string
		switch d.name {
		case "pg-hour":
			tablename = tb + "_" + time.Now().Format("06010215")
		case "pg-day":
			tablename = tb + "_" + time.Now().Format("060102") + "00"
		case "pg-week":
			_, week := time.Now().ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			tablename = tb + "_" + time.Now().Format("0601") + weekstr
		case "pg-month":
			tablename = tb + "_" + time.Now().Format("0601")
		default:
			tablename = tb
		}

		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM public.tableinfo where _tablename=$1", tablename).Scan(&exist_count)
		if err != nil {
			log.Fatal(err)
		}

		if exist_count == 0 {
			tx := d.db.MustBegin()
			if d.name[:2] == "pg" {
				switch tb {
				case "realtimeperf":
					tx.MustExec(fmt.Sprintf(data.RealtimeperfPgStmt, tablename))
				case "realtimecpu":
					tx.MustExec(fmt.Sprintf(data.RealtimecpuPgStmt, tablename))
				case "realtimedisk":
					tx.MustExec(fmt.Sprintf(data.RealtimediskPgStmt, tablename))
				case "realtimenet":
					tx.MustExec(fmt.Sprintf(data.RealtimenetPgStmt, tablename))
				case "realtimepid":
					tx.MustExec(fmt.Sprintf(data.RealtimepidPgStmt, tablename))
				case "realtimeproc":
					tx.MustExec(fmt.Sprintf(data.RealtimeprocPgStmt, tablename))
				}
			} else {
				switch tb {
				case "realtimeperf":
					tx.MustExec(data.RealtimeperfTsStmt)
				case "realtimecpu":
					tx.MustExec(data.RealtimecpuTsStmt)
				case "realtimedisk":
					tx.MustExec(data.RealtimediskTsStmt)
				case "realtimenet":
					tx.MustExec(data.RealtimenetTsStmt)
				case "realtimepid":
					tx.MustExec(data.RealtimepidTsStmt)
				case "realtimeproc":
					tx.MustExec(data.RealtimeprocTsStmt)
				}
			}
			tx.MustExec(data.InsertTableinfo, tablename, time.Now().Unix())
			tx.Commit()
		}
	}
}

func (d *DBHandler) DemoHostSetting(arr *data.AgentinfoArr) {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM agentinfo where _enabled=1").Scan(&exist_count)
	if err != nil {
		log.Fatal(err)
	}

	if exist_count != d.demo.HostCount {
		tx := d.db.MustBegin()
		tx.MustExec(data.DeleteAgentinfo)
		tx.MustExec(data.DemoInsertAgentinfoUnnest, arr.GetArgs()...)
		tx.Commit()
	}
}

func (d *DBHandler) DemoHostStateChange(agent_str string) {
	tx := d.db.MustBegin()
	timestamp := time.Now().Unix()

	// connected=0인 agent를 1로 초기화
	tx.MustExec(data.DemoUpdateAgentinfoReset, timestamp)
	tx.MustExec(fmt.Sprintf(data.DemoUpdateAgentinfoState, 0, timestamp, agent_str, 1, timestamp))
	tx.Commit()
}

func (d *DBHandler) DemoBptUpdate(arr *data.LastrealtimeperfArr) {
	tx := d.db.MustBegin()
	tx.MustExec(data.DeleteLastrealtimeperf)

	tx.MustExec(data.DemoInsertLastrealtimeperf, arr.GetArgs()...)
	tx.Commit()
}

func (d *DBHandler) DBClose() {
	defer d.db.Close()
}

func DBConnection(dbinfo DbInfo) *sqlx.DB {
	conn := dbinfo.Datasource()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(5)
	return db
}

func DBInit(dbinfo DbInfo, info *data.AgentinfoArr) *DBHandler {
	d := &DBHandler{
		name: dbinfo.Name,
		db:   DBConnection(dbinfo),
	}
	d.CheckTable()
	d.DemoHostSetting(info)

	return d
}
