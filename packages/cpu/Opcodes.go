package GBCPU

// This struct is to load functions from an array based on opcode
type Opcode_function_loader struct {
	eightBitFuncArray   [255]func(uint8, uint8) uint8
	eightbitparam1      [255]uint8
	eightbitparam2      [255]uint8
	sixteenBitFuncArray [255]func(uint16, uint16) uint16
	sixteenbitparam1    [255]uint16
	sixteenbitparam2    [255]uint16
}

// Function takes Opcode_function_loader and the CPU and loades the loader
// with all the functions and params that will be called by Opcodes
func loadLoader(loader *Opcode_function_loader, cpu *CPU) {

	for i := 0; i <= 255; i++ {

		if i < 30 && i%8 == 6 {
			loader.eightBitFuncArray[i] = LDn
		}

		if (i >= 0x77 && i <= 0x7F) || (i >= 0x40 && i <= 0x75) || i == 36 {

			loader.eightBitFuncArray[i] = LDr

			if i >= 0x78 && i <= 0x7F {
				loader.eightbitparam1[i] = cpu.registerA
			} else if i >= 0x40 && i >= 0x47 {
				loader.eightbitparam1[i] = cpu.registerB
			} else if i >= 0x48 && i <= 0x4F {
				loader.eightbitparam1[i] = cpu.registerC
			} else if i >= 0x50 && i <= 0x57 {
				loader.eightbitparam1[i] = cpu.registerD
			} else if i >= 0x58 && i <= 0x5F {
				loader.eightbitparam1[i] = cpu.registerE
			} else if i >= 0x60 && i <= 0x67 {
				loader.eightbitparam1[i] = cpu.registerH
			} else if i >= 0x68 && i <= 0x6F {
				loader.eightbitparam1[i] = cpu.registerL
			} else if (i >= 0x70 && i <= 0x75) || i == 36 {
				loader.eightbitparam1[i] = cpu.registerHL
			}

			remainder := i % 8

			switch remainder {
			case 0:
				loader.eightbitparam2[i] = cpu.registerB
				break
			case 1:
				loader.eightbitparam2[i] = cpu.registerC
				break
			case 2:
				loader.eightbitparam2[i] = cpu.registerD
				break
			case 3:
				loader.eightbitparam2[i] = cpu.registerE
				break
			case 4:
				loader.eightbitparam2[i] = cpu.registerH
				break
			case 5:
				loader.eightbitparam2[i] = cpu.registerL
				break
			case 6:
				if i == 36 {
					// This is a temporary param for immediate value,
					// replace registryA with that later
					loader.eightbitparam2[i] = cpu.registerA
				} else {
					loader.eightbitparam2[i] = cpu.registerC
				}
				break
			case 7:
				loader.eightbitparam2[i] = cpu.registerA
				break
			}

		}

	}
}

// LDn loads a value from a register nn into another register
// or immediate value n

func LDn(nn uint8, n uint8) uint8 {
	n = nn
	return n
}

// LDr loads a value from a register r2 into another register
// or immediate value r1
func LDr(r1 uint8, r2 uint8) uint8 {
	r1 = r2
	return 0
}

/*
Takes in an opcode and runs the function associated with that code
param: an 8 bit or 16 bit value (16 bit has to begin at 0xCB00 and ends at 0xCBFF)
*/
func ReadOpcode(opcode uint16) {
	cpu := NewCPU()
	loader := new(Opcode_function_loader)
	loadLoader(loader, cpu)
	if opcode > 255 {
		function := loader.sixteenBitFuncArray[opcode-0xCB00]
		first := loader.sixteenbitparam1[opcode-0xCB00]
		second := loader.sixteenbitparam2[opcode-0xCB00]
		function(first, second)
	} else {
		function := loader.eightBitFuncArray[opcode]
		first := loader.eightbitparam1[opcode]
		second := loader.eightbitparam2[opcode]
		function(first, second)
	}
}
