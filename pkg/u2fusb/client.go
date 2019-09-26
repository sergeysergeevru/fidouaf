package u2fusb

import (
	"fmt"
	"github.com/karalabe/hid"
	"unsafe"
)

type InitializationPacket struct {
	CID  uint32
	Cmd  U2fCommand
	Len  uint16
	Data []byte
}

type ContinuationPacket struct {
	CID  uint32
	Seq  uint8
	Data []byte
}

func (p *InitializationPacket) Bytes() []byte {
	var bytePacket []byte
	cidByte := (*[4]uint8)(unsafe.Pointer(&p.CID))
	lenByte := (*[2]uint8)(unsafe.Pointer(&p.Len))
	bytePacket = cidByte[:]
	bytePacket = append(bytePacket, byte(p.Cmd), lenByte[1], lenByte[0],)
	bytePacket = append(bytePacket, p.Data...)
	fmt.Println(bytePacket)

	return bytePacket
}

type U2fDevice struct {
	device *hid.Device
}

func NewDevice(device *hid.Device) *U2fDevice {
	return &U2fDevice{device}
}

func IsU2fDevice(d hid.DeviceInfo) bool {
	if d.Usage == FidoUsage && d.UsagePage == FidoUsagePage {
		return true
	}
	return false
}

func DiscoverU2fDevices() []hid.DeviceInfo {
	var u2fDevices []hid.DeviceInfo
	devs := hid.Enumerate(0,0)
	for _, d := range devs {
		if IsU2fDevice(d) {
			u2fDevices = append(u2fDevices, d)
		}
	}
	return u2fDevices
}

func (d *U2fDevice)U2FHidInit()  {
	cid := uint32(d.device.VendorID)<<16|uint32(d.device.ProductID)
	p := &InitializationPacket{
		CID: cid,
		Cmd:U2FHidInit,
		Len:8,
		Data: []byte{0,0,0,0,0,0,0,0},
	}
	fmt.Println(d.device.Write(p.Bytes()))
	t := make([]byte, 64)
	fmt.Println(d.device.Read(t))
	fmt.Println(t)
}

func (d *U2fDevice)U2FHidPing()  {
	cid := uint32(d.device.VendorID)<<16|uint32(d.device.ProductID)
	p := &InitializationPacket{
		CID: cid,
		Cmd:U2FHidPing,
		Len:8,
		Data: []byte{1,2,3,4,5,6,7,8},
	}
	fmt.Println(d.device.Write(p.Bytes()))
	t := make([]byte, 64)
	fmt.Println(d.device.Read(t))
	fmt.Println(t)
}
