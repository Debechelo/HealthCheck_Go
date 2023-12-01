package main

import "fmt"

const PassStatus = "pass"
const FailStatus = "fail"

func main() {
	check := Checker{}
	check.Add(CreateGoMetrClient("host1", 1),
		CreateGoMetrClient("host2", 1),
		CreateGoMetrClient("host3", 1),
		CreateGoMetrClient("host4", 1),
		CreateGoMetrClient("host5", 1),
		CreateGoMetrClient("host6", 1))

	fmt.Println(check)
	check.Check()
}
