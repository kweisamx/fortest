package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"store"
	"strconv"
	"strings"
	"time"
)

type NodeInfo struct {
	CPUUsage    float64
	MemoryUsage float64
	PodNum      int
}

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
						// beacause idle maybe be a big number
						// so we need to set two case
						// one has exportential num, other not
						if strings.Contains(s[i], "e+") {
							s := strings.Split(x[len(x)-1], "e+")
							value, _ := strconv.ParseFloat(s[0], 64)
							ten, _ := strconv.ParseFloat(s[1], 64)
							cpu[0].idle = value * math.Pow(10, ten)
						} else {
							cpu[0].idle, _ = strconv.ParseFloat(x[len(x)-1], 64)
						}
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
						if strings.Contains(s[i], "e+") {
							s := strings.Split(x[len(x)-1], "e+")
							value, _ := strconv.ParseFloat(s[0], 64)
							ten, _ := strconv.ParseFloat(s[1], 64)
							cpu[1].idle = value * math.Pow(10, ten)
						} else {
							cpu[1].idle, _ = strconv.ParseFloat(x[len(x)-1], 64)
						}
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
						if strings.Contains(s[i], "e+") {
							s := strings.Split(x[len(x)-1], "e+")
							value, _ := strconv.ParseFloat(s[0], 64)
							ten, _ := strconv.ParseFloat(s[1], 64)
							cpu[2].idle = value * math.Pow(10, ten)
						} else {
							cpu[2].idle, _ = strconv.ParseFloat(x[len(x)-1], 64)
						}
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
						if strings.Contains(s[i], "e+") {
							s := strings.Split(x[len(x)-1], "e+")
							value, _ := strconv.ParseFloat(s[0], 64)
							ten, _ := strconv.ParseFloat(s[1], 64)
							cpu[3].idle = value * math.Pow(10, ten)
						} else {
							cpu[3].idle, _ = strconv.ParseFloat(x[len(x)-1], 64)
						}
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
	return ((mem.memtotal - mem.memfree - mem.membuf - mem.memcah) / mem.memtotal) * 100
}
func warning(value float64, limit float64, element string, node string) {
	if value > limit {
		fmt.Println(node, element, value, ",more than limit")
	} else {
		fmt.Println(node, element, value, ",fine")
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func algo(Node1 NodeInfo, Node2 NodeInfo, Node3 NodeInfo, a float64) int32 {
	n1 := 0
	n2 := 0
	n3 := 0
	if Node1.PodNum > 0 {
		n1 = 1
	}
	if Node2.PodNum > 0 {
		n2 = 1
	}
	if Node3.PodNum > 0 {
		n3 = 1
	}
	if (n1 + n2 + n3) == 0 {
		return 1
	}
	avg_cpu := (Node1.CPUUsage*float64(n1) + Node2.CPUUsage*float64(n2) + Node3.CPUUsage*float64(n3)) / float64(n1+n2+n3)
	avg_mem := (Node1.MemoryUsage*float64(n1) + Node2.MemoryUsage*float64(n2) + Node3.MemoryUsage*float64(n3)) / float64(n1+n2+n3)

	if Node1.CPUUsage > 50 || Node2.CPUUsage > 50 || Node1.CPUUsage > 50 || Node1.MemoryUsage > 30 || Node2.MemoryUsage > 30 || Node3.MemoryUsage > 30 {
		if int32((avg_cpu*a+avg_mem*(1-a))/10) < 1 {
			return 1
		}
		return int32((avg_cpu*a + avg_mem*(1-a)) / 10)
	}
	return 1
}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("please input 3 args, ex: /bin [replica name] [refresh time] [a vale]")
	}
	storeTime, _ := strconv.Atoi(os.Args[2])
	a, _ := strconv.ParseFloat(os.Args[3], 64)
	store.CsvInit(storeTime, int(a*10))
	alltime := 0
	url_m1 := "http://140.113.207.81:9100/metrics"
	url_m2 := "http://140.113.207.82:9100/metrics"
	url_m3 := "http://140.113.207.83:9100/metrics"

	Node1 := NodeInfo{CPUUsage: 0, MemoryUsage: 0, PodNum: 0}
	Node2 := NodeInfo{CPUUsage: 0, MemoryUsage: 0, PodNum: 0}
	Node3 := NodeInfo{CPUUsage: 0, MemoryUsage: 0, PodNum: 0}
	m1_cpu := make([]cpuinfo, 4)
	m1_num := 4
	var m1_mem meminfo

	m2_cpu := make([]cpuinfo, 4)
	m2_num := 4
	var m2_mem meminfo

	m3_cpu := make([]cpuinfo, 4)
	m3_num := 4
	var m3_mem meminfo

	movingAvg := make([]int32, 1)
	var total int32

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "default"

	rc := os.Args[1]

	for {
		Node1.PodNum = 0
		Node2.PodNum = 0
		Node3.PodNum = 0
		p, _ := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		for _, e := range p.Items {
			fmt.Printf(e.Spec.NodeName + ": ")
			fmt.Println(e.ObjectMeta.Name)
			if strings.Contains(e.ObjectMeta.Name, "stress-gin") {
				if strings.Contains(e.Spec.NodeName, "1") {
					Node1.PodNum += 1
				} else if strings.Contains(e.Spec.NodeName, "2") {
					Node2.PodNum += 1
				} else {
					Node3.PodNum += 1
				}
			}
		}
		fmt.Printf("node1: %d, node2: %d, node3: %d\n", Node1.PodNum, Node2.PodNum, Node3.PodNum)
		r, _ := clientset.ExtensionsV1beta1().Deployments(namespace).GetScale(rc, metav1.GetOptions{})
		fmt.Println("rc num is", r.Spec.Replicas)
		Node1.CPUUsage = get_cpu_load(url_m1, m1_cpu, m1_num)
		Node1.MemoryUsage = get_mem_load(url_m1, m1_mem)

		warning(Node1.CPUUsage, 70, "cpu", "m1")
		warning(Node1.MemoryUsage, 70, "mem", "m1")

		Node2.CPUUsage = get_cpu_load(url_m2, m2_cpu, m2_num)
		Node2.MemoryUsage = get_mem_load(url_m2, m2_mem)

		warning(Node2.CPUUsage, 70, "cpu", "m2")
		warning(Node2.MemoryUsage, 70, "mem", "m2")

		Node3.CPUUsage = get_cpu_load(url_m3, m3_cpu, m3_num)
		Node3.MemoryUsage = get_mem_load(url_m3, m3_mem)

		warning(Node3.CPUUsage, 70, "cpu", "m3")
		warning(Node3.MemoryUsage, 70, "mem", "m3")

		avg_cpu := (Node1.CPUUsage + Node2.CPUUsage + Node3.CPUUsage) / 3
		avg_mem := (Node1.MemoryUsage + Node2.MemoryUsage + Node3.MemoryUsage) / 3
		fmt.Println("\n\n", "avg_cpu", avg_cpu, "avg_mem", avg_mem)
		refreshTime, _ := strconv.ParseFloat(os.Args[2], 64)

		time.Sleep(time.Duration(int(1)) * time.Second)

		num := algo(Node1, Node2, Node3, a)

		// add num to moving average array
		//movingAvg = movingAvg[1:]
		//movingAvg = append(movingAvg, num)
		total = 0

		// Calculate moving average
		//for _, value:= range movingAvg{
		//   total += value
		//}
		if alltime%(int(refreshTime)) == 0 {
			// Set pod number
			if alltime == 0 {
				r.Spec.Replicas = 0
			} else {
				r.Spec.Replicas = num
			}
		}
		alltime += 1
		if alltime%20 == 0 {
			store.CsvStore(alltime, int(r.Spec.Replicas), Node1.CPUUsage, Node2.CPUUsage, Node3.CPUUsage, Node1.MemoryUsage, Node2.MemoryUsage, Node3.MemoryUsage, storeTime, int(a*10))
		}
		_, err := clientset.ExtensionsV1beta1().Deployments(namespace).UpdateScale(rc, r)
		fmt.Println("Moving Average Queue", movingAvg)
		fmt.Println("Alltime", alltime)
		fmt.Println("Change num to", total/1)
		fmt.Println(err)
		if alltime == 600 {
			fmt.Println("test finish")
			os.Exit(0)
		}
	}
}
