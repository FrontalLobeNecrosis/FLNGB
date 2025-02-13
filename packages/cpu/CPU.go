package GBCPU

type CPU struct {
	registerA     uint8
	registerF     uint8
	registerB     uint8
	registerC     uint8
	registerD     uint8
	registerE     uint8
	registerH     uint8
	registerL     uint8
	flagresgister uint8
	registerSP    uint16
	registerPC    uint16
}
