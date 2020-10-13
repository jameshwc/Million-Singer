package stat

import (
	"github.com/jameshwc/Million-Singer/pkg/log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type CPU struct {
	Total, User, System, Idle uint64
}

type Mem struct {
	Used, Available, Total uint64
}

type Disk struct {
	Used, Total uint64
}

func (m *Mem) Usage() float64 {
	return 1 - float64(m.Available)/float64(m.Total)
}

func (d *Disk) Usage() float64 {
	return float64(d.Used) / float64(d.Total)
}

type Server struct {
	CPU
	Mem
	Disk
}

func CalCpuUsage() CPU {
	var st CPU
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Error("stat read fail")
		return st
	}
	for _, s := range stat.CPUStats {
		st.Total += s.IOWait + s.IRQ + s.Idle + s.Nice + s.SoftIRQ + s.Steal + s.System + s.User
		st.User += s.User
		st.System += s.System
		st.Idle += s.Idle
	}
	return st
}

func CalMemUsage() Mem {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Error("stat read fail")
	}
	return Mem{Used: stat.MemTotal - stat.MemAvailable, Available: stat.MemAvailable, Total: stat.MemTotal}
}

func CalDiskUsage() Disk {
	stat, err := linuxproc.ReadDisk("/")
	if err != nil {
		log.Error("stat read fail")
	}
	return Disk{Used: stat.Used, Total: stat.All}
}

func GetServer() *Server {
	return &Server{CalCpuUsage(), CalMemUsage(), CalDiskUsage()}
}
