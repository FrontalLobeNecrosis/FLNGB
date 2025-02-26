package GBCPU

// Struct to emulate the CPU and it's registry
type CPU struct {
	registerA  uint8
	registerF  uint8
	registerB  uint8
	registerC  uint8
	registerD  uint8
	registerE  uint8
	registerH  uint8
	registerL  uint8
	registerAF uint16
	registerBC uint16
	registerDE uint16
	registerHL uint16
	registerSP uint16
	registerPC uint16
	cycles     uint32
	halted     bool
}

/*
Initializes the CPU struct with the proper values
also note that values and functions that are public start their name
capaltilized. This is not one of those.
*/
func initCPU(cpu *CPU) *CPU {
	cpu.registerA = 01
	cpu.registerF = 0
	cpu.registerB = 0xFF
	cpu.registerC = 0x13
	cpu.registerD = 0
	cpu.registerE = 0xC1
	cpu.registerH = 0x84
	cpu.registerL = 03
	cpu.registerAF = (uint16(cpu.registerA) << 8) | uint16(cpu.registerF)
	cpu.registerBC = (uint16(cpu.registerB) << 8) | uint16(cpu.registerC)
	cpu.registerDE = (uint16(cpu.registerD) << 8) | uint16(cpu.registerE)
	cpu.registerHL = (uint16(cpu.registerH) << 8) | uint16(cpu.registerL)
	cpu.registerSP = 0xFFFE
	cpu.registerPC = 0x0100
	cpu.cycles = 0
	cpu.halted = false
	return cpu
}

// Makes a new CPU struct and initializes it
func NewCPU() *CPU {
	cpu := new(CPU)
	cpu = initCPU(cpu)
	return cpu
}

// Makes sure paired and single registers are properly equivilent,
// this is to be used if one of the single registries were changed
func SingleToPaired(cpu *CPU) {
	if cpu.registerAF != ((uint16(cpu.registerA) << 8) | uint16(cpu.registerF)) {
		cpu.registerAF = ((uint16(cpu.registerA) << 8) | uint16(cpu.registerF))
	}
	if cpu.registerBC != ((uint16(cpu.registerB) << 8) | uint16(cpu.registerC)) {
		cpu.registerBC = ((uint16(cpu.registerB) << 8) | uint16(cpu.registerC))
	}
	if cpu.registerDE != ((uint16(cpu.registerD) << 8) | uint16(cpu.registerE)) {
		cpu.registerDE = ((uint16(cpu.registerD) << 8) | uint16(cpu.registerE))
	}
	if cpu.registerHL != ((uint16(cpu.registerH) << 8) | uint16(cpu.registerL)) {
		cpu.registerHL = ((uint16(cpu.registerH) << 8) | uint16(cpu.registerL))
	}
}

// Makes sure paired and single registers are properly equivilent,
// this is to be used if one of the paired registries were changed
func PairedToSingle(cpu *CPU) {
	if cpu.registerAF != ((uint16(cpu.registerA) << 8) | uint16(cpu.registerF)) {
		cpu.registerA = uint8((cpu.registerAF - uint16(cpu.registerF)) >> 8)
		cpu.registerF = uint8(cpu.registerAF - (uint16(cpu.registerA) << 8))
	}
	if cpu.registerBC != ((uint16(cpu.registerB) << 8) | uint16(cpu.registerC)) {
		cpu.registerB = uint8((cpu.registerBC - uint16(cpu.registerC)) >> 8)
		cpu.registerC = uint8(cpu.registerBC - (uint16(cpu.registerB) << 8))
	}
	if cpu.registerDE != ((uint16(cpu.registerD) << 8) | uint16(cpu.registerE)) {
		cpu.registerD = uint8((cpu.registerDE - uint16(cpu.registerE)) >> 8)
		cpu.registerE = uint8(cpu.registerDE - (uint16(cpu.registerD) << 8))
	}
	if cpu.registerHL != ((uint16(cpu.registerH) << 8) | uint16(cpu.registerL)) {
		cpu.registerH = uint8((cpu.registerHL - uint16(cpu.registerL)) >> 8)
		cpu.registerL = uint8(cpu.registerHL - (uint16(cpu.registerH) << 8))
	}
}

// Initializes memory and passes it out as a slice
func initMemory() []uint8 {
	memory := [0x10000]uint8{}
	for i := 0; i < len(memory); i++ {
		memory[i] = 0
	}
	return memory[:]
}

// Calls intialize memory so that you can grab a new memory
// and not accidentaly initialize your already existing memory
func NewMemory() []uint8 {
	memory := initMemory()
	return memory
}

func GetMemoryAndIncrement(memory []uint8, address *uint16) uint8 {
	*address++
	return memory[*address]
}

func GetMemoryAndDeincrement(memory []uint8, address *uint16) uint8 {
	*address--
	return memory[*address]
}

func Write16bToMemory(r uint16, value uint16, memory []uint8) {
	memory[r] = uint8((value & 0xFF00) >> 8)
	r--
	memory[r] = uint8(value & 0xFF)
}

func Read16bFromMemory(r1 uint16, r2 uint16, memory []uint8) {
	var value uint16
	value = uint16(memory[r1])
	r1++
	value += uint16(memory[r1]) << 8
	r1++
	r2 = value
}
