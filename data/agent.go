package data

import (
	"time"
)

type AgentHostAgentInfo struct {
	AgentName    string
	AgentID      string
	Model        string
	Serial       string
	Ip           string
	Os           string
	Agentversion string
	ProcessCount int
	ProcessClock int
	MemorySize   int
	SwapMemory   int
}

type AgentRealTimePerf struct {
	AgentID          string
	Agenttime        time.Time
	User             int
	Sys              int
	Wait             int
	Idle             int
	ProcessorCount   int
	RunQueue         int
	BlockQueue       int
	WaitQueue        int
	PQueue           int
	PCRateUser       int
	PCRateSys        int
	MemorySize       int
	MemoryUsed       int
	MemoryPinned     int
	MemorySys        int
	MemoryUser       int
	MemoryCache      int
	Avm              int
	PagingspaceIn    int
	PaingSpaceOut    int
	FileSystemIn     int
	FileSystmeOut    int
	MemoryScan       int
	MemoryFreed      int
	SwapSize         int
	SwapUsed         int
	SwapActive       int
	Fork             int
	Exec             int
	Interupt         int
	SystemCall       int
	ContextSwitch    int
	Semaphore        int
	Msg              int
	DiskReadWrite    int
	DiskIOPS         int
	NetworkReadWrite int
	TopCommandID     int
	TopCommandCount  int
	TopUserID        int
	TopCPU           int
	TopDiskID        int
	TopvgID          int
	TOPBusy          int
	MaxPID           int
	ThreadCount      int
	PIDCount         int
	LinuxBuffer      int
	LinuxCached      int
	Linuxsrec        int
	Memused_mb       int
	IRQ              int
	SoftIRQ          int
	Swapused_MB      int
	DUSM             int
}

type AgentRealTimePID struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []AgentRealTimePIDInner
}

type AgentRealTimePIDInner struct {
	Pid        int
	Ppid       int
	Uid        int
	Cmdname    string
	Username   string
	Argname    string
	Usr        int
	Sys        int
	Usrsys     int
	Sz         int
	Rss        int
	Vmem       int
	Chario     int
	Processcnt int
	Threadcnt  int
	Handlecnt  int
	Stime      int
	Pvbytes    int
	Pgpool     int
}

type AgentRealTimeDisk struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []AgentRealTimeDiskInner
}

type AgentRealTimeDiskInner struct {
	Ioname       string
	Readrate     int
	Writerate    int
	Iops         int
	Busy         int
	Descname     string
	Readsvctime  int
	Writesvctime int
}

type AgentRealTimeNet struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []AgentRealTimeNetInner
}

type AgentRealTimeNetInner struct {
	Ioname    string
	Readrate  int
	Writerate int
	Readiops  int
	Writeiops int
	Errorps   int
	Collision int
}
