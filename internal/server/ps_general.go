package server

import (
	"fmt"
	"time"

	pscpu "github.com/shirou/gopsutil/v3/cpu"
	psdisk "github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	psnet "github.com/shirou/gopsutil/v3/net"
)

func (s *GeneralServer) genps() {
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	stats, err := psnet.IOCounters(false)
	if err != nil {
		panic(err)
	}
	for _, av := range stats {
		fmt.Println(av)
	}
	cpus, err := pscpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	for _, cv := range cpus {
		fmt.Println(cv)
	}
	cdisk, err := psdisk.Partitions(false)
	if err != nil {
		panic(err)
	}
	for _, dv := range cdisk {
		ds, err := psdisk.Usage(dv.Mountpoint)
		if err != nil {
			panic(err)
		}
		fmt.Println(ds)
	}
}
