package GBCPU

// Struct to emulate the CPU and it's registry
type CPU struct {
	registerA  uint16
	registerF  uint16
	registerB  uint16
	registerC  uint16
	registerD  uint16
	registerE  uint16
	registerH  uint16
	registerL  uint16
	registerAF uint16
	registerBC uint16
	registerDE uint16
	registerHL uint16
	registerSP uint16
	registerPC uint16
	cycles     uint64
	halted     bool
	interrupts bool
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
	cpu.interrupts = false
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
		cpu.registerA = (cpu.registerAF & 0xFF00) >> 8
		cpu.registerF = (cpu.registerAF & 0xFF)
	}
	if cpu.registerBC != ((uint16(cpu.registerB) << 8) | uint16(cpu.registerC)) {
		cpu.registerB = (cpu.registerBC & 0xFF00) >> 8
		cpu.registerC = (cpu.registerBC & 0xFF)
	}
	if cpu.registerDE != ((uint16(cpu.registerD) << 8) | uint16(cpu.registerE)) {
		cpu.registerD = (cpu.registerDE & 0xFF00) >> 8
		cpu.registerE = (cpu.registerDE & 0xFF)
	}
	if cpu.registerHL != ((uint16(cpu.registerH) << 8) | uint16(cpu.registerL)) {
		cpu.registerH = (cpu.registerHL & 0xFF00) >> 8
		cpu.registerL = (cpu.registerHL & 0xFF)
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

func Write16bToMemory(r uint16, value uint16, memory []uint8) {
	r--
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

// Sets the zero flag on register F
func SetZFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF | 0b10000000
}

// Resets the zero flag on register F
func ResetZFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF ^ 0b10000000
}

// Sets the subtraction flag on register F
func SetNFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF | 0b01000000
}

// Resets the subtraction flag on register F
func ResetNFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF ^ 0b01000000
}

// Sets the half carry flag on register F
func SetHFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF | 0b00100000
}

// Resets the half carry flag on register F
func ResetHFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF ^ 0b00100000
}

// Sets the carry flag on register F
func SetCFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF | 0b00010000
}

// Resets the carry flag on register F
func ResetCFlag(cpu *CPU) {
	cpu.registerF = cpu.registerF ^ 0b00010000
}

// Checks if zero flag is set
func IsZFlagSet(cpu *CPU) bool {
	return (cpu.registerF & 0b10000000) == 0b10000000
}

// Checks if subtraction flag is set
func IsNFlagSet(cpu *CPU) bool {
	return (cpu.registerF & 0b01000000) == 0b01000000
}

// Checks if half carry flag is set
func IsHFlagSet(cpu *CPU) bool {
	return (cpu.registerF & 0b00100000) == 0b00100000
}

// Checks if carry flag is set
func IsCFlagSet(cpu *CPU) bool {
	return (cpu.registerF & 0b00010000) == 0b00010000
}
