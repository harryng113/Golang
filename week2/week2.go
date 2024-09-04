package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	YOB  int
	Job  string
}

func (p Person) CalculateAge() int {
	return 2024 - p.YOB
}
func (p Person) FitForJob() bool {
	if p.YOB%len(p.Name) == 0 {
		return true
	}
	return false
}

func bai22(s string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		m[ch]++
	}
	return m
}

func bai23(s []int) (int, int, int, float64, []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	mn, mx := s[0], s[len(s)-1]
	sum := 0
	for _, v := range s {
		sum += v
	}
	var tbc float64
	tbc = float64(sum) / float64(len(s))
	return mn, mx, sum, tbc, s
}
func bai24() []Person {
	var sli []Person
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var p Person
	p = Person{"", 0, ""}
	for scanner.Scan() {
		s := scanner.Text()
		cur := 0
		tmp := ""
		for i := 0; i < len(s); i++ {
			ch := string(s[i])
			if ch == "|" {
				tmp = strings.TrimSpace(tmp)
				cur = (cur + 1) % 3
				if cur == 1 {
					p.Name = strings.ToUpper(tmp)
				} else {
					p.YOB, _ = strconv.Atoi(tmp)
				}
				tmp = ""
			} else {
				tmp = tmp + ch
			}
		}
		tmp = strings.TrimSpace(tmp)
		p.Job = strings.ToLower(tmp)
		sli = append(sli, p)

	}
	return sli
}
func main() {
	p := Person{"Hieu", 2002, "Student"}
	fmt.Println(p.CalculateAge())
	fmt.Println(p.FitForJob())
	m := bai22("aaabbbccdd")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(bai23([]int{5, 4, 3, 2, 1}))
	sli := bai24()
	for _, v := range sli {
		fmt.Println(v.Name, v.YOB, v.Job)
	}
}
