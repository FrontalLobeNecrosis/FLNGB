package GBCPU

// This struct is to call functions from an array based on opcode
type Opcode_function_caller struct {
	eightBitFuncArray   [255]func(uint16, uint16, *CPU, []uint8)
	eightbitparam1      [255]uint16
	eightbitparam2      [255]uint16
	sixteenBitFuncArray [255]func(uint16, uint16)
	sixteenbitparam1    [255]uint16
	sixteenbitparam2    [255]uint16
}

// Function makes an Opcode_function_caller and takes a CPU struct and loades the
// caller with all the functions and params that will be called by opcodes
func initCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := new(Opcode_function_caller)

	for i := 0; i <= 255; i++ {

		if i <= 0x31 && i%16 == 1 {
			caller.eightBitFuncArray[i] = LD16b
			caller.eightbitparam2[i] = immediateValue

			switch i {
			case 0x01:
				caller.eightbitparam1[i] = cpu.registerBC
				break
			case 0x11:
				caller.eightbitparam1[i] = cpu.registerDE
				break
			case 0x21:
				caller.eightbitparam1[i] = cpu.registerHL
				break
			case 0x31:
				caller.eightbitparam1[i] = cpu.registerSP
				break
			}
		}

		if i <= 0x3E && i%8 == 6 {
			caller.eightBitFuncArray[i] = LDn
			caller.eightbitparam2[i] = immediateValue

			switch i {
			case 0x06:
				caller.eightbitparam1[i] = uint16(cpu.registerB)
				break
			case 0x0E:
				caller.eightbitparam1[i] = uint16(cpu.registerC)
				break
			case 0x16:
				caller.eightbitparam1[i] = uint16(cpu.registerD)
				break
			case 0x1E:
				caller.eightbitparam1[i] = uint16(cpu.registerE)
				break
			case 0x26:
				caller.eightbitparam1[i] = uint16(cpu.registerH)
				break
			case 0x2E:
				caller.eightbitparam1[i] = uint16(cpu.registerL)
				break
			case 0x36:
				caller.eightbitparam1[i] = uint16(memory[cpu.registerHL])
				break
			case 0x3E:
				caller.eightbitparam1[i] = uint16(cpu.registerA)
				break
			}
		}

		if (i <= 0xF5) && (i >= 0xC5) && (i%16 == 5) {
			caller.eightBitFuncArray[i] = PUSH
			caller.eightbitparam1[i] = cpu.registerSP

			switch i {
			case 0xC5:
				caller.eightbitparam2[i] = cpu.registerBC
				break
			case 0xD5:
				caller.eightbitparam2[i] = cpu.registerDE
				break
			case 0xE5:
				caller.eightbitparam2[i] = cpu.registerHL
				break
			case 0xF5:
				caller.eightbitparam2[i] = cpu.registerAF
				break
			}
		}

		if (i <= 0xF1) && (i >= 0xC1) && (i%16 == 1) {
			caller.eightBitFuncArray[i] = POP
			caller.eightbitparam1[i] = cpu.registerSP

			switch i {
			case 0xC1:
				caller.eightbitparam2[i] = cpu.registerBC
				break
			case 0xD1:
				caller.eightbitparam2[i] = cpu.registerDE
				break
			case 0xE1:
				caller.eightbitparam2[i] = cpu.registerHL
				break
			case 0xF1:
				caller.eightbitparam2[i] = cpu.registerAF
				break
			}
		}

		if (i >= 0x77 && i <= 0x7F) || (i >= 0x40 && i <= 0x75) {

			caller.eightBitFuncArray[i] = LDr

			if i >= 0x78 && i <= 0x7F {
				caller.eightbitparam1[i] = uint16(cpu.registerA)
			} else if i >= 0x40 && i >= 0x47 {
				caller.eightbitparam1[i] = uint16(cpu.registerB)
			} else if i >= 0x48 && i <= 0x4F {
				caller.eightbitparam1[i] = uint16(cpu.registerC)
			} else if i >= 0x50 && i <= 0x57 {
				caller.eightbitparam1[i] = uint16(cpu.registerD)
			} else if i >= 0x58 && i <= 0x5F {
				caller.eightbitparam1[i] = uint16(cpu.registerE)
			} else if i >= 0x60 && i <= 0x67 {
				caller.eightbitparam1[i] = uint16(cpu.registerH)
			} else if i >= 0x68 && i <= 0x6F {
				caller.eightbitparam1[i] = uint16(cpu.registerL)
			} else if i >= 0x70 && i <= 0x75 {
				value := uint16(memory[cpu.registerHL])
				caller.eightbitparam1[i] = value
			}

			remainder := i % 8

			switch remainder {
			case 0:
				caller.eightbitparam2[i] = uint16(cpu.registerB)
				break
			case 1:
				caller.eightbitparam2[i] = uint16(cpu.registerC)
				break
			case 2:
				caller.eightbitparam2[i] = uint16(cpu.registerD)
				break
			case 3:
				caller.eightbitparam2[i] = uint16(cpu.registerE)
				break
			case 4:
				caller.eightbitparam2[i] = uint16(cpu.registerH)
				break
			case 5:
				caller.eightbitparam2[i] = uint16(cpu.registerL)
				break
			case 6:
				caller.eightbitparam2[i] = uint16(memory[cpu.registerHL])
				break
			case 7:
				caller.eightbitparam2[i] = uint16(cpu.registerA)
				break
			}
		}

		if i >= 0x80 && i <= 0x8F || i == 0xC6 || i == 0xCE {
			caller.eightBitFuncArray[i] = ADD
			caller.eightbitparam1[i] = uint16(cpu.registerA)

			var carry uint8 = 0
			if i > 0x87 || i == 0xCE {
				carry = cpu.registerF & 0b00010000
			}
			remainder := i % 8

			switch remainder {
			case 0:
				caller.eightbitparam2[i] = uint16(cpu.registerB + carry)
				break
			case 1:
				caller.eightbitparam2[i] = uint16(cpu.registerC + carry)
				break
			case 2:
				caller.eightbitparam2[i] = uint16(cpu.registerD + carry)
				break
			case 3:
				caller.eightbitparam2[i] = uint16(cpu.registerE + carry)
				break
			case 4:
				caller.eightbitparam2[i] = uint16(cpu.registerH + carry)
				break
			case 5:
				caller.eightbitparam2[i] = uint16(cpu.registerL + carry)
				break
			case 6:
				if i <= 0xCE {
					caller.eightbitparam2[i] = immediateValue + uint16(carry)
				} else {
					caller.eightbitparam2[i] = uint16(memory[cpu.registerHL] + carry)
				}
				break
			case 7:
				caller.eightbitparam2[i] = uint16(cpu.registerA + carry)
				break
			}
		}

		switch i {
		case 0x02:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[cpu.registerBC])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x08:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = cpu.registerSP
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x0A:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[cpu.registerBC])
			break
		case 0x12:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[cpu.registerDE])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x1A:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[cpu.registerDE])
			break
		case 0x22:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
		case 0x2A:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
		case 0x32:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndDeincrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x3A:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(GetMemoryAndDeincrement(memory, &cpu.registerHL))
			break
		case 0xE0:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[0xFF00+immediateValue])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0xE2:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[0xFF00+uint16(cpu.registerC)])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0xEA:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[immediateValue])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0xF0:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			caller.eightbitparam1[i] = uint16(memory[0xFF00+immediateValue])
		case 0xF2:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[0xFF00+uint16(cpu.registerC)])
			break
		case 0xF8:
			caller.eightBitFuncArray[i] = LDHL
			caller.eightbitparam2[i] = cpu.registerSP + immediateValue
			caller.eightbitparam1[i] = cpu.registerHL
			break
		case 0xF9:
			caller.eightBitFuncArray[i] = LD16b
			caller.eightbitparam1[i] = cpu.registerSP
			caller.eightbitparam2[i] = cpu.registerHL
			break
		case 0xFA:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[immediateValue])
			break
		}

	}
	return caller
}

// Makes a new caller and initializes it with the
// proper functions and params at the proper opcode location
func NewCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := initCaller(cpu, memory, immediateValue)
	return caller
}

// LDn loads a value from a register nn into another register
// or immediate value n
// params:
// 			nn, a register to have a value written to
// 			n, a register, memory addres, or an 8 bit immediate value to have a value read
func LDn(nn uint16, n uint16, cpu *CPU, memory []uint8) {
	nn = n
}

// LDr loads a value from a register r2 into another register
// or immediate value r1
// params:
// 			r2, a register to write to
// 			r1, a register, memory addres, or an 8 bit immediate value being read from
func LDr(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	r1 = r2
}

// LD16b is like the other LD finctions but intended only for use with 16 bit
// values this is the reason every other function has the 16 bit params
// params:
// 			r, a register to write to
// 			value, a paired register or an 16 bit immediate value being read from
func LD16b(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
}

// LDHL was made for the spcific case where the flag register needs to be edited
func LDHL(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
	cpu.registerF = cpu.registerF & 0b00110000
}

// Pushes 16 bits worth of values onto the memory stack
func PUSH(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r--
	Write16bToMemory(r, value, memory)
}

// Pops 16 bit worth of values off of the memory stack
func POP(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	Read16bFromMemory(r1, r2, memory)
}

func ADD(r uint16, value uint16, cpu *CPU, memory []uint8) {
	temp := value
	if value > 0xFF {
		temp = uint16(memory[value])
	}
	result := r + temp
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF -= 0b01000000
	}
	if (r&0xF)+(temp&0xF) > 0xF {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (result & 0xFF00) > 0 {
		cpu.registerF = cpu.registerF | 0b00010000
	}
	r = result & 0xFF
}

// Takes in an opcode and runs the function with appropriate params associated with that code
//
// params:
// 			opcode, can be 8 or 16 bit value 16 bit has to begin at 0xCB00 and ends at 0xCBFF
// 			and might be followed by an 8 or 16 bit immediate value
// 			cpu, where the registers are read from and written to
// 			memory, An arrray of 8 bit integers that is 0x10000 addresses long
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
		function(first, second, cpu, memory)
	}

}
