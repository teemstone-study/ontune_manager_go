package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type TableSetArray interface {
	SetData(data interface{}, agentid int, ids ...int)
}

type TableSetArrayInner interface {
	SetData(data interface{}, agentid int, agenttime time.Time, ids ...int)
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

func (r *RealtimeperfPgArray) SetData(data interface{}, agentid int, ids ...int) {
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

func (r *RealtimeperfPgArray) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePerfUnnest, tablename, "int")
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

func (a *RealtimeperfPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.User[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Wait[i]))
		d = append(d, strconv.Itoa(a.Idle[i]))
		d = append(d, strconv.Itoa(a.Processorcount[i]))
		d = append(d, strconv.Itoa(a.Runqueue[i]))
		d = append(d, strconv.Itoa(a.Blockqueue[i]))
		d = append(d, strconv.Itoa(a.Waitqueue[i]))
		d = append(d, strconv.Itoa(a.Pqueue[i]))
		d = append(d, strconv.Itoa(a.Pcrateuser[i]))
		d = append(d, strconv.Itoa(a.Pcratesys[i]))
		d = append(d, strconv.Itoa(a.Memorysize[i]))
		d = append(d, strconv.Itoa(a.Memoryused[i]))
		d = append(d, strconv.Itoa(a.Memorypinned[i]))
		d = append(d, strconv.Itoa(a.Memorysys[i]))
		d = append(d, strconv.Itoa(a.Memoryuser[i]))
		d = append(d, strconv.Itoa(a.Memorycache[i]))
		d = append(d, strconv.Itoa(a.Avm[i]))
		d = append(d, strconv.Itoa(a.Pagingspacein[i]))
		d = append(d, strconv.Itoa(a.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(a.Filesystemin[i]))
		d = append(d, strconv.Itoa(a.Filesystemout[i]))
		d = append(d, strconv.Itoa(a.Memoryscan[i]))
		d = append(d, strconv.Itoa(a.Memoryfreed[i]))
		d = append(d, strconv.Itoa(a.Swapsize[i]))
		d = append(d, strconv.Itoa(a.Swapused[i]))
		d = append(d, strconv.Itoa(a.Swapactive[i]))
		d = append(d, strconv.Itoa(a.Fork[i]))
		d = append(d, strconv.Itoa(a.Exec[i]))
		d = append(d, strconv.Itoa(a.Interupt[i]))
		d = append(d, strconv.Itoa(a.Systemcall[i]))
		d = append(d, strconv.Itoa(a.Contextswitch[i]))
		d = append(d, strconv.Itoa(a.Semaphore[i]))
		d = append(d, strconv.Itoa(a.Msg[i]))
		d = append(d, strconv.Itoa(a.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(a.Diskiops[i]))
		d = append(d, strconv.Itoa(a.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(a.Networkiops[i]))
		d = append(d, strconv.Itoa(a.Topcommandid[i]))
		d = append(d, strconv.Itoa(a.Topcommandcount[i]))
		d = append(d, strconv.Itoa(a.Topuserid[i]))
		d = append(d, strconv.Itoa(a.Topcpu[i]))
		d = append(d, strconv.Itoa(a.Topdiskid[i]))
		d = append(d, strconv.Itoa(a.Topvgid[i]))
		d = append(d, strconv.Itoa(a.Topbusy[i]))
		d = append(d, strconv.Itoa(a.Maxpid[i]))
		d = append(d, strconv.Itoa(a.Threadcount[i]))
		d = append(d, strconv.Itoa(a.Pidcount[i]))
		d = append(d, strconv.Itoa(a.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(a.Linuxcached[i]))
		d = append(d, strconv.Itoa(a.Linuxsrec[i]))
		d = append(d, strconv.Itoa(a.Memused_Mb[i]))
		d = append(d, strconv.Itoa(a.Irq[i]))
		d = append(d, strconv.Itoa(a.Softirq[i]))
		d = append(d, strconv.Itoa(a.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(a.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
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

func (r *RealtimeperfTsArray) SetData(data interface{}, agentid int, ids ...int) {
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

func (r *RealtimeperfTsArray) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimePerfUnnest, tablename, "timestamptz")
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

func (a *RealtimeperfTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.User[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Wait[i]))
		d = append(d, strconv.Itoa(a.Idle[i]))
		d = append(d, strconv.Itoa(a.Processorcount[i]))
		d = append(d, strconv.Itoa(a.Runqueue[i]))
		d = append(d, strconv.Itoa(a.Blockqueue[i]))
		d = append(d, strconv.Itoa(a.Waitqueue[i]))
		d = append(d, strconv.Itoa(a.Pqueue[i]))
		d = append(d, strconv.Itoa(a.Pcrateuser[i]))
		d = append(d, strconv.Itoa(a.Pcratesys[i]))
		d = append(d, strconv.Itoa(a.Memorysize[i]))
		d = append(d, strconv.Itoa(a.Memoryused[i]))
		d = append(d, strconv.Itoa(a.Memorypinned[i]))
		d = append(d, strconv.Itoa(a.Memorysys[i]))
		d = append(d, strconv.Itoa(a.Memoryuser[i]))
		d = append(d, strconv.Itoa(a.Memorycache[i]))
		d = append(d, strconv.Itoa(a.Avm[i]))
		d = append(d, strconv.Itoa(a.Pagingspacein[i]))
		d = append(d, strconv.Itoa(a.Pagingspaceout[i]))
		d = append(d, strconv.Itoa(a.Filesystemin[i]))
		d = append(d, strconv.Itoa(a.Filesystemout[i]))
		d = append(d, strconv.Itoa(a.Memoryscan[i]))
		d = append(d, strconv.Itoa(a.Memoryfreed[i]))
		d = append(d, strconv.Itoa(a.Swapsize[i]))
		d = append(d, strconv.Itoa(a.Swapused[i]))
		d = append(d, strconv.Itoa(a.Swapactive[i]))
		d = append(d, strconv.Itoa(a.Fork[i]))
		d = append(d, strconv.Itoa(a.Exec[i]))
		d = append(d, strconv.Itoa(a.Interupt[i]))
		d = append(d, strconv.Itoa(a.Systemcall[i]))
		d = append(d, strconv.Itoa(a.Contextswitch[i]))
		d = append(d, strconv.Itoa(a.Semaphore[i]))
		d = append(d, strconv.Itoa(a.Msg[i]))
		d = append(d, strconv.Itoa(a.Diskreadwrite[i]))
		d = append(d, strconv.Itoa(a.Diskiops[i]))
		d = append(d, strconv.Itoa(a.Networkreadwrite[i]))
		d = append(d, strconv.Itoa(a.Networkiops[i]))
		d = append(d, strconv.Itoa(a.Topcommandid[i]))
		d = append(d, strconv.Itoa(a.Topcommandcount[i]))
		d = append(d, strconv.Itoa(a.Topuserid[i]))
		d = append(d, strconv.Itoa(a.Topcpu[i]))
		d = append(d, strconv.Itoa(a.Topdiskid[i]))
		d = append(d, strconv.Itoa(a.Topvgid[i]))
		d = append(d, strconv.Itoa(a.Topbusy[i]))
		d = append(d, strconv.Itoa(a.Maxpid[i]))
		d = append(d, strconv.Itoa(a.Threadcount[i]))
		d = append(d, strconv.Itoa(a.Pidcount[i]))
		d = append(d, strconv.Itoa(a.Linuxbuffer[i]))
		d = append(d, strconv.Itoa(a.Linuxcached[i]))
		d = append(d, strconv.Itoa(a.Linuxsrec[i]))
		d = append(d, strconv.Itoa(a.Memused_Mb[i]))
		d = append(d, strconv.Itoa(a.Irq[i]))
		d = append(d, strconv.Itoa(a.Softirq[i]))
		d = append(d, strconv.Itoa(a.Swapused_Mb[i]))
		d = append(d, strconv.Itoa(a.Dusm[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
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

func (r *RealtimecpuPgArray) SetData(data interface{}, agentid int, ids ...int) {
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

func (r *RealtimecpuPgArray) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeCpuUnnest, tablename, "int")
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

func (a *RealtimecpuPgArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Index[i]))
		d = append(d, strconv.Itoa(a.User[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Wait[i]))
		d = append(d, strconv.Itoa(a.Idle[i]))
		d = append(d, strconv.Itoa(a.Runqueue[i]))
		d = append(d, strconv.Itoa(a.Fork[i]))
		d = append(d, strconv.Itoa(a.Exec[i]))
		d = append(d, strconv.Itoa(a.Interupt[i]))
		d = append(d, strconv.Itoa(a.Systemcall[i]))
		d = append(d, strconv.Itoa(a.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
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

func (r *RealtimecpuTsArray) SetData(data interface{}, agentid int, ids ...int) {
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

func (r *RealtimecpuTsArray) GetInsertStmt(tablename string) string {
	return fmt.Sprintf(InsertRealtimeCpuUnnest, tablename, "timestamptz")
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

func (a *RealtimecpuTsArray) GetArrString() []string {
	arr := make([]string, 0)
	size := len(a.Agentid)

	for i := 0; i < size; i++ {
		d := make([]string, 0)
		d = append(d, fmt.Sprintf("%v", a.Ontunetime[i]))
		d = append(d, strconv.Itoa(a.Agenttime[i]))
		d = append(d, strconv.Itoa(a.Agentid[i]))
		d = append(d, strconv.Itoa(a.Index[i]))
		d = append(d, strconv.Itoa(a.User[i]))
		d = append(d, strconv.Itoa(a.Sys[i]))
		d = append(d, strconv.Itoa(a.Wait[i]))
		d = append(d, strconv.Itoa(a.Idle[i]))
		d = append(d, strconv.Itoa(a.Runqueue[i]))
		d = append(d, strconv.Itoa(a.Fork[i]))
		d = append(d, strconv.Itoa(a.Exec[i]))
		d = append(d, strconv.Itoa(a.Interupt[i]))
		d = append(d, strconv.Itoa(a.Systemcall[i]))
		d = append(d, strconv.Itoa(a.Contextswitch[i]))

		arr = append(arr, strings.Join(d, ","))
	}

	return arr
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

func (r *RealtimediskPgArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimediskPgArray) GetArrString() []string {
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

func (r *RealtimediskTsArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimediskTsArray) GetArrString() []string {
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
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetPgArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimenetPgArray) GetArrString() []string {
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
	r.Agentid = append(r.Agentid, ids[0])
	r.Ionameid = append(r.Ionameid, ids[1])
	r.Readrate = append(r.Readrate, d.Readrate)
	r.Writerate = append(r.Writerate, d.Writerate)
	r.Readiops = append(r.Readiops, d.Readiops)
	r.Writeiops = append(r.Writeiops, d.Writeiops)
	r.Errorps = append(r.Errorps, d.Errorps)
	r.Collision = append(r.Collision, d.Collision)
}

func (r *RealtimenetTsArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimenetTsArray) GetArrString() []string {
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

func (r *RealtimepidPgArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimepidPgArray) GetArrString() []string {
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

func (r *RealtimepidTsArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimepidTsArray) GetArrString() []string {
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

func (r *RealtimeprocPgArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimeprocPgArray) GetArrString() []string {
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

func (r *RealtimeprocTsArray) GetInsertStmt(tablename string) string {
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

func (a *RealtimeprocTsArray) GetArrString() []string {
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
