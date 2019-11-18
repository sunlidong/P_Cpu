package main

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/mem"
)

func main() {
    v, _ := mem.VirtualMemory()
    c, _ := cpu.Info()
    d, _ := disk.Usage("/")
    n, _ := host.Info()
rep:=strconv.FormatFloat(v.UsedPercent, 'E', -1, 64)
    fmt.Printf("        Mem       : %v GB  Free: %v MB Usage:%f%%\n", v.Total/4/1024/1024/1024/1024, v.Free/4/1024/1024/1024, v.UsedPercent)
    if len(c) > 1 {
        for _, sub_cpu := range c {
            modelname := sub_cpu.ModelName
            cores := sub_cpu.Cores
            fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)
        }
    } else {
        sub_cpu := c[0]
        modelname := sub_cpu.ModelName
        cores := sub_cpu.Cores
        fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)

    }
    fmt.Printf("        HD        : %v GB  Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
    fmt.Printf("        OS        : %v   %v  \n", n.OS, n.PlatformVersion)
    fmt.Printf("        Hostname  : %v  \n", n.Hostname)
	
	fmt.Println("rep==>",rep)

}
