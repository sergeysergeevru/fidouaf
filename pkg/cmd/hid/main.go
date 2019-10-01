package main

import (
	"fmt"
	"github.com/sergeysergeevru/fidouaf/pkg/u2fusb"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	u2fDevices := u2fusb.DiscoverU2fDevices()
	fmt.Printf("%d device(s) found \n", len(u2fDevices))
	u2fD := u2fDevices[0]
	device, err := u2fD.Open()
	defer device.Close()
	if err != nil {
		log.Fatal(err)
	}
	u2fusbDevice := u2fusb.NewDevice(device)
	u2fusbDevice.U2FHidInit()
}