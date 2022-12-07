package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Table interface {
	SetData(data interface{}, agentid int)
	GetArgs() []interface{}
	GetString() string
	GetInsertStmt(tablename string) string
}

type TableSetArr interface {
	SetData(data interface{}, agenttime time.Time, ids ...int)
}

type TableGet interface {
	GetInsertStmt(tablename string) string
	GetArgs() []interface{}
	GetArrString() []string
}

type Tableinfo struct {
	Tablename   string `db:"_tablename"`
	Version     int    `db:"_version"`
	Createdtime int    `db:"_createdtime"`
	Updatetime  int    `db:"_updatetime"`
	Durationmin int    `db:"_durationmin"`
}

type Agentinfo struct {
	Agentid          int    `db:"_agentid"`
	Hostname         string `db:"_hostname"`
	Hostnameext      string `db:"_hostnameext"`
	Enabled          int    `db:"_enabled"`
	Connected        int    `db:"_connected"`
	Updated          int    `db:"_updated"`
	Shorttermbasic   int    `db:"_shorttermbasic"`
	Shorttermproc    int    `db:"_shorttermproc"`
	Shorttermio      int    `db:"_shorttermio"`
	Shorttermcpu     int    `db:"_shorttermcpu"`
	Longtermbasic    int    `db:"_longtermbasic"`
	Longtermproc     int    `db:"_longtermproc"`
	Longtermio       int    `db:"_longtermio"`
	Longtermcpu      int    `db:"_longtermcpu"`
	Group            string `db:"_group"`
	Ipaddress        string `db:"_ipaddress"`
	Pscommand        string `db:"_pscommand"`
	Logevent         string `db:"_logevent"`
	Processevent     string `db:"_processevent"`
	Timecheck        int    `db:"_timecheck"`
	Disconnectedtime int64  `db:"_disconnectedtime"`
	Skipdatatypes    int    `db:"_skipdatatypes"`
	Virbasicperf     int    `db:"_virbasicperf"`
	Hypervisor       int    `db:"_hypervisor"`
	Serviceevent     string `db:"_serviceevent"`
	Installdate      int64  `db:"_installdate"`
	Ibmpcrate        int    `db:"_ibmpcrate"`
	Updatedtime      int64  `db:"_updatedtime"`
	Os               string `db:"_os"`
	Fw               string `db:"_fw"`
	Agentversion     string `db:"_agentversion"`
	Model            string `db:"_model"`
	Serial           string `db:"_serial"`
	Processorcount   int    `db:"_processorcount"`
	Processorclock   int    `db:"_processorclock"`
	Memorysize       int    `db:"_memorysize"`
	Swapsize         int    `db:"_swapsize"`
	Poolid           int    `db:"_poolid"`
	Replication      int    `db:"_replication"`
	Smt              int    `db:"_smt"`
	Micropar         int    `db:"_micropar"`
	Capped           int    `db:"_capped"`
	Ec               int    `db:"_ec"`
	Virtualcpu       int    `db:"_virtualcpu"`
	Weight           int    `db:"_weight"`
	Cpupool          int    `db:"_cpupool"`
	Ams              int    `db:"_ams"`
	Allip            string `db:"_allip"`
	Numanodecount    int    `db:"_numanodecount"`
	Btime            int64  `db:"_btime"`
}

type Lastrealtimeperf struct {
	Ontunetime    int64  `db:"_ontunetime"`
	Agentid       int    `db:"_agentid"`
	Hostname      string `db:"_hostname"`
	User          int    `db:"_user"`
	Sys           int    `db:"_sys"`
	Wait          int    `db:"_wait"`
	Idle          int    `db:"_idle"`
	Memoryused    int    `db:"_memoryused"`
	Filecache     int    `db:"_filecache"`
	Memorysize    int    `db:"_memorysize"`
	Avm           int    `db:"_avm"`
	Swapused      int    `db:"_swapused"`
	Swapsize      int    `db:"_swapsize"`
	Diskiorate    int    `db:"_diskiorate"`
	Networkiorate int    `db:"_networkiorate"`
	Topproc       string `db:"_topproc"`
	Topuser       string `db:"_topuser"`
	Topproccount  int    `db:"_topproccount"`
	Topcpu        int    `db:"_topcpu"`
	Topdisk       string `db:"_topdisk"`
	Topvg         string `db:"_topvg"`
	Topbusy       int    `db:"_topbusy"`
	Maxcpu        int    `db:"_maxcpu"`
	Maxmem        int    `db:"_maxmem"`
	Maxswap       int    `db:"_maxswap"`
	Maxdisk       int    `db:"_maxdisk"`
	Diskiops      int    `db:"_diskiops"`
	Networkiops   int    `db:"_networkiops"`
	Dummy01       int    `db:"_dummy01"`
	Dummy02       int    `db:"_dummy02"`
	Dummy03       int    `db:"_dummy03"`
	Dummy04       int    `db:"_dummy04"`
	Dummy05       int    `db:"_dummy05"`
	Dummy06       int    `db:"_dummy06"`
	Dummy07       int    `db:"_dummy07"`
	Dummy08       int    `db:"_dummy08"`
	Dummy09       int    `db:"_dummy09"`
	Dummy10       int    `db:"_dummy10"`
	Dummy11       int    `db:"_dummy11"`
	Dummy12       int    `db:"_dummy12"`
	Dummy13       int    `db:"_dummy13"`
	Dummy14       int    `db:"_dummy14"`
	Dummy15       int    `db:"_dummy15"`
	Dummy16       int    `db:"_dummy16"`
	Dummy17       int    `db:"_dummy17"`
	Dummy18       int    `db:"_dummy18"`
	Dummy19       int    `db:"_dummy19"`
	Dummy20       int    `db:"_dummy20"`
	Dummy21       int    `db:"_dummy21"`
	Dummy22       int    `db:"_dummy22"`
	Dummy23       int    `db:"_dummy23"`
	Dummy24       int    `db:"_dummy24"`
	Dummy25       int    `db:"_dummy25"`
	Dummy26       int    `db:"_dummy26"`
	Dummy27       int    `db:"_dummy27"`
	Dummy28       int    `db:"_dummy28"`
	Dummy29       int    `db:"_dummy29"`
	Dummy30       int    `db:"_dummy30"`
}

func (r *Lastrealtimeperf) SetData(data interface{}, agentid int) {
	d := data.(*AgentRealTimePerf)
	r.Ontunetime = d.Agenttime.Unix()
	r.Agentid = agentid
	r.Hostname = d.AgentID
	r.User = d.User
	r.Sys = d.Sys
	r.Wait = d.Wait
	r.Idle = d.Idle
	r.Memoryused = d.MemoryUsed
	r.Filecache = d.MemoryCache
	r.Memorysize = d.MemorySize
	r.Avm = d.Avm
	r.Swapused = d.SwapUsed
	r.Swapsize = d.SwapSize
	r.Diskiorate = d.DiskReadWrite
	r.Networkiorate = d.NetworkReadWrite
	r.Topproc = ""
	r.Topuser = ""
	r.Topproccount = 0
	r.Topcpu = 0
	r.Topdisk = ""
	r.Topbusy = 0
	r.Maxcpu = 0
	r.Maxmem = 0
	r.Maxswap = 0
	r.Maxdisk = 0
	r.Diskiops = d.DiskIOPS
	r.Networkiops = 0
}

func (r *Lastrealtimeperf) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, r.Ontunetime)
	data = append(data, r.Agentid)
	data = append(data, r.Hostname)
	data = append(data, r.User)
	data = append(data, r.Sys)
	data = append(data, r.Wait)
	data = append(data, r.Idle)
	data = append(data, r.Memoryused)
	data = append(data, r.Filecache)
	data = append(data, r.Memorysize)
	data = append(data, r.Avm)
	data = append(data, r.Swapused)
	data = append(data, r.Swapsize)
	data = append(data, r.Diskiorate)
	data = append(data, r.Networkiorate)
	data = append(data, r.Topproc)
	data = append(data, r.Topuser)
	data = append(data, r.Topproccount)
	data = append(data, r.Topcpu)
	data = append(data, r.Topdisk)
	data = append(data, r.Topbusy)
	data = append(data, r.Maxcpu)
	data = append(data, r.Maxmem)
	data = append(data, r.Maxswap)
	data = append(data, r.Maxdisk)
	data = append(data, r.Diskiops)
	data = append(data, r.Networkiops)

	return data
}

func (r *Lastrealtimeperf) GetString() string {
	data := r.GetArgs()
	str_data := make([]string, 0)
	for _, d := range data {
		str_data = append(str_data, fmt.Sprintf("%v", d))
	}

	return strings.Join(str_data, ",")
}

func (r *Lastrealtimeperf) GetInsertStmt(tablename string) string {
	return InsertLastRealtimePerf
}

type RealtimeperfPg struct {
	Ontunetime       int64 `db:"_ontunetime"`
	Agenttime        int   `db:"_agenttime"`
	Agentid          int   `db:"_agentid"`
	User             int   `db:"_user"`
	Sys              int   `db:"_sys"`
	Wait             int   `db:"_wait"`
	Idle             int   `db:"_idle"`
	Processorcount   int   `db:"_processorcount"`
	Runqueue         int   `db:"_runqueue"`
	Blockqueue       int   `db:"_blockqueue"`
	Waitqueue        int   `db:"_waitqueue"`
	Pqueue           int   `db:"_pqueue"`
	Pcrateuser       int   `db:"_pcrateuser"`
	Pcratesys        int   `db:"_pcratesys"`
	Memorysize       int   `db:"_memorysize"`
	Memoryused       int   `db:"_memoryused"`
	Memorypinned     int   `db:"_memorypinned"`
	Memorysys        int   `db:"_memorysys"`
	Memoryuser       int   `db:"_memoryuser"`
	Memorycache      int   `db:"_memorycache"`
	Avm              int   `db:"_avm"`
	Pagingspacein    int   `db:"_pagingspacein"`
	Pagingspaceout   int   `db:"_pagingspaceout"`
	Filesystemin     int   `db:"_filesystemin"`
	Filesystemout    int   `db:"_filesystemout"`
	Memoryscan       int   `db:"_memoryscan"`
	Memoryfreed      int   `db:"_memoryfreed"`
	Swapsize         int   `db:"_swapsize"`
	Swapused         int   `db:"_swapused"`
	Swapactive       int   `db:"_swapactive"`
	Fork             int   `db:"_fork"`
	Exec             int   `db:"_exec"`
	Interupt         int   `db:"_interupt"`
	Systemcall       int   `db:"_systemcall"`
	Contextswitch    int   `db:"_contextswitch"`
	Semaphore        int   `db:"_semaphore"`
	Msg              int   `db:"_msg"`
	Diskreadwrite    int   `db:"_diskreadwrite"`
	Diskiops         int   `db:"_diskiops"`
	Networkreadwrite int   `db:"_networkreadwrite"`
	Networkiops      int   `db:"_networkiops"`
	Topcommandid     int   `db:"_topcommandid"`
	Topcommandcount  int   `db:"_topcommandcount"`
	Topuserid        int   `db:"_topuserid"`
	Topcpu           int   `db:"_topcpu"`
	Topdiskid        int   `db:"_topdiskid"`
	Topvgid          int   `db:"_topvgid"`
	Topbusy          int   `db:"_topbusy"`
	Maxpid           int   `db:"_maxpid"`
	Threadcount      int   `db:"_threadcount"`
	Pidcount         int   `db:"_pidcount"`
	Linuxbuffer      int   `db:"_linuxbuffer"`
	Linuxcached      int   `db:"_linuxcached"`
	Linuxsrec        int   `db:"_linuxsrec"`
	Memused_Mb       int   `db:"_memused_Mb"`
	Irq              int   `db:"_irq"`
	Softirq          int   `db:"_softirq"`
	Swapused_Mb      int   `db:"_swapused_Mb"`
	Dusm             int   `db:"_dusm"`
}

func (r *RealtimeperfPg) SetData(data interface{}, agentid int) {
	d := data.(*AgentRealTimePerf)
	r.Ontunetime = d.Agenttime.Unix()
	r.Agenttime = int(d.Agenttime.Unix())
	r.Agentid = agentid
	r.User = d.User
	r.Sys = d.Sys
	r.Wait = d.Wait
	r.Idle = d.Idle
	r.Processorcount = d.ProcessorCount
	r.Runqueue = d.RunQueue
	r.Blockqueue = d.BlockQueue
	r.Waitqueue = d.WaitQueue
	r.Pqueue = d.PQueue
	r.Pcrateuser = d.PCRateUser
	r.Pcratesys = d.PCRateSys
	r.Memorysize = d.MemorySize
	r.Memoryused = d.MemoryUsed
	r.Memorypinned = d.MemoryPinned
	r.Memorysys = d.MemorySys
	r.Memoryuser = d.MemoryUser
	r.Memorycache = d.MemoryCache
	r.Avm = d.Avm
	r.Pagingspacein = d.PagingspaceIn
	r.Pagingspaceout = d.PaingSpaceOut
	r.Filesystemin = d.FileSystemIn
	r.Filesystemout = d.FileSystmeOut
	r.Memoryscan = d.MemoryScan
	r.Memoryfreed = d.MemoryFreed
	r.Swapsize = d.SwapSize
	r.Swapused = d.SwapUsed
	r.Swapactive = d.SwapActive
	r.Fork = d.Fork
	r.Exec = d.Exec
	r.Interupt = d.Interupt
	r.Systemcall = d.SystemCall
	r.Contextswitch = d.ContextSwitch
	r.Semaphore = d.Semaphore
	r.Msg = d.Msg
	r.Diskreadwrite = d.DiskReadWrite
	r.Diskiops = d.DiskIOPS
	r.Networkreadwrite = d.NetworkReadWrite
	r.Networkiops = 0
	r.Topcommandid = d.TopCommandID
	r.Topcommandcount = d.TopCommandCount
	r.Topuserid = d.TopUserID
	r.Topcpu = d.TopCPU
	r.Topdiskid = d.TopDiskID
	r.Topvgid = d.TopvgID
	r.Topbusy = d.TOPBusy
	r.Maxpid = d.MaxPID
	r.Threadcount = d.ThreadCount
	r.Pidcount = d.PIDCount
	r.Linuxbuffer = d.LinuxBuffer
	r.Linuxcached = d.LinuxCached
	r.Linuxsrec = d.Linuxsrec
	r.Memused_Mb = d.Memused_mb
	r.Irq = d.IRQ
	r.Softirq = d.SoftIRQ
	r.Swapused_Mb = d.Swapused_MB
	r.Dusm = d.DUSM
}

func (r *RealtimeperfPg) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, r.Ontunetime)
	data = append(data, r.Agenttime)
	data = append(data, r.Agentid)
	data = append(data, r.User)
	data = append(data, r.Sys)
	data = append(data, r.Wait)
	data = append(data, r.Idle)
	data = append(data, r.Processorcount)
	data = append(data, r.Runqueue)
	data = append(data, r.Blockqueue)
	data = append(data, r.Waitqueue)
	data = append(data, r.Pqueue)
	data = append(data, r.Pcrateuser)
	data = append(data, r.Pcratesys)
	data = append(data, r.Memorysize)
	data = append(data, r.Memoryused)
	data = append(data, r.Memorypinned)
	data = append(data, r.Memorysys)
	data = append(data, r.Memoryuser)
	data = append(data, r.Memorycache)
	data = append(data, r.Avm)
	data = append(data, r.Pagingspacein)
	data = append(data, r.Pagingspaceout)
	data = append(data, r.Filesystemin)
	data = append(data, r.Filesystemout)
	data = append(data, r.Memoryscan)
	data = append(data, r.Memoryfreed)
	data = append(data, r.Swapsize)
	data = append(data, r.Swapused)
	data = append(data, r.Swapactive)
	data = append(data, r.Fork)
	data = append(data, r.Exec)
	data = append(data, r.Interupt)
	data = append(data, r.Systemcall)
	data = append(data, r.Contextswitch)
	data = append(data, r.Semaphore)
	data = append(data, r.Msg)
	data = append(data, r.Diskreadwrite)
	data = append(data, r.Diskiops)
	data = append(data, r.Networkreadwrite)
	data = append(data, r.Networkiops)
	data = append(data, r.Topcommandid)
	data = append(data, r.Topcommandcount)
	data = append(data, r.Topuserid)
	data = append(data, r.Topcpu)
	data = append(data, r.Topdiskid)
	data = append(data, r.Topvgid)
	data = append(data, r.Topbusy)
	data = append(data, r.Maxpid)
	data = append(data, r.Threadcount)
	data = append(data, r.Pidcount)
	data = append(data, r.Linuxbuffer)
	data = append(data, r.Linuxcached)
	data = append(data, r.Linuxsrec)
	data = append(data, r.Memused_Mb)
	data = append(data, r.Irq)
	data = append(data, r.Softirq)
	data = append(data, r.Swapused_Mb)
	data = append(data, r.Dusm)

	return data
}

func (r *RealtimeperfPg) GetString() string {
	data := r.GetArgs()
	str_data := make([]string, 0)
	for _, d := range data {
		str_data = append(str_data, fmt.Sprintf("%v", d))
	}

	return strings.Join(str_data, ",")
}

func (r *RealtimeperfPg) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePerf, tablename)
}

type RealtimeperfTs struct {
	Ontunetime       time.Time `db:"_ontunetime"`
	Agenttime        int       `db:"_agenttime"`
	Agentid          int       `db:"_agentid"`
	User             int       `db:"_user"`
	Sys              int       `db:"_sys"`
	Wait             int       `db:"_wait"`
	Idle             int       `db:"_idle"`
	Processorcount   int       `db:"_processorcount"`
	Runqueue         int       `db:"_runqueue"`
	Blockqueue       int       `db:"_blockqueue"`
	Waitqueue        int       `db:"_waitqueue"`
	Pqueue           int       `db:"_pqueue"`
	Pcrateuser       int       `db:"_pcrateuser"`
	Pcratesys        int       `db:"_pcratesys"`
	Memorysize       int       `db:"_memorysize"`
	Memoryused       int       `db:"_memoryused"`
	Memorypinned     int       `db:"_memorypinned"`
	Memorysys        int       `db:"_memorysys"`
	Memoryuser       int       `db:"_memoryuser"`
	Memorycache      int       `db:"_memorycache"`
	Avm              int       `db:"_avm"`
	Pagingspacein    int       `db:"_pagingspacein"`
	Pagingspaceout   int       `db:"_pagingspaceout"`
	Filesystemin     int       `db:"_filesystemin"`
	Filesystemout    int       `db:"_filesystemout"`
	Memoryscan       int       `db:"_memoryscan"`
	Memoryfreed      int       `db:"_memoryfreed"`
	Swapsize         int       `db:"_swapsize"`
	Swapused         int       `db:"_swapused"`
	Swapactive       int       `db:"_swapactive"`
	Fork             int       `db:"_fork"`
	Exec             int       `db:"_exec"`
	Interupt         int       `db:"_interupt"`
	Systemcall       int       `db:"_systemcall"`
	Contextswitch    int       `db:"_contextswitch"`
	Semaphore        int       `db:"_semaphore"`
	Msg              int       `db:"_msg"`
	Diskreadwrite    int       `db:"_diskreadwrite"`
	Diskiops         int       `db:"_diskiops"`
	Networkreadwrite int       `db:"_networkreadwrite"`
	Networkiops      int       `db:"_networkiops"`
	Topcommandid     int       `db:"_topcommandid"`
	Topcommandcount  int       `db:"_topcommandcount"`
	Topuserid        int       `db:"_topuserid"`
	Topcpu           int       `db:"_topcpu"`
	Topdiskid        int       `db:"_topdiskid"`
	Topvgid          int       `db:"_topvgid"`
	Topbusy          int       `db:"_topbusy"`
	Maxpid           int       `db:"_maxpid"`
	Threadcount      int       `db:"_threadcount"`
	Pidcount         int       `db:"_pidcount"`
	Linuxbuffer      int       `db:"_linuxbuffer"`
	Linuxcached      int       `db:"_linuxcached"`
	Linuxsrec        int       `db:"_linuxsrec"`
	Memused_Mb       int       `db:"_memused_Mb"`
	Irq              int       `db:"_irq"`
	Softirq          int       `db:"_softirq"`
	Swapused_Mb      int       `db:"_swapused_Mb"`
	Dusm             int       `db:"_dusm"`
}

func (r *RealtimeperfTs) SetData(data interface{}, agentid int) {
	d := data.(*AgentRealTimePerf)
	r.Ontunetime = d.Agenttime
	r.Agenttime = int(d.Agenttime.Unix())
	r.Agentid = agentid
	r.User = d.User
	r.Sys = d.Sys
	r.Wait = d.Wait
	r.Idle = d.Idle
	r.Processorcount = d.ProcessorCount
	r.Runqueue = d.RunQueue
	r.Blockqueue = d.BlockQueue
	r.Waitqueue = d.WaitQueue
	r.Pqueue = d.PQueue
	r.Pcrateuser = d.PCRateUser
	r.Pcratesys = d.PCRateSys
	r.Memorysize = d.MemorySize
	r.Memoryused = d.MemoryUsed
	r.Memorypinned = d.MemoryPinned
	r.Memorysys = d.MemorySys
	r.Memoryuser = d.MemoryUser
	r.Memorycache = d.MemoryCache
	r.Avm = d.Avm
	r.Pagingspacein = d.PagingspaceIn
	r.Pagingspaceout = d.PaingSpaceOut
	r.Filesystemin = d.FileSystemIn
	r.Filesystemout = d.FileSystmeOut
	r.Memoryscan = d.MemoryScan
	r.Memoryfreed = d.MemoryFreed
	r.Swapsize = d.SwapSize
	r.Swapused = d.SwapUsed
	r.Swapactive = d.SwapActive
	r.Fork = d.Fork
	r.Exec = d.Exec
	r.Interupt = d.Interupt
	r.Systemcall = d.SystemCall
	r.Contextswitch = d.ContextSwitch
	r.Semaphore = d.Semaphore
	r.Msg = d.Msg
	r.Diskreadwrite = d.DiskReadWrite
	r.Diskiops = d.DiskIOPS
	r.Networkreadwrite = d.NetworkReadWrite
	r.Networkiops = 0
	r.Topcommandid = d.TopCommandID
	r.Topcommandcount = d.TopCommandCount
	r.Topuserid = d.TopUserID
	r.Topcpu = d.TopCPU
	r.Topdiskid = d.TopDiskID
	r.Topvgid = d.TopvgID
	r.Topbusy = d.TOPBusy
	r.Maxpid = d.MaxPID
	r.Threadcount = d.ThreadCount
	r.Pidcount = d.PIDCount
	r.Linuxbuffer = d.LinuxBuffer
	r.Linuxcached = d.LinuxCached
	r.Linuxsrec = d.Linuxsrec
	r.Memused_Mb = d.Memused_mb
	r.Irq = d.IRQ
	r.Softirq = d.SoftIRQ
	r.Swapused_Mb = d.Swapused_MB
	r.Dusm = d.DUSM
}

func (r *RealtimeperfTs) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, r.Ontunetime)
	data = append(data, r.Agenttime)
	data = append(data, r.Agentid)
	data = append(data, r.User)
	data = append(data, r.Sys)
	data = append(data, r.Wait)
	data = append(data, r.Idle)
	data = append(data, r.Processorcount)
	data = append(data, r.Runqueue)
	data = append(data, r.Blockqueue)
	data = append(data, r.Waitqueue)
	data = append(data, r.Pqueue)
	data = append(data, r.Pcrateuser)
	data = append(data, r.Pcratesys)
	data = append(data, r.Memorysize)
	data = append(data, r.Memoryused)
	data = append(data, r.Memorypinned)
	data = append(data, r.Memorysys)
	data = append(data, r.Memoryuser)
	data = append(data, r.Memorycache)
	data = append(data, r.Avm)
	data = append(data, r.Pagingspacein)
	data = append(data, r.Pagingspaceout)
	data = append(data, r.Filesystemin)
	data = append(data, r.Filesystemout)
	data = append(data, r.Memoryscan)
	data = append(data, r.Memoryfreed)
	data = append(data, r.Swapsize)
	data = append(data, r.Swapused)
	data = append(data, r.Swapactive)
	data = append(data, r.Fork)
	data = append(data, r.Exec)
	data = append(data, r.Interupt)
	data = append(data, r.Systemcall)
	data = append(data, r.Contextswitch)
	data = append(data, r.Semaphore)
	data = append(data, r.Msg)
	data = append(data, r.Diskreadwrite)
	data = append(data, r.Diskiops)
	data = append(data, r.Networkreadwrite)
	data = append(data, r.Networkiops)
	data = append(data, r.Topcommandid)
	data = append(data, r.Topcommandcount)
	data = append(data, r.Topuserid)
	data = append(data, r.Topcpu)
	data = append(data, r.Topdiskid)
	data = append(data, r.Topvgid)
	data = append(data, r.Topbusy)
	data = append(data, r.Maxpid)
	data = append(data, r.Threadcount)
	data = append(data, r.Pidcount)
	data = append(data, r.Linuxbuffer)
	data = append(data, r.Linuxcached)
	data = append(data, r.Linuxsrec)
	data = append(data, r.Memused_Mb)
	data = append(data, r.Irq)
	data = append(data, r.Softirq)
	data = append(data, r.Swapused_Mb)
	data = append(data, r.Dusm)

	return data
}

func (r *RealtimeperfTs) GetString() string {
	data := r.GetArgs()
	str_data := make([]string, 0)
	for _, d := range data {
		str_data = append(str_data, fmt.Sprintf("%v", d))
	}

	return strings.Join(str_data, ",")
}

func (r *RealtimeperfTs) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePerf, tablename)
}

type RealtimecpuPg struct {
	Ontunetime    int64 `db:"_ontunetime"`
	Agenttime     int   `db:"_agenttime"`
	Agentid       int   `db:"_agentid"`
	Index         int   `db:"_index"`
	User          int   `db:"_user"`
	Sys           int   `db:"_sys"`
	Wait          int   `db:"_wait"`
	Idle          int   `db:"_idle"`
	Runqueue      int   `db:"_runqueue"`
	Fork          int   `db:"_fork"`
	Exec          int   `db:"_exec"`
	Interupt      int   `db:"_interupt"`
	Systemcall    int   `db:"_systemcall"`
	Contextswitch int   `db:"_contextswitch"`
}

func (r *RealtimecpuPg) SetData(data interface{}, agentid int) {
	d := data.(*AgentRealTimePerf)
	r.Ontunetime = d.Agenttime.Unix()
	r.Agenttime = int(d.Agenttime.Unix())
	r.Agentid = agentid
	r.Index = 0
	r.User = d.User
	r.Sys = d.Sys
	r.Wait = d.Wait
	r.Idle = d.Idle
	r.Runqueue = d.RunQueue
	r.Fork = d.Fork
	r.Exec = d.Exec
	r.Interupt = d.Interupt
	r.Systemcall = d.SystemCall
	r.Contextswitch = d.ContextSwitch
}

func (r *RealtimecpuPg) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, r.Ontunetime)
	data = append(data, r.Agenttime)
	data = append(data, r.Agentid)
	data = append(data, r.Index)
	data = append(data, r.User)
	data = append(data, r.Sys)
	data = append(data, r.Wait)
	data = append(data, r.Idle)
	data = append(data, r.Runqueue)
	data = append(data, r.Fork)
	data = append(data, r.Exec)
	data = append(data, r.Interupt)
	data = append(data, r.Systemcall)
	data = append(data, r.Contextswitch)

	return data
}

func (r *RealtimecpuPg) GetString() string {
	data := r.GetArgs()
	str_data := make([]string, 0)
	for _, d := range data {
		str_data = append(str_data, fmt.Sprintf("%v", d))
	}

	return strings.Join(str_data, ",")
}

func (r *RealtimecpuPg) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeCpu, tablename)
}

type RealtimecpuTs struct {
	Ontunetime    time.Time `db:"_ontunetime"`
	Agenttime     int       `db:"_agenttime"`
	Agentid       int       `db:"_agentid"`
	Index         int       `db:"_index"`
	User          int       `db:"_user"`
	Sys           int       `db:"_sys"`
	Wait          int       `db:"_wait"`
	Idle          int       `db:"_idle"`
	Runqueue      int       `db:"_runqueue"`
	Fork          int       `db:"_fork"`
	Exec          int       `db:"_exec"`
	Interupt      int       `db:"_interupt"`
	Systemcall    int       `db:"_systemcall"`
	Contextswitch int       `db:"_contextswitch"`
}

func (r *RealtimecpuTs) SetData(data interface{}, agentid int) {
	d := data.(*AgentRealTimePerf)
	r.Ontunetime = d.Agenttime
	r.Agenttime = int(d.Agenttime.Unix())
	r.Agentid = agentid
	r.Index = 0
	r.User = d.User
	r.Sys = d.Sys
	r.Wait = d.Wait
	r.Idle = d.Idle
	r.Runqueue = d.RunQueue
	r.Fork = d.Fork
	r.Exec = d.Exec
	r.Interupt = d.Interupt
	r.Systemcall = d.SystemCall
	r.Contextswitch = d.ContextSwitch
}

func (r *RealtimecpuTs) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, r.Ontunetime)
	data = append(data, r.Agenttime)
	data = append(data, r.Agentid)
	data = append(data, r.Index)
	data = append(data, r.User)
	data = append(data, r.Sys)
	data = append(data, r.Wait)
	data = append(data, r.Idle)
	data = append(data, r.Runqueue)
	data = append(data, r.Fork)
	data = append(data, r.Exec)
	data = append(data, r.Interupt)
	data = append(data, r.Systemcall)
	data = append(data, r.Contextswitch)

	return data
}

func (r *RealtimecpuTs) GetString() string {
	data := r.GetArgs()
	str_data := make([]string, 0)
	for _, d := range data {
		str_data = append(str_data, fmt.Sprintf("%v", d))
	}

	return strings.Join(str_data, ",")
}

func (r *RealtimecpuTs) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeCpu, tablename)
}

type RealtimediskPg struct {
	Ontunetime   int64 `db:"_ontunetime"`
	Agenttime    int   `db:"_agenttime"`
	Agentid      int   `db:"_agentid"`
	Ionameid     int   `db:"_ionameid"`
	Readrate     int   `db:"_readrate"`
	Writerate    int   `db:"_writerate"`
	Iops         int   `db:"_iops"`
	Busy         int   `db:"_busy"`
	Descid       int   `db:"_descid"`
	Readsvctime  int   `db:"_readsvctime"`
	Writesvctime int   `db:"_writesvctime"`
}

type RealtimediskPgArr struct {
	Ontunetime   []int64
	Agenttime    []int
	Agentid      []int
	Ionameid     []int
	Readrate     []int
	Writerate    []int
	Iops         []int
	Busy         []int
	Descid       []int
	Readsvctime  []int
	Writesvctime []int
}

func (r *RealtimediskPgArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeDiskInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Iops = append(r.Iops, d.Iops)
	r.Busy = append(r.Busy, d.Busy)
	r.Descid = append(r.Descid, ids[2])
	r.Readsvctime = append(r.Readsvctime, d.Readsvctime)
	r.Writesvctime = append(r.Writesvctime, d.Writesvctime)
}

func (r *RealtimediskPgArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeDisk, tablename, "int")
}

func (r *RealtimediskPgArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Ionameid))
	data = append(data, pq.Array(r.Readrate))
	data = append(data, pq.Array(r.Writerate))
	data = append(data, pq.Array(r.Iops))
	data = append(data, pq.Array(r.Busy))
	data = append(data, pq.Array(r.Descid))
	data = append(data, pq.Array(r.Readsvctime))
	data = append(data, pq.Array(r.Writesvctime))

	return data
}

func (a *RealtimediskPgArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Ionameid[i]))
		d = append(d, strconv.Itoa(a.Readrate[i]))
		d = append(d, strconv.Itoa(a.Writerate[i]))
		d = append(d, strconv.Itoa(a.Iops[i]))
		d = append(d, strconv.Itoa(a.Busy[i]))
		d = append(d, strconv.Itoa(a.Descid[i]))
		d = append(d, strconv.Itoa(a.Readsvctime[i]))
		d = append(d, strconv.Itoa(a.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimediskTs struct {
	Ontunetime   time.Time `db:"_ontunetime"`
	Agenttime    int       `db:"_agenttime"`
	Agentid      int       `db:"_agentid"`
	Ionameid     int       `db:"_ionameid"`
	Readrate     int       `db:"_readrate"`
	Writerate    int       `db:"_writerate"`
	Iops         int       `db:"_iops"`
	Busy         int       `db:"_busy"`
	Descid       int       `db:"_descid"`
	Readsvctime  int       `db:"_readsvctime"`
	Writesvctime int       `db:"_writesvctime"`
}

type RealtimediskTsArr struct {
	Ontunetime   []time.Time
	Agenttime    []int
	Agentid      []int
	Ionameid     []int
	Readrate     []int
	Writerate    []int
	Iops         []int
	Busy         []int
	Descid       []int
	Readsvctime  []int
	Writesvctime []int
}

func (r *RealtimediskTsArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeDiskInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Iops = append(r.Iops, d.Iops)
	r.Busy = append(r.Busy, d.Busy)
	r.Descid = append(r.Descid, ids[2])
	r.Readsvctime = append(r.Readsvctime, d.Readsvctime)
	r.Writesvctime = append(r.Writesvctime, d.Writesvctime)
}

func (r *RealtimediskTsArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeDisk, tablename, "timestamptz")
}

func (r *RealtimediskTsArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Ionameid))
	data = append(data, pq.Array(r.Readrate))
	data = append(data, pq.Array(r.Writerate))
	data = append(data, pq.Array(r.Iops))
	data = append(data, pq.Array(r.Busy))
	data = append(data, pq.Array(r.Descid))
	data = append(data, pq.Array(r.Readsvctime))
	data = append(data, pq.Array(r.Writesvctime))

	return data
}

func (a *RealtimediskTsArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Ionameid[i]))
		d = append(d, strconv.Itoa(a.Readrate[i]))
		d = append(d, strconv.Itoa(a.Writerate[i]))
		d = append(d, strconv.Itoa(a.Iops[i]))
		d = append(d, strconv.Itoa(a.Busy[i]))
		d = append(d, strconv.Itoa(a.Descid[i]))
		d = append(d, strconv.Itoa(a.Readsvctime[i]))
		d = append(d, strconv.Itoa(a.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimenetPg struct {
	Ontunetime int64 `db:"_ontunetime"`
	Agenttime  int   `db:"_agenttime"`
	Agentid    int   `db:"_agentid"`
	Ionameid   int   `db:"_ionameid"`
	Readrate   int   `db:"_readrate"`
	Writerate  int   `db:"_writerate"`
	Readiops   int   `db:"_readiops"`
	Writeiops  int   `db:"_writeiops"`
	Errorps    int   `db:"_errorps"`
	Collision  int   `db:"_collision"`
}

type RealtimenetPgArr struct {
	Ontunetime []int64
	Agenttime  []int
	Agentid    []int
	Ionameid   []int
	Readrate   []int
	Writerate  []int
	Readiops   []int
	Writeiops  []int
	Errorps    []int
	Collision  []int
}

func (r *RealtimenetPgArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeNetInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetPgArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeNet, tablename, "int")
}

func (r *RealtimenetPgArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Ionameid))
	data = append(data, pq.Array(r.Readrate))
	data = append(data, pq.Array(r.Writerate))
	data = append(data, pq.Array(r.Readiops))
	data = append(data, pq.Array(r.Writeiops))
	data = append(data, pq.Array(r.Errorps))
	data = append(data, pq.Array(r.Collision))

	return data
}

func (a *RealtimenetPgArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Ionameid[i]))
		d = append(d, strconv.Itoa(a.Readrate[i]))
		d = append(d, strconv.Itoa(a.Writerate[i]))
		d = append(d, strconv.Itoa(a.Readiops[i]))
		d = append(d, strconv.Itoa(a.Writeiops[i]))
		d = append(d, strconv.Itoa(a.Errorps[i]))
		d = append(d, strconv.Itoa(a.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimenetTs struct {
	Ontunetime time.Time `db:"_ontunetime"`
	Agenttime  int       `db:"_agenttime"`
	Agentid    int       `db:"_agentid"`
	Ionameid   int       `db:"_ionameid"`
	Readrate   int       `db:"_readrate"`
	Writerate  int       `db:"_writerate"`
	Readiops   int       `db:"_readiops"`
	Writeiops  int       `db:"_writeiops"`
	Errorps    int       `db:"_errorps"`
	Collision  int       `db:"_collision"`
}

type RealtimenetTsArr struct {
	Ontunetime []time.Time
	Agenttime  []int
	Agentid    []int
	Ionameid   []int
	Readrate   []int
	Writerate  []int
	Readiops   []int
	Writeiops  []int
	Errorps    []int
	Collision  []int
}

func (r *RealtimenetTsArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeNetInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetTsArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeNet, tablename, "timestamptz")
}

func (r *RealtimenetTsArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Ionameid))
	data = append(data, pq.Array(r.Readrate))
	data = append(data, pq.Array(r.Writerate))
	data = append(data, pq.Array(r.Readiops))
	data = append(data, pq.Array(r.Writeiops))
	data = append(data, pq.Array(r.Errorps))
	data = append(data, pq.Array(r.Collision))

	return data
}

func (a *RealtimenetTsArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Ionameid[i]))
		d = append(d, strconv.Itoa(a.Readrate[i]))
		d = append(d, strconv.Itoa(a.Writerate[i]))
		d = append(d, strconv.Itoa(a.Readiops[i]))
		d = append(d, strconv.Itoa(a.Writeiops[i]))
		d = append(d, strconv.Itoa(a.Errorps[i]))
		d = append(d, strconv.Itoa(a.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimepidPg struct {
	Ontunetime int64 `db:"_ontunetime"`
	Agenttime  int   `db:"_agenttime"`
	Agentid    int   `db:"_agentid"`
	Pid        int   `db:"_pid"`
	Ppid       int   `db:"_ppid"`
	Uid        int   `db:"_uid"`
	Cmdid      int   `db:"_cmdid"`
	Userid     int   `db:"_userid"`
	Argid      int   `db:"_argid"`
	Usr        int   `db:"_usr"`
	Sys        int   `db:"_sys"`
	Usrsys     int   `db:"_usrsys"`
	Sz         int   `db:"_sz"`
	Rss        int   `db:"_rss"`
	Vmem       int   `db:"_vmem"`
	Chario     int   `db:"_chario"`
	Processcnt int   `db:"_processcnt"`
	Threadcnt  int   `db:"_threadcnt"`
	Handlecnt  int   `db:"_handlecnt"`
	Stime      int   `db:"_stime"`
	Pvbytes    int   `db:"_pvbytes"`
	Pgpool     int   `db:"_pgpool"`
}

type RealtimepidPgArr struct {
	Ontunetime []int64
	Agenttime  []int
	Agentid    []int
	Pid        []int
	Ppid       []int
	Uid        []int
	Cmdid      []int
	Userid     []int
	Argid      []int
	Usr        []int
	Sys        []int
	Usrsys     []int
	Sz         []int
	Rss        []int
	Vmem       []int
	Chario     []int
	Processcnt []int
	Threadcnt  []int
	Handlecnt  []int
	Stime      []int
	Pvbytes    []int
	Pgpool     []int
}

func (r *RealtimepidPgArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Pid = append(r.Pid, d.Pid)
	r.Ppid = append(r.Ppid, d.Ppid)
	r.Uid = append(r.Uid, d.Uid)
	r.Cmdid = append(r.Cmdid, ids[1])
	r.Userid = append(r.Userid, ids[2])
	r.Argid = append(r.Argid, ids[3])
	r.Usr = append(r.Usr, d.Usr)
	r.Sys = append(r.Sys, d.Sys)
	r.Usrsys = append(r.Usrsys, d.Usrsys)
	r.Sz = append(r.Sz, d.Sz)
	r.Rss = append(r.Rss, d.Rss)
	r.Vmem = append(r.Vmem, d.Vmem)
	r.Chario = append(r.Chario, d.Chario)
	r.Processcnt = append(r.Processcnt, d.Processcnt)
	r.Threadcnt = append(r.Threadcnt, d.Threadcnt)
	r.Handlecnt = append(r.Handlecnt, d.Handlecnt)
	r.Stime = append(r.Stime, d.Stime)
	r.Pvbytes = append(r.Pvbytes, d.Pvbytes)
	r.Pgpool = append(r.Pgpool, d.Pgpool)
}

func (r *RealtimepidPgArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePid, tablename, "int")
}

func (r *RealtimepidPgArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Pid))
	data = append(data, pq.Array(r.Ppid))
	data = append(data, pq.Array(r.Uid))
	data = append(data, pq.Array(r.Cmdid))
	data = append(data, pq.Array(r.Userid))
	data = append(data, pq.Array(r.Argid))
	data = append(data, pq.Array(r.Usr))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Usrsys))
	data = append(data, pq.Array(r.Sz))
	data = append(data, pq.Array(r.Rss))
	data = append(data, pq.Array(r.Vmem))
	data = append(data, pq.Array(r.Chario))
	data = append(data, pq.Array(r.Processcnt))
	data = append(data, pq.Array(r.Threadcnt))
	data = append(data, pq.Array(r.Handlecnt))
	data = append(data, pq.Array(r.Stime))
	data = append(data, pq.Array(r.Pvbytes))
	data = append(data, pq.Array(r.Pgpool))

	return data
}

func (a *RealtimepidPgArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Pid[i]))
		d = append(d, strconv.Itoa(a.Ppid[i]))
		d = append(d, strconv.Itoa(a.Uid[i]))
		d = append(d, strconv.Itoa(a.Cmdid[i]))
		d = append(d, strconv.Itoa(a.Userid[i]))
		d = append(d, strconv.Itoa(a.Argid[i]))
		d = append(d, strconv.Itoa(a.Usr[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Usrsys[i]))
		d = append(d, strconv.Itoa(a.Sz[i]))
		d = append(d, strconv.Itoa(a.Rss[i]))
		d = append(d, strconv.Itoa(a.Vmem[i]))
		d = append(d, strconv.Itoa(a.Chario[i]))
		d = append(d, strconv.Itoa(a.Processcnt[i]))
		d = append(d, strconv.Itoa(a.Threadcnt[i]))
		d = append(d, strconv.Itoa(a.Handlecnt[i]))
		d = append(d, strconv.Itoa(a.Stime[i]))
		d = append(d, strconv.Itoa(a.Pvbytes[i]))
		d = append(d, strconv.Itoa(a.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimepidTs struct {
	Ontunetime time.Time `db:"_ontunetime"`
	Agenttime  int       `db:"_agenttime"`
	Agentid    int       `db:"_agentid"`
	Pid        int       `db:"_pid"`
	Ppid       int       `db:"_ppid"`
	Uid        int       `db:"_uid"`
	Cmdid      int       `db:"_cmdid"`
	Userid     int       `db:"_userid"`
	Argid      int       `db:"_argid"`
	Usr        int       `db:"_usr"`
	Sys        int       `db:"_sys"`
	Usrsys     int       `db:"_usrsys"`
	Sz         int       `db:"_sz"`
	Rss        int       `db:"_rss"`
	Vmem       int       `db:"_vmem"`
	Chario     int       `db:"_chario"`
	Processcnt int       `db:"_processcnt"`
	Threadcnt  int       `db:"_threadcnt"`
	Handlecnt  int       `db:"_handlecnt"`
	Stime      int       `db:"_stime"`
	Pvbytes    int       `db:"_pvbytes"`
	Pgpool     int       `db:"_pgpool"`
}

type RealtimepidTsArr struct {
	Ontunetime []time.Time
	Agenttime  []int
	Agentid    []int
	Pid        []int
	Ppid       []int
	Uid        []int
	Cmdid      []int
	Userid     []int
	Argid      []int
	Usr        []int
	Sys        []int
	Usrsys     []int
	Sz         []int
	Rss        []int
	Vmem       []int
	Chario     []int
	Processcnt []int
	Threadcnt  []int
	Handlecnt  []int
	Stime      []int
	Pvbytes    []int
	Pgpool     []int
}

func (r *RealtimepidTsArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Pid = append(r.Pid, d.Pid)
	r.Ppid = append(r.Ppid, d.Ppid)
	r.Uid = append(r.Uid, d.Uid)
	r.Cmdid = append(r.Cmdid, ids[1])
	r.Userid = append(r.Userid, ids[2])
	r.Argid = append(r.Argid, ids[3])
	r.Usr = append(r.Usr, d.Usr)
	r.Sys = append(r.Sys, d.Sys)
	r.Usrsys = append(r.Usrsys, d.Usrsys)
	r.Sz = append(r.Sz, d.Sz)
	r.Rss = append(r.Rss, d.Rss)
	r.Vmem = append(r.Vmem, d.Vmem)
	r.Chario = append(r.Chario, d.Chario)
	r.Processcnt = append(r.Processcnt, d.Processcnt)
	r.Threadcnt = append(r.Threadcnt, d.Threadcnt)
	r.Handlecnt = append(r.Handlecnt, d.Handlecnt)
	r.Stime = append(r.Stime, d.Stime)
	r.Pvbytes = append(r.Pvbytes, d.Pvbytes)
	r.Pgpool = append(r.Pgpool, d.Pgpool)
}

func (r *RealtimepidTsArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePid, tablename, "timestamptz")
}

func (r *RealtimepidTsArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Pid))
	data = append(data, pq.Array(r.Ppid))
	data = append(data, pq.Array(r.Uid))
	data = append(data, pq.Array(r.Cmdid))
	data = append(data, pq.Array(r.Userid))
	data = append(data, pq.Array(r.Argid))
	data = append(data, pq.Array(r.Usr))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Usrsys))
	data = append(data, pq.Array(r.Sz))
	data = append(data, pq.Array(r.Rss))
	data = append(data, pq.Array(r.Vmem))
	data = append(data, pq.Array(r.Chario))
	data = append(data, pq.Array(r.Processcnt))
	data = append(data, pq.Array(r.Threadcnt))
	data = append(data, pq.Array(r.Handlecnt))
	data = append(data, pq.Array(r.Stime))
	data = append(data, pq.Array(r.Pvbytes))
	data = append(data, pq.Array(r.Pgpool))

	return data
}

func (a *RealtimepidTsArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Pid[i]))
		d = append(d, strconv.Itoa(a.Ppid[i]))
		d = append(d, strconv.Itoa(a.Uid[i]))
		d = append(d, strconv.Itoa(a.Cmdid[i]))
		d = append(d, strconv.Itoa(a.Userid[i]))
		d = append(d, strconv.Itoa(a.Argid[i]))
		d = append(d, strconv.Itoa(a.Usr[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Usrsys[i]))
		d = append(d, strconv.Itoa(a.Sz[i]))
		d = append(d, strconv.Itoa(a.Rss[i]))
		d = append(d, strconv.Itoa(a.Vmem[i]))
		d = append(d, strconv.Itoa(a.Chario[i]))
		d = append(d, strconv.Itoa(a.Processcnt[i]))
		d = append(d, strconv.Itoa(a.Threadcnt[i]))
		d = append(d, strconv.Itoa(a.Handlecnt[i]))
		d = append(d, strconv.Itoa(a.Stime[i]))
		d = append(d, strconv.Itoa(a.Pvbytes[i]))
		d = append(d, strconv.Itoa(a.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimeprocPg struct {
	Ontunetime int64 `db:"_ontunetime"`
	Agenttime  int   `db:"_agenttime"`
	Agentid    int   `db:"_agentid"`
	Cmdid      int   `db:"_cmdid"`
	Userid     int   `db:"_userid"`
	Usr        int   `db:"_usr"`
	Sys        int   `db:"_sys"`
	Usrsys     int   `db:"_usrsys"`
	Sz         int   `db:"_sz"`
	Rss        int   `db:"_rss"`
	Vmem       int   `db:"_vmem"`
	Chario     int   `db:"_chario"`
	Processcnt int   `db:"_processcnt"`
	Threadcnt  int   `db:"_threadcnt"`
	Pvbytes    int   `db:"_pvbytes"`
	Pgpool     int   `db:"_pgpool"`
}

type RealtimeprocPgArr struct {
	Ontunetime []int64
	Agenttime  []int
	Agentid    []int
	Cmdid      []int
	Userid     []int
	Usr        []int
	Sys        []int
	Usrsys     []int
	Sz         []int
	Rss        []int
	Vmem       []int
	Chario     []int
	Processcnt []int
	Threadcnt  []int
	Pvbytes    []int
	Pgpool     []int
}

func (r *RealtimeprocPgArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Cmdid = append(r.Cmdid, ids[1])
	r.Userid = append(r.Userid, ids[2])
	r.Usr = append(r.Usr, d.Usr)
	r.Sys = append(r.Sys, d.Sys)
	r.Usrsys = append(r.Usrsys, d.Usrsys)
	r.Sz = append(r.Sz, d.Sz)
	r.Rss = append(r.Rss, d.Rss)
	r.Vmem = append(r.Vmem, d.Vmem)
	r.Chario = append(r.Chario, d.Chario)
	r.Processcnt = append(r.Processcnt, d.Processcnt)
	r.Threadcnt = append(r.Threadcnt, d.Threadcnt)
	r.Pvbytes = append(r.Pvbytes, d.Pvbytes)
	r.Pgpool = append(r.Pgpool, d.Pgpool)
}

func (r *RealtimeprocPgArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeProc, tablename, "int")
}

func (r *RealtimeprocPgArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Cmdid))
	data = append(data, pq.Array(r.Userid))
	data = append(data, pq.Array(r.Usr))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Usrsys))
	data = append(data, pq.Array(r.Sz))
	data = append(data, pq.Array(r.Rss))
	data = append(data, pq.Array(r.Vmem))
	data = append(data, pq.Array(r.Chario))
	data = append(data, pq.Array(r.Processcnt))
	data = append(data, pq.Array(r.Threadcnt))
	data = append(data, pq.Array(r.Pvbytes))
	data = append(data, pq.Array(r.Pgpool))

	return data
}

func (a *RealtimeprocPgArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Cmdid[i]))
		d = append(d, strconv.Itoa(a.Userid[i]))
		d = append(d, strconv.Itoa(a.Usr[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Usrsys[i]))
		d = append(d, strconv.Itoa(a.Sz[i]))
		d = append(d, strconv.Itoa(a.Rss[i]))
		d = append(d, strconv.Itoa(a.Vmem[i]))
		d = append(d, strconv.Itoa(a.Chario[i]))
		d = append(d, strconv.Itoa(a.Processcnt[i]))
		d = append(d, strconv.Itoa(a.Threadcnt[i]))
		d = append(d, strconv.Itoa(a.Pvbytes[i]))
		d = append(d, strconv.Itoa(a.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimeprocTs struct {
	Ontunetime time.Time `db:"_ontunetime"`
	Agenttime  int       `db:"_agenttime"`
	Agentid    int       `db:"_agentid"`
	Cmdid      int       `db:"_cmdid"`
	Userid     int       `db:"_userid"`
	Usr        int       `db:"_usr"`
	Sys        int       `db:"_sys"`
	Usrsys     int       `db:"_usrsys"`
	Sz         int       `db:"_sz"`
	Rss        int       `db:"_rss"`
	Vmem       int       `db:"_vmem"`
	Chario     int       `db:"_chario"`
	Processcnt int       `db:"_processcnt"`
	Threadcnt  int       `db:"_threadcnt"`
	Pvbytes    int       `db:"_pvbytes"`
	Pgpool     int       `db:"_pgpool"`
}

type RealtimeprocTsArr struct {
	Ontunetime []time.Time
	Agenttime  []int
	Agentid    []int
	Cmdid      []int
	Userid     []int
	Usr        []int
	Sys        []int
	Usrsys     []int
	Sz         []int
	Rss        []int
	Vmem       []int
	Chario     []int
	Processcnt []int
	Threadcnt  []int
	Pvbytes    []int
	Pgpool     []int
}

func (r *RealtimeprocTsArr) SetData(data interface{}, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, ids[0])
	r.Cmdid = append(r.Cmdid, ids[1])
	r.Userid = append(r.Userid, ids[2])
	r.Usr = append(r.Usr, d.Usr)
	r.Sys = append(r.Sys, d.Sys)
	r.Usrsys = append(r.Usrsys, d.Usrsys)
	r.Sz = append(r.Sz, d.Sz)
	r.Rss = append(r.Rss, d.Rss)
	r.Vmem = append(r.Vmem, d.Vmem)
	r.Chario = append(r.Chario, d.Chario)
	r.Processcnt = append(r.Processcnt, d.Processcnt)
	r.Threadcnt = append(r.Threadcnt, d.Threadcnt)
	r.Pvbytes = append(r.Pvbytes, d.Pvbytes)
	r.Pgpool = append(r.Pgpool, d.Pgpool)
}

func (r *RealtimeprocTsArr) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeProc, tablename, "timestamptz")
}

func (r *RealtimeprocTsArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Cmdid))
	data = append(data, pq.Array(r.Userid))
	data = append(data, pq.Array(r.Usr))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Usrsys))
	data = append(data, pq.Array(r.Sz))
	data = append(data, pq.Array(r.Rss))
	data = append(data, pq.Array(r.Vmem))
	data = append(data, pq.Array(r.Chario))
	data = append(data, pq.Array(r.Processcnt))
	data = append(data, pq.Array(r.Threadcnt))
	data = append(data, pq.Array(r.Pvbytes))
	data = append(data, pq.Array(r.Pgpool))

	return data
}

func (a *RealtimeprocTsArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Cmdid[i]))
		d = append(d, strconv.Itoa(a.Userid[i]))
		d = append(d, strconv.Itoa(a.Usr[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Usrsys[i]))
		d = append(d, strconv.Itoa(a.Sz[i]))
		d = append(d, strconv.Itoa(a.Rss[i]))
		d = append(d, strconv.Itoa(a.Vmem[i]))
		d = append(d, strconv.Itoa(a.Chario[i]))
		d = append(d, strconv.Itoa(a.Processcnt[i]))
		d = append(d, strconv.Itoa(a.Threadcnt[i]))
		d = append(d, strconv.Itoa(a.Pvbytes[i]))
		d = append(d, strconv.Itoa(a.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type SimpleStruct struct {
	Id   int    `db:"_id"`
	Name string `db:"_name"`
}

type DbInsert interface {
	SetData(data *interface{})
	GetArgs() []interface{}
}

type AgentinfoArr struct {
	Agentid          []int
	Hostname         []string
	Hostnameext      []string
	Enabled          []int
	Connected        []int
	Updated          []int
	Shorttermbasic   []int
	Shorttermproc    []int
	Shorttermio      []int
	Shorttermcpu     []int
	Longtermbasic    []int
	Longtermproc     []int
	Longtermio       []int
	Longtermcpu      []int
	Group            []string
	Ipaddress        []string
	Pscommand        []string
	Logevent         []string
	Processevent     []string
	Timecheck        []int
	Disconnectedtime []int64
	Skipdatatypes    []int
	Virbasicperf     []int
	Hypervisor       []int
	Serviceevent     []string
	Installdate      []int64
	Ibmpcrate        []int
	Updatedtime      []int64
	Os               []string
	Fw               []string
	Agentversion     []string
	Model            []string
	Serial           []string
	Processorcount   []int
	Processorclock   []int
	Memorysize       []int
	Swapsize         []int
	Poolid           []int
	Replication      []int
	Smt              []int
	Micropar         []int
	Capped           []int
	Ec               []int
	Virtualcpu       []int
	Weight           []int
	Cpupool          []int
	Ams              []int
	Allip            []string
	Numanodecount    []int
	Btime            []int64
}

func (a *AgentinfoArr) SetData(data interface{}) {
	d := data.(Agentinfo)
	a.Agentid = append(a.Agentid, d.Agentid)
	a.Hostname = append(a.Hostname, d.Hostname)
	a.Hostnameext = append(a.Hostnameext, d.Hostnameext)
	a.Enabled = append(a.Enabled, d.Enabled)
	a.Connected = append(a.Connected, d.Connected)
	a.Updated = append(a.Updated, d.Updated)
	a.Shorttermbasic = append(a.Shorttermbasic, d.Shorttermbasic)
	a.Shorttermproc = append(a.Shorttermproc, d.Shorttermproc)
	a.Shorttermio = append(a.Shorttermio, d.Shorttermio)
	a.Shorttermcpu = append(a.Shorttermcpu, d.Shorttermcpu)
	a.Longtermbasic = append(a.Longtermbasic, d.Longtermbasic)
	a.Longtermproc = append(a.Longtermproc, d.Longtermproc)
	a.Longtermio = append(a.Longtermio, d.Longtermio)
	a.Longtermcpu = append(a.Longtermcpu, d.Longtermcpu)
	a.Group = append(a.Group, d.Group)
	a.Ipaddress = append(a.Ipaddress, d.Ipaddress)
	a.Pscommand = append(a.Pscommand, d.Pscommand)
	a.Logevent = append(a.Logevent, d.Logevent)
	a.Processevent = append(a.Processevent, d.Processevent)
	a.Timecheck = append(a.Timecheck, d.Timecheck)
	a.Disconnectedtime = append(a.Disconnectedtime, d.Disconnectedtime)
	a.Skipdatatypes = append(a.Skipdatatypes, d.Skipdatatypes)
	a.Virbasicperf = append(a.Virbasicperf, d.Virbasicperf)
	a.Hypervisor = append(a.Hypervisor, d.Hypervisor)
	a.Serviceevent = append(a.Serviceevent, d.Serviceevent)
	a.Installdate = append(a.Installdate, d.Installdate)
	a.Ibmpcrate = append(a.Ibmpcrate, d.Ibmpcrate)
	a.Updatedtime = append(a.Updatedtime, d.Updatedtime)
	a.Os = append(a.Os, d.Os)
	a.Fw = append(a.Fw, d.Fw)
	a.Agentversion = append(a.Agentversion, d.Agentversion)
	a.Model = append(a.Model, d.Model)
	a.Serial = append(a.Serial, d.Serial)
	a.Processorcount = append(a.Processorcount, d.Processorcount)
	a.Processorclock = append(a.Processorclock, d.Processorclock)
	a.Memorysize = append(a.Memorysize, d.Memorysize)
	a.Swapsize = append(a.Swapsize, d.Swapsize)
	a.Poolid = append(a.Poolid, d.Poolid)
	a.Replication = append(a.Replication, d.Replication)
	a.Smt = append(a.Smt, d.Smt)
	a.Micropar = append(a.Micropar, d.Micropar)
	a.Capped = append(a.Capped, d.Capped)
	a.Ec = append(a.Ec, d.Ec)
	a.Virtualcpu = append(a.Virtualcpu, d.Virtualcpu)
	a.Weight = append(a.Weight, d.Weight)
	a.Cpupool = append(a.Cpupool, d.Cpupool)
	a.Ams = append(a.Ams, d.Ams)
	a.Allip = append(a.Allip, d.Allip)
	a.Numanodecount = append(a.Numanodecount, d.Numanodecount)
	a.Btime = append(a.Btime, d.Btime)
}

func (a *AgentinfoArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.StringArray(a.Hostnameext))
	data = append(data, pq.Array(a.Enabled))
	data = append(data, pq.Array(a.Connected))
	data = append(data, pq.Array(a.Updated))
	data = append(data, pq.Array(a.Shorttermbasic))
	data = append(data, pq.Array(a.Shorttermproc))
	data = append(data, pq.Array(a.Shorttermio))
	data = append(data, pq.Array(a.Shorttermcpu))
	data = append(data, pq.Array(a.Longtermbasic))
	data = append(data, pq.Array(a.Longtermproc))
	data = append(data, pq.Array(a.Longtermio))
	data = append(data, pq.Array(a.Longtermcpu))
	data = append(data, pq.StringArray(a.Group))
	data = append(data, pq.StringArray(a.Ipaddress))
	data = append(data, pq.StringArray(a.Pscommand))
	data = append(data, pq.StringArray(a.Logevent))
	data = append(data, pq.StringArray(a.Processevent))
	data = append(data, pq.Array(a.Timecheck))
	data = append(data, pq.Array(a.Disconnectedtime))
	data = append(data, pq.Array(a.Skipdatatypes))
	data = append(data, pq.Array(a.Virbasicperf))
	data = append(data, pq.Array(a.Hypervisor))
	data = append(data, pq.StringArray(a.Serviceevent))
	data = append(data, pq.Array(a.Installdate))
	data = append(data, pq.Array(a.Ibmpcrate))
	data = append(data, pq.Array(a.Updatedtime))
	data = append(data, pq.StringArray(a.Os))
	data = append(data, pq.StringArray(a.Fw))
	data = append(data, pq.StringArray(a.Agentversion))
	data = append(data, pq.StringArray(a.Model))
	data = append(data, pq.StringArray(a.Serial))
	data = append(data, pq.Array(a.Processorcount))
	data = append(data, pq.Array(a.Processorclock))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Poolid))
	data = append(data, pq.Array(a.Replication))
	data = append(data, pq.Array(a.Smt))
	data = append(data, pq.Array(a.Micropar))
	data = append(data, pq.Array(a.Capped))
	data = append(data, pq.Array(a.Ec))
	data = append(data, pq.Array(a.Virtualcpu))
	data = append(data, pq.Array(a.Weight))
	data = append(data, pq.Array(a.Cpupool))
	data = append(data, pq.Array(a.Ams))
	data = append(data, pq.StringArray(a.Allip))
	data = append(data, pq.Array(a.Numanodecount))
	data = append(data, pq.Array(a.Btime))

	return data
}

func (a *AgentinfoArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, a.Hostname[i])
		d = append(d, a.Hostnameext[i])
		d = append(d, strconv.Itoa(a.Enabled[i]))
		d = append(d, strconv.Itoa(a.Connected[i]))
		d = append(d, strconv.Itoa(a.Updated[i]))
		d = append(d, strconv.Itoa(a.Shorttermbasic[i]))
		d = append(d, strconv.Itoa(a.Shorttermproc[i]))
		d = append(d, strconv.Itoa(a.Shorttermio[i]))
		d = append(d, strconv.Itoa(a.Shorttermcpu[i]))
		d = append(d, strconv.Itoa(a.Longtermbasic[i]))
		d = append(d, strconv.Itoa(a.Longtermproc[i]))
		d = append(d, strconv.Itoa(a.Longtermio[i]))
		d = append(d, strconv.Itoa(a.Longtermcpu[i]))
		d = append(d, a.Group[i])
		d = append(d, a.Ipaddress[i])
		d = append(d, a.Pscommand[i])
		d = append(d, a.Logevent[i])
		d = append(d, a.Processevent[i])
		d = append(d, strconv.Itoa(a.Timecheck[i]))
		d = append(d, fmt.Sprintf("%v", a.Disconnectedtime[i]))
		d = append(d, strconv.Itoa(a.Skipdatatypes[i]))
		d = append(d, strconv.Itoa(a.Virbasicperf[i]))
		d = append(d, strconv.Itoa(a.Hypervisor[i]))
		d = append(d, a.Serviceevent[i])
		d = append(d, fmt.Sprintf("%v", a.Installdate[i]))
		d = append(d, strconv.Itoa(a.Ibmpcrate[i]))
		d = append(d, fmt.Sprintf("%v", a.Updatedtime[i]))
		d = append(d, a.Os[i])
		d = append(d, a.Fw[i])
		d = append(d, a.Agentversion[i])
		d = append(d, a.Model[i])
		d = append(d, a.Serial[i])
		d = append(d, strconv.Itoa(a.Processorcount[i]))
		d = append(d, strconv.Itoa(a.Processorclock[i]))
		d = append(d, strconv.Itoa(a.Memorysize[i]))
		d = append(d, strconv.Itoa(a.Swapsize[i]))
		d = append(d, strconv.Itoa(a.Poolid[i]))
		d = append(d, strconv.Itoa(a.Replication[i]))
		d = append(d, strconv.Itoa(a.Smt[i]))
		d = append(d, strconv.Itoa(a.Micropar[i]))
		d = append(d, strconv.Itoa(a.Capped[i]))
		d = append(d, strconv.Itoa(a.Ec[i]))
		d = append(d, strconv.Itoa(a.Virtualcpu[i]))
		d = append(d, strconv.Itoa(a.Weight[i]))
		d = append(d, strconv.Itoa(a.Cpupool[i]))
		d = append(d, strconv.Itoa(a.Ams[i]))
		d = append(d, a.Allip[i])
		d = append(d, strconv.Itoa(a.Numanodecount[i]))
		d = append(d, fmt.Sprintf("%v", a.Btime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type LastrealtimeperfArr struct {
	Ontunetime    []int64
	Agentid       []int
	Hostname      []string
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Memoryused    []int
	Filecache     []int
	Memorysize    []int
	Avm           []int
	Swapused      []int
	Swapsize      []int
	Diskiorate    []int
	Networkiorate []int
	Topproc       []string
	Topuser       []string
	Topproccount  []int
	Topcpu        []int
	Topdisk       []string
	Topvg         []string
	Topbusy       []int
	Maxcpu        []int
	Maxmem        []int
	Maxswap       []int
	Maxdisk       []int
	Diskiops      []int
	Networkiops   []int
	Dummy01       []int
	Dummy02       []int
	Dummy03       []int
	Dummy04       []int
	Dummy05       []int
	Dummy06       []int
	Dummy07       []int
	Dummy08       []int
	Dummy09       []int
	Dummy10       []int
	Dummy11       []int
	Dummy12       []int
	Dummy13       []int
	Dummy14       []int
	Dummy15       []int
	Dummy16       []int
	Dummy17       []int
	Dummy18       []int
	Dummy19       []int
	Dummy20       []int
	Dummy21       []int
	Dummy22       []int
	Dummy23       []int
	Dummy24       []int
	Dummy25       []int
	Dummy26       []int
	Dummy27       []int
	Dummy28       []int
	Dummy29       []int
	Dummy30       []int
}

func (l *LastrealtimeperfArr) SetData(data interface{}) {
	d := data.(Lastrealtimeperf)
	l.Ontunetime = append(l.Ontunetime, d.Ontunetime)
	l.Agentid = append(l.Agentid, d.Agentid)
	l.Hostname = append(l.Hostname, d.Hostname)
	l.User = append(l.User, d.User)
	l.Sys = append(l.Sys, d.Sys)
	l.Wait = append(l.Wait, d.Wait)
	l.Idle = append(l.Idle, d.Idle)
	l.Memoryused = append(l.Memoryused, d.Memoryused)
	l.Filecache = append(l.Filecache, d.Filecache)
	l.Memorysize = append(l.Memorysize, d.Memorysize)
	l.Avm = append(l.Avm, d.Avm)
	l.Swapused = append(l.Swapused, d.Swapused)
	l.Swapsize = append(l.Swapsize, d.Swapsize)
	l.Diskiorate = append(l.Diskiorate, d.Diskiorate)
	l.Networkiorate = append(l.Networkiorate, d.Networkiorate)
	l.Topproc = append(l.Topproc, d.Topproc)
	l.Topuser = append(l.Topuser, d.Topuser)
	l.Topproccount = append(l.Topproccount, d.Topproccount)
	l.Topcpu = append(l.Topcpu, d.Topcpu)
	l.Topdisk = append(l.Topdisk, d.Topdisk)
	l.Topvg = append(l.Topvg, d.Topvg)
	l.Topbusy = append(l.Topbusy, d.Topbusy)
	l.Maxcpu = append(l.Maxcpu, d.Maxcpu)
	l.Maxmem = append(l.Maxmem, d.Maxmem)
	l.Maxswap = append(l.Maxswap, d.Maxswap)
	l.Maxdisk = append(l.Maxdisk, d.Maxdisk)
	l.Diskiops = append(l.Diskiops, d.Diskiops)
	l.Networkiops = append(l.Networkiops, d.Networkiops)
	l.Dummy01 = append(l.Dummy01, d.Dummy01)
	l.Dummy02 = append(l.Dummy02, d.Dummy02)
	l.Dummy03 = append(l.Dummy03, d.Dummy03)
	l.Dummy04 = append(l.Dummy04, d.Dummy04)
	l.Dummy05 = append(l.Dummy05, d.Dummy05)
	l.Dummy06 = append(l.Dummy06, d.Dummy06)
	l.Dummy07 = append(l.Dummy07, d.Dummy07)
	l.Dummy08 = append(l.Dummy08, d.Dummy08)
	l.Dummy09 = append(l.Dummy09, d.Dummy09)
	l.Dummy10 = append(l.Dummy10, d.Dummy10)
	l.Dummy11 = append(l.Dummy11, d.Dummy11)
	l.Dummy12 = append(l.Dummy12, d.Dummy12)
	l.Dummy13 = append(l.Dummy13, d.Dummy13)
	l.Dummy14 = append(l.Dummy14, d.Dummy14)
	l.Dummy15 = append(l.Dummy15, d.Dummy15)
	l.Dummy16 = append(l.Dummy16, d.Dummy16)
	l.Dummy17 = append(l.Dummy17, d.Dummy17)
	l.Dummy18 = append(l.Dummy18, d.Dummy18)
	l.Dummy19 = append(l.Dummy19, d.Dummy19)
	l.Dummy20 = append(l.Dummy20, d.Dummy20)
	l.Dummy21 = append(l.Dummy21, d.Dummy21)
	l.Dummy22 = append(l.Dummy22, d.Dummy22)
	l.Dummy23 = append(l.Dummy23, d.Dummy23)
	l.Dummy24 = append(l.Dummy24, d.Dummy24)
	l.Dummy25 = append(l.Dummy25, d.Dummy25)
	l.Dummy26 = append(l.Dummy26, d.Dummy26)
	l.Dummy27 = append(l.Dummy27, d.Dummy27)
	l.Dummy28 = append(l.Dummy28, d.Dummy28)
	l.Dummy29 = append(l.Dummy29, d.Dummy29)
	l.Dummy30 = append(l.Dummy30, d.Dummy30)
}

func (l *LastrealtimeperfArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(l.Ontunetime))
	data = append(data, pq.Array(l.Agentid))
	data = append(data, pq.StringArray(l.Hostname))
	data = append(data, pq.Array(l.User))
	data = append(data, pq.Array(l.Sys))
	data = append(data, pq.Array(l.Wait))
	data = append(data, pq.Array(l.Idle))
	data = append(data, pq.Array(l.Memoryused))
	data = append(data, pq.Array(l.Filecache))
	data = append(data, pq.Array(l.Memorysize))
	data = append(data, pq.Array(l.Avm))
	data = append(data, pq.Array(l.Swapused))
	data = append(data, pq.Array(l.Swapsize))
	data = append(data, pq.Array(l.Diskiorate))
	data = append(data, pq.Array(l.Networkiorate))
	data = append(data, pq.StringArray(l.Topproc))
	data = append(data, pq.StringArray(l.Topuser))
	data = append(data, pq.Array(l.Topproccount))
	data = append(data, pq.Array(l.Topcpu))
	data = append(data, pq.StringArray(l.Topdisk))
	data = append(data, pq.StringArray(l.Topvg))
	data = append(data, pq.Array(l.Topbusy))
	data = append(data, pq.Array(l.Maxcpu))
	data = append(data, pq.Array(l.Maxmem))
	data = append(data, pq.Array(l.Maxswap))
	data = append(data, pq.Array(l.Maxdisk))
	data = append(data, pq.Array(l.Diskiops))
	data = append(data, pq.Array(l.Networkiops))
	data = append(data, pq.Array(l.Dummy01))
	data = append(data, pq.Array(l.Dummy02))
	data = append(data, pq.Array(l.Dummy03))
	data = append(data, pq.Array(l.Dummy04))
	data = append(data, pq.Array(l.Dummy05))
	data = append(data, pq.Array(l.Dummy06))
	data = append(data, pq.Array(l.Dummy07))
	data = append(data, pq.Array(l.Dummy08))
	data = append(data, pq.Array(l.Dummy09))
	data = append(data, pq.Array(l.Dummy10))
	data = append(data, pq.Array(l.Dummy11))
	data = append(data, pq.Array(l.Dummy12))
	data = append(data, pq.Array(l.Dummy13))
	data = append(data, pq.Array(l.Dummy14))
	data = append(data, pq.Array(l.Dummy15))
	data = append(data, pq.Array(l.Dummy16))
	data = append(data, pq.Array(l.Dummy17))
	data = append(data, pq.Array(l.Dummy18))
	data = append(data, pq.Array(l.Dummy19))
	data = append(data, pq.Array(l.Dummy20))
	data = append(data, pq.Array(l.Dummy21))
	data = append(data, pq.Array(l.Dummy22))
	data = append(data, pq.Array(l.Dummy23))
	data = append(data, pq.Array(l.Dummy24))
	data = append(data, pq.Array(l.Dummy25))
	data = append(data, pq.Array(l.Dummy26))
	data = append(data, pq.Array(l.Dummy27))
	data = append(data, pq.Array(l.Dummy28))
	data = append(data, pq.Array(l.Dummy29))
	data = append(data, pq.Array(l.Dummy30))

	return data
}

func (l *LastrealtimeperfArr) GetArray() []Lastrealtimeperf {
	arr := make([]Lastrealtimeperf, 0)
	size := len(l.Agentid)

	for i := 0; i < size; i++ {
		arr = append(arr, Lastrealtimeperf{
			Ontunetime:    l.Ontunetime[i],
			Agentid:       l.Agentid[i],
			Hostname:      l.Hostname[i],
			User:          l.User[i],
			Sys:           l.Sys[i],
			Wait:          l.Wait[i],
			Idle:          l.Idle[i],
			Memoryused:    l.Memoryused[i],
			Filecache:     l.Filecache[i],
			Memorysize:    l.Memorysize[i],
			Avm:           l.Avm[i],
			Swapused:      l.Swapused[i],
			Swapsize:      l.Swapsize[i],
			Diskiorate:    l.Diskiorate[i],
			Networkiorate: l.Networkiorate[i],
			Topproc:       l.Topproc[i],
			Topuser:       l.Topuser[i],
			Topproccount:  l.Topproccount[i],
			Topcpu:        l.Topcpu[i],
			Topdisk:       l.Topdisk[i],
			Topvg:         l.Topvg[i],
			Topbusy:       l.Topbusy[i],
			Maxcpu:        l.Maxcpu[i],
			Maxmem:        l.Maxmem[i],
			Maxswap:       l.Maxswap[i],
			Maxdisk:       l.Maxdisk[i],
			Diskiops:      l.Diskiops[i],
			Networkiops:   l.Networkiops[i],
			Dummy01:       l.Dummy01[i],
			Dummy02:       l.Dummy02[i],
			Dummy03:       l.Dummy03[i],
			Dummy04:       l.Dummy04[i],
			Dummy05:       l.Dummy05[i],
			Dummy06:       l.Dummy06[i],
			Dummy07:       l.Dummy07[i],
			Dummy08:       l.Dummy08[i],
			Dummy09:       l.Dummy09[i],
			Dummy10:       l.Dummy10[i],
			Dummy11:       l.Dummy11[i],
			Dummy12:       l.Dummy12[i],
			Dummy13:       l.Dummy13[i],
			Dummy14:       l.Dummy14[i],
			Dummy15:       l.Dummy15[i],
			Dummy16:       l.Dummy16[i],
			Dummy17:       l.Dummy17[i],
			Dummy18:       l.Dummy18[i],
			Dummy19:       l.Dummy19[i],
			Dummy20:       l.Dummy20[i],
			Dummy21:       l.Dummy21[i],
			Dummy22:       l.Dummy22[i],
			Dummy23:       l.Dummy23[i],
			Dummy24:       l.Dummy24[i],
			Dummy25:       l.Dummy25[i],
			Dummy26:       l.Dummy26[i],
			Dummy27:       l.Dummy27[i],
			Dummy28:       l.Dummy28[i],
			Dummy29:       l.Dummy29[i],
			Dummy30:       l.Dummy30[i],
		})
	}

	return arr
}

func (a *LastrealtimeperfArr) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, a.Hostname[i])
		d = append(d, strconv.Itoa(a.User[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Wait[i]))
		d = append(d, strconv.Itoa(a.Idle[i]))
		d = append(d, strconv.Itoa(a.Memoryused[i]))
		d = append(d, strconv.Itoa(a.Filecache[i]))
		d = append(d, strconv.Itoa(a.Memorysize[i]))
		d = append(d, strconv.Itoa(a.Avm[i]))
		d = append(d, strconv.Itoa(a.Swapused[i]))
		d = append(d, strconv.Itoa(a.Swapsize[i]))
		d = append(d, strconv.Itoa(a.Diskiorate[i]))
		d = append(d, strconv.Itoa(a.Networkiorate[i]))
		d = append(d, a.Topproc[i])
		d = append(d, a.Topuser[i])
		d = append(d, strconv.Itoa(a.Topproccount[i]))
		d = append(d, strconv.Itoa(a.Topcpu[i]))
		d = append(d, a.Topdisk[i])
		d = append(d, a.Topvg[i])
		d = append(d, strconv.Itoa(a.Topbusy[i]))
		d = append(d, strconv.Itoa(a.Maxcpu[i]))
		d = append(d, strconv.Itoa(a.Maxmem[i]))
		d = append(d, strconv.Itoa(a.Maxswap[i]))
		d = append(d, strconv.Itoa(a.Maxdisk[i]))
		d = append(d, strconv.Itoa(a.Diskiops[i]))
		d = append(d, strconv.Itoa(a.Networkiops[i]))
		d = append(d, strconv.Itoa(a.Dummy01[i]))
		d = append(d, strconv.Itoa(a.Dummy02[i]))
		d = append(d, strconv.Itoa(a.Dummy03[i]))
		d = append(d, strconv.Itoa(a.Dummy04[i]))
		d = append(d, strconv.Itoa(a.Dummy05[i]))
		d = append(d, strconv.Itoa(a.Dummy06[i]))
		d = append(d, strconv.Itoa(a.Dummy07[i]))
		d = append(d, strconv.Itoa(a.Dummy08[i]))
		d = append(d, strconv.Itoa(a.Dummy09[i]))
		d = append(d, strconv.Itoa(a.Dummy10[i]))
		d = append(d, strconv.Itoa(a.Dummy11[i]))
		d = append(d, strconv.Itoa(a.Dummy12[i]))
		d = append(d, strconv.Itoa(a.Dummy13[i]))
		d = append(d, strconv.Itoa(a.Dummy14[i]))
		d = append(d, strconv.Itoa(a.Dummy15[i]))
		d = append(d, strconv.Itoa(a.Dummy16[i]))
		d = append(d, strconv.Itoa(a.Dummy17[i]))
		d = append(d, strconv.Itoa(a.Dummy18[i]))
		d = append(d, strconv.Itoa(a.Dummy19[i]))
		d = append(d, strconv.Itoa(a.Dummy20[i]))
		d = append(d, strconv.Itoa(a.Dummy21[i]))
		d = append(d, strconv.Itoa(a.Dummy22[i]))
		d = append(d, strconv.Itoa(a.Dummy23[i]))
		d = append(d, strconv.Itoa(a.Dummy24[i]))
		d = append(d, strconv.Itoa(a.Dummy25[i]))
		d = append(d, strconv.Itoa(a.Dummy26[i]))
		d = append(d, strconv.Itoa(a.Dummy27[i]))
		d = append(d, strconv.Itoa(a.Dummy28[i]))
		d = append(d, strconv.Itoa(a.Dummy29[i]))
		d = append(d, strconv.Itoa(a.Dummy30[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}
