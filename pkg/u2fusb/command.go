package u2fusb

const (
	U2FHidPing  U2fCommand = 0x80 | 1
	U2FHidMsg   U2fCommand = 0x80 | 3
	U2FHidLock  U2fCommand = 0x80 | 4
	U2FHidInit  U2fCommand = 0x80 | 6
	U2FHidWink  U2fCommand = 0x80 | 8
	U2FHidSync  U2fCommand = 0x80 | 0x3c
	U2FHidError U2fCommand = 0x80 | 0x3f

	FidoUsagePage = 0xf1d0
	FidoUsage     = 0x1
)

type U2fCommand uint8
