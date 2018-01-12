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

func main() {
	var cpu [4]cpuinfo
	for {
		resp, err := http.Get("http://140.113.207.82:9100/metrics")
		if err != nil {
			fmt.Println("Get failed")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		allString := string(body)
		s := strings.Split(allString, "\n")
		fmt.Println(cpu[0].idle_old)
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
							s := strings.Split(x[len(x)-1], "e+0")
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
					fmt.Println(cpu[0])
				}
				//fmt.Println("########")
				if strings.Contains(s[i], "cpu1") {
					for j := 0; j < 10; j++ {
						x := strings.Split(s[i], "} ")
						switch j {
						case 0:
							cpu[1].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 1:
							cpu[1].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 2:
							s := strings.Split(x[len(x)-1], "e+0")
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
					fmt.Println(cpu[1])
				}
				//fmt.Println("########")
				if strings.Contains(s[i], "cpu2") {
					for j := 0; j < 10; j++ {
						x := strings.Split(s[i], "} ")
						switch j {
						case 0:
							cpu[2].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 1:
							cpu[2].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 2:
							s := strings.Split(x[len(x)-1], "e+0")
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
					fmt.Println(cpu[0])
				}
				//fmt.Println("########")
				if strings.Contains(s[i], "cpu3") {
					for j := 0; j < 10; j++ {
						x := strings.Split(s[i], "} ")
						switch j {
						case 0:
							cpu[3].guest, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 1:
							cpu[3].guest_nice, _ = strconv.ParseFloat(x[len(x)-1], 64)
						case 2:
							s := strings.Split(x[len(x)-1], "e+0")
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
					fmt.Println(cpu[3])
				}
			}
		}

		// calcute cpu usage
		for i := 0; i < 4; i++ {
			userfrme := cpu[i].user - cpu[i].user_old
			nicefrme := cpu[i].nice - cpu[i].nice_old
			idlefrme := cpu[i].idle - cpu[i].idle_old
			iowaitfrme := cpu[i].iowait - cpu[i].iowait_old
			irqfrme := cpu[i].irq - cpu[i].irq_old
			stealfrme := cpu[i].steal - cpu[i].steal_old
			softirqfrme := cpu[i].softirq - cpu[i].softirq_old
			systemfrme := cpu[i].system - cpu[i].system_old
			total := userfrme + systemfrme + nicefrme + idlefrme + iowaitfrme + softirqfrme + stealfrme + irqfrme
			//fmt.Println("######", cpu[i].idle, cpu[i].idle_old, idlefrme)
			cpuload := (userfrme + systemfrme + nicefrme) / (total + 0.01) * 100
			fmt.Println("total:", cpuload)
			cpu[i].user_old = cpu[i].user
			cpu[i].nice_old = cpu[i].nice
			cpu[i].idle_old = cpu[i].idle
			cpu[i].iowait_old = cpu[i].iowait
			cpu[i].irq_old = cpu[i].irq
			cpu[i].steal_old = cpu[i].steal
			cpu[i].softirq_old = cpu[i].softirq
			cpu[i].system_old = cpu[i].system

		}
		time.Sleep(1 * time.Second)
	}
}
