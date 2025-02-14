package GBCPU

// Struct to emulate the CPU and it's registry
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
	cycles        uint32
}

/*
Initializes the CPU struct with the proper values
also note that values and functions that are public start their name
capaltilized. This is not one of those.
*/
func initCPU() *CPU {
	cpu := new(CPU)
	cpu.registerA = 0
	cpu.registerF = 0
	cpu.registerB = 0
	cpu.registerC = 0
	cpu.registerD = 0
	cpu.registerE = 0
	cpu.registerH = 0
	cpu.registerL = 0
	cpu.flagresgister = 0b00000000
	cpu.registerSP = 0x0100
	cpu.registerPC = 0xFFFE
	cpu.cycles = 0
	return cpu
}
