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
	name  string
	db    *sqlx.DB
	demo  DemoInfo
	hosts map[string]int
}

var (
	info_tables       = [...]string{"tableinfo", "agentinfo", "lastrealtimeperf", "deviceid", "descid"}
	metric_ref_tables = [...]string{"proccmd", "procuserid", "procargid"}
	metric_tables     = [...]string{"realtimeperf", "realtimecpu", "realtimedisk", "realtimenet", "realtimepid", "realtimeproc"}
)

func (d *DBHandler) CheckTable() {
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
		tablename := d.GetTablename("metric_ref", tb)

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
		tablename := d.GetTablename("metric", tb)

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
	err := d.db.QueryRow("SELECT COUNT(*) FROM agentinfo where _enabled=1 and _hostname like 'Dummy%'").Scan(&exist_count)
	if err != nil {
		log.Fatal(err)
	}

	if exist_count < d.demo.HostCount {
		tx := d.db.MustBegin()
		tx.MustExec(data.DeleteAgentinfoDummy)
		tx.MustExec(data.InsertAgentinfoUnnest, arr.GetArgs()...)
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

func (d *DBHandler) GetHostnames() map[string]int {
	hostnames := make(map[string]int, 0)
	rows, err := d.db.Query("SELECT _agentid, _hostname FROM agentinfo where _enabled=1")
	if err != nil {
		log.Println("Query Error")
		return hostnames
	}
	defer rows.Close()

	for rows.Next() {
		var agentid int
		var hostname string
		err := rows.Scan(&agentid, &hostname)
		if err != nil {
			log.Println("Query Error")
			return hostnames
		}
		hostnames[hostname] = agentid
	}

	return hostnames
}

func (d *DBHandler) SetHost(cshost *data.AgentHostAgentInfo) {
	/*--------------------------
	AgentID -> Hostname
	Agentname -> Hostnameext
	Model -> Model
	Serial -> Serial
	Ip -> Allip, Ipaddress
	Os -> Os
	Agentversion -> Agentversion
	ProcessCount -> Processorcount
	ProcessClock -> Processorclock
	MemoriSize -> Memorysize
	SwapMemory -> Swapsize
	----------------------------*/
	fmt.Printf("agentinfo update %s\n", cshost.AgentID)

	if _, ok := d.hosts[cshost.AgentID]; ok {
		// Already Exists
		ts := time.Now().Unix()
		tx := d.db.MustBegin()
		tx.MustExec(data.UpdateAgentinfo, cshost.AgentName, cshost.Ip, cshost.Model, cshost.Serial, cshost.Os, cshost.Agentversion, cshost.ProcessCount, cshost.ProcessClock, cshost.MemorySize, cshost.SwapMemory, ts, cshost.AgentID)
		tx.Commit()
	} else {
		// Not Exists
		d.SetHostinfo(cshost)
	}
}

func (d *DBHandler) SetHostinfo(new_agent *data.AgentHostAgentInfo) int {
	ts := time.Now().Unix()
	var max_count int
	err := d.db.QueryRow("SELECT max(_agentid) FROM agentinfo where _enabled=1").Scan(&max_count)
	if err != nil {
		max_count = 0
	}

	// Add Hostinfo
	new_agentid := max_count + 1
	d.hosts[new_agent.AgentID] = new_agentid

	// Insert Logic
	var agentinfo_arr data.AgentinfoArr
	new_data := &data.Agentinfo{
		Agentid:          new_agentid,
		Hostname:         new_agent.AgentID,
		Hostnameext:      new_agent.AgentName,
		Enabled:          1,
		Connected:        1,
		Updated:          1,
		Shorttermbasic:   2,
		Shorttermproc:    5,
		Shorttermio:      5,
		Shorttermcpu:     5,
		Longtermbasic:    600,
		Longtermproc:     600,
		Longtermio:       600,
		Longtermcpu:      600,
		Group:            "-",
		Ipaddress:        new_agent.Ip,
		Pscommand:        "-",
		Logevent:         "-",
		Processevent:     "-",
		Timecheck:        1,
		Disconnectedtime: ts,
		Skipdatatypes:    0,
		Virbasicperf:     1,
		Hypervisor:       0,
		Serviceevent:     "-",
		Installdate:      ts,
		Ibmpcrate:        0,
		Updatedtime:      ts,
		Os:               new_agent.Os,
		Fw:               "-",
		Agentversion:     new_agent.Agentversion,
		Model:            new_agent.Model,
		Serial:           new_agent.Serial,
		Processorcount:   new_agent.ProcessCount,
		Processorclock:   new_agent.ProcessClock,
		Memorysize:       new_agent.MemorySize,
		Swapsize:         new_agent.SwapMemory,
		Poolid:           -1,
		Replication:      0,
		Smt:              0,
		Micropar:         0,
		Capped:           0,
		Ec:               -1,
		Virtualcpu:       0,
		Weight:           0,
		Cpupool:          0,
		Ams:              0,
		Allip:            new_agent.Ip,
		Numanodecount:    0,
		Btime:            0,
	}
	agentinfo_arr.SetData(*new_data)

	tx := d.db.MustBegin()
	tx.MustExec(data.InsertAgentinfoUnnest, agentinfo_arr.GetArgs()...)
	tx.Commit()

	return new_agentid
}

func (d *DBHandler) SetPerf(csperf *data.AgentRealTimePerf) {
	fmt.Printf("realtimeperf update %s\n", csperf.AgentID)
	var agentid int

	// Check Agent
	if _, ok := d.hosts[csperf.AgentID]; ok {
		agentid = d.hosts[csperf.AgentID]
	} else {
		// Agent 정보보다 Perf 정보가 먼저 들어올 때 미존재하는 Agent면 Error 발생 가능성이 있으므로, Agent 정보를 먼저 구성해야 함
		// Agent 정보는 AgentID만 넘기고 나머지는 빈 정보로 구성
		// 어차피 Consumer Agent 정보가 들어오면 이미 존재하는 AgentID이므로 업데이트됨
		agentid = d.SetHostinfo(&data.AgentHostAgentInfo{
			AgentName:    csperf.AgentID,
			AgentID:      csperf.AgentID,
			Model:        "",
			Serial:       "",
			Ip:           "",
			Os:           "",
			Agentversion: "",
			ProcessCount: 0,
			ProcessClock: 0,
			MemorySize:   0,
			SwapMemory:   0,
		})
	}

	// Insert Perf
	fmt.Println(agentid)
	tablename, dbtype := d.GetTableinfo("realtimeperf")
	tablename2, _ := d.GetTableinfo("realtimecpu")

	if dbtype == "pg" {
		perf_data := data.RealtimeperfPg{}
		perf_data.SetData(csperf, agentid)

		cpu_data := data.RealtimecpuPg{}
		cpu_data.SetData(csperf, agentid)

		tx := d.db.MustBegin()
		var err error
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimePerf, tablename), perf_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimeCpu, tablename2), cpu_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		tx.Commit()
	} else {
		perf_data := data.RealtimeperfTs{}
		perf_data.SetData(csperf, agentid)

		cpu_data := data.RealtimecpuPg{}
		cpu_data.SetData(csperf, agentid)

		tx := d.db.MustBegin()
		var err error
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimePerf, tablename), perf_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimeCpu, tablename2), cpu_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		tx.Commit()
	}
}

func (d *DBHandler) SetPid(cspid *data.AgentRealTimePID) {
	fmt.Printf("realtimeperf update %s\n", cspid.AgentID)
	var agentid int

	// Check Agent
	if _, ok := d.hosts[cspid.AgentID]; ok {
		agentid = d.hosts[cspid.AgentID]
	} else {
		// Agent 정보보다 Perf 정보가 먼저 들어올 때 미존재하는 Agent면 Error 발생 가능성이 있으므로, Agent 정보를 먼저 구성해야 함
		// Agent 정보는 AgentID만 넘기고 나머지는 빈 정보로 구성
		// 어차피 Consumer Agent 정보가 들어오면 이미 존재하는 AgentID이므로 업데이트됨
		agentid = d.SetHostinfo(&data.AgentHostAgentInfo{
			AgentName:    cspid.AgentID,
			AgentID:      cspid.AgentID,
			Model:        "",
			Serial:       "",
			Ip:           "",
			Os:           "",
			Agentversion: "",
			ProcessCount: 0,
			ProcessClock: 0,
			MemorySize:   0,
			SwapMemory:   0,
		})
	}

	// Insert Perf
	fmt.Println(agentid)
	tablename, dbtype := d.GetTableinfo("realtimepid")
	tablename2, _ := d.GetTableinfo("realtimeproc")
	fmt.Println(tablename)

	if dbtype == "pg" {
		pid_data := data.RealtimepidPgArr{}
		proc_data := data.RealtimeprocPgArr{}
		for _, p := range cspid.PerfList {
			cmdid, userid, argid := d.GetProcid(&p)
			pid_data.SetData(p, cspid.Agenttime, agentid, cmdid, userid, argid)
			proc_data.SetData(p, cspid.Agenttime, agentid, cmdid, userid, argid)
		}

		tx := d.db.MustBegin()
		var err error
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimePidPg, tablename), pid_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimeProcPg, tablename2), proc_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		tx.Commit()
	} else {
		pid_data := data.RealtimepidTsArr{}
		proc_data := data.RealtimeprocTsArr{}
		for _, p := range cspid.PerfList {
			cmdid, userid, argid := d.GetProcid(&p)
			pid_data.SetData(p, cspid.Agenttime, agentid, cmdid, userid, argid)
			proc_data.SetData(p, cspid.Agenttime, agentid, cmdid, userid, argid)
		}

		var err error
		tx := d.db.MustBegin()
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimePidTs, tablename), pid_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		_, err = tx.Exec(fmt.Sprintf(data.InsertRealtimeProcTs, tablename2), proc_data.GetArgs()...)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
		tx.Commit()
	}
}

func (d *DBHandler) GetProcid(cspidinner *data.AgentRealTimePIDInner) (int, int, int) {
	var tablename string
	var err error

	var cmdid int
	tablename, _ = d.GetTableinfo("proccmd")
	err = d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Cmdname).Scan(&cmdid)
	if err != nil {
		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.InsertSimpleTable, tablename), cspidinner.Cmdname)
		tx.Commit()

		err := d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Cmdname).Scan(&cmdid)
		if err != nil {
			log.Fatal(err)
		}
	}

	var userid int
	tablename, _ = d.GetTableinfo("procuserid")
	err = d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Username).Scan(&userid)
	if err != nil {
		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.InsertSimpleTable, tablename), cspidinner.Username)
		tx.Commit()

		err := d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Username).Scan(&userid)
		if err != nil {
			log.Fatal(err)
		}
	}

	var argid int
	tablename, _ = d.GetTableinfo("procargid")
	err = d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Argname).Scan(&argid)
	if err != nil {
		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.InsertSimpleTable, tablename), cspidinner.Argname)
		tx.Commit()

		err := d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), cspidinner.Argname).Scan(&argid)
		if err != nil {
			log.Fatal(err)
		}
	}

	return cmdid, userid, argid
}

func (d *DBHandler) SetDisk(csperf *data.AgentRealTimeDisk) {

}

func (d *DBHandler) SetNet(csperf *data.AgentRealTimeNet) {

}

func (d *DBHandler) GetTableinfo(tablename string) (string, string) {
	var dbtype string
	if d.name[:2] == "pg" {
		dbtype = "pg"
	} else {
		dbtype = "ts"
	}

	for _, tb := range info_tables {
		if tablename == tb {
			return tablename, dbtype
		}
	}

	for _, tb := range metric_ref_tables {
		if tablename == tb {
			return d.GetTablename("metric_ref", tb), dbtype
		}
	}

	for _, tb := range metric_tables {
		if tablename == tb {
			return d.GetTablename("metric", tb), dbtype
		}
	}

	return "", ""
}

func (d *DBHandler) GetTablename(tableinfo string, tb string) string {
	if tableinfo == "metric_ref" {
		switch d.name {
		case "pg-hour", "pg-day":
			return tb + "_" + time.Now().Format("060102") + "00"
		case "pg-week":
			_, week := time.Now().ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			return tb + "_" + time.Now().Format("0601") + weekstr
		case "pg-month":
			return tb + "_" + time.Now().Format("0601")
		default:
			return tb
		}
	} else if tableinfo == "metric" {
		switch d.name {
		case "pg-hour":
			return tb + "_" + time.Now().Format("06010215")
		case "pg-day":
			return tb + "_" + time.Now().Format("060102") + "00"
		case "pg-week":
			_, week := time.Now().ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			return tb + "_" + time.Now().Format("0601") + weekstr
		case "pg-month":
			return tb + "_" + time.Now().Format("0601")
		default:
			return tb
		}
	} else {
		return ""
	}
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

func DBInit(dbinfo DbInfo, demoinfo DemoInfo, info *data.AgentinfoArr) *DBHandler {
	d := &DBHandler{
		name: dbinfo.Name,
		db:   DBConnection(dbinfo),
		demo: demoinfo,
	}
	d.CheckTable()
	d.DemoHostSetting(info)
	d.hosts = d.GetHostnames()

	return d
}
