package mynes

import (
	"encoding/gob"
	"fmt"
)

type Mapper interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
	Step()
	Save(encoder *gob.Encoder) error
	Load(decoder *gob.Decoder) error
}

type Mapper2 struct {
	*Cartridge
	prgBanks int
	prgBank1 int
	prgBank2 int
}

// Load implements Mapper.
// Subtle: this method shadows the method (*Cartridge).Load of Mapper2.Cartridge.
func (m *Mapper2) Load(decoder *gob.Decoder) error {
	panic("unimplemented")
}

// Read implements Mapper.
func (m *Mapper2) Read(address uint16) byte {
	panic("unimplemented")
}

// Save implements Mapper.
// Subtle: this method shadows the method (*Cartridge).Save of Mapper2.Cartridge.
func (m *Mapper2) Save(encoder *gob.Encoder) error {
	panic("unimplemented")
}

// Step implements Mapper.
func (m *Mapper2) Step() {
	panic("unimplemented")
}

// Write implements Mapper.
func (m *Mapper2) Write(address uint16, value byte) {
	panic("unimplemented")
}

func NewMapper2(cartridge *Cartridge) Mapper {
	prgBanks := len(cartridge.PRG) / 0x4000
	prgBank1 := 0
	prgBank2 := prgBanks - 1
	return &Mapper2{cartridge, prgBanks, prgBank1, prgBank2}
}

func NewMapper(console *Console) (Mapper, error) {
	cartridge := console.Cartridge
	switch cartridge.Mapper {
	case 0:
		return NewMapper2(cartridge), nil
	}
	err := fmt.Errorf("unsupported mapper: %d", cartridge.Mapper)
	return nil, err
}
