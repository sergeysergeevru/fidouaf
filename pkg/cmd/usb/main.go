package main

import (
	"fmt"
	"github.com/google/gousb"
	"log"
)

func main()  {
	ctx := gousb.NewContext()
	defer ctx.Close()
	fmt.Println("test")
	ctx.Debug(0)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		fmt.Println(desc.String(), desc.Class, desc.SubClass)
		if desc.Vendor == 0x096e {
			return true
		}
		return false
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d devices opened\n", len(devs))
	}
	fmt.Println(devs)

	for _, d := range devs {
		defer d.Close()
	}
	dev := devs[0]

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(2): %v", dev, err)
	}
	defer cfg.Close()
	intf, err := cfg.Interface(1, 0)
	if err != nil {
		log.Fatalf("%s.Interface(3, 0): %v", cfg, err)
	}
	defer intf.Close()
}