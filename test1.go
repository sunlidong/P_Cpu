package main

import (
    "bytes"
    //"fmt"
    "log"
    "os/exec"
    "strconv"
    "strings"
    "syscall"
    "runtime"
)

type Process struct {
    pid int
    cpu float64
}


 var num float64
func main1() {
    cmd := exec.Command("ps", "aux")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    processes := make([]*Process, 0)
    for {
        line, err := out.ReadString('\n')
        if err != nil {
            break
        }
        tokens := strings.Split(line, " ")
        ft := make([]string, 0)
        for _, t := range tokens {
            if t != "" && t != "\t" {
                ft = append(ft, t)
            }
        }
        log.Println(len(ft), ft)
        pid, err := strconv.Atoi(ft[1])
        if err != nil {
            continue
        }
        cpu, err := strconv.ParseFloat(ft[2], 64)
        if err != nil {
            log.Fatal(err)
        }
        processes = append(processes, &Process{pid, cpu})
    }

    for _, p := range processes {
        log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
        //
        num+=p.cpu
    }

    //
    log.Println("=====================================>",num)
}

import (

)

type MemStatus struct {
    All  uint32 `json:"all"`
    Used uint32 `json:"used"`
    Free uint32 `json:"free"`
    Self uint64 `json:"self"`
}

func MemStat() MemStatus {
    //自身占用
    memStat := new(runtime.MemStats)
    runtime.ReadMemStats(memStat)
    mem := MemStatus{}
    mem.Self = memStat.Alloc

    //系统占用,仅linux/mac下有效
    //system memory usage
    sysInfo := new(syscall.Sysinfo_t)
    err := syscall.Sysinfo(sysInfo)
    if err == nil {
        mem.All = sysInfo.Totalram * uint32(syscall.Getpagesize())
        mem.Free = sysInfo.Freeram * uint32(syscall.Getpagesize())
        mem.Used = mem.All - mem.Free
    }
    return mem
}
