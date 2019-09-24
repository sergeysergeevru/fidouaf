package main

import (
	"fmt"
	"github.com/karalabe/hid"
	"unsafe"
)

func main() {
	devs := hid.Enumerate(0,0)
	fsIndex := 0;
	for ind, d := range devs {
		fmt.Println(ind, d.Manufacturer, d.VendorID)
		if d.VendorID == 2414 {
			fsIndex = ind
		}

	}
	d := devs[fsIndex]
	device, err := d.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("test sdf ds")
	//              CID |CMD     |BCN |data
	sr := []byte{0,0,0,1, 0x80|4, 0,1, 10}

	sr[0],sr[1] = regToMem(device.VendorID)
	sr[2],sr[3] = regToMem(device.ProductID)
	fmt.Println(sr)
	fmt.Println(device.Manufacturer)
	fmt.Println(device.Write(sr))
	t := make([]byte, 64)
	fmt.Println(device.Read(t))
	fmt.Println(t)

}

func regToMem(reg uint16) (uint8, uint8) {
	arr := (*[2]uint8)(unsafe.Pointer(&reg))

	return arr[0], arr[1]
}