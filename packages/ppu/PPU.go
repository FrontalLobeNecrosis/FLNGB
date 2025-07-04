package GBPPU

type colour uint8

const (
	black colour = iota
	light_grey
	dark_grey
	white
)

type PPU struct {
	size   uint16
	pallet []colour
}
