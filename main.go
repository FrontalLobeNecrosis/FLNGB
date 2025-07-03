package main

import (
	GBCPU "github.com/FrontalLobeNecrosis/FLNGB/packages/cpu"
	//GBPPU "github.com/FrontalLobeNecrosis/FLNGB/packages/ppu"
)

func main() {
	var quit bool = false
	cpu := GBCPU.NewCPU()
	memory := GBCPU.NewMemory()
	var immediateValue uint16 = 0
	var opcode uint16 = 0
	functionCaller := GBCPU.NewCaller(cpu, memory, &immediateValue)
	for !quit {
		GBCPU.ReadOpcode(functionCaller, opcode, cpu, memory)
	}

}
