package GBCPU

// This struct is to call functions from an array based on opcode
// 8 bit function and params array is for 8 bit opcodes
// 16 bit functions and params are for 16 bit opcodes
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
				caller.eightbitparam1[i] = uint16(memory[cpu.registerHL])
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

		if i >= 0x80 && i <= 0xA7 || i == 0xC6 || i == 0xCE || i == 0xD6 ||
			i == 0xDE || i == 0xE6 || i == 0xEE || i == 0xF6 || i == 0xFE {
			if i <= 0xB8 || i == 0xFE {
				caller.eightBitFuncArray[i] = CP
			} else if i <= 0xB0 || i == 0xF6 {
				caller.eightBitFuncArray[i] = OR
			} else if i <= 0xA8 || i == 0xEE {
				caller.eightBitFuncArray[i] = XOR
			} else if i <= 0xA0 || i == 0xE6 {
				caller.eightBitFuncArray[i] = AND
			} else if i <= 0x8F || i == 0xC6 || i == 0xCE {
				caller.eightBitFuncArray[i] = ADD
			} else {
				caller.eightBitFuncArray[i] = SUB
			}
			caller.eightbitparam1[i] = uint16(cpu.registerA)

			var carry uint8 = 0
			if i > 0x87 && i < 0x90 || i > 0x97 && i < 0xA0 || i == 0xCE || i == 0xDE {
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
				if i >= 0xC6 {
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

		// TODO: Find a pattern to replace switch case stack with
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
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
		case 0x2A:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
		case 0x32:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndDeincrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x3A:
			// TODO: Fix the way the memory and HL register is accessed and updated
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
			caller.eightbitparam1[i] = uint16(memory[0xFF00+immediateValue])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
		case 0xF2:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[0xFF00+uint16(cpu.registerC)])
			break
		case 0xF8:
			caller.eightBitFuncArray[i] = LDFlag
			caller.eightbitparam1[i] = cpu.registerHL
			caller.eightbitparam2[i] = cpu.registerSP + immediateValue
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
//
// params:
// 			cpu, a CPU struct containing registers to write to and read from
// 			memory, an array of 8 bit values with the size of 0xFFFF
// 			immediateValue, the immediate value included in the opcode, can be 8 or 16 bit
func NewCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := initCaller(cpu, memory, immediateValue)
	return caller
}

// LDn loads a value from a register or immediate value, into a register
//
// params:
// 			nn, a register to have a value written to
// 			n, a register, memory addres, or an 8 bit immediate value to write from
func LDn(nn uint16, n uint16, cpu *CPU, memory []uint8) {
	nn = n
}

// LDr loads a value from a register r2 into another register
// or immediate value r1
//
// params:
// 			r1, a register to write to
// 			r2, a register, memory addres, or an 8 bit immediate value being read from
func LDr(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	r1 = r2
}

// LD16b is like the other LD finctions but intended only for use with 16 bit
// values (like the paired registers) this is the reason every other function has the 16 bit params
//
// params:
// 			r, a register to write to
// 			value, a paired register or an 16 bit immediate value being read from
func LD16b(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
}

// LDFlag was made for the spcific case where the flag register needs to be edited,
// this is the reason all functions have the CPU param
//
// params:
// 			r, a register to write to
// 			value, a paired register or an 16 bit immediate value being read from
func LDFlag(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
	cpu.registerF = cpu.registerF & 0b00110000
}

// Pushes 16 bits worth of values onto the memory stack by SP register
//
// params:
// 			r, SP register (stack pointer)
// 			value, value to be pushed onto the stack. either paired register or immediate value
// 			memory, an array of 8 bit values with the size of 0xFFFF, will store the pushed value
func PUSH(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r--
	Write16bToMemory(r, value, memory)
}

// Pops 16 bit worth of values off of the memory stack by SP register
//
// params:
// 			r1, SP register (stack pointer)
// 			r2, paired register to push value onto
// 			memory, an array of 8 bit values with the size of 0xFFFF, will have values poped off it
func POP(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	Read16bFromMemory(r1, r2, memory)
}

// Adds a value to the arithmitic registry (register A)
//
// params:
// 			A, arithmitic register (register A)
// 			n, a value to be added to arithmitic register, can be value from memory,
// 				   other registers, or immediate value
// 			cpu, CPU struct to edit flag register (register F)
// 			memory, an array of 8 bit values with the size of 0xFFFF
func ADD(A uint16, n uint16, cpu *CPU, memory []uint8) {
	temp := n
	if n > 0xFF {
		temp = uint16(memory[n])
	}
	result := A + temp
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (A&0b111)+(temp&0b111) >= 0xF {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (A&0b1111111)+(temp&0b1111111) >= 0xF0 {
		cpu.registerF = cpu.registerF | 0b00010000
	}
	A = result & 0xFF
}

// Subtracts a value from the arithmitic register (register A)
//
// params:
// 			A, arithmitic register (register A)
// 			n, a value to be added to arithmitic register, can be value from memory,
// 				   other registers, or immediate value
// 			cpu, CPU struct to edit flag register (register F)
// 			memory, an array of 8 bit values with the size of 0xFFFF
func SUB(A uint16, n uint16, cpu *CPU, memory []uint8) {
	temp := n
	if n > 0xFF {
		temp = uint16(memory[n])
	}
	result := A - temp
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) != 0b01000000 {
		cpu.registerF = cpu.registerF | 0b01000000
	}
	if (A&0b111)-(temp&0b111) >= 0 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (A&0b1111111)-(temp&0b1111111) >= 0 {
		cpu.registerF = cpu.registerF | 0b00010000
	}
	A = result & 0xFF
}

// Logically and a value with the arithmetic register (register A)
//
// params:
// 			A, arithmitic register (register A)
// 			n, a value to be added to arithmitic register, can be value from memory,
// 				   other registers, or immediate value
// 			cpu, CPU struct to edit flag register (register F)
// 			memory, an array of 8 bit values with the size of 0xFFFF
func AND(A uint16, n uint16, cpu *CPU, memory []uint8) {
	result := A & n
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (cpu.registerF & 0b00100000) != 0b00100000 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		cpu.registerF = cpu.registerF ^ 0b00010000
	}
	A = result & 0xFF
}

// Logically or a value with the arithmetic register (register A)
//
// params:
// 			A, arithmitic register (register A)
// 			n, a value to be added to arithmitic register, can be value from memory,
// 				   other registers, or immediate value
// 			cpu, CPU struct to edit flag register (register F)
// 			memory, an array of 8 bit values with the size of 0xFFFF
func OR(A uint16, n uint16, cpu *CPU, memory []uint8) {
	result := A | n
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		cpu.registerF = cpu.registerF ^ 0b00100000
	}
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		cpu.registerF = cpu.registerF ^ 0b00010000
	}
	A = result & 0xFF
}

// Logically xor a value with the arithmetic register (register A)
//
// params:
// 			A, arithmitic register (register A)
// 			n, a value to be added to arithmitic register, can be value from memory,
// 				   other registers, or immediate value
// 			cpu, CPU struct to edit flag register (register F)
// 			memory, an array of 8 bit values with the size of 0xFFFF
func XOR(A uint16, n uint16, cpu *CPU, memory []uint8) {
	result := A ^ n
	if result&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		cpu.registerF = cpu.registerF ^ 0b00100000
	}
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		cpu.registerF = cpu.registerF ^ 0b00010000
	}
	A = result & 0xFF
}

// Takes in an opcode and runs the function with appropriate params associated with that code
//
// params:
// 			opcode, can be 8 or 16 bit value 16 bit has to begin at 0xCB00 and ends at 0xCBFF
// 					and might be followed by an 8 or 16 bit immediate value
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
