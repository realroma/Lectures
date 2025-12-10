package main

import (
	"fmt"
	"time"
)

var str string
var l int

type Checker struct {
	items []interface{}
}

func (Checker *Checker) Add(items Checkable) {
	Checker.items = append(Checker.items, items)
}

func Check(c *Checker) {
	for _, value := range c.items {
		a, _ := value.(*GoMetrClient)
		a.Health()
	}
}

type Checkable interface {
	Measurable
	Ping()
	GetID()
	Health()
}

type Measurable interface {
	GetMetrics()
}

func GetMertics() {

}

type GoMetrClient struct {
	url     string
	timeOut time.Duration
}

type HealthCheck struct {
	ServisID string
	status   string
}

func (GMC *GoMetrClient) Ping() {

}

func (GMC *GoMetrClient) GetID() {

}

func (GMC *GoMetrClient) Health() {
	select {
	case <-time.After(GMC.timeOut * time.Second):
		break
	default:
		fmt.Println("Nothing.")
	}
	a := getHealth()
	str = str + "GetID: " + a.ServisID + " Health: " + a.status + "\n"
	fmt.Println(str)
}

func (GMC *GoMetrClient) GetMetrics() {

}

func getHealth() *HealthCheck {
	a := new(HealthCheck)
	switch l {
	case 0:
		a = &HealthCheck{"1", "true"}
	case 1:
		a = &HealthCheck{"2", "false"}
	case 2:
		a = &HealthCheck{"3", "true"}
	}
	l++
	return a
}

func main() {
	items := &Checker{}
	var Checkable1 Checkable
	var Checkable2 Checkable
	var Checkable3 Checkable
	Checkable1 = &GoMetrClient{}
	Checkable2 = &GoMetrClient{}
	Checkable3 = &GoMetrClient{}
	items.Add(Checkable1)
	items.Add(Checkable2)
	items.Add(Checkable3)
	Check(items)
}
