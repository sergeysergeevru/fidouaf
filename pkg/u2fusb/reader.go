package u2fusb

import (
	"encoding/binary"
	"fmt"
)

func (d *U2fDevice) Read() {
	pack := &InitializationPacket{}
	cid := make([]byte, 4)
	//cid = (*[4]byte)(unsafe.Pointer(&pack.CID))
	//(*[4]uint8)(unsafe.Pointer(&p.CID))
	d.device.Read(cid)
	pack.CID = binary.LittleEndian.Uint32(cid)
	command := make([]byte, 1)
	d.device.Read(command)
	pack.Cmd = U2fCommand(command[0])
	fmt.Print(pack)
}


