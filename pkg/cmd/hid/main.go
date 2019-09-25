package main

import (
	"fmt"
	"github.com/karalabe/hid"
	"github.com/sergeysergeevru/fidouaf/pkg/u2fusb"
	"log"
	"unsafe"
)

func main() {
	u2fDevices := u2fusb.DiscoverU2fDevices()
	fmt.Printf("%d device(s) found \n", len(u2fDevices))
	u2fD := u2fDevices[0]
	device, err := u2fD.Open()
	defer device.Close()
	if err != nil {
		log.Fatal(err)
	}

	//var cid uint32
	//cid = uint32(device.VendorID)
	//fmt.Printf("%x, %x\n", cid<<16|uint32(device.VendorID))
	u2fusbDevice := u2fusb.NewDevice(device)
	u2fusbDevice.U2FHidInit()


	return
	devs := hid.Enumerate(0,0)
	fsIndex := 0;
	for ind, d := range devs {
		//fmt.Println(ind, d.Manufacturer, d.VendorID, d.Usage, d.UsagePage)
		fmt.Printf("%d, %s, %X, %X \n",ind, d.Manufacturer,  d.Usage, d.UsagePage)
		if d.VendorID == 2414 {
			fsIndex = ind
		}

	}
	return
	d := devs[fsIndex]
	device, err = d.Open()
	if err != nil {
		panic(err)
	}
	defer device.Close()
	fmt.Println("test sdf ds")
	//              CID |CMD     |BCN |data
	sr := []byte{0,0,0,1, byte(u2fusb.U2FHidMsg), 0,7, 0x00,0x03,0x00,0x00,0x00,0x0,0x00}

	sr[0],sr[1] = regToMem(device.VendorID)
	sr[2],sr[3] = regToMem(device.ProductID)
	fmt.Println(sr)
	fmt.Println(device.Manufacturer)
	fmt.Println(device.Write(sr))
	t := make([]byte, 64)
	fmt.Println(device.Read(t))
	fmt.Println(t)
	fmt.Println("String: ",string(t))

}

func regToMem(reg uint16) (uint8, uint8) {
	arr := (*[2]uint8)(unsafe.Pointer(&reg))

	return arr[0], arr[1]
}