package main

import (
	"fmt"

	"github.com/siliconsagekz/govkz-go"
	"github.com/siliconsagekz/govkz-go/services"
)

func main() {
	govinit := govkzgo.Init()

	go_h :=  services.NewStat(govinit)

	compone, err := go_h.GetInfoByIinOrBin("010123000150", "kz"); if err != nil {
		fmt.Print(err)
	}

	fmt.Print(compone)
}