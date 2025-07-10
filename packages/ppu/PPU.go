package GBPPU

type colour uint8

const (
	black colour = iota
	light_grey
	dark_grey
	white
	VRAM_START = 0x8000
	VRAM_END   = 0x9FFF
	VRAM_SIZE  = VRAM_END - VRAM_START + 1
)

type PPU struct {
	pallet [VRAM_END - VRAM_START + 1]colour
}

func initPPU(ppu *PPU) {
	for i := 0; uint16(i) <= VRAM_SIZE; i++ {
		ppu.pallet[i] = 0
	}
}

func NewPPU() *PPU {
	ppu := new(PPU)
	initPPU(ppu)
	return ppu
}

func ReadPallet(memory []uint8) {

}
