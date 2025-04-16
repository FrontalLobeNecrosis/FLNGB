package GBCPU

// This struct is to call functions from an array based on opcode
// 8 bit function and params array is for 8 bit opcodes
// 16 bit functions and params are for 16 bit opcodes
type Opcode_function_caller struct {
	// TODO: Figure out if these need to hold pointers
	eightBitFuncArray   [256]func(uint16, uint16, *CPU, []uint8)
	eightbitparam1      [256]uint16
	eightbitparam2      [256]uint16
	sixteenBitFuncArray [256]func(uint16, uint16, *CPU, []uint8)
	sixteenbitparam1    [256]uint16
	sixteenbitparam2    [256]uint16
}

// Function makes an Opcode_function_caller and takes a CPU struct and loades the
// caller with all the functions and params that will be called by opcodes
func CallerLoader(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := new(Opcode_function_caller)

	for i := 0; i <= 0xFF; i++ {

		// This initial set of blocks if for the sixteen bit array
		// should contain all calls needed between 0 - 255
		caller.sixteenbitparam2[i] = 0
		if i < 0x08 {
			caller.sixteenBitFuncArray[i] = RLCn
		} else if i < 0x10 {
			caller.sixteenBitFuncArray[i] = RRCn
		} else if i < 0x18 {
			caller.sixteenBitFuncArray[i] = RLn
		} else if i < 0x20 {
			caller.sixteenBitFuncArray[i] = RRn
		} else if i < 0x28 {
			caller.sixteenBitFuncArray[i] = SLA
		} else if i < 0x30 {
			caller.sixteenBitFuncArray[i] = SRA
		} else if i < 0x38 {
			caller.sixteenBitFuncArray[i] = SWAP
		} else if i < 0x40 {
			caller.sixteenBitFuncArray[i] = SRL
		} else if i < 0x80 {
			caller.sixteenBitFuncArray[i] = BIT
		} else if i < 0xC0 {
			caller.sixteenBitFuncArray[i] = RES
		} else {
			caller.sixteenBitFuncArray[i] = SET
		}

		remainder := i % 8
		if i < 0x40 {
			switch remainder {
			case 0:
				caller.sixteenbitparam1[i] = uint16(cpu.registerB)
				break
			case 1:
				caller.sixteenbitparam1[i] = uint16(cpu.registerC)
				break
			case 2:
				caller.sixteenbitparam1[i] = uint16(cpu.registerD)
				break
			case 3:
				caller.sixteenbitparam1[i] = uint16(cpu.registerE)
				break
			case 4:
				caller.sixteenbitparam1[i] = uint16(cpu.registerH)
				break
			case 5:
				caller.sixteenbitparam1[i] = uint16(cpu.registerL)
				break
			case 6:
				caller.sixteenbitparam1[i] = uint16(memory[cpu.registerHL])
				break
			case 7:
				caller.sixteenbitparam1[i] = uint16(cpu.registerA)
				break
			}
		} else {
			switch remainder {
			case 0:
				caller.sixteenbitparam2[i] = uint16(cpu.registerB)
				break
			case 1:
				caller.sixteenbitparam2[i] = uint16(cpu.registerC)
				break
			case 2:
				caller.sixteenbitparam2[i] = uint16(cpu.registerD)
				break
			case 3:
				caller.sixteenbitparam2[i] = uint16(cpu.registerE)
				break
			case 4:
				caller.sixteenbitparam2[i] = uint16(cpu.registerH)
				break
			case 5:
				caller.sixteenbitparam2[i] = uint16(cpu.registerL)
				break
			case 6:
				caller.sixteenbitparam2[i] = uint16(memory[cpu.registerHL])
				break
			case 7:
				caller.sixteenbitparam2[i] = uint16(cpu.registerA)
				break
			}

			if i < 0x48 || (i >= 0x80 && i < 0x88) || (i >= 0xC0 && i < 0xC8) {
				caller.sixteenbitparam1[i] = 0
			} else if i < 0x50 || (i >= 0x88 && i < 0x90) || (i >= 0xC8 && i < 0xD0) {
				caller.sixteenbitparam1[i] = 1
			} else if i < 0x58 || (i >= 0x90 && i < 0x98) || (i >= 0xD0 && i < 0xD8) {
				caller.sixteenbitparam1[i] = 2
			} else if i < 0x60 || (i >= 0x98 && i < 0xA0) || (i >= 0xD8 && i < 0xE0) {
				caller.sixteenbitparam1[i] = 3
			} else if i < 0x68 || (i >= 0xA0 && i < 0xA8) || (i >= 0xE0 && i < 0xE8) {
				caller.sixteenbitparam1[i] = 4
			} else if i < 0x70 || (i >= 0xA8 && i < 0xB0) || (i >= 0xE8 && i < 0xF0) {
				caller.sixteenbitparam1[i] = 5
			} else if i < 0x78 || (i >= 0xB0 && i < 0xB8) || (i >= 0xF0 && i < 0xF8) {
				caller.sixteenbitparam1[i] = 6
			} else if i < 0x80 || (i >= 0xB8 && i < 0xC0) || (i >= 0xF8 && i <= 0xFF) {
				caller.sixteenbitparam1[i] = 7
			}
		}
		// Sixteen bit callers end here

		if i <= 0x3B && (i%16 == 1 || i%16 == 9 || i%16 == 3 || i%16 == 11) {

			if i <= 0xF {
				caller.eightbitparam1[i] = cpu.registerBC
			} else if i <= 0x1F {
				caller.eightbitparam1[i] = cpu.registerDE
			} else if i <= 0x2F {
				caller.eightbitparam1[i] = cpu.registerHL
			} else if i <= 0x3F {
				caller.eightbitparam1[i] = cpu.registerSP
			}

			remainder := i % 16
			if remainder == 1 {
				caller.eightBitFuncArray[i] = LD16b
				caller.eightbitparam2[i] = immediateValue
			} else if remainder == 3 {
				caller.eightBitFuncArray[i] = INC16b
				caller.eightbitparam2[i] = cpu.registerHL
			} else if remainder == 9 {
				caller.eightBitFuncArray[i] = ADD16b
				caller.eightbitparam1[i] = cpu.registerHL
				switch i {
				case 0x09:
					caller.eightbitparam2[i] = cpu.registerBC
					break
				case 0x19:
					caller.eightbitparam2[i] = cpu.registerDE
					break
				case 0x29:
					caller.eightbitparam2[i] = cpu.registerHL
					break
				case 0x39:
					caller.eightbitparam2[i] = cpu.registerSP
					break
				}
			} else if remainder == 0xB {
				caller.eightBitFuncArray[i] = DEC16b
				caller.eightbitparam2[i] = cpu.registerHL
			}

		}

		if i <= 0x3E && (i%8 == 6 || i%8 == 4) {
			remainder := i % 8
			caller.eightbitparam2[i] = immediateValue
			if i <= 0x8 {
				caller.eightbitparam1[i] = uint16(cpu.registerB)
			} else if i <= 0xF {
				caller.eightbitparam1[i] = uint16(cpu.registerC)
			} else if i <= 0x18 {
				caller.eightbitparam1[i] = uint16(cpu.registerD)
			} else if i <= 0x1F {
				caller.eightbitparam1[i] = uint16(cpu.registerE)
			} else if i <= 0x28 {
				caller.eightbitparam1[i] = uint16(cpu.registerH)
			} else if i <= 0x2F {
				caller.eightbitparam1[i] = uint16(cpu.registerL)
			} else if i <= 0x38 {
				caller.eightbitparam1[i] = uint16(memory[cpu.registerHL])
			} else {
				caller.eightbitparam1[i] = uint16(cpu.registerA)
			}

			if remainder == 4 {
				caller.eightBitFuncArray[i] = INC
			} else if remainder == 5 {
				caller.eightBitFuncArray[i] = DEC
			} else if remainder == 6 {
				caller.eightBitFuncArray[i] = LDn
			}
		}

		if (i >= 0xC5) && (i <= 0xF5) && (i%16 == 5) {
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

		if (i >= 0xC1) && (i <= 0xF1) && (i%16 == 1) {
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
		case 0x00:
			caller.eightBitFuncArray[i] = NOP
			caller.eightbitparam1[i] = uint16(immediateValue)
			caller.eightbitparam2[i] = uint16(immediateValue)
			break
		case 0x02:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[cpu.registerBC])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x07:
			caller.eightBitFuncArray[i] = RLCA
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = 0
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
		case 0x0F:
			caller.eightBitFuncArray[i] = RRCA
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = 0
			break
		case 0x12:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(memory[cpu.registerDE])
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x17:
			caller.eightBitFuncArray[i] = RLA
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = 0
			break
		case 0x18:
			caller.eightBitFuncArray[i] = JR
			caller.eightbitparam1[i] = immediateValue
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x1A:
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(memory[cpu.registerDE])
			break
		case 0x1F:
			caller.eightBitFuncArray[i] = RRA
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = 0
			break
		case 0x20:
			caller.eightBitFuncArray[i] = JRcc
			caller.eightbitparam1[i] = 1
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x22:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x28:
			caller.eightBitFuncArray[i] = JRcc
			caller.eightbitparam1[i] = 2
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x2A:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(GetMemoryAndIncrement(memory, &cpu.registerHL))
			break
		case 0x2F:
			caller.eightBitFuncArray[i] = CPL
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x30:
			caller.eightBitFuncArray[i] = JRcc
			caller.eightbitparam1[i] = 3
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x32:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(GetMemoryAndDeincrement(memory, &cpu.registerHL))
			caller.eightbitparam2[i] = uint16(cpu.registerA)
			break
		case 0x37:
			caller.eightBitFuncArray[i] = SCF
			caller.eightbitparam1[i] = immediateValue
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x38:
			caller.eightBitFuncArray[i] = JRcc
			caller.eightbitparam1[i] = 4
			caller.eightbitparam2[i] = immediateValue
			break
		case 0x3A:
			// TODO: Fix the way the memory and HL register is accessed and updated
			caller.eightBitFuncArray[i] = LDr
			caller.eightbitparam1[i] = uint16(cpu.registerA)
			caller.eightbitparam2[i] = uint16(GetMemoryAndDeincrement(memory, &cpu.registerHL))
			break
		case 0x3F:
			caller.eightBitFuncArray[i] = CCF
			caller.eightbitparam1[i] = immediateValue
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xC2:
			caller.eightBitFuncArray[i] = JPcc
			caller.eightbitparam1[i] = 1
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xC3:
			caller.eightBitFuncArray[i] = JP
			caller.eightbitparam1[i] = immediateValue
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xC4:
			caller.eightBitFuncArray[i] = CALLcc
			caller.eightbitparam1[i] = 1
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xCA:
			caller.eightBitFuncArray[i] = JPcc
			caller.eightbitparam1[i] = 2
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xCC:
			caller.eightBitFuncArray[i] = CALLcc
			caller.eightbitparam1[i] = 2
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xCD:
			caller.eightBitFuncArray[i] = CALL
			caller.eightbitparam1[i] = immediateValue
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xD2:
			caller.eightBitFuncArray[i] = JPcc
			caller.eightbitparam1[i] = 3
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xD4:
			caller.eightBitFuncArray[i] = CALLcc
			caller.eightbitparam1[i] = 3
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xDA:
			caller.eightBitFuncArray[i] = JPcc
			caller.eightbitparam1[i] = 4
			caller.eightbitparam2[i] = immediateValue
			break
		case 0xDC:
			caller.eightBitFuncArray[i] = CALLcc
			caller.eightbitparam1[i] = 4
			caller.eightbitparam2[i] = immediateValue
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
		case 0xE8:
			caller.eightBitFuncArray[i] = ADDSP
			caller.eightbitparam1[i] = uint16(cpu.registerSP)
			caller.eightbitparam2[i] = uint16(immediateValue)
		case 0xE9:
			caller.eightBitFuncArray[i] = JPHL
			caller.eightbitparam1[i] = cpu.registerHL
			caller.eightbitparam2[i] = immediateValue
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

		if i == 0x76 {
			caller.eightBitFuncArray[i] = HALT
			caller.eightbitparam1[i] = 0
			caller.eightbitparam2[i] = 0
		}

		if i == 0xF3 {
			caller.eightBitFuncArray[i] = DI
			caller.eightbitparam1[i] = 0
			caller.eightbitparam2[i] = 0
		}

		if i == 0xFB {
			caller.eightBitFuncArray[i] = EI
			caller.eightbitparam1[i] = 0
			caller.eightbitparam2[i] = 0
		}

	}
	return caller
}

// Makes a new caller and initializes it with the
// proper functions and params at the proper opcode location
//
// params:
//
//	cpu, a CPU struct containing registers to write to and read from
//	memory, an array of 8 bit values with the size of 0xFFFF
//	immediateValue, the immediate value included in the opcode, can be 8 or 16 bit
func NewCaller(cpu *CPU, memory []uint8, immediateValue uint16) *Opcode_function_caller {
	caller := CallerLoader(cpu, memory, immediateValue)
	return caller
}

// LDn loads a value from a register or immediate value, into a register
//
// params:
//
//	nn, a register to have a value written to
//	n, a register, memory addres, or an 8 bit immediate value to write from
func LDn(nn uint16, n uint16, cpu *CPU, memory []uint8) {
	nn = n
}

// LDr loads a value from a register r2 into another register
// or immediate value r1
//
// params:
//
//	r1, a register to write to
//	r2, a register, memory addres, or an 8 bit immediate value being read from
func LDr(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	r1 = r2
}

// LD16b is like the other LD finctions but intended only for use with 16 bit
// values (like the paired registers) this is the reason every other function has the 16 bit params
//
// params:
//
//	r, a register to write to
//	value, a paired register or an 16 bit immediate value being read from
func LD16b(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
}

// LDFlag was made for the spcific case where the flag register needs to be edited,
// this is the reason all functions have the CPU param
//
// params:
//
//	r, a register to write to
//	value, a paired register or an 16 bit immediate value being read from
func LDFlag(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r = value
	cpu.registerF = cpu.registerF & 0b00110000
}

// Pushes 16 bits worth of values onto the memory stack by SP register
//
// params:
//
//	r, SP register (stack pointer)
//	value, value to be pushed onto the stack. either paired register or immediate value
//	memory, an array of 8 bit values with the size of 0xFFFF, will store the pushed value
func PUSH(r uint16, value uint16, cpu *CPU, memory []uint8) {
	r--
	Write16bToMemory(r, value, memory)
}

// Pops 16 bit worth of values off of the memory stack by SP register
//
// params:
//
//	r1, SP register (stack pointer)
//	r2, paired register to push value onto
//	memory, an array of 8 bit values with the size of 0xFFFF, will have values poped off it
func POP(r1 uint16, r2 uint16, cpu *CPU, memory []uint8) {
	Read16bFromMemory(r1, r2, memory)
}

// Adds a value to the arithmitic registry (register A)
//
// params:
//
//	A, arithmitic register (register A)
//	n, a value to be added to arithmitic register, can be value from memory,
//		   other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
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
	if (A&0b111)+(temp&0b111) >= 0b1000 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (A&0b1111111)+(temp&0b1111111) >= 0b10000000 {
		cpu.registerF = cpu.registerF | 0b00010000
	}
	A = result & 0xFF
}

// Adds a value in a paired register to paired register HL
//
// params:
//
//	HL, paired register HL
//	n, any paired register
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func ADD16b(HL uint16, n uint16, cpu *CPU, memory []uint8) {
	result := HL + n
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (HL&0b11111111111)+(n&0b11111111111) >= 0b100000000000 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	if (HL&0b111111111111111)+(n&0b111111111111111) >= 0b1000000000000000 {
		cpu.registerF = cpu.registerF | 0b00010000
	}
	HL = result & 0xFFFF
}

// Adds an 8 bit signed immediate value to the SP register (Stack Pointer)
//
// params:
//
//	SP, SP register (Stack Pointer)
//	n, 8 bit signed immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func ADDSP(SP uint16, n uint16, cpu *CPU, memory []uint8) {
	//TODO: Figure out if signed immediate value would
	// 		subtract from SP if it was negative
	result := SP + n
	if cpu.registerF&0b10000000 == 0b10000000 {
		cpu.registerF = cpu.registerF ^ 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	// TODO: Figure out what set or reset acoording to operation means
	// 		 in this context for flag H and C
	SP = result & 0xFFFF
}

// Subtracts a value from the arithmitic register (register A)
//
// params:
//
//	A, arithmitic register (register A)
//	n, a value to be subtracted from arithmitic register,
//		can be value from memory, other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func SUB(A uint16, n uint16, cpu *CPU, memory []uint8) {
	temp := n
	if n > 0xFF {
		temp = uint16(memory[n])
	}
	result := A - temp
	if A == temp {
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
//
//	A, arithmitic register (register A)
//	n, a value to be added to arithmitic register, can be value from memory,
//		   other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
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
//
//	A, arithmitic register (register A)
//	n, a value to be added to arithmitic register, can be value from memory,
//		   other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
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
//
//	A, arithmitic register (register A)
//	n, a value to be added to arithmitic register, can be value from memory,
//		   other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
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

// Compares a value with the arithmitic register (register A), importantly this will not
// be entered into register A instead the results are discarded otherwise this is the same as SUB
//
// params:
//
//	A, arithmitic register (register A)
//	n, a value to be added to arithmitic register, can be value from memory,
//		   other registers, or immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func CP(A uint16, n uint16, cpu *CPU, memory []uint8) {
	temp := n
	if n > 0xFF {
		temp = uint16(memory[n])
	}
	if A == temp {
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
}

// Increments a non-paired register or value in memory
//
// params:
//
//	n, a non paired register or a spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func INC(n uint16, none uint16, cpu *CPU, memory []uint8) {
	temp := n + 1
	if temp&0xFF == 0 {
		cpu.registerF = cpu.registerF | 0b10000000
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (n&0b111)+(temp&0b111) >= 0b1000 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	n++
}

// Increments a 16 bit register
//
// params:
//
//	n, a 16 bit register
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0xFFFF
func INC16b(nn uint16, none uint16, cpu *CPU, memory []uint8) {
	nn++
}

// Deincrements a non paired register or value in memory
//
// params:
//
//	n, a non paired register or a spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func DEC(n uint16, none uint16, cpu *CPU, memory []uint8) {
	temp := n - 1
	if temp&0xFF == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) != 0b01000000 {
		SetNFlag(cpu)
	}
	if (n&0b111)-(temp&0b111) >= 0 {
		SetHFlag(cpu)
	}
	n--
}

// Deincrements a 16 bit register
//
// params:
//
//	n, a 16 bit register
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func DEC16b(nn uint16, none uint16, cpu *CPU, memory []uint8) {
	nn--
}

// Swaps the lower and upper bits of n
//
// params:
//
//	n, a non paired register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
func SWAP(n uint16, none uint16, cpu *CPU, mempry []uint8) {
	lower := n & 0xF
	upper := n & 0xF0
	lower = lower << 4
	upper = upper >> 4
	n = lower | upper
	if n == 0 {
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
}

// TODO: implement DAA function
func DAA(n uint16, none uint16, cpu *CPU, memory []uint8) {
	return
}

// Compliments/flips every bit in the arithemtic register (registerA)
//
// params:
//
//	n, registerA
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func CPL(n uint16, none uint16, cpu *CPU, memory []uint8) {
	n = n ^ 0b11111111
	if (cpu.registerF & 0b01000000) != 0b01000000 {
		cpu.registerF = cpu.registerF | 0b01000000
	}
	if (cpu.registerF & 0b00100000) != 0b00100000 {
		cpu.registerF = cpu.registerF | 0b00100000
	}
	cpu.cycles += 4
}

// Compliments/flips the carry flag in the flag register (registerF)
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func CCF(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.registerF = cpu.registerF ^ 0b00010000
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		cpu.registerF = cpu.registerF ^ 0b00100000
	}
	cpu.cycles += 4
}

// Sets the carry flag in the flag register (registerF)
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func SCF(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.registerF = cpu.registerF | 0b00010000
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		cpu.registerF = cpu.registerF ^ 0b01000000
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		cpu.registerF = cpu.registerF ^ 0b00100000
	}
	cpu.cycles += 4
}

// No operation, does nothing
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func NOP(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.cycles += 4
	cpu.registerPC++
}

// Halts the cpu and prevents it from performing instructions
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func HALT(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.cycles += 4
	cpu.halted = true
}

// TODO: Implement stop instruction
func STOP(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	return
}

// Disables interrupts, not immediately but after instruction
// after DI is executed
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func DI(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.cycles += 4
	cpu.interrupts = false
}

// Enables interrupts, not immediately but after instruction
// after EI is executed
//
// params:
//
//	none1, not used in this function
//	none2, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func EI(none1 uint16, none2 uint16, cpu *CPU, memory []uint8) {
	cpu.cycles += 4
	cpu.interrupts = true
}

// Rotates register A left, old bit 7 becomes carry flag
//
// params:
//
//	n, register A
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RLCA(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b10000000
	n = (n << 1) & 0xFF
	if carry > 0 {
		SetCFlag(cpu)
		n = n | 0b00000001
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 4
}

// Rotates register A left through carry flag,
// old bit 7 becomes carry flag
//
// params:
//
//	n, register A
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RLA(n uint16, none2 uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b10000000
	n = (n << 1) & 0xFF
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		n = n | 0x01
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 4
}

// Rotates register A right, old bit 0 becomes carry flag
//
// params:
//
//	n, register A
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RRCA(n uint16, none2 uint16, cpu *CPU, memory []uint8) {
	carry := n & 0x01
	n = n >> 1
	if carry > 0 {
		SetCFlag(cpu)
		n = n | 0b10000000
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 4
}

// Rotates register A right through carry flag,
// old bit 0 becomes carry flag
//
// params:
//
//	n, register A
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RRA(n uint16, none2 uint16, cpu *CPU, memory []uint8) {
	carry := n & 0x01
	n = n >> 1
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		n = n | 0b10000000
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 4
}

// Rotates a register or a spot in memory left,
// old bit 7 becomes carry flag
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RLCn(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b10000000
	n = (n << 1) & 0xFF
	if carry > 0 {
		SetCFlag(cpu)
		n = n | 0b00000001
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 8
}

// Rotates a register or a spot in memory left
// through carry flag, old bit 7 becomes carry flag
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RLn(n uint16, none2 uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b10000000
	n = (n << 1) & 0xFF
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		n = n | 0x01
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 4
}

// Rotates a register or a spot in memory right,
// old bit 0 becomes carry flag
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RRCn(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b00000001
	n = (n >> 1) & 0xFF
	if carry > 0 {
		SetCFlag(cpu)
		n = n | 0b10000000
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 8
}

// Rotates a register or a spot in memory right
// through carry flag, old bit 0 becomes carry flag
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RRn(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := n & 0b00000001
	n = (n >> 1) & 0xFF
	if (cpu.registerF & 0b00010000) == 0b00010000 {
		n = n | 0b10000000
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b00010000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	cpu.cycles += 8
}

// Shifts a register or a value in memory
// left into carry, LSB of is set to 0
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func SLA(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := uint8(n) & 0b10000000
	n = (n << 1) & 0xFF
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b0001000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
}

// Shifts a register or a value in memory
// right into carry, MSB doesn't change
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func SRA(n uint16, none uint16, cpu *CPU, memory []uint8) {
	MSB := uint8(n) & 0b10000000
	carry := uint8(n) & 1
	n = (n >> 1) & 0xFF
	n = n | uint16(MSB)
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b0001000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
}

// Shifts a register or a value in memory
// right into carry, MSB is set to 0
//
// params:
//
//	n, a register or spot in memory
//	none, not used in this function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func SRL(n uint16, none uint16, cpu *CPU, memory []uint8) {
	carry := uint8(n) & 1
	n = (n >> 1) & 0xFF
	if n == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	if (cpu.registerF & 0b00100000) == 0b00100000 {
		ResetHFlag(cpu)
	}
	if carry > 0 {
		SetCFlag(cpu)
	} else {
		if (cpu.registerF & 0b0001000) == 0b00010000 {
			ResetCFlag(cpu)
		}
	}
}

// Tests a bit in a register or value in memory
//
// params:
//
//	b, the position of the bit to be examined
//	r, a register or spot in memory
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func BIT(b uint16, r uint16, cpu *CPU, memory []uint8) {
	var bit uint8
	switch b {
	case 0:
		bit = 0b00000001
		break
	case 1:
		bit = 0b00000010
		break
	case 2:
		bit = 0b00000100
		break
	case 3:
		bit = 0b00001000
		break
	case 4:
		bit = 0b00010000
		break
	case 5:
		bit = 0b00100000
		break
	case 6:
		bit = 0b01000000
		break
	case 7:
		bit = 0b10000000
		break
	}
	if (r & uint16(bit)) == 0 {
		SetZFlag(cpu)
	}
	if (cpu.registerF & 0b01000000) == 0b01000000 {
		ResetNFlag(cpu)
	}
	SetHFlag(cpu)
}

// Sets a bit in a register or value in memory
//
// params:
//
//	b, the position of the bit to be set
//	r, a register or spot in memory
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func SET(b uint16, r uint16, cpu *CPU, memory []uint8) {
	var bit uint8
	switch b {
	case 0:
		bit = 0b00000001
		break
	case 1:
		bit = 0b00000010
		break
	case 2:
		bit = 0b00000100
		break
	case 3:
		bit = 0b00001000
		break
	case 4:
		bit = 0b00010000
		break
	case 5:
		bit = 0b00100000
		break
	case 6:
		bit = 0b01000000
		break
	case 7:
		bit = 0b10000000
		break
	}
	r = r | uint16(bit)
}

// Resets a bit in a register or value in memory
//
// params:
//
//	b, the position of the bit to be reset
//	r, a register or spot in memory
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func RES(b uint16, r uint16, cpu *CPU, memory []uint8) {
	var bit uint8
	switch b {
	case 0:
		bit = 0b00000001
		break
	case 1:
		bit = 0b00000010
		break
	case 2:
		bit = 0b00000100
		break
	case 3:
		bit = 0b00001000
		break
	case 4:
		bit = 0b00010000
		break
	case 5:
		bit = 0b00100000
		break
	case 6:
		bit = 0b01000000
		break
	case 7:
		bit = 0b10000000
		break
	}
	if (r & uint16(bit)) > 0 {
		r = r ^ uint16(bit)
	}
}

// Jumps to memory address pointed to by an immediate value
//
// params:
//
//	nn, 16 bit immediate value
//	none, not used in function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func JP(nn uint16, none uint16, cpu *CPU, memory []uint8) {
	cpu.registerPC = nn
	cpu.cycles += 12
}

// Jumps to memory address pointed to by an immediate
// value if certain conditions are true
//
// params:
//
//	cc, condition to check
//	nn, 16 bit immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func JPcc(cc uint16, nn uint16, cpu *CPU, memory []uint8) {
	switch cc {
	case 1:
		if !IsZFlagSet(cpu) {
			cpu.registerPC = nn
		}
		break
	case 2:
		if IsZFlagSet(cpu) {
			cpu.registerPC = nn
		}
		break
	case 3:
		if !IsCFlagSet(cpu) {
			cpu.registerPC = nn
		}
		break
	case 4:
		if IsCFlagSet(cpu) {
			cpu.registerPC = nn
		}
		break
	}
	cpu.cycles += 12
}

// Jumps to memory address pointed to by register HL
//
// params:
//
//	r, register HL
//	none, not used in function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func JPHL(r uint16, none uint16, cpu *CPU, memory []uint8) {
	cpu.registerPC = r
	cpu.cycles += 4
}

// Jumps to memory address by adding a signed immediate
// value to current address
//
// params:
//
//	n, 8 bit signed immediate value
//	none, not used in function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func JR(n uint16, none uint16, cpu *CPU, memory []uint8) {
	if int8(n) >= 0 {
		cpu.registerPC += n
	} else {
		cpu.registerPC -= (n - 1) ^ 0b11111111
	}
	cpu.cycles += 8
}

// Jumps to memory address pointed to by a signed immediate
// value if certain conditions are true
//
// params:
//
//	cc, condition to check
//	nn, 8 bit signed immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func JRcc(cc uint16, n uint16, cpu *CPU, memory []uint8) {
	cpu.registerPC += n
	switch cc {
	case 1:
		if !IsZFlagSet(cpu) {
			if int8(n) >= 0 {
				cpu.registerPC += n
			} else {
				cpu.registerPC -= (n - 1) ^ 0b11111111
			}
		}
		break
	case 2:
		if IsZFlagSet(cpu) {
			if int8(n) >= 0 {
				cpu.registerPC += n
			} else {
				cpu.registerPC -= (n - 1) ^ 0b11111111
			}
		}
		break
	case 3:
		if !IsCFlagSet(cpu) {
			if int8(n) >= 0 {
				cpu.registerPC += n
			} else {
				cpu.registerPC -= (n - 1) ^ 0b11111111
			}
		}
		break
	case 4:
		if IsCFlagSet(cpu) {
			if int8(n) >= 0 {
				cpu.registerPC += n
			} else {
				cpu.registerPC -= (n - 1) ^ 0b11111111
			}
		}
		break
	}
	cpu.cycles += 8
}

// Push address of next instruction onto stack and
// then jumps to address at immediate value
//
// params:
//
//	nn, 16 bit immediate value
//	none, not used in function
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func CALL(nn uint16, none uint16, cpu *CPU, memory []uint8) {
	cpu.registerPC++
	PUSH(cpu.registerSP, cpu.registerPC, cpu, memory)
	cpu.registerPC = nn
	cpu.cycles += 12
}

// Push address of next instruction onto stack and
// then jumps to address at immediate value if condition is true
//
// params:
//
//	cc, condition to check
//	nn, 16 bit immediate value
//	cpu, CPU struct to edit flag register (register F)
//	memory, an array of 8 bit values with the size of 0x10000
func CALLcc(cc uint16, nn uint16, cpu *CPU, memory []uint8) {
	switch cc {
	case 1:
		if !IsZFlagSet(cpu) {
			CALL(nn, nn, cpu, memory)
		}
		break
	case 2:
		if IsZFlagSet(cpu) {
			CALL(nn, nn, cpu, memory)
		}
		break
	case 3:
		if !IsCFlagSet(cpu) {
			CALL(nn, nn, cpu, memory)
		}
		break
	case 4:
		if IsCFlagSet(cpu) {
			CALL(nn, nn, cpu, memory)
		}
		break
	}
}

// Takes in an opcode and runs the function with appropriate params associated with that code
//
// params:
//
//	opcode, can be 8 or 16 bit value 16 bit has to begin at 0xCB00 and ends at 0xCBFF
//			and might be followed by an 8 or 16 bit immediate value
//	cpu, where the registers are read from and written to
//	memory, An array of 8 bit integers that is 0x10000 addresses long
func ReadOpcode(opcode uint32, cpu *CPU, memory []uint8) {

	var immediateValue uint16

	if (opcode > 0xFF) && ((opcode&0xCB00) != 0xCB00) || opcode > 0xFFFF {
		if opcode > 0xFFFF {
			immediateValue = uint16(opcode & 0xFFFF)
			opcode = opcode >> 16
		} else {
			immediateValue = uint16(opcode) & 0xFF
			opcode = opcode >> 8
		}
	}

	caller := NewCaller(cpu, memory, immediateValue)

	if (opcode > 255) && ((opcode & 0xCB00) == 0xCB00) {
		function := caller.sixteenBitFuncArray[opcode^0xCB00]
		first := caller.sixteenbitparam1[opcode^0xCB00]
		second := caller.sixteenbitparam2[opcode^0xCB00]
		function(first, second, cpu, memory)
	} else {
		function := caller.eightBitFuncArray[opcode]
		first := caller.eightbitparam1[opcode]
		second := caller.eightbitparam2[opcode]
		function(first, second, cpu, memory)
	}

}
