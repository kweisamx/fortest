package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type cpuinfo struct {
	guest          float64
	guest_nice     float64
	idle           float64
	iowait         float64
	irq            float64
	nice           float64
	softirq        float64
	steal          float64
	system         float64
	user           float64
	guest_old      float64
	guest_nice_0ld float64
	idle_old       float64
	iowait_old     float64
	irq_old        float64
	nice_old       float64
	softirq_old    float64
	steal_old      float64
	system_old     float64
	user_old       float64
}
type meminfo struct {
	memtotal float64
	memfree  float64
	membuf   float64
	memcah   float64
}

func get_cpu_load(url string, cpu []cpuinfo, cpu_num int) float64 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	allString := string(body)
	s := strings.Split(allString, "\n")
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "node_cpu{cpu=") {
			if strings.Contains(s[i], "cpu0") {
				for j := 0; j < 10; j++ {
					x := strings.Split(s[i], "} ")
					switch j {
					case 0:
						cpu[0].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 1:
						cpu[0].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 2:
						s := strings.Split(x[len(x)-1], "e+")
						value, _ := strconv.ParseFloat(s[0], 64)
						ten, _ := strconv.ParseFloat(s[1], 64)
						cpu[0].idle = value * math.Pow(10, ten)
					case 3:
						cpu[0].iowait, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 4:
						cpu[0].irq, _ = strconv.ParseFloat(strings.TrimSpace(x[len(x)-1]), 64)
					case 5:
						cpu[0].nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 6:
						cpu[0].softirq, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 7:
						cpu[0].steal, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 8:
						cpu[0].system, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 9:
						cpu[0].user, _ = strconv.ParseFloat(x[len(x)-1], 64)

					}
					i++
				}
				//fmt.Println(cpu[0])
			}
			if strings.Contains(s[i], "cpu1") {
				for j := 0; j < 10; j++ {
					x := strings.Split(s[i], "} ")
					switch j {
					case 0:
						cpu[1].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 1:
						cpu[1].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 2:
						s := strings.Split(x[len(x)-1], "e+")
						value, _ := strconv.ParseFloat(s[0], 64)
						ten, _ := strconv.ParseFloat(s[1], 64)
						cpu[1].idle = value * math.Pow(10, ten)
					case 3:
						cpu[1].iowait, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 4:
						cpu[1].irq, _ = strconv.ParseFloat(strings.TrimSpace(x[len(x)-1]), 64)
					case 5:
						cpu[1].nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 6:
						cpu[1].softirq, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 7:
						cpu[1].steal, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 8:
						cpu[1].system, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 9:
						cpu[1].user, _ = strconv.ParseFloat(x[len(x)-1], 64)

					}
					i++
				}
			}
			if strings.Contains(s[i], "cpu2") {
				for j := 0; j < 10; j++ {
					x := strings.Split(s[i], "} ")
					switch j {
					case 0:
						cpu[2].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 1:
						cpu[2].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 2:
						s := strings.Split(x[len(x)-1], "e+")
						value, _ := strconv.ParseFloat(s[0], 64)
						ten, _ := strconv.ParseFloat(s[1], 64)
						cpu[2].idle = value * math.Pow(10, ten)
					case 3:
						cpu[2].iowait, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 4:
						cpu[2].irq, _ = strconv.ParseFloat(strings.TrimSpace(x[len(x)-1]), 64)
					case 5:
						cpu[2].nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 6:
						cpu[2].softirq, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 7:
						cpu[2].steal, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 8:
						cpu[2].system, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 9:
						cpu[2].user, _ = strconv.ParseFloat(x[len(x)-1], 64)

					}
					i++
				}
			}
			if strings.Contains(s[i], "cpu3") {
				for j := 0; j < 10; j++ {
					x := strings.Split(s[i], "} ")
					switch j {
					case 0:
						cpu[3].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 1:
						cpu[3].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 2:
						s := strings.Split(x[len(x)-1], "e+")
						value, _ := strconv.ParseFloat(s[0], 64)
						ten, _ := strconv.ParseFloat(s[1], 64)
						cpu[3].idle = value * math.Pow(10, ten)
					case 3:
						cpu[3].iowait, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 4:
						cpu[3].irq, _ = strconv.ParseFloat(strings.TrimSpace(x[len(x)-1]), 64)
					case 5:
						cpu[3].nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 6:
						cpu[3].softirq, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 7:
						cpu[3].steal, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 8:
						cpu[3].system, _ = strconv.ParseFloat(x[len(x)-1], 64)
					case 9:
						cpu[3].user, _ = strconv.ParseFloat(x[len(x)-1], 64)

					}
					i++
				}
			}
		}
	}
	var total float64
	average := 0.0
	// calcute cpu usage
	for i := 0; i < cpu_num; i++ {
		userfrme := cpu[i].user - cpu[i].user_old
		nicefrme := cpu[i].nice - cpu[i].nice_old
		idlefrme := cpu[i].idle - cpu[i].idle_old
		iowaitfrme := cpu[i].iowait - cpu[i].iowait_old
		irqfrme := cpu[i].irq - cpu[i].irq_old
		stealfrme := cpu[i].steal - cpu[i].steal_old
		softirqfrme := cpu[i].softirq - cpu[i].softirq_old
		systemfrme := cpu[i].system - cpu[i].system_old
		total = userfrme + systemfrme + nicefrme + idlefrme + iowaitfrme + softirqfrme + stealfrme + irqfrme
		cpuload := (userfrme + systemfrme + nicefrme) / (total + 0.01) * 100
		//fmt.Println("total:", cpuload)
		cpu[i].user_old = cpu[i].user
		cpu[i].nice_old = cpu[i].nice
		cpu[i].idle_old = cpu[i].idle
		cpu[i].iowait_old = cpu[i].iowait
		cpu[i].irq_old = cpu[i].irq
		cpu[i].steal_old = cpu[i].steal
		cpu[i].softirq_old = cpu[i].softirq
		cpu[i].system_old = cpu[i].system
		average += cpuload
	}
	return average / float64(cpu_num)
}
func get_mem_load(url string, mem meminfo) float64 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	allString := string(body)
	s := strings.Split(allString, "\n")
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "node_memory_MemTotal") && !strings.Contains(s[i], "#") {
			x := strings.Split(s[i], " ")
			a := strings.Split(x[len(x)-1], "e+")
			value, _ := strconv.ParseFloat(a[0], 64)
			ten, _ := strconv.ParseFloat(a[1], 64)
			mem.memtotal = value * math.Pow(10, ten)
		}
		if strings.Contains(s[i], "node_memory_MemFree") && !strings.Contains(s[i], "#") {
			x := strings.Split(s[i], " ")
			a := strings.Split(x[len(x)-1], "e+")
			value, _ := strconv.ParseFloat(a[0], 64)
			ten, _ := strconv.ParseFloat(a[1], 64)
			mem.memfree = value * math.Pow(10, ten)
		}
		if strings.Contains(s[i], "node_memory_Cached") && !strings.Contains(s[i], "#") {
			x := strings.Split(s[i], " ")
			a := strings.Split(x[len(x)-1], "e+")
			value, _ := strconv.ParseFloat(a[0], 64)
			ten, _ := strconv.ParseFloat(a[1], 64)
			mem.memcah = value * math.Pow(10, ten)
		}
		if strings.Contains(s[i], "node_memory_Buffers") && !strings.Contains(s[i], "#") {
			x := strings.Split(s[i], " ")
			a := strings.Split(x[len(x)-1], "e+")
			value, _ := strconv.ParseFloat(a[0], 64)
			ten, _ := strconv.ParseFloat(a[1], 64)
			mem.membuf = value * math.Pow(10, ten)
		}
	}
	return (mem.memtotal - mem.memfree - mem.membuf - mem.memcah) / mem.memtotal
}
func warning(vaule float64, limit float64, element string, node string){
    if (vaule > limit){
        fmt.Println("warning, the",node," ", element, "is more than limit")
    }else{
        fmt.Println(node, element,"fine")
    }

}
func main() {
	m2_cpu := make([]cpuinfo, 4)
	m2_num := 2
	var m2_mem meminfo

	master_cpu := make([]cpuinfo, 4)
	master_num := 4
	var master_mem meminfo

	url_m2 := "http://140.113.207.82:9100/metrics"
	url_master := "http://140.113.207.84:9100/metrics"
	for {

		a := get_cpu_load(url_m2, m2_cpu, m2_num)
		a_mem := get_mem_load(url_m2, m2_mem)
		//fmt.Println("m2 cpu_load:", a, "m2_mem_load", a_mem*100)

        warning(a, 70, "cpu", "m2")
        warning(a_mem, 70, "mem", "m2")

		b := get_cpu_load(url_master, master_cpu, master_num)
		b_mem := get_mem_load(url_master, master_mem)

        warning(b, 70, "cpu", "master")
        warning(b_mem, 70, "mem", "master")

		//fmt.Println("master cpu_load:", b, "master_mem_load", b_mem*100)
        fmt.Println("ave_cpu", (a+b)/2, "ave_mem", (a_mem + b_mem)/2)
		time.Sleep(2 * time.Second)
	}
}
