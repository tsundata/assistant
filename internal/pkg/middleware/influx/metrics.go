package influx

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tsundata/assistant/internal/pkg/util"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

func PushGoServerMetrics(in influxdb2.Client, service, org, bucket string) {
	numPauseQuantiles := 5
	hostname, _ := os.Hostname()
	writeAPI := in.WriteAPI(org, bucket)
	for range time.Tick(5 * time.Second) {
		start := time.Now()
		numGoroutines := runtime.NumGoroutine()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		gcStats := &debug.GCStats{
			PauseQuantiles: make([]time.Duration, numPauseQuantiles),
		}
		debug.ReadGCStats(gcStats)
		duration := time.Since(start)

		metricsCostNs := uint64(duration.Nanoseconds())
		goroutineNum := uint64(numGoroutines)
		alloc := memStats.Alloc
		heapAlloc := memStats.HeapAlloc
		connsNum := util.GetSocketCount()
		fdsNum := util.GetFDCount()
		gcNum := uint64(gcStats.NumGC)

		var pauseQuantiles strings.Builder
		for idx := 0; idx < numPauseQuantiles; idx++ {
			percent := idx * 25
			pauseQuantiles.WriteString(fmt.Sprintf("%s=%d",
				fmt.Sprintf("gc_pause_ns_%d", percent),
				uint64(gcStats.PauseQuantiles[idx].Nanoseconds())))
			if idx != numPauseQuantiles-1 {
				pauseQuantiles.WriteString(",")
			}
		}

		writeAPI.WriteRecord(fmt.Sprintf("server_stat,host=%s,service=%s metrics_cost_ns=%d,goroutine_num=%d,alloc=%d,heap_alloc=%d,conns_num=%d,fds_num=%d,gc_num=%d,%s",
			hostname, service, metricsCostNs, goroutineNum, alloc, heapAlloc, connsNum, fdsNum, gcNum, pauseQuantiles.String()))
		writeAPI.Flush()
	}
}
