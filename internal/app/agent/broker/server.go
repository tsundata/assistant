package broker

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

type Server struct {
	broker *Broker
}

func NewServer(broker *Broker) (*Server, error) {
	s := &Server{}
	s.broker = broker
	return s, nil
}

func (b *Server) Run() {
	t := time.NewTicker(time.Second)
	// host
	h, err := host.Info()
	if err != nil {
		b.broker.logger.Error(err.Error())
		return
	}
	writeAPI := b.broker.influx.WriteAPI(b.broker.Org, b.broker.Bucket)
	for range t.C {
		// mem
		v, err := mem.VirtualMemory()
		if err != nil {
			b.broker.logger.Error(err.Error())
			continue
		}
		// swap
		s, err := mem.SwapMemory()
		if err != nil {
			b.broker.logger.Error(err.Error())
			continue
		}
		// cpu
		c, err := cpu.Percent(time.Second, false)
		if err != nil {
			b.broker.logger.Error(err.Error())
			continue
		}
		// disk
		d, err := disk.Usage("/")
		if err != nil {
			b.broker.logger.Error(err.Error())
			continue
		}

		for i, ci := range c {
			writeAPI.WriteRecord(fmt.Sprintf("cpu_stat,host=%s,cpu=%d cpu_used=%f", h.Hostname, i+1, ci))
		}
		writeAPI.WriteRecord(fmt.Sprintf("mem_stat,host=%s mem_total=%v,mem_free=%v,mem_used=%f", h.Hostname, v.Total, v.Free, v.UsedPercent))
		writeAPI.WriteRecord(fmt.Sprintf("disk_stat,host=%s disk_total=%d,disk_free=%d,disk_used=%f", h.Hostname, d.Total, d.Free, d.UsedPercent))
		writeAPI.WriteRecord(fmt.Sprintf("swap_stat,host=%s swap_total=%d,swap_free=%d,swap_used=%f", h.Hostname, s.Total, s.Free, s.UsedPercent))
		writeAPI.Flush()
	}
}
