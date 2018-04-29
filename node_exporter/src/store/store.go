package store

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func CsvInit(RefreshTime int, a int) {
	filename := fmt.Sprintf("%d_time_%d_a.csv", RefreshTime, a)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.Write([]string{"nowtime", "m1_cpu", "m2_cpu", "m3_cpu", "m1_mem", "m2_mem", "m3_mem", "pod", "refreshtime"})
	w.Flush()
}
func CsvStore(nowtime int, pod int, m1_cpu float64, m2_cpu float64, m3_cpu float64, m1_mem float64, m2_mem float64, m3_mem float64, RefreshTime int, a int) {
	filename := fmt.Sprintf("%d_time_%d_a.csv", RefreshTime, a)
	fmt.Println(filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	s := make([]string, 9)
	s[0] = strconv.Itoa(nowtime)
	s[1], s[2], s[3] = FloatToString(m1_cpu), FloatToString(m2_cpu), FloatToString(m3_cpu)
	s[4], s[5], s[6] = FloatToString(m1_mem), FloatToString(m2_mem), FloatToString(m3_mem)
	s[7] = strconv.Itoa(pod)
	s[8] = strconv.Itoa(RefreshTime)
	fmt.Println(s)
	w.Write(s)
	w.Flush()
}
