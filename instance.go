package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/shirou/gopsutil/v3/process"
)

type Instance struct {
	thisCmdLineRun       string
	thisCmdLineRunQuoted string
	exePathname          string
	runningProcess       *process.Process
}

func (i *Instance) Init() {
	i.exePathname, _ = filepath.Abs(os.Args[0])
	i.thisCmdLineRun = i.exePathname + " " + RUN_COMMAND
	i.thisCmdLineRunQuoted = "\"" + i.exePathname + "\" " + RUN_COMMAND
}

func (i *Instance) Refresh() error {
	i.runningProcess = nil

	pids, err := process.Pids()

	if err != nil {
		return err
	}

	for _, ipid := range pids {
		iprocess, err := process.NewProcess(ipid)

		if err != nil {
			return err
		}

		cmd, err := iprocess.Cmdline()

		if err != nil {
			return err
		}

		if cmd == i.thisCmdLineRun || cmd == i.thisCmdLineRunQuoted {
			i.runningProcess = iprocess
			return nil
		}
	}

	return nil
}

func (i *Instance) IsRunnnig() bool {
	return i.runningProcess != nil
}

func (i *Instance) Kill() error {
	if i.runningProcess == nil {
		return errors.New("process is not running")
	}

	i.runningProcess.Kill()

	return nil
}

func (i *Instance) GetExePathname() string {
	return i.exePathname
}
