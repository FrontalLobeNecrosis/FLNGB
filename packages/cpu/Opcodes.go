package GBCPU

// This struct is to call functions from an array based on opcode
type Opcode_function_caller struct {
	eightBitFuncArray   [255]func(uint8, uint8)
	eightbitparam1      [255]uint8
	eightbitparam2      [255]uint8
	sixteenBitFuncArray [255]func(uint16, uint16)
	sixteenbitparam1    [255]uint16
	sixteenbitparam2    [255]uint16
}

// Function makes an Opcode_function_caller and takes a CPU struct and loades the
// caller with all the functions and params that will be called by Opcodes
func initCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := new(Opcode_function_caller)

	for i := 0; i <= 255; i++ {

		if i < 37 && i%8 == 6 {
			caller.eightBitFuncArray[i] = LDn
			caller.eightbitparam2[i] = uint8(immediateValue)
		}

		if (i >= 0x77 && i <= 0x7F) || (i >= 0x40 && i <= 0x75) {

			caller.eightBitFuncArray[i] = LDr

			if i >= 0x78 && i <= 0x7F {
				caller.eightbitparam1[i] = cpu.registerA
			} else if i >= 0x40 && i >= 0x47 {
				caller.eightbitparam1[i] = cpu.registerB
			} else if i >= 0x48 && i <= 0x4F {
				caller.eightbitparam1[i] = cpu.registerC
			} else if i >= 0x50 && i <= 0x57 {
				caller.eightbitparam1[i] = cpu.registerD
			} else if i >= 0x58 && i <= 0x5F {
				caller.eightbitparam1[i] = cpu.registerE
			} else if i >= 0x60 && i <= 0x67 {
				caller.eightbitparam1[i] = cpu.registerH
			} else if i >= 0x68 && i <= 0x6F {
				caller.eightbitparam1[i] = cpu.registerL
			} else if (i >= 0x70 && i <= 0x75) || i == 36 {
				value := memory[cpu.registerHL]
				caller.eightbitparam1[i] = value
			}

			remainder := i % 8

			switch remainder {
			case 0:
				caller.eightbitparam2[i] = cpu.registerB
				break
			case 1:
				caller.eightbitparam2[i] = cpu.registerC
				break
			case 2:
				caller.eightbitparam2[i] = cpu.registerD
				break
			case 3:
				caller.eightbitparam2[i] = cpu.registerE
				break
			case 4:
				caller.eightbitparam2[i] = cpu.registerH
				break
			case 5:
				caller.eightbitparam2[i] = cpu.registerL
				break
			case 6:
				caller.eightbitparam2[i] = cpu.registerC
				break
			case 7:
				caller.eightbitparam2[i] = cpu.registerA
				break
			}

		}

	}
	return caller
}

func NewCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := initCaller(cpu, memory, immediateValue)
	return caller
}

// LDn loads a value from a register nn into another register
// or immediate value n
// param: nn, a register to have a value read from
// 		  n, a register or an 8 bit immediate value to have a value written to
func LDn(nn uint8, n uint8) {
	nn = n
}

// LDr loads a value from a register r2 into another register
// or immediate value r1
// param: r2, a register to have a value read from
// 		  r1, a register or immediate value to have a value written to
func LDr(r1 uint8, r2 uint8) {
	r1 = r2
}

// Takes in an opcode and runs the function with appropriate params associated with that code
// param: an opcode that can be 8 or 16 bit value (16 bit has to begin at 0xCB00 and ends at 0xCBFF)
// and might be followed by an 8 or 16 bit immediate value
func ReadOpcode(opcode uint32, cpu *CPU, memory []uint8) {

	var immediateValue uint16

	if (opcode > 0xFF) && ((opcode&0xCB00) != 0xCB00) || opcode > 0xFFFF {
		if opcode > 0xFFFF {
			immediateValue = uint16(opcode)
			opcode = opcode >> 16
		} else {
			immediateValue = uint16(opcode) & 0xFF
			opcode = opcode >> 8
		}
	}

	caller := NewCaller(cpu, memory, immediateValue)

	if (opcode > 255) && ((opcode & 0xCB00) == 0xCB00) {
		function := caller.sixteenBitFuncArray[opcode-0xCB00]
		first := caller.sixteenbitparam1[opcode-0xCB00]
		second := caller.sixteenbitparam2[opcode-0xCB00]
		function(first, second)
	} else {
		function := caller.eightBitFuncArray[opcode]
		first := caller.eightbitparam1[opcode]
		second := caller.eightbitparam2[opcode]
		function(first, second)
	}

}
