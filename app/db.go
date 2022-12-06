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
	name            string
	db              *sqlx.DB
	demo            DemoInfo
	hosts           map[string]int
	ids             map[string]map[string]int
	metric_ref_flag string
	metric_flag     string
}

var (
	info_tables       = [...]string{"tableinfo", "agentinfo", "lastrealtimeperf", "deviceid", "descid"}
	metric_ref_tables = [...]string{"proccmd", "procuserid", "procargid"}
	metric_tables     = [...]string{"realtimeperf", "realtimecpu", "realtimedisk", "realtimenet", "realtimepid", "realtimeproc"}
)

func (d *DBHandler) CheckTable() {
	d.CheckTableInfo()
	d.CheckTableMetricref()
	d.CheckTableMetric()
}

func (d *DBHandler) CheckTableInfo() {
	// Info Table Check
	for _, tb := range info_tables {
		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM pg_tables where tablename=$1", tb).Scan(&exist_count)
		ErrorFatal(err)

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
}

func (d *DBHandler) CheckTableMetricref() {
	// Metric Reference Table Check
	for _, tb := range metric_ref_tables {
		tablename := d.GetTablename(tb)

		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM public.tableinfo where _tablename=$1", tablename).Scan(&exist_count)
		ErrorFatal(err)

		if exist_count == 0 {
			tx := d.db.MustBegin()
			tx.MustExec(fmt.Sprintf(data.ProcStmt, tablename))
			tx.MustExec(data.InsertTableinfo, tablename, time.Now().Unix())
			tx.Commit()
		}
	}
}

func (d *DBHandler) CheckTableMetric() {
	// Metric Table Check
	for _, tb := range metric_tables {
		tablename := d.GetTablename(tb)

		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM public.tableinfo where _tablename=$1", tablename).Scan(&exist_count)
		ErrorFatal(err)

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

func (d *DBHandler) CheckTableInterval() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		aftertime := time.Now().Add(1 * time.Second)
		metric_ref := d.GetTableFlag(aftertime, "metric_ref", metric_ref_tables[0])
		if metric_ref != d.metric_ref_flag {
			d.CheckTableMetricref()

			// 현재 테이블의 값을 이후 테이블로 복사
			tx := d.db.MustBegin()
			for _, s := range metric_ref_tables {
				_, err := tx.Exec(fmt.Sprintf(data.InsertPrevData, d.GetCustomTablename(aftertime, s), d.GetTablename(s)))
				if err != nil {
					d.CheckTableMetricref()
					tx.MustExec(fmt.Sprintf(data.InsertPrevData, d.GetCustomTablename(aftertime, s), d.GetTablename(s)))
				}
			}
			tx.Commit()
		}

		metric := d.GetTableFlag(time.Now().Add(1*time.Second), "metric", metric_tables[0])
		if metric != d.metric_flag {
			d.CheckTableMetric()
		}
	}
}

func (d *DBHandler) DemoHostSetting(arr *data.AgentinfoArr) {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM agentinfo where _enabled=1 and _hostname like 'Dummy%'").Scan(&exist_count)
	ErrorFatal(err)

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
	tx.MustExec(data.DeleteLastrealtimeperfAll, d.demo.HostCount)

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

func (d *DBHandler) GetNames(tb string) map[string]int {
	names := make(map[string]int, 0)
	tablename := d.GetTablename(tb)

	rows, err := d.db.Query(fmt.Sprintf("SELECT _id, _name FROM %s", tablename))
	if err != nil {
		log.Fatal()
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Query Error")
			return names
		}
		names[name] = id
	}

	return names
}

func (d *DBHandler) SetHost(cshost *data.AgentHostAgentInfo, agentinfo_arr *data.AgentinfoArr) {
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
	//fmt.Printf("agentinfo update %s\n", cshost.AgentID)

	if _, ok := d.hosts[cshost.AgentID]; ok {
		// Already Exists
		ts := time.Now().Unix()
		tx := d.db.MustBegin()
		tx.MustExec(data.UpdateAgentinfo, cshost.AgentName, cshost.Ip, cshost.Model, cshost.Serial, cshost.Os, cshost.Agentversion, cshost.ProcessCount, cshost.ProcessClock, cshost.MemorySize, cshost.SwapMemory, ts, cshost.AgentID)
		tx.Commit()

		agent_data := []data.Agentinfo{}
		d.db.Select(&agent_data, "SELECT * FROM agentinfo where _enabled=1 and _hostname=$1", cshost.AgentName)

		for _, a := range agent_data {
			agentinfo_arr.SetData(a)
		}
	} else {
		// Not Exists
		d.SetHostinfo(cshost, agentinfo_arr)
	}
}

func (d *DBHandler) SetHostinfo(new_agent *data.AgentHostAgentInfo, agentinfo_arr *data.AgentinfoArr) int {
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

func (d *DBHandler) SetEmptyAgentinfo(agentname string) int {
	var agentid int
	if _, ok := d.hosts[agentname]; ok {
		agentid = d.hosts[agentname]
	} else {
		// Agent 정보보다 Perf 정보가 먼저 들어올 때 미존재하는 Agent면 Error 발생 가능성이 있으므로, Agent 정보를 먼저 구성해야 함
		// Agent 정보는 AgentID만 넘기고 나머지는 빈 정보로 구성
		// 어차피 Consumer Agent 정보가 들어오면 이미 존재하는 AgentID이므로 업데이트됨
		agentinfo_arr := data.AgentinfoArr{}
		agentid = d.SetHostinfo(&data.AgentHostAgentInfo{
			AgentName:    agentname,
			AgentID:      agentname,
			Model:        "",
			Serial:       "",
			Ip:           "",
			Os:           "",
			Agentversion: "",
			ProcessCount: 0,
			ProcessClock: 0,
			MemorySize:   0,
			SwapMemory:   0,
		}, &agentinfo_arr)
	}

	return agentid
}

func (d *DBHandler) SetPerf(csperf *data.AgentRealTimePerf, agentid int, tables ...data.Table) {
	// Set Lastrealtimeperf
	for _, t := range tables {
		t.SetData(csperf, agentid)
	}
}

func (d *DBHandler) InsertPerf(agentid int, tables ...data.Table) {
	var err error
	tx := d.db.MustBegin()

	for _, t := range tables {
		var tablename string
		switch t.(type) {
		case *data.Lastrealtimeperf:
			tablename = d.GetTablename("lastrealtimeperf")
			tx.MustExec(data.DeleteLastrealtimeperf, agentid)
		case *data.RealtimeperfPg, *data.RealtimeperfTs:
			tablename = d.GetTablename("realtimeperf")
		case *data.RealtimecpuPg, *data.RealtimecpuTs:
			tablename = d.GetTablename("realtimecpu")
		}

		_, err = tx.Exec(t.GetInsertStmt(tablename), t.GetArgs()...)
		ErrorTx(err, tx)
	}
	tx.Commit()
}

func (d *DBHandler) SetPid(cspid *data.AgentRealTimePID, agentid int, tables ...data.TableArr) {
	for _, p := range cspid.PerfList {
		cmdid, userid, argid := d.GetProcId(&p)
		for _, t := range tables {
			t.SetData(p, cspid.Agenttime, agentid, cmdid, userid, argid)
		}
	}
}

func (d *DBHandler) SetDisk(csdisk *data.AgentRealTimeDisk, agentid int, tables ...data.TableArr) {
	for _, p := range csdisk.PerfList {
		ionameid, descid := d.GetDeviceId(p.Ioname, p.Descname)
		for _, t := range tables {
			t.SetData(p, csdisk.Agenttime, agentid, ionameid, descid)
		}
	}
}

func (d *DBHandler) SetNet(csnet *data.AgentRealTimeNet, agentid int, tables ...data.TableArr) {
	for _, p := range csnet.PerfList {
		ionameid, _ := d.GetDeviceId(p.Ioname, "")
		for _, t := range tables {
			t.SetData(p, csnet.Agenttime, agentid, ionameid)
		}
	}
}

func (d *DBHandler) InsertTableArr(tables ...data.TableArr) {
	var err error
	tx := d.db.MustBegin()

	for _, t := range tables {
		var tablename string
		switch t.(type) {
		case *data.RealtimepidPgArr, *data.RealtimepidTsArr:
			tablename = d.GetTablename("realtimepid")
		case *data.RealtimeprocPgArr, *data.RealtimeprocTsArr:
			tablename = d.GetTablename("realtimeproc")
		case *data.RealtimediskPgArr, *data.RealtimediskTsArr:
			tablename = d.GetTablename("realtimedisk")
		case *data.RealtimenetPgArr, *data.RealtimenetTsArr:
			tablename = d.GetTablename("realtimenet")
		}

		_, err = tx.Exec(t.GetInsertStmt(tablename), t.GetArgs()...)
		ErrorTx(err, tx)
	}
	tx.Commit()
}

func (d *DBHandler) GetProcId(cspidinner *data.AgentRealTimePIDInner) (int, int, int) {
	cmdid := d.GetId(cspidinner.Cmdname, "proccmd")
	userid := d.GetId(cspidinner.Username, "procuserid")
	argid := d.GetId(cspidinner.Argname, "procargid")

	return cmdid, userid, argid
}

func (d *DBHandler) GetDeviceId(ioname string, descname string) (int, int) {
	ioid := d.GetId(ioname, "deviceid")
	descid := d.GetId(descname, "descid")

	return ioid, descid
}

func (d *DBHandler) GetId(name string, flag string) int {
	var id int
	if _, ok := d.ids[flag][name]; ok {
		id = d.ids[flag][name]
	} else {
		tablename := d.GetTablename(flag)

		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.InsertSimpleTable, tablename), name)
		tx.Commit()

		err := d.db.QueryRow(fmt.Sprintf("SELECT _id FROM %s where _name=$1", tablename), name).Scan(&id)
		ErrorFatal(err)

		d.ids[flag][name] = id
	}

	return id
}

func (d *DBHandler) GetTabletype(tablename string) string {
	if d.name[:2] == "pg" {
		return "pg"
	} else {
		return "ts"
	}
}

func (d *DBHandler) GetTablename(tablename string) string {
	for _, tb := range info_tables {
		if tablename == tb {
			return tablename
		}
	}

	for _, tb := range metric_ref_tables {
		if tablename == tb {
			flag := d.GetTableFlag(time.Now(), "metric_ref", tb)

			if flag == "" {
				return tb
			} else {
				return tb + "_" + flag
			}
		}
	}

	for _, tb := range metric_tables {
		if tablename == tb {
			flag := d.GetTableFlag(time.Now(), "metric", tb)

			if flag == "" {
				return tb
			} else {
				return tb + "_" + flag
			}
		}
	}

	return ""
}

func (d *DBHandler) GetCustomTablename(timevalue time.Time, tb string) string {
	flag := d.GetTableFlag(timevalue, "metric_ref", tb)

	if flag == "" {
		return tb
	} else {
		return tb + "_" + flag
	}
}

func (d *DBHandler) GetTableFlag(timevalue time.Time, tableinfo string, tb string) string {
	if tableinfo == "metric_ref" {
		switch d.name {
		case "pg-hour", "pg-day":
			return timevalue.Format("060102") + "00"
		case "pg-week":
			_, week := timevalue.ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			return timevalue.Format("0601") + weekstr
		case "pg-month":
			return timevalue.Format("0601")
		default:
			return ""
		}
	} else if tableinfo == "metric" {
		switch d.name {
		case "pg-hour":
			return timevalue.Format("06010215")
		case "pg-day":
			return timevalue.Format("060102") + "00"
		case "pg-week":
			_, week := timevalue.ISOWeek()
			weekstr := strconv.Itoa(week)
			if week < 10 {
				weekstr = "0" + weekstr
			}

			return timevalue.Format("0601") + weekstr
		case "pg-month":
			return timevalue.Format("0601")
		default:
			return ""
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
	ErrorFatal(err)
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
	d.metric_ref_flag = d.GetTableFlag(time.Now(), "metric_ref", metric_ref_tables[0])
	d.metric_flag = d.GetTableFlag(time.Now(), "metric", metric_tables[0])

	d.hosts = d.GetHostnames()
	d.ids = make(map[string]map[string]int)
	d.ids["proccmd"] = d.GetNames("proccmd")
	d.ids["procuserid"] = d.GetNames("procuserid")
	d.ids["procargid"] = d.GetNames("procargid")
	d.ids["deviceid"] = d.GetNames("deviceid")
	d.ids["descid"] = d.GetNames("descid")

	go d.CheckTableInterval()

	return d
}
