package runmetrics_test

import (
	"context"
	"fmt"
	"log"
	"sort"
	
	"github.com/gozelle/opencensus/metric/metricdata"
	"github.com/gozelle/opencensus/metric/metricexport"
	"github.com/gozelle/opencensus/plugin/runmetrics"
)

type printExporter struct {
}

func (l *printExporter) ExportMetrics(ctx context.Context, data []*metricdata.Metric) error {
	mapData := make(map[string]metricdata.Metric, 0)
	
	for _, v := range data {
		mapData[v.Descriptor.Name] = *v
	}
	
	mapKeys := make([]string, 0, len(mapData))
	for key := range mapData {
		mapKeys = append(mapKeys, key)
	}
	sort.Strings(mapKeys)
	
	// for the sake of a simple example, we cannot use the real value here
	simpleVal := func(v interface{}) int { return 42 }
	
	for _, k := range mapKeys {
		v := mapData[k]
		fmt.Printf("%s %d\n", k, simpleVal(v.TimeSeries[0].Points[0].Value))
	}
	
	return nil
}

func ExampleEnable() {
	
	// Enable collection of runtime metrics and supply options
	err := runmetrics.Enable(runmetrics.RunMetricOptions{
		EnableCPU:    true,
		EnableMemory: true,
		Prefix:       "mayapp/",
	})
	if err != nil {
		log.Fatal(err)
	}
	
	// Use your reader/exporter to extract values
	// This part is not specific to runtime metrics and only here to make it a complete example.
	metricexport.NewReader().ReadAndExport(&printExporter{})
	
	// output:
	// mayapp/process/cpu_cgo_calls 42
	// mayapp/process/cpu_goroutines 42
	// mayapp/process/gc_cpu_fraction 42
	// mayapp/process/gc_sys 42
	// mayapp/process/heap_alloc 42
	// mayapp/process/heap_idle 42
	// mayapp/process/heap_inuse 42
	// mayapp/process/heap_objects 42
	// mayapp/process/heap_release 42
	// mayapp/process/last_gc_finished_timestamp 42
	// mayapp/process/memory_alloc 42
	// mayapp/process/memory_frees 42
	// mayapp/process/memory_lookups 42
	// mayapp/process/memory_malloc 42
	// mayapp/process/next_gc_heap_size 42
	// mayapp/process/num_forced_gc 42
	// mayapp/process/num_gc 42
	// mayapp/process/other_sys 42
	// mayapp/process/pause_total 42
	// mayapp/process/stack_inuse 42
	// mayapp/process/stack_mcache_inuse 42
	// mayapp/process/stack_mspan_inuse 42
	// mayapp/process/sys_heap 42
	// mayapp/process/sys_memory_alloc 42
	// mayapp/process/sys_stack 42
	// mayapp/process/sys_stack_mcache 42
	// mayapp/process/sys_stack_mspan 42
	// mayapp/process/total_memory_alloc 42
}
