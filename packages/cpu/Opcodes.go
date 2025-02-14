package GBCPU

// This struct is to load functions from an array based on opcode
type Opcode_function_loader struct {
	eitghtBitFuncArray  [255]*func()
	sixteenBitFuncArray [255]*func()
}

/*
LD loads a value from a register nn into a another register
or immediate value n
*/
func LD(nn uint8, n uint8) uint8 {
	n = nn
	return n
}

/*
Takes in an opcode and runs the function associated with that code
param: an 8 bit or 16 bit value (16 bit has to begin at 0xCB00 and ends at 0xCBFF)
*/
func ReadOpcode(opcode uint16) {
	loader := new(Opcode_function_loader)
	if opcode > 255 {
		loader.sixteenBitFuncArray[opcode-51968]
	} else {
		loader.eitghtBitFuncArray[opcode]
	}
}
