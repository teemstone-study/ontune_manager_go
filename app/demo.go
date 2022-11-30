package app

import (
	"fmt"
	"manager/data"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetDemoAgentinfo(arr chan<- *data.AgentinfoArr, hostcount int) {
	ts := time.Now().Unix()

	var agentinfo_arr data.AgentinfoArr
	for i := 0; i < hostcount; i++ {
		demo_data := &data.Agentinfo{
			Agentid:          i + 1,
			Hostname:         fmt.Sprintf("DummyAgent%d", i+1),
			Hostnameext:      fmt.Sprintf("DummyAgent%d", i+1),
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
			Ipaddress:        "localhost",
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
			Os:               "-",
			Fw:               "-",
			Agentversion:     "V1",
			Model:            "-",
			Serial:           "-",
			Processorcount:   rand.Intn(4),
			Processorclock:   rand.Intn(3200),
			Memorysize:       rand.Intn(32000),
			Swapsize:         0,
			Poolid:           -1,
			Replication:      0,
			Smt:              0,
			Micropar:         0,
			Capped:           0,
			Ec:               -1,
			Virtualcpu:       rand.Intn(4),
			Weight:           0,
			Cpupool:          0,
			Ams:              0,
			Allip:            "localhost",
			Numanodecount:    0,
			Btime:            0,
		}
		agentinfo_arr.SetData(*demo_data)
	}
	arr <- &agentinfo_arr
}

func GetDemoChangeStateAgent(agents chan<- []int, agent_str chan<- string, demo DemoInfo) {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		random_agent_map := make(map[int]struct{})
		random_agents := make([]int, 0)
		random_agent_str := make([]string, 0)
		for i := 0; i < demo.HostChangeCount; i++ {
			agentid := rand.Intn(demo.HostCount)
			if _, ok := random_agent_map[agentid]; ok {
				continue
			} else {
				random_agent_map[agentid] = struct{}{}
				random_agents = append(random_agents, agentid)
				random_agent_str = append(random_agent_str, strconv.Itoa(agentid))
			}
		}
		agents <- random_agents
		agent_str <- strings.Join(random_agent_str, ",")
		time.Sleep(time.Second * time.Duration(demo.Interval))
	}
}

func GetDemoLastrealtimeperf(lrtp chan<- *data.LastrealtimeperfArr, demo DemoInfo) {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		random_agent_map := make(map[int]struct{})
		ts := time.Now().Unix()

		var lastrealtimeperf_arr data.LastrealtimeperfArr
		for i := 0; i < demo.BptCount; i++ {
			agentid := rand.Intn(demo.HostCount)
			if _, ok := random_agent_map[agentid]; ok {
				continue
			} else {
				random_agent_map[agentid] = struct{}{}
				demo_data := &data.Lastrealtimeperf{
					Ontunetime:    ts,
					Agentid:       agentid,
					Hostname:      fmt.Sprintf("DummyAgent%d", agentid),
					User:          rand.Intn(100),
					Sys:           rand.Intn(100),
					Wait:          rand.Intn(100),
					Idle:          rand.Intn(100),
					Memoryused:    rand.Intn(10000),
					Filecache:     rand.Intn(10000),
					Memorysize:    rand.Intn(10000),
					Avm:           rand.Intn(10000),
					Swapused:      rand.Intn(10000),
					Swapsize:      rand.Intn(10000),
					Diskiorate:    rand.Intn(10000),
					Networkiorate: rand.Intn(10000),
					Topproc:       "ontuned",
					Topuser:       "root",
					Topproccount:  rand.Intn(10),
					Topcpu:        rand.Intn(10000),
					Topdisk:       "-",
					Topvg:         "-",
					Topbusy:       0,
					Maxcpu:        rand.Intn(100),
					Maxmem:        rand.Intn(10000),
					Maxswap:       rand.Intn(10000),
					Maxdisk:       rand.Intn(100),
					Diskiops:      rand.Intn(100),
					Networkiops:   rand.Intn(100),
					Dummy01:       rand.Intn(100),
					Dummy02:       rand.Intn(100),
					Dummy03:       rand.Intn(100),
					Dummy04:       rand.Intn(100),
					Dummy05:       rand.Intn(100),
					Dummy06:       rand.Intn(100),
					Dummy07:       rand.Intn(100),
					Dummy08:       rand.Intn(100),
					Dummy09:       rand.Intn(100),
					Dummy10:       rand.Intn(100),
					Dummy11:       rand.Intn(100),
					Dummy12:       rand.Intn(100),
					Dummy13:       rand.Intn(100),
					Dummy14:       rand.Intn(100),
					Dummy15:       rand.Intn(100),
					Dummy16:       rand.Intn(100),
					Dummy17:       rand.Intn(100),
					Dummy18:       rand.Intn(100),
					Dummy19:       rand.Intn(100),
					Dummy20:       rand.Intn(100),
					Dummy21:       rand.Intn(100),
					Dummy22:       rand.Intn(100),
					Dummy23:       rand.Intn(100),
					Dummy24:       rand.Intn(100),
					Dummy25:       rand.Intn(100),
					Dummy26:       rand.Intn(100),
					Dummy27:       rand.Intn(100),
					Dummy28:       rand.Intn(100),
					Dummy29:       rand.Intn(100),
					Dummy30:       rand.Intn(100),
				}
				lastrealtimeperf_arr.SetData(*demo_data)
			}
		}

		lrtp <- &lastrealtimeperf_arr
		time.Sleep(time.Second * time.Duration(demo.Interval))
	}
}
