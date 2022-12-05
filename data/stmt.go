package data

var TableinfoStmt = `
CREATE TABLE IF NOT EXISTS public.tableinfo (			
	_tablename varchar(64) NOT NULL PRIMARY KEY,		
	_version int4 NULL,		
	_createdtime int4 NULL,		
	_updatetime int4 NULL,		
	_durationmin int4 NULL		
);			
CREATE INDEX tableinfo_idx ON public.tableinfo USING btree (_createdtime);			
`

var AgentinfoStmt = `
CREATE TABLE IF NOT EXISTS public.agentinfo (	
	_agentid serial not null primary key,
	_hostname text NULL,
	_hostnameext text NULL,
	_enabled int4 NULL,
	_connected int4 NULL,
	_updated int4 NULL,
	_shorttermbasic int4 NULL,
	_shorttermproc int4 NULL,
	_shorttermio int4 NULL,
	_shorttermcpu int4 NULL,
	_longtermbasic int4 NULL,
	_longtermproc int4 NULL,
	_longtermio int4 NULL,
	_longtermcpu int4 NULL,
	_group text NULL,
	_ipaddress text NULL,
	_pscommand text NULL,
	_logevent text NULL,
	_processevent text NULL,
	_timecheck int4 NULL,
	_disconnectedtime int4 NULL,
	_skipdatatypes int4 NULL,
	_virbasicperf int4 NULL,
	_hypervisor int4 NULL,
	_serviceevent text NULL,
	_installdate int4 NULL DEFAULT 0,
	_ibmpcrate int4 NULL DEFAULT 0,
	_updatedtime int8 NULL DEFAULT 0,
	_os text NULL,
	_fw text NULL,
	_agentversion text NULL,
	_model text NULL,
	_serial text NULL,
	_processorcount int4 NULL,
	_processorclock int4 NULL,
	_memorysize int4 NULL,
	_swapsize int4 NULL,
	_poolid int4 NULL,
	_replication int4 NULL,
	_smt int4 NULL,
	_micropar int4 NULL,
	_capped int4 NULL,
	_ec int4 NULL,
	_virtualcpu int4 NULL,
	_weight int4 NULL,
	_cpupool int4 NULL,
	_ams int4 NULL,
	_allip text NULL,
	_numanodecount int4 NULL DEFAULT 0,
	_btime int8 NULL DEFAULT 0
);	
`

var LastrealtimeperfStmt = `
CREATE UNLOGGED TABLE public.lastrealtimeperf (			
	_ontunetime int4 NULL,		
	_agentid int4 NOT NULL,		
	_hostname text NULL,		
	_user int4 NULL,		
	_sys int4 NULL,		
	_wait int4 NULL,		
	_idle int4 NULL,		
	_memoryused int4 NULL,		
	_filecache int4 NULL,		
	_memorysize int4 NULL,		
	_avm int4 NULL,		
	_swapused int4 NULL,		
	_swapsize int4 NULL,		
	_diskiorate int4 NULL,		
	_networkiorate int4 NULL,		
	_topproc text NULL,		
	_topuser text NULL,		
	_topproccount int4 NULL,		
	_topcpu int4 NULL,		
	_topdisk text NULL,		
	_topvg text NULL,		
	_topbusy int4 NULL,		
	_maxcpu int4 NULL,		
	_maxmem int4 NULL,		
	_maxswap int4 NULL,		
	_maxdisk int4 NULL,		
	_diskiops int4 NULL,		
	_networkiops int4 NULL,
	_dummy01 int4 NULL,		
	_dummy02 int4 NULL,		
	_dummy03 int4 NULL,		
	_dummy04 int4 NULL,		
	_dummy05 int4 NULL,		
	_dummy06 int4 NULL,		
	_dummy07 int4 NULL,		
	_dummy08 int4 NULL,		
	_dummy09 int4 NULL,		
	_dummy10 int4 NULL,		
	_dummy11 int4 NULL,		
	_dummy12 int4 NULL,		
	_dummy13 int4 NULL,		
	_dummy14 int4 NULL,		
	_dummy15 int4 NULL,		
	_dummy16 int4 NULL,		
	_dummy17 int4 NULL,		
	_dummy18 int4 NULL,		
	_dummy19 int4 NULL,		
	_dummy20 int4 NULL,		
	_dummy21 int4 NULL,		
	_dummy22 int4 NULL,		
	_dummy23 int4 NULL,		
	_dummy24 int4 NULL,		
	_dummy25 int4 NULL,		
	_dummy26 int4 NULL,		
	_dummy27 int4 NULL,		
	_dummy28 int4 NULL,		
	_dummy29 int4 NULL,		
	_dummy30 int4 NULL
);			
CREATE INDEX lastrealtimeperf_pkey ON public.lastrealtimeperf USING btree (_agentid);			
`

var DeviceidStmt = `
CREATE TABLE IF NOT EXISTS public.deviceid (	
	_id serial NOT NULL PRIMARY KEY,
	_name text NULL
);	
`

var DescidStmt = `
CREATE TABLE IF NOT EXISTS public.descid (	
	_id serial NOT NULL PRIMARY KEY,
	_name text NULL
);	
`

var ProcStmt = `
CREATE TABLE IF NOT EXISTS public.%s (
	_id serial NOT NULL PRIMARY KEY,	
	_name text NULL	
);		
`

var RealtimePgPrefix = `
CREATE TABLE IF NOT EXISTS public.%s (	
	_ontunetime int4 NULL,
`

var RealtimeperfStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_user int4 NULL,
_sys int4 NULL,
_wait int4 NULL,
_idle int4 NULL,
_processorcount int4 NULL,
_runqueue int4 NULL,
_blockqueue int4 NULL,
_waitqueue int4 NULL,
_pqueue int4 NULL,
_pcrateuser int4 NULL,
_pcratesys int4 NULL,
_memorysize int4 NULL,
_memoryused int4 NULL,
_memorypinned int4 NULL,
_memorysys int4 NULL,
_memoryuser int4 NULL,
_memorycache int4 NULL,
_avm int4 NULL,
_pagingspacein int4 NULL,
_pagingspaceout int4 NULL,
_filesystemin int4 NULL,
_filesystemout int4 NULL,
_memoryscan int4 NULL,
_memoryfreed int4 NULL,
_swapsize int4 NULL,
_swapused int4 NULL,
_swapactive int4 NULL,
_fork int4 NULL,
_exec int4 NULL,
_interupt int4 NULL,
_systemcall int4 NULL,
_contextswitch int4 NULL,
_semaphore int4 NULL,
_msg int4 NULL,
_diskreadwrite int4 NULL,
_diskiops int4 NULL,
_networkreadwrite int4 NULL,
_networkiops int4 NULL,
_topcommandid int4 NULL,
_topcommandcount int4 NULL,
_topuserid int4 NULL,
_topcpu int4 NULL,
_topdiskid int4 NULL,
_topvgid int4 NULL,
_topbusy int4 NULL,
_maxpid int4 NULL,
_threadcount int4 NULL,
_pidcount int4 NULL,
_linuxbuffer int4 NULL,
_linuxcached int4 NULL,
_linuxsrec int4 NULL,
_memused_mb int4 NULL,
_irq int4 NULL,
_softirq int4 NULL,
_swapused_mb int4 NULL,
_dusm int4 NULL,
PRIMARY KEY(_agentid, _ontunetime)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimeperfPgStmt = RealtimePgPrefix + RealtimeperfStmt

var RealtimeperfTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimeperf (	
	_ontunetime timestamptz NOT NULL,
` + RealtimeperfStmt + `
select create_hypertable('realtimeperf','_ontunetime', chunk_time_interval => interval '1 day');
`

var RealtimecpuStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_index int4 NULL,
_user int4 NULL,
_sys int4 NULL,
_wait int4 NULL,
_idle int4 NULL,
_runqueue int4 NULL,
_fork int4 NULL,
_exec int4 NULL,
_interupt int4 NULL,
_systemcall int4 NULL,
_contextswitch int4 NULL,
PRIMARY KEY(_agentid, _ontunetime)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimecpuPgStmt = RealtimePgPrefix + RealtimecpuStmt

var RealtimecpuTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimecpu (	
	_ontunetime timestamptz NOT NULL,
` + RealtimecpuStmt + `
select create_hypertable('realtimecpu','_ontunetime', chunk_time_interval => interval '1 day');
`

var RealtimediskStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_ionameid int4 NULL,
_readrate int4 NULL,
_writerate int4 NULL,
_iops int4 NULL,
_busy int4 NULL,
_descid int4 NULL,
_readsvctime int4 NULL,
_writesvctime int4 NULL,
PRIMARY KEY(_agentid, _ontunetime)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimediskPgStmt = RealtimePgPrefix + RealtimediskStmt

var RealtimediskTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimedisk (	
	_ontunetime timestamptz NOT NULL,
` + RealtimediskStmt + `
select create_hypertable('realtimedisk','_ontunetime', chunk_time_interval => interval '1 day');
`

var RealtimenetStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_ionameid int4 NULL,
_readrate int4 NULL,
_writerate int4 NULL,
_readiops int4 NULL,
_writeiops int4 NULL,
_errorps int4 NULL,
_collision int4 NULL,
PRIMARY KEY(_agentid, _ontunetime)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimenetPgStmt = RealtimePgPrefix + RealtimenetStmt

var RealtimenetTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimenet (	
	_ontunetime timestamptz NOT NULL,
` + RealtimenetStmt + `
select create_hypertable('realtimenet','_ontunetime', chunk_time_interval => interval '1 day');
`

var RealtimepidStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_pid int4 NULL,
_ppid int4 NULL,
_uid int4 NULL,
_cmdid int4 NULL,
_userid int4 NULL,
_argid int4 NULL,
_usr int4 NULL,
_sys int4 NULL,
_usrsys int4 NULL,
_sz int4 NULL,
_rss int4 NULL,
_vmem int4 NULL,
_chario int4 NULL,
_processcnt int4 NULL,
_threadcnt int4 NULL,
_handlecnt int4 NULL,
_stime int4 NULL,
_pvbytes int4 NULL,
_pgpool int4 NULL,
PRIMARY KEY(_agentid, _ontunetime, _cmdid, _userid, _argid)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimepidPgStmt = RealtimePgPrefix + RealtimepidStmt

var RealtimepidTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimepid (	
	_ontunetime timestamptz NOT NULL,
` + RealtimepidStmt + `
select create_hypertable('realtimepid','_ontunetime', chunk_time_interval => interval '1 day');
`

var RealtimeprocStmt = `
_agenttime int4 NULL,
_agentid int4 NULL,
_cmdid int4 NULL,
_userid int4 NULL,
_usr int4 NULL,
_sys int4 NULL,
_usrsys int4 NULL,
_sz int4 NULL,
_rss int4 NULL,
_vmem int4 NULL,
_chario int4 NULL,
_processcnt int4 NULL,
_threadcnt int4 NULL,
_pvbytes int4 NULL,
_pgpool int4 NULL,
PRIMARY KEY(_agentid, _ontunetime, _cmdid, _userid)
)	
WITH (	
autovacuum_enabled=false
);	
`

var RealtimeprocPgStmt = RealtimePgPrefix + RealtimeprocStmt

var RealtimeprocTsStmt = `
CREATE TABLE IF NOT EXISTS public.realtimeproc (	
	_ontunetime timestamptz NOT NULL,
` + RealtimeprocStmt + `
select create_hypertable('realtimeproc','_ontunetime', chunk_time_interval => interval '1 day');
`

var InsertTableinfo = `
INSERT INTO tableinfo values ($1, 0, $2, $2, 0);
`

var DeleteAgentinfoDummy = `
DELETE FROM agentinfo where _hostname like 'Dummy%';
`

var DeleteLastrealtimeperf = `
DELETE FROM lastrealtimeperf;
`

var DemoInsertAgentinfoPrefix = `
INSERT INTO agentinfo values 
`

var DemoInsertAgentinfoPostfix = `
(%d, 'DummyHost%d', 'DummyHost%d', 1, 1, 1, 2, 5, 5, 5, 600, 600, 600, 600,
 null, 'localhost', null, null, null, 1, 0, 0, 1, 0, null, %d, 0, 0,
 'Windows 11 Home', 'Unknown', 'V4', null, null, 1, 3200, 4095, 0, -1, 0, 0,
 0, 0, 0, 1, 0, 0, 0, 'localhost', 0, 0)
`

var DemoUpdateAgentinfoState = `
UPDATE agentinfo 
   set _connected=%d, _updatedtime=%d 
 where _agentid in (%s) and _enabled=1 
   and _connected=%d
   and _updatedtime<%d
`

var DemoUpdateAgentinfoReset = `
UPDATE agentinfo 
   set _connected=1, _updatedtime=$1
 where _enabled=1 
   and _connected=0
`

var DemoInsertLastrealtimeperf = `
INSERT INTO lastrealtimeperf 
(select * from unnest($1::int[], $2::int[], $3::text[], $4::int[], $5::int[], $6::int[], $7::int[], 
	$8::int[], $9::int[], $10::int[], $11::int[], $12::int[], $13::int[], $14::int[], $15::int[], 
	$16::text[], $17::text[], $18::int[], $19::int[], $20::text[], $21::text[], $22::int[], 
	$23::int[], $24::int[], $25::int[], $26::int[], $27::int[], $28::int[],
	$29::int[], $30::int[], $31::int[], $32::int[], $33::int[], $34::int[], $35::int[], $36::int[],
	$37::int[], $38::int[], $39::int[], $40::int[], $41::int[], $42::int[], $43::int[], $44::int[],
	$45::int[], $46::int[], $47::int[], $48::int[], $49::int[], $50::int[], $51::int[], $52::int[],
	$53::int[], $54::int[], $55::int[], $56::int[], $57::int[], $58::int[]))
`

var InsertAgentinfoUnnest = `
INSERT INTO agentinfo 
(select * from unnest($1::int[], $2::text[], $3::text[], $4::int[], $5::int[], $6::int[],
	$7::int[], $8::int[], $9::int[], $10::int[], $11::int[], $12::int[], $13::int[], $14::int[], 
	$15::text[], $16::text[], $17::text[], $18::text[], $19::text[], 
	$20::int[], $21::int[], $22::int[], $23::int[], $24::int[],
	$25::text[], $26::int[], $27::int[], $28::int[], $29::text[], $30::text[], $31::text[], $32::text[], $33::text[], 
	$34::int[], $35::int[], $36::int[], $37::int[], $38::int[], $39::int[], $40::int[], $41::int[], 
	$42::int[], $43::int[], $44::int[], $45::int[], $46::int[], $47::int[], $48::text[], $49::int[], $50::int[]))
`

var UpdateAgentinfo = `
UPDATE agentinfo
   SET _hostnameext=$1, _ipaddress=$2, _allip=$2, _model=$3, _serial=$4,
       _os=$5, _agentversion=$6, _processorcount=$7, _processorclock=$8,
	   _memorysize=$9, _swapsize=$10, _updatedtime=$11
 WHERE _hostname=$12
`

var InsertRealtimePerf = `
INSERT INTO %s
VALUES (
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
	$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,
	$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59
)
`

var InsertRealtimeCpu = `
INSERT INTO %s
VALUES (
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
)
`

var InsertSimpleTable = `
INSERT INTO %s (_name) values ($1)
`

var InsertRealtimePidPg = `
INSERT INTO %s
(select * from unnest(
	$1::int[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[],$12::int[],$13::int[],$14::int[],
	$15::int[],$16::int[],$17::int[],$18::int[],$19::int[],$20::int[],$21::int[],$22::int[]
))
`

var InsertRealtimePidTs = `
INSERT INTO %s
(select * from unnest(
	$1::timestamptz[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[],$12::int[],$13::int[],$14::int[],
	$15::int[],$16::int[],$17::int[],$18::int[],$19::int[],$20::int[],$21::int[],$22::int[]
))
`

var InsertRealtimeProcPg = `
INSERT INTO %s
(select * from unnest(
	$1::int[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[],$12::int[],$13::int[],$14::int[],
	$15::int[],$16::int[]
))
`

var InsertRealtimeProcTs = `
INSERT INTO %s
(select * from unnest(
	$1::timestamptz[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[],$12::int[],$13::int[],$14::int[],
	$15::int[],$16::int[]
))
`

var InsertRealtimeDiskPg = `
INSERT INTO %s
(select * from unnest(
	$1::int[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[]
))
`

var InsertRealtimeDiskTs = `
INSERT INTO %s
(select * from unnest(
	$1::timestamptz[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[],$11::int[]
))
`

var InsertRealtimeNetPg = `
INSERT INTO %s
(select * from unnest(
	$1::int[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[]
))
`

var InsertRealtimeNetTs = `
INSERT INTO %s
(select * from unnest(
	$1::timestamptz[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],
	$8::int[],$9::int[],$10::int[]
))
`
