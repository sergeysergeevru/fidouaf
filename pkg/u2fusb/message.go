package u2fusb

type RequestMessageFrame struct {
	CLA byte
	INS byte
	P1 byte
	P2 byte
	Lc byte
	Data []byte
	Le byte
}