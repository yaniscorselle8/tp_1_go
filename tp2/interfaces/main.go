package main

import (
	"fmt"
	"strconv"
	"time"
)

type IPAddr [4]byte

//converts IPAddr to String while using fmt.Sprintf
/*func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}*/

//converts IPAddr to String while using strcnv.Itoa
func (a IPAddr) String() string {
	return strconv.Itoa(int(a[0])) + "." +
		strconv.Itoa(int(a[1])) + "." +
		strconv.Itoa(int(a[2])) + "." +
		strconv.Itoa(int(a[3]))
}

func (e myError) Error() string {
	return fmt.Sprintf("Error happenned at time : %s. Log : %s ", time.Now(), "bug")
}

type myError struct {
	When time.Time
	What string
}

func run() error {
	e := myError{time.Now(), "bug"}
	return e
}

func PrintIt(input interface{}) {
	switch input.(type) {
	case int:
		fmt.Println("Value is int")
	case string:
		fmt.Println("Value is string")
	default:
		fmt.Println("I don't know about type")
	}
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	err := run()
	if err != nil {
		fmt.Println(err)
	}
	PrintIt(1)
	PrintIt("un")

}
