package main

import (
	"os"

	"github.com/shirou/gopsutil/v3/process"
)

type Instance struct {
	thisCmdLineRun       string
	thisCmdLineRunQuoted string
	runningProcess       *process.Process
}

func (i *Instance) Init() {
	i.thisCmdLineRun = os.Args[0] + " " + RUN_COMMAND
	i.thisCmdLineRunQuoted = "\"" + os.Args[0] + "\" " + RUN_COMMAND
}

func (i *Instance) IsRunnnig() bool {
	i.runningProcess = nil

	pids, err := process.Pids()

	if err != nil {
		LogDialog{}.Panicln(APP_NAME, err.Error())
	}

	for _, ipid := range pids {
		iprocess, err := process.NewProcess(ipid)

		if err != nil {
			LogDialog{}.Panicln(APP_NAME, err.Error())
			continue
		}

		cmd, err := iprocess.Cmdline()

		if err != nil {
			LogDialog{}.Panicln(APP_NAME, err.Error())
			continue
		}

		if cmd == i.thisCmdLineRun || cmd == i.thisCmdLineRunQuoted {
			i.runningProcess = iprocess
			return true
		}
	}

	return false
}

func (i *Instance) Kill() {
	if i.runningProcess == nil {
		LogDialog{}.Panicln(APP_NAME, "Process is not running.")
	}

	i.runningProcess.Kill()
}
