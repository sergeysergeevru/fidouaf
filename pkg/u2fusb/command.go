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

type U2fCommand byte

func (c U2fCommand)IsError() bool {
	return c == U2FHidError
}

func getCmdName(code U2fCommand) string  {
	switch code {
	case U2FHidPing:
		return "U2FHID_PING"
	case U2FHidMsg:
		return "U2FHID_MSG"
	case U2FHidInit:
		return "U2FHID_INIT"
	case U2FHidError:
		return "U2FHID_ERROR"
	case U2FHidWink:
		return "U2FHID_WINK"
		//todo:: lock cmd
	}
	return ""
}
