package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type TableSetArray interface {
	SetData(data interface{}, agentid int, strs ...string)
}

type TableSetArrayInner interface {
	SetData(data interface{}, agentid int, agenttime time.Time, ids ...int)
}

type LastrealtimeperfArray struct {
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

func (r *LastrealtimeperfArray) SetData(data interface{}, agentid int, strs ...string) {
	d := data.(AgentRealTimePerf)
	r.Ontunetime = append(r.Ontunetime, d.Agenttime.Unix())
	r.Agentid = append(r.Agentid, agentid)
	r.Hostname = append(r.Hostname, "")
	r.User = append(r.User, d.User)
	r.Sys = append(r.Sys, d.Sys)
	r.Wait = append(r.Wait, d.Wait)
	r.Idle = append(r.Idle, d.Idle)
	r.Memoryused = append(r.Memoryused, d.MemoryUsed)
	r.Filecache = append(r.Filecache, d.MemoryCache)
	r.Memorysize = append(r.Memorysize, d.MemorySize)
	r.Avm = append(r.Avm, d.Avm)
	r.Swapused = append(r.Swapused, d.SwapUsed)
	r.Swapsize = append(r.Swapsize, d.SwapSize)
	r.Diskiorate = append(r.Diskiorate, d.DiskReadWrite)
	r.Networkiorate = append(r.Networkiorate, d.NetworkReadWrite)
	r.Topproc = append(r.Topproc, "")
	r.Topuser = append(r.Topuser, "")
	r.Topproccount = append(r.Topproccount, 0)
	r.Topcpu = append(r.Topcpu, 0)
	r.Topdisk = append(r.Topdisk, "")
	r.Topvg = append(r.Topvg, "")
	r.Topbusy = append(r.Topbusy, 0)
	r.Maxcpu = append(r.Maxcpu, 0)
	r.Maxmem = append(r.Maxmem, 0)
	r.Maxswap = append(r.Maxswap, 0)
	r.Maxdisk = append(r.Maxdisk, 0)
	r.Diskiops = append(r.Diskiops, d.DiskIOPS)
	r.Networkiops = append(r.Networkiops)
	r.Dummy01 = append(r.Dummy01, 0)
	r.Dummy02 = append(r.Dummy02, 0)
	r.Dummy03 = append(r.Dummy03, 0)
	r.Dummy04 = append(r.Dummy04, 0)
	r.Dummy05 = append(r.Dummy05, 0)
	r.Dummy06 = append(r.Dummy06, 0)
	r.Dummy07 = append(r.Dummy07, 0)
	r.Dummy08 = append(r.Dummy08, 0)
	r.Dummy09 = append(r.Dummy09, 0)
	r.Dummy10 = append(r.Dummy10, 0)
	r.Dummy11 = append(r.Dummy11, 0)
	r.Dummy12 = append(r.Dummy12, 0)
	r.Dummy13 = append(r.Dummy13, 0)
	r.Dummy14 = append(r.Dummy14, 0)
	r.Dummy15 = append(r.Dummy15, 0)
	r.Dummy16 = append(r.Dummy16, 0)
	r.Dummy17 = append(r.Dummy17, 0)
	r.Dummy18 = append(r.Dummy18, 0)
	r.Dummy19 = append(r.Dummy19, 0)
	r.Dummy20 = append(r.Dummy20, 0)
	r.Dummy21 = append(r.Dummy21, 0)
	r.Dummy22 = append(r.Dummy22, 0)
	r.Dummy23 = append(r.Dummy23, 0)
	r.Dummy24 = append(r.Dummy24, 0)
	r.Dummy25 = append(r.Dummy25, 0)
	r.Dummy26 = append(r.Dummy26, 0)
	r.Dummy27 = append(r.Dummy27, 0)
	r.Dummy28 = append(r.Dummy28, 0)
	r.Dummy29 = append(r.Dummy29, 0)
	r.Dummy30 = append(r.Dummy30, 0)
}

func (r *LastrealtimeperfArray) GetInsertStmt(tablename string, timetype string) string {
	return InsertLastrealtimeperf
}

func (r *LastrealtimeperfArray) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.StringArray(r.Hostname))
	data = append(data, pq.Array(r.User))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Wait))
	data = append(data, pq.Array(r.Idle))
	data = append(data, pq.Array(r.Memoryused))
	data = append(data, pq.Array(r.Filecache))
	data = append(data, pq.Array(r.Memorysize))
	data = append(data, pq.Array(r.Avm))
	data = append(data, pq.Array(r.Swapused))
	data = append(data, pq.Array(r.Swapsize))
	data = append(data, pq.Array(r.Diskiorate))
	data = append(data, pq.Array(r.Networkiorate))
	data = append(data, pq.StringArray(r.Topproc))
	data = append(data, pq.StringArray(r.Topuser))
	data = append(data, pq.Array(r.Topproccount))
	data = append(data, pq.Array(r.Topcpu))
	data = append(data, pq.StringArray(r.Topdisk))
	data = append(data, pq.StringArray(r.Topvg))
	data = append(data, pq.Array(r.Topbusy))
	data = append(data, pq.Array(r.Maxcpu))
	data = append(data, pq.Array(r.Maxmem))
	data = append(data, pq.Array(r.Maxswap))
	data = append(data, pq.Array(r.Maxdisk))
	data = append(data, pq.Array(r.Diskiops))
	data = append(data, pq.Array(r.Networkiops))
	data = append(data, pq.Array(r.Dummy01))
	data = append(data, pq.Array(r.Dummy02))
	data = append(data, pq.Array(r.Dummy03))
	data = append(data, pq.Array(r.Dummy04))
	data = append(data, pq.Array(r.Dummy05))
	data = append(data, pq.Array(r.Dummy06))
	data = append(data, pq.Array(r.Dummy07))
	data = append(data, pq.Array(r.Dummy08))
	data = append(data, pq.Array(r.Dummy09))
	data = append(data, pq.Array(r.Dummy10))
	data = append(data, pq.Array(r.Dummy11))
	data = append(data, pq.Array(r.Dummy12))
	data = append(data, pq.Array(r.Dummy13))
	data = append(data, pq.Array(r.Dummy14))
	data = append(data, pq.Array(r.Dummy15))
	data = append(data, pq.Array(r.Dummy16))
	data = append(data, pq.Array(r.Dummy17))
	data = append(data, pq.Array(r.Dummy18))
	data = append(data, pq.Array(r.Dummy19))
	data = append(data, pq.Array(r.Dummy20))
	data = append(data, pq.Array(r.Dummy21))
	data = append(data, pq.Array(r.Dummy22))
	data = append(data, pq.Array(r.Dummy23))
	data = append(data, pq.Array(r.Dummy24))
	data = append(data, pq.Array(r.Dummy25))
	data = append(data, pq.Array(r.Dummy26))
	data = append(data, pq.Array(r.Dummy27))
	data = append(data, pq.Array(r.Dummy28))
	data = append(data, pq.Array(r.Dummy29))
	data = append(data, pq.Array(r.Dummy30))

	return data
}

func (r *LastrealtimeperfArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, r.Hostname[i])
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Filecache[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Diskiorate[i]))
		d = append(d, strconv.Itoa(r.Networkiorate[i]))
		d = append(d, r.Topproc[i])
		d = append(d, r.Topuser[i])
		d = append(d, strconv.Itoa(r.Topproccount[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, r.Topdisk[i])
		d = append(d, r.Topvg[i])
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxcpu[i]))
		d = append(d, strconv.Itoa(r.Maxmem[i]))
		d = append(d, strconv.Itoa(r.Maxswap[i]))
		d = append(d, strconv.Itoa(r.Maxdisk[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Dummy01[i]))
		d = append(d, strconv.Itoa(r.Dummy02[i]))
		d = append(d, strconv.Itoa(r.Dummy03[i]))
		d = append(d, strconv.Itoa(r.Dummy04[i]))
		d = append(d, strconv.Itoa(r.Dummy05[i]))
		d = append(d, strconv.Itoa(r.Dummy06[i]))
		d = append(d, strconv.Itoa(r.Dummy07[i]))
		d = append(d, strconv.Itoa(r.Dummy08[i]))
		d = append(d, strconv.Itoa(r.Dummy09[i]))
		d = append(d, strconv.Itoa(r.Dummy10[i]))
		d = append(d, strconv.Itoa(r.Dummy11[i]))
		d = append(d, strconv.Itoa(r.Dummy12[i]))
		d = append(d, strconv.Itoa(r.Dummy13[i]))
		d = append(d, strconv.Itoa(r.Dummy14[i]))
		d = append(d, strconv.Itoa(r.Dummy15[i]))
		d = append(d, strconv.Itoa(r.Dummy16[i]))
		d = append(d, strconv.Itoa(r.Dummy17[i]))
		d = append(d, strconv.Itoa(r.Dummy18[i]))
		d = append(d, strconv.Itoa(r.Dummy19[i]))
		d = append(d, strconv.Itoa(r.Dummy20[i]))
		d = append(d, strconv.Itoa(r.Dummy21[i]))
		d = append(d, strconv.Itoa(r.Dummy22[i]))
		d = append(d, strconv.Itoa(r.Dummy23[i]))
		d = append(d, strconv.Itoa(r.Dummy24[i]))
		d = append(d, strconv.Itoa(r.Dummy25[i]))
		d = append(d, strconv.Itoa(r.Dummy26[i]))
		d = append(d, strconv.Itoa(r.Dummy27[i]))
		d = append(d, strconv.Itoa(r.Dummy28[i]))
		d = append(d, strconv.Itoa(r.Dummy29[i]))
		d = append(d, strconv.Itoa(r.Dummy30[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *LastrealtimeperfArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, r.Hostname[i])
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Filecache[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Diskiorate[i]))
		d = append(d, strconv.Itoa(r.Networkiorate[i]))
		d = append(d, r.Topproc[i])
		d = append(d, r.Topuser[i])
		d = append(d, strconv.Itoa(r.Topproccount[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, r.Topdisk[i])
		d = append(d, r.Topvg[i])
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxcpu[i]))
		d = append(d, strconv.Itoa(r.Maxmem[i]))
		d = append(d, strconv.Itoa(r.Maxswap[i]))
		d = append(d, strconv.Itoa(r.Maxdisk[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Dummy01[i]))
		d = append(d, strconv.Itoa(r.Dummy02[i]))
		d = append(d, strconv.Itoa(r.Dummy03[i]))
		d = append(d, strconv.Itoa(r.Dummy04[i]))
		d = append(d, strconv.Itoa(r.Dummy05[i]))
		d = append(d, strconv.Itoa(r.Dummy06[i]))
		d = append(d, strconv.Itoa(r.Dummy07[i]))
		d = append(d, strconv.Itoa(r.Dummy08[i]))
		d = append(d, strconv.Itoa(r.Dummy09[i]))
		d = append(d, strconv.Itoa(r.Dummy10[i]))
		d = append(d, strconv.Itoa(r.Dummy11[i]))
		d = append(d, strconv.Itoa(r.Dummy12[i]))
		d = append(d, strconv.Itoa(r.Dummy13[i]))
		d = append(d, strconv.Itoa(r.Dummy14[i]))
		d = append(d, strconv.Itoa(r.Dummy15[i]))
		d = append(d, strconv.Itoa(r.Dummy16[i]))
		d = append(d, strconv.Itoa(r.Dummy17[i]))
		d = append(d, strconv.Itoa(r.Dummy18[i]))
		d = append(d, strconv.Itoa(r.Dummy19[i]))
		d = append(d, strconv.Itoa(r.Dummy20[i]))
		d = append(d, strconv.Itoa(r.Dummy21[i]))
		d = append(d, strconv.Itoa(r.Dummy22[i]))
		d = append(d, strconv.Itoa(r.Dummy23[i]))
		d = append(d, strconv.Itoa(r.Dummy24[i]))
		d = append(d, strconv.Itoa(r.Dummy25[i]))
		d = append(d, strconv.Itoa(r.Dummy26[i]))
		d = append(d, strconv.Itoa(r.Dummy27[i]))
		d = append(d, strconv.Itoa(r.Dummy28[i]))
		d = append(d, strconv.Itoa(r.Dummy29[i]))
		d = append(d, strconv.Itoa(r.Dummy30[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimeperfPgArray struct {
	Ontunetime       []int64
	Agenttime        []int
	Agentid          []int
	User             []int
	Sys              []int
	Wait             []int
	Idle             []int
	Processorcount   []int
	Runqueue         []int
	Blockqueue       []int
	Waitqueue        []int
	Pqueue           []int
	Pcrateuser       []int
	Pcratesys        []int
	Memorysize       []int
	Memoryused       []int
	Memorypinned     []int
	Memorysys        []int
	Memoryuser       []int
	Memorycache      []int
	Avm              []int
	Pagingspacein    []int
	Pagingspaceout   []int
	Filesystemin     []int
	Filesystemout    []int
	Memoryscan       []int
	Memoryfreed      []int
	Swapsize         []int
	Swapused         []int
	Swapactive       []int
	Fork             []int
	Exec             []int
	Interupt         []int
	Systemcall       []int
	Contextswitch    []int
	Semaphore        []int
	Msg              []int
	Diskreadwrite    []int
	Diskiops         []int
	Networkreadwrite []int
	Networkiops      []int
	Topcommandid     []int
	Topcommandcount  []int
	Topuserid        []int
	Topcpu           []int
	Topdiskid        []int
	Topvgid          []int
	Topbusy          []int
	Maxpid           []int
	Threadcount      []int
	Pidcount         []int
	Linuxbuffer      []int
	Linuxcached      []int
	Linuxsrec        []int
	Memused_Mb       []int
	Irq              []int
	Softirq          []int
	Swapused_Mb      []int
	Dusm             []int
}

func (r *RealtimeperfPgArray) SetData(data interface{}, agentid int, strs ...string) {
	d := data.(AgentRealTimePerf)
	r.Ontunetime = append(r.Ontunetime, d.Agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(d.Agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.User = append(r.User, d.User)
	r.Sys = append(r.Sys, d.Sys)
	r.Wait = append(r.Wait, d.Wait)
	r.Idle = append(r.Idle, d.Idle)
	r.Processorcount = append(r.Processorcount, d.ProcessorCount)
	r.Runqueue = append(r.Runqueue, d.RunQueue)
	r.Blockqueue = append(r.Blockqueue, d.BlockQueue)
	r.Waitqueue = append(r.Waitqueue, d.WaitQueue)
	r.Pqueue = append(r.Pqueue, d.PQueue)
	r.Pcrateuser = append(r.Pcrateuser, d.PCRateUser)
	r.Pcratesys = append(r.Pcratesys, d.PCRateSys)
	r.Memorysize = append(r.Memorysize, d.MemorySize)
	r.Memoryused = append(r.Memoryused, d.MemoryUsed)
	r.Memorypinned = append(r.Memorypinned, d.MemoryPinned)
	r.Memorysys = append(r.Memorysys, d.MemorySys)
	r.Memoryuser = append(r.Memoryuser, d.MemoryUser)
	r.Memorycache = append(r.Memorycache, d.MemoryCache)
	r.Avm = append(r.Avm, d.Avm)
	r.Pagingspacein = append(r.Pagingspacein, d.PagingspaceIn)
	r.Pagingspaceout = append(r.Pagingspaceout, d.PaingSpaceOut)
	r.Filesystemin = append(r.Filesystemin, d.FileSystemIn)
	r.Filesystemout = append(r.Filesystemout, d.FileSystmeOut)
	r.Memoryscan = append(r.Memoryscan, d.MemoryScan)
	r.Memoryfreed = append(r.Memoryfreed, d.MemoryFreed)
	r.Swapsize = append(r.Swapsize, d.SwapSize)
	r.Swapused = append(r.Swapused, d.SwapUsed)
	r.Swapactive = append(r.Swapactive, d.SwapActive)
	r.Fork = append(r.Fork, d.Fork)
	r.Exec = append(r.Exec, d.Exec)
	r.Interupt = append(r.Interupt, d.Interupt)
	r.Systemcall = append(r.Systemcall, d.SystemCall)
	r.Contextswitch = append(r.Contextswitch, d.ContextSwitch)
	r.Semaphore = append(r.Semaphore, d.Semaphore)
	r.Msg = append(r.Msg, d.Msg)
	r.Diskreadwrite = append(r.Diskreadwrite, d.DiskReadWrite)
	r.Diskiops = append(r.Diskiops, d.DiskIOPS)
	r.Networkreadwrite = append(r.Networkreadwrite, d.NetworkReadWrite)
	r.Networkiops = append(r.Networkiops, 0)
	r.Topcommandid = append(r.Topcommandid, d.TopCommandID)
	r.Topcommandcount = append(r.Topcommandcount, d.TopCommandCount)
	r.Topuserid = append(r.Topuserid, d.TopUserID)
	r.Topcpu = append(r.Topcpu, d.TopCPU)
	r.Topdiskid = append(r.Topdiskid, d.TopDiskID)
	r.Topvgid = append(r.Topvgid, d.TopvgID)
	r.Topbusy = append(r.Topbusy, d.TOPBusy)
	r.Maxpid = append(r.Maxpid, d.MaxPID)
	r.Threadcount = append(r.Threadcount, d.ThreadCount)
	r.Pidcount = append(r.Pidcount, d.PIDCount)
	r.Linuxbuffer = append(r.Linuxbuffer, d.LinuxBuffer)
	r.Linuxcached = append(r.Linuxcached, d.LinuxCached)
	r.Linuxsrec = append(r.Linuxsrec, d.Linuxsrec)
	r.Memused_Mb = append(r.Memused_Mb, d.Memused_mb)
	r.Irq = append(r.Irq, d.IRQ)
	r.Softirq = append(r.Softirq, d.SoftIRQ)
	r.Swapused_Mb = append(r.Swapused_Mb, d.Swapused_MB)
	r.Dusm = append(r.Dusm, d.DUSM)
}

func (r *RealtimeperfPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimePerf, tablename, "int")
}

func (r *RealtimeperfPgArray) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.User))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Wait))
	data = append(data, pq.Array(r.Idle))
	data = append(data, pq.Array(r.Processorcount))
	data = append(data, pq.Array(r.Runqueue))
	data = append(data, pq.Array(r.Blockqueue))
	data = append(data, pq.Array(r.Waitqueue))
	data = append(data, pq.Array(r.Pqueue))
	data = append(data, pq.Array(r.Pcrateuser))
	data = append(data, pq.Array(r.Pcratesys))
	data = append(data, pq.Array(r.Memorysize))
	data = append(data, pq.Array(r.Memoryused))
	data = append(data, pq.Array(r.Memorypinned))
	data = append(data, pq.Array(r.Memorysys))
	data = append(data, pq.Array(r.Memoryuser))
	data = append(data, pq.Array(r.Memorycache))
	data = append(data, pq.Array(r.Avm))
	data = append(data, pq.Array(r.Pagingspacein))
	data = append(data, pq.Array(r.Pagingspaceout))
	data = append(data, pq.Array(r.Filesystemin))
	data = append(data, pq.Array(r.Filesystemout))
	data = append(data, pq.Array(r.Memoryscan))
	data = append(data, pq.Array(r.Memoryfreed))
	data = append(data, pq.Array(r.Swapsize))
	data = append(data, pq.Array(r.Swapused))
	data = append(data, pq.Array(r.Swapactive))
	data = append(data, pq.Array(r.Fork))
	data = append(data, pq.Array(r.Exec))
	data = append(data, pq.Array(r.Interupt))
	data = append(data, pq.Array(r.Systemcall))
	data = append(data, pq.Array(r.Contextswitch))
	data = append(data, pq.Array(r.Semaphore))
	data = append(data, pq.Array(r.Msg))
	data = append(data, pq.Array(r.Diskreadwrite))
	data = append(data, pq.Array(r.Diskiops))
	data = append(data, pq.Array(r.Networkreadwrite))
	data = append(data, pq.Array(r.Networkiops))
	data = append(data, pq.Array(r.Topcommandid))
	data = append(data, pq.Array(r.Topcommandcount))
	data = append(data, pq.Array(r.Topuserid))
	data = append(data, pq.Array(r.Topcpu))
	data = append(data, pq.Array(r.Topdiskid))
	data = append(data, pq.Array(r.Topvgid))
	data = append(data, pq.Array(r.Topbusy))
	data = append(data, pq.Array(r.Maxpid))
	data = append(data, pq.Array(r.Threadcount))
	data = append(data, pq.Array(r.Pidcount))
	data = append(data, pq.Array(r.Linuxbuffer))
	data = append(data, pq.Array(r.Linuxcached))
	data = append(data, pq.Array(r.Linuxsrec))
	data = append(data, pq.Array(r.Memused_Mb))
	data = append(data, pq.Array(r.Irq))
	data = append(data, pq.Array(r.Softirq))
	data = append(data, pq.Array(r.Swapused_Mb))
	data = append(data, pq.Array(r.Dusm))
	return data
}

func (r *RealtimeperfPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Processorcount[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Blockqueue[i]))
		d = append(d, strconv.Itoa(r.Waitqueue[i]))
		d = append(d, strconv.Itoa(r.Pqueue[i]))
		d = append(d, strconv.Itoa(r.Pcrateuser[i]))
		d = append(d, strconv.Itoa(r.Pcratesys[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Memorypinned[i]))
		d = append(d, strconv.Itoa(r.Memorysys[i]))
		d = append(d, strconv.Itoa(r.Memoryuser[i]))
		d = append(d, strconv.Itoa(r.Memorycache[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Pagingspacein[i]))
		d = append(d, strconv.Itoa(r.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(r.Filesystemin[i]))
		d = append(d, strconv.Itoa(r.Filesystemout[i]))
		d = append(d, strconv.Itoa(r.Memoryscan[i]))
		d = append(d, strconv.Itoa(r.Memoryfreed[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapactive[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))
		d = append(d, strconv.Itoa(r.Semaphore[i]))
		d = append(d, strconv.Itoa(r.Msg[i]))
		d = append(d, strconv.Itoa(r.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Topcommandid[i]))
		d = append(d, strconv.Itoa(r.Topcommandcount[i]))
		d = append(d, strconv.Itoa(r.Topuserid[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, strconv.Itoa(r.Topdiskid[i]))
		d = append(d, strconv.Itoa(r.Topvgid[i]))
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxpid[i]))
		d = append(d, strconv.Itoa(r.Threadcount[i]))
		d = append(d, strconv.Itoa(r.Pidcount[i]))
		d = append(d, strconv.Itoa(r.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(r.Linuxcached[i]))
		d = append(d, strconv.Itoa(r.Linuxsrec[i]))
		d = append(d, strconv.Itoa(r.Memused_Mb[i]))
		d = append(d, strconv.Itoa(r.Irq[i]))
		d = append(d, strconv.Itoa(r.Softirq[i]))
		d = append(d, strconv.Itoa(r.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(r.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimeperfPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Processorcount[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Blockqueue[i]))
		d = append(d, strconv.Itoa(r.Waitqueue[i]))
		d = append(d, strconv.Itoa(r.Pqueue[i]))
		d = append(d, strconv.Itoa(r.Pcrateuser[i]))
		d = append(d, strconv.Itoa(r.Pcratesys[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Memorypinned[i]))
		d = append(d, strconv.Itoa(r.Memorysys[i]))
		d = append(d, strconv.Itoa(r.Memoryuser[i]))
		d = append(d, strconv.Itoa(r.Memorycache[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Pagingspacein[i]))
		d = append(d, strconv.Itoa(r.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(r.Filesystemin[i]))
		d = append(d, strconv.Itoa(r.Filesystemout[i]))
		d = append(d, strconv.Itoa(r.Memoryscan[i]))
		d = append(d, strconv.Itoa(r.Memoryfreed[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapactive[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))
		d = append(d, strconv.Itoa(r.Semaphore[i]))
		d = append(d, strconv.Itoa(r.Msg[i]))
		d = append(d, strconv.Itoa(r.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Topcommandid[i]))
		d = append(d, strconv.Itoa(r.Topcommandcount[i]))
		d = append(d, strconv.Itoa(r.Topuserid[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, strconv.Itoa(r.Topdiskid[i]))
		d = append(d, strconv.Itoa(r.Topvgid[i]))
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxpid[i]))
		d = append(d, strconv.Itoa(r.Threadcount[i]))
		d = append(d, strconv.Itoa(r.Pidcount[i]))
		d = append(d, strconv.Itoa(r.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(r.Linuxcached[i]))
		d = append(d, strconv.Itoa(r.Linuxsrec[i]))
		d = append(d, strconv.Itoa(r.Memused_Mb[i]))
		d = append(d, strconv.Itoa(r.Irq[i]))
		d = append(d, strconv.Itoa(r.Softirq[i]))
		d = append(d, strconv.Itoa(r.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(r.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimeperfTsArray struct {
	Ontunetime       []time.Time
	Agenttime        []int
	Agentid          []int
	User             []int
	Sys              []int
	Wait             []int
	Idle             []int
	Processorcount   []int
	Runqueue         []int
	Blockqueue       []int
	Waitqueue        []int
	Pqueue           []int
	Pcrateuser       []int
	Pcratesys        []int
	Memorysize       []int
	Memoryused       []int
	Memorypinned     []int
	Memorysys        []int
	Memoryuser       []int
	Memorycache      []int
	Avm              []int
	Pagingspacein    []int
	Pagingspaceout   []int
	Filesystemin     []int
	Filesystemout    []int
	Memoryscan       []int
	Memoryfreed      []int
	Swapsize         []int
	Swapused         []int
	Swapactive       []int
	Fork             []int
	Exec             []int
	Interupt         []int
	Systemcall       []int
	Contextswitch    []int
	Semaphore        []int
	Msg              []int
	Diskreadwrite    []int
	Diskiops         []int
	Networkreadwrite []int
	Networkiops      []int
	Topcommandid     []int
	Topcommandcount  []int
	Topuserid        []int
	Topcpu           []int
	Topdiskid        []int
	Topvgid          []int
	Topbusy          []int
	Maxpid           []int
	Threadcount      []int
	Pidcount         []int
	Linuxbuffer      []int
	Linuxcached      []int
	Linuxsrec        []int
	Memused_Mb       []int
	Irq              []int
	Softirq          []int
	Swapused_Mb      []int
	Dusm             []int
}

func (r *RealtimeperfTsArray) SetData(data interface{}, agentid int, strs ...string) {
	d := data.(AgentRealTimePerf)
	r.Ontunetime = append(r.Ontunetime, d.Agenttime)
	r.Agenttime = append(r.Agenttime, int(d.Agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.User = append(r.User, d.User)
	r.Sys = append(r.Sys, d.Sys)
	r.Wait = append(r.Wait, d.Wait)
	r.Idle = append(r.Idle, d.Idle)
	r.Processorcount = append(r.Processorcount, d.ProcessorCount)
	r.Runqueue = append(r.Runqueue, d.RunQueue)
	r.Blockqueue = append(r.Blockqueue, d.BlockQueue)
	r.Waitqueue = append(r.Waitqueue, d.WaitQueue)
	r.Pqueue = append(r.Pqueue, d.PQueue)
	r.Pcrateuser = append(r.Pcrateuser, d.PCRateUser)
	r.Pcratesys = append(r.Pcratesys, d.PCRateSys)
	r.Memorysize = append(r.Memorysize, d.MemorySize)
	r.Memoryused = append(r.Memoryused, d.MemoryUsed)
	r.Memorypinned = append(r.Memorypinned, d.MemoryPinned)
	r.Memorysys = append(r.Memorysys, d.MemorySys)
	r.Memoryuser = append(r.Memoryuser, d.MemoryUser)
	r.Memorycache = append(r.Memorycache, d.MemoryCache)
	r.Avm = append(r.Avm, d.Avm)
	r.Pagingspacein = append(r.Pagingspacein, d.PagingspaceIn)
	r.Pagingspaceout = append(r.Pagingspaceout, d.PaingSpaceOut)
	r.Filesystemin = append(r.Filesystemin, d.FileSystemIn)
	r.Filesystemout = append(r.Filesystemout, d.FileSystmeOut)
	r.Memoryscan = append(r.Memoryscan, d.MemoryScan)
	r.Memoryfreed = append(r.Memoryfreed, d.MemoryFreed)
	r.Swapsize = append(r.Swapsize, d.SwapSize)
	r.Swapused = append(r.Swapused, d.SwapUsed)
	r.Swapactive = append(r.Swapactive, d.SwapActive)
	r.Fork = append(r.Fork, d.Fork)
	r.Exec = append(r.Exec, d.Exec)
	r.Interupt = append(r.Interupt, d.Interupt)
	r.Systemcall = append(r.Systemcall, d.SystemCall)
	r.Contextswitch = append(r.Contextswitch, d.ContextSwitch)
	r.Semaphore = append(r.Semaphore, d.Semaphore)
	r.Msg = append(r.Msg, d.Msg)
	r.Diskreadwrite = append(r.Diskreadwrite, d.DiskReadWrite)
	r.Diskiops = append(r.Diskiops, d.DiskIOPS)
	r.Networkreadwrite = append(r.Networkreadwrite, d.NetworkReadWrite)
	r.Networkiops = append(r.Networkiops, 0)
	r.Topcommandid = append(r.Topcommandid, d.TopCommandID)
	r.Topcommandcount = append(r.Topcommandcount, d.TopCommandCount)
	r.Topuserid = append(r.Topuserid, d.TopUserID)
	r.Topcpu = append(r.Topcpu, d.TopCPU)
	r.Topdiskid = append(r.Topdiskid, d.TopDiskID)
	r.Topvgid = append(r.Topvgid, d.TopvgID)
	r.Topbusy = append(r.Topbusy, d.TOPBusy)
	r.Maxpid = append(r.Maxpid, d.MaxPID)
	r.Threadcount = append(r.Threadcount, d.ThreadCount)
	r.Pidcount = append(r.Pidcount, d.PIDCount)
	r.Linuxbuffer = append(r.Linuxbuffer, d.LinuxBuffer)
	r.Linuxcached = append(r.Linuxcached, d.LinuxCached)
	r.Linuxsrec = append(r.Linuxsrec, d.Linuxsrec)
	r.Memused_Mb = append(r.Memused_Mb, d.Memused_mb)
	r.Irq = append(r.Irq, d.IRQ)
	r.Softirq = append(r.Softirq, d.SoftIRQ)
	r.Swapused_Mb = append(r.Swapused_Mb, d.Swapused_MB)
	r.Dusm = append(r.Dusm, d.DUSM)
}

func (r *RealtimeperfTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimePerf, tablename, "timestamptz")
}

func (r *RealtimeperfTsArray) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.User))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Wait))
	data = append(data, pq.Array(r.Idle))
	data = append(data, pq.Array(r.Processorcount))
	data = append(data, pq.Array(r.Runqueue))
	data = append(data, pq.Array(r.Blockqueue))
	data = append(data, pq.Array(r.Waitqueue))
	data = append(data, pq.Array(r.Pqueue))
	data = append(data, pq.Array(r.Pcrateuser))
	data = append(data, pq.Array(r.Pcratesys))
	data = append(data, pq.Array(r.Memorysize))
	data = append(data, pq.Array(r.Memoryused))
	data = append(data, pq.Array(r.Memorypinned))
	data = append(data, pq.Array(r.Memorysys))
	data = append(data, pq.Array(r.Memoryuser))
	data = append(data, pq.Array(r.Memorycache))
	data = append(data, pq.Array(r.Avm))
	data = append(data, pq.Array(r.Pagingspacein))
	data = append(data, pq.Array(r.Pagingspaceout))
	data = append(data, pq.Array(r.Filesystemin))
	data = append(data, pq.Array(r.Filesystemout))
	data = append(data, pq.Array(r.Memoryscan))
	data = append(data, pq.Array(r.Memoryfreed))
	data = append(data, pq.Array(r.Swapsize))
	data = append(data, pq.Array(r.Swapused))
	data = append(data, pq.Array(r.Swapactive))
	data = append(data, pq.Array(r.Fork))
	data = append(data, pq.Array(r.Exec))
	data = append(data, pq.Array(r.Interupt))
	data = append(data, pq.Array(r.Systemcall))
	data = append(data, pq.Array(r.Contextswitch))
	data = append(data, pq.Array(r.Semaphore))
	data = append(data, pq.Array(r.Msg))
	data = append(data, pq.Array(r.Diskreadwrite))
	data = append(data, pq.Array(r.Diskiops))
	data = append(data, pq.Array(r.Networkreadwrite))
	data = append(data, pq.Array(r.Networkiops))
	data = append(data, pq.Array(r.Topcommandid))
	data = append(data, pq.Array(r.Topcommandcount))
	data = append(data, pq.Array(r.Topuserid))
	data = append(data, pq.Array(r.Topcpu))
	data = append(data, pq.Array(r.Topdiskid))
	data = append(data, pq.Array(r.Topvgid))
	data = append(data, pq.Array(r.Topbusy))
	data = append(data, pq.Array(r.Maxpid))
	data = append(data, pq.Array(r.Threadcount))
	data = append(data, pq.Array(r.Pidcount))
	data = append(data, pq.Array(r.Linuxbuffer))
	data = append(data, pq.Array(r.Linuxcached))
	data = append(data, pq.Array(r.Linuxsrec))
	data = append(data, pq.Array(r.Memused_Mb))
	data = append(data, pq.Array(r.Irq))
	data = append(data, pq.Array(r.Softirq))
	data = append(data, pq.Array(r.Swapused_Mb))
	data = append(data, pq.Array(r.Dusm))
	return data
}

func (r *RealtimeperfTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Processorcount[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Blockqueue[i]))
		d = append(d, strconv.Itoa(r.Waitqueue[i]))
		d = append(d, strconv.Itoa(r.Pqueue[i]))
		d = append(d, strconv.Itoa(r.Pcrateuser[i]))
		d = append(d, strconv.Itoa(r.Pcratesys[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Memorypinned[i]))
		d = append(d, strconv.Itoa(r.Memorysys[i]))
		d = append(d, strconv.Itoa(r.Memoryuser[i]))
		d = append(d, strconv.Itoa(r.Memorycache[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Pagingspacein[i]))
		d = append(d, strconv.Itoa(r.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(r.Filesystemin[i]))
		d = append(d, strconv.Itoa(r.Filesystemout[i]))
		d = append(d, strconv.Itoa(r.Memoryscan[i]))
		d = append(d, strconv.Itoa(r.Memoryfreed[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapactive[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))
		d = append(d, strconv.Itoa(r.Semaphore[i]))
		d = append(d, strconv.Itoa(r.Msg[i]))
		d = append(d, strconv.Itoa(r.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Topcommandid[i]))
		d = append(d, strconv.Itoa(r.Topcommandcount[i]))
		d = append(d, strconv.Itoa(r.Topuserid[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, strconv.Itoa(r.Topdiskid[i]))
		d = append(d, strconv.Itoa(r.Topvgid[i]))
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxpid[i]))
		d = append(d, strconv.Itoa(r.Threadcount[i]))
		d = append(d, strconv.Itoa(r.Pidcount[i]))
		d = append(d, strconv.Itoa(r.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(r.Linuxcached[i]))
		d = append(d, strconv.Itoa(r.Linuxsrec[i]))
		d = append(d, strconv.Itoa(r.Memused_Mb[i]))
		d = append(d, strconv.Itoa(r.Irq[i]))
		d = append(d, strconv.Itoa(r.Softirq[i]))
		d = append(d, strconv.Itoa(r.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(r.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimeperfTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Processorcount[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Blockqueue[i]))
		d = append(d, strconv.Itoa(r.Waitqueue[i]))
		d = append(d, strconv.Itoa(r.Pqueue[i]))
		d = append(d, strconv.Itoa(r.Pcrateuser[i]))
		d = append(d, strconv.Itoa(r.Pcratesys[i]))
		d = append(d, strconv.Itoa(r.Memorysize[i]))
		d = append(d, strconv.Itoa(r.Memoryused[i]))
		d = append(d, strconv.Itoa(r.Memorypinned[i]))
		d = append(d, strconv.Itoa(r.Memorysys[i]))
		d = append(d, strconv.Itoa(r.Memoryuser[i]))
		d = append(d, strconv.Itoa(r.Memorycache[i]))
		d = append(d, strconv.Itoa(r.Avm[i]))
		d = append(d, strconv.Itoa(r.Pagingspacein[i]))
		d = append(d, strconv.Itoa(r.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(r.Filesystemin[i]))
		d = append(d, strconv.Itoa(r.Filesystemout[i]))
		d = append(d, strconv.Itoa(r.Memoryscan[i]))
		d = append(d, strconv.Itoa(r.Memoryfreed[i]))
		d = append(d, strconv.Itoa(r.Swapsize[i]))
		d = append(d, strconv.Itoa(r.Swapused[i]))
		d = append(d, strconv.Itoa(r.Swapactive[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))
		d = append(d, strconv.Itoa(r.Semaphore[i]))
		d = append(d, strconv.Itoa(r.Msg[i]))
		d = append(d, strconv.Itoa(r.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(r.Diskiops[i]))
		d = append(d, strconv.Itoa(r.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(r.Networkiops[i]))
		d = append(d, strconv.Itoa(r.Topcommandid[i]))
		d = append(d, strconv.Itoa(r.Topcommandcount[i]))
		d = append(d, strconv.Itoa(r.Topuserid[i]))
		d = append(d, strconv.Itoa(r.Topcpu[i]))
		d = append(d, strconv.Itoa(r.Topdiskid[i]))
		d = append(d, strconv.Itoa(r.Topvgid[i]))
		d = append(d, strconv.Itoa(r.Topbusy[i]))
		d = append(d, strconv.Itoa(r.Maxpid[i]))
		d = append(d, strconv.Itoa(r.Threadcount[i]))
		d = append(d, strconv.Itoa(r.Pidcount[i]))
		d = append(d, strconv.Itoa(r.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(r.Linuxcached[i]))
		d = append(d, strconv.Itoa(r.Linuxsrec[i]))
		d = append(d, strconv.Itoa(r.Memused_Mb[i]))
		d = append(d, strconv.Itoa(r.Irq[i]))
		d = append(d, strconv.Itoa(r.Softirq[i]))
		d = append(d, strconv.Itoa(r.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(r.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimecpuPgArray struct {
	Ontunetime    []int64
	Agenttime     []int
	Agentid       []int
	Index         []int
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Runqueue      []int
	Fork          []int
	Exec          []int
	Interupt      []int
	Systemcall    []int
	Contextswitch []int
}

func (r *RealtimecpuPgArray) SetData(data interface{}, agentid int, strs ...string) {
	d := data.(AgentRealTimePerf)
	r.Ontunetime = append(r.Ontunetime, d.Agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(d.Agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Index = append(r.Index, 0)
	r.User = append(r.User, d.User)
	r.Sys = append(r.Sys, d.Sys)
	r.Wait = append(r.Wait, d.Wait)
	r.Idle = append(r.Idle, d.Idle)
	r.Runqueue = append(r.Runqueue, d.RunQueue)
	r.Fork = append(r.Fork, d.Fork)
	r.Exec = append(r.Exec, d.Exec)
	r.Interupt = append(r.Interupt, d.Interupt)
	r.Systemcall = append(r.Systemcall, d.SystemCall)
	r.Contextswitch = append(r.Contextswitch, d.ContextSwitch)
}

func (r *RealtimecpuPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeCpu, tablename, "int")
}

func (r *RealtimecpuPgArray) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Index))
	data = append(data, pq.Array(r.User))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Wait))
	data = append(data, pq.Array(r.Idle))
	data = append(data, pq.Array(r.Runqueue))
	data = append(data, pq.Array(r.Fork))
	data = append(data, pq.Array(r.Exec))
	data = append(data, pq.Array(r.Interupt))
	data = append(data, pq.Array(r.Systemcall))
	data = append(data, pq.Array(r.Contextswitch))
	return data
}

func (r *RealtimecpuPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Index[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimecpuPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Index[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimecpuTsArray struct {
	Ontunetime    []time.Time
	Agenttime     []int
	Agentid       []int
	Index         []int
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Runqueue      []int
	Fork          []int
	Exec          []int
	Interupt      []int
	Systemcall    []int
	Contextswitch []int
}

func (r *RealtimecpuTsArray) SetData(data interface{}, agentid int, strs ...string) {
	d := data.(AgentRealTimePerf)
	r.Ontunetime = append(r.Ontunetime, d.Agenttime)
	r.Agenttime = append(r.Agenttime, int(d.Agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Index = append(r.Index, 0)
	r.User = append(r.User, d.User)
	r.Sys = append(r.Sys, d.Sys)
	r.Wait = append(r.Wait, d.Wait)
	r.Idle = append(r.Idle, d.Idle)
	r.Runqueue = append(r.Runqueue, d.RunQueue)
	r.Fork = append(r.Fork, d.Fork)
	r.Exec = append(r.Exec, d.Exec)
	r.Interupt = append(r.Interupt, d.Interupt)
	r.Systemcall = append(r.Systemcall, d.SystemCall)
	r.Contextswitch = append(r.Contextswitch, d.ContextSwitch)
}

func (r *RealtimecpuTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeCpu, tablename, "timestamptz")
}

func (r *RealtimecpuTsArray) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(r.Ontunetime))
	data = append(data, pq.Array(r.Agenttime))
	data = append(data, pq.Array(r.Agentid))
	data = append(data, pq.Array(r.Index))
	data = append(data, pq.Array(r.User))
	data = append(data, pq.Array(r.Sys))
	data = append(data, pq.Array(r.Wait))
	data = append(data, pq.Array(r.Idle))
	data = append(data, pq.Array(r.Runqueue))
	data = append(data, pq.Array(r.Fork))
	data = append(data, pq.Array(r.Exec))
	data = append(data, pq.Array(r.Interupt))
	data = append(data, pq.Array(r.Systemcall))
	data = append(data, pq.Array(r.Contextswitch))
	return data
}

func (r *RealtimecpuTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Index[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimecpuTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Index[i]))
		d = append(d, strconv.Itoa(r.User[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Wait[i]))
		d = append(d, strconv.Itoa(r.Idle[i]))
		d = append(d, strconv.Itoa(r.Runqueue[i]))
		d = append(d, strconv.Itoa(r.Fork[i]))
		d = append(d, strconv.Itoa(r.Exec[i]))
		d = append(d, strconv.Itoa(r.Interupt[i]))
		d = append(d, strconv.Itoa(r.Systemcall[i]))
		d = append(d, strconv.Itoa(r.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimediskPgArray struct {
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

func (r *RealtimediskPgArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeDiskInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Ionameid = append(r.Ionameid, ids[0])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Iops = append(r.Iops, d.Iops)
	r.Busy = append(r.Busy, d.Busy)
	r.Descid = append(r.Descid, ids[1])
	r.Readsvctime = append(r.Readsvctime, d.Readsvctime)
	r.Writesvctime = append(r.Writesvctime, d.Writesvctime)
}

func (r *RealtimediskPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeDisk, tablename, "int")
}

func (r *RealtimediskPgArray) GetArgs() []interface{} {
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

func (r *RealtimediskPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Iops[i]))
		d = append(d, strconv.Itoa(r.Busy[i]))
		d = append(d, strconv.Itoa(r.Descid[i]))
		d = append(d, strconv.Itoa(r.Readsvctime[i]))
		d = append(d, strconv.Itoa(r.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimediskPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Iops[i]))
		d = append(d, strconv.Itoa(r.Busy[i]))
		d = append(d, strconv.Itoa(r.Descid[i]))
		d = append(d, strconv.Itoa(r.Readsvctime[i]))
		d = append(d, strconv.Itoa(r.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimediskTsArray struct {
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

func (r *RealtimediskTsArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeDiskInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Ionameid = append(r.Ionameid, ids[0])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Iops = append(r.Iops, d.Iops)
	r.Busy = append(r.Busy, d.Busy)
	r.Descid = append(r.Descid, ids[1])
	r.Readsvctime = append(r.Readsvctime, d.Readsvctime)
	r.Writesvctime = append(r.Writesvctime, d.Writesvctime)
}

func (r *RealtimediskTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeDisk, tablename, "timestamptz")
}

func (r *RealtimediskTsArray) GetArgs() []interface{} {
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

func (r *RealtimediskTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Iops[i]))
		d = append(d, strconv.Itoa(r.Busy[i]))
		d = append(d, strconv.Itoa(r.Descid[i]))
		d = append(d, strconv.Itoa(r.Readsvctime[i]))
		d = append(d, strconv.Itoa(r.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

func (r *RealtimediskTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Iops[i]))
		d = append(d, strconv.Itoa(r.Busy[i]))
		d = append(d, strconv.Itoa(r.Descid[i]))
		d = append(d, strconv.Itoa(r.Readsvctime[i]))
		d = append(d, strconv.Itoa(r.Writesvctime[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimenetPgArray struct {
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

func (r *RealtimenetPgArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeNetInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Ionameid = append(r.Ionameid, ids[0])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeNet, tablename, "int")
}

func (r *RealtimenetPgArray) GetArgs() []interface{} {
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

func (r *RealtimenetPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Readiops[i]))
		d = append(d, strconv.Itoa(r.Writeiops[i]))
		d = append(d, strconv.Itoa(r.Errorps[i]))
		d = append(d, strconv.Itoa(r.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

func (r *RealtimenetPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Readiops[i]))
		d = append(d, strconv.Itoa(r.Writeiops[i]))
		d = append(d, strconv.Itoa(r.Errorps[i]))
		d = append(d, strconv.Itoa(r.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimenetTsArray struct {
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

func (r *RealtimenetTsArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimeNetInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Ionameid = append(r.Ionameid, ids[0])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeNet, tablename, "timestamptz")
}

func (r *RealtimenetTsArray) GetArgs() []interface{} {
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

func (r *RealtimenetTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Readiops[i]))
		d = append(d, strconv.Itoa(r.Writeiops[i]))
		d = append(d, strconv.Itoa(r.Errorps[i]))
		d = append(d, strconv.Itoa(r.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimenetTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Ionameid[i]))
		d = append(d, strconv.Itoa(r.Readrate[i]))
		d = append(d, strconv.Itoa(r.Writerate[i]))
		d = append(d, strconv.Itoa(r.Readiops[i]))
		d = append(d, strconv.Itoa(r.Writeiops[i]))
		d = append(d, strconv.Itoa(r.Errorps[i]))
		d = append(d, strconv.Itoa(r.Collision[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimepidPgArray struct {
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

func (r *RealtimepidPgArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Pid = append(r.Pid, d.Pid)
	r.Ppid = append(r.Ppid, d.Ppid)
	r.Uid = append(r.Uid, d.Uid)
	r.Cmdid = append(r.Cmdid, ids[0])
	r.Userid = append(r.Userid, ids[1])
	r.Argid = append(r.Argid, ids[2])
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

func (r *RealtimepidPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimePid, tablename, "int")
}

func (r *RealtimepidPgArray) GetArgs() []interface{} {
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

func (r *RealtimepidPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Pid[i]))
		d = append(d, strconv.Itoa(r.Ppid[i]))
		d = append(d, strconv.Itoa(r.Uid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Argid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Handlecnt[i]))
		d = append(d, strconv.Itoa(r.Stime[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimepidPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Pid[i]))
		d = append(d, strconv.Itoa(r.Ppid[i]))
		d = append(d, strconv.Itoa(r.Uid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Argid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Handlecnt[i]))
		d = append(d, strconv.Itoa(r.Stime[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimepidTsArray struct {
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

func (r *RealtimepidTsArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Pid = append(r.Pid, d.Pid)
	r.Ppid = append(r.Ppid, d.Ppid)
	r.Uid = append(r.Uid, d.Uid)
	r.Cmdid = append(r.Cmdid, ids[0])
	r.Userid = append(r.Userid, ids[1])
	r.Argid = append(r.Argid, ids[2])
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

func (r *RealtimepidTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimePid, tablename, "timestamptz")
}

func (r *RealtimepidTsArray) GetArgs() []interface{} {
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

func (r *RealtimepidTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Pid[i]))
		d = append(d, strconv.Itoa(r.Ppid[i]))
		d = append(d, strconv.Itoa(r.Uid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Argid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Handlecnt[i]))
		d = append(d, strconv.Itoa(r.Stime[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

func (r *RealtimepidTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Pid[i]))
		d = append(d, strconv.Itoa(r.Ppid[i]))
		d = append(d, strconv.Itoa(r.Uid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Argid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Handlecnt[i]))
		d = append(d, strconv.Itoa(r.Stime[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

type RealtimeprocPgArray struct {
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

func (r *RealtimeprocPgArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime.Unix())
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Cmdid = append(r.Cmdid, ids[0])
	r.Userid = append(r.Userid, ids[1])
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

func (r *RealtimeprocPgArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeProc, tablename, "int")
}

func (r *RealtimeprocPgArray) GetArgs() []interface{} {
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

func (r *RealtimeprocPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimeprocPgArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}

type RealtimeprocTsArray struct {
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

func (r *RealtimeprocTsArray) SetData(data interface{}, agentid int, agenttime time.Time, ids ...int) {
	d := data.(AgentRealTimePIDInner)
	r.Ontunetime = append(r.Ontunetime, agenttime)
	r.Agenttime = append(r.Agenttime, int(agenttime.Unix()))
	r.Agentid = append(r.Agentid, agentid)
	r.Cmdid = append(r.Cmdid, ids[0])
	r.Userid = append(r.Userid, ids[1])
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

func (r *RealtimeprocTsArray) GetInsertStmt(tablename string, timetype string) string {
	return fmt.Sprintf(InsertRealtimeProc, tablename, "timestamptz")
}

func (r *RealtimeprocTsArray) GetArgs() []interface{} {
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

func (r *RealtimeprocTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
}

func (r *RealtimeprocTsArray) GetString() string {
	arr := make([]string, 0)
	size := len(r.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", r.Ontunetime[i]))
		d = append(d, strconv.Itoa(r.Agenttime[i]))
		d = append(d, strconv.Itoa(r.Agentid[i]))
		d = append(d, strconv.Itoa(r.Cmdid[i]))
		d = append(d, strconv.Itoa(r.Userid[i]))
		d = append(d, strconv.Itoa(r.Usr[i]))
		d = append(d, strconv.Itoa(r.Sys[i]))
		d = append(d, strconv.Itoa(r.Usrsys[i]))
		d = append(d, strconv.Itoa(r.Sz[i]))
		d = append(d, strconv.Itoa(r.Rss[i]))
		d = append(d, strconv.Itoa(r.Vmem[i]))
		d = append(d, strconv.Itoa(r.Chario[i]))
		d = append(d, strconv.Itoa(r.Processcnt[i]))
		d = append(d, strconv.Itoa(r.Threadcnt[i]))
		d = append(d, strconv.Itoa(r.Pvbytes[i]))
		d = append(d, strconv.Itoa(r.Pgpool[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return strings.Join(arr, ",")
}
