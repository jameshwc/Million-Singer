package stat

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type CPU struct {
	Total, User, System, Idle uint64
}

func CalCpuUsage() *CPU {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var st CPU
	for _, s := range stat.CPUStats {
		st.Total += s.IOWait + s.IRQ + s.Idle + s.Nice + s.SoftIRQ + s.Steal + s.System + s.User
		st.User += s.User
		st.System += s.System
		st.Idle += s.Idle
	}
	return &st
}
