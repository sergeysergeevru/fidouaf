package main

import (
	"fmt"
	"github.com/google/gousb"
)

func main()  {
	ctx := gousb.NewContext()
	defer ctx.Close()
	fmt.Println("test")
	ctx.Debug(0)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		//fmt.Println(desc.Vendor)

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
	//dev := devs[0]

}