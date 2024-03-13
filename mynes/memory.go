package mynes

import "log"

type Memory interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
}

type cpuMemory struct {
	console *Console
}

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}
}

func (mem *cpuMemory) Read(address uint16) byte {
	switch {
	case address < 0x2000:
		return mem.console.RAM[address%0x0800]
	case address < 0x4000:
	case address == 0x4014:
	case address == 0x4015:
	case address == 0x4016:
	case address == 0x4017:
	case address < 0x6000:
		// TODO: I/O registers
	case address >= 0x6000:
		return mem.console.Mapper.Read(address)
	default:
		log.Fatalf("unhandled cpu memory read at address: 0x%04X", address)
	}
	return 0
}

func (mem *cpuMemory) Write(address uint16, value byte) {
	switch {
	case address < 0x2000:
		mem.console.RAM[address%0x0800] = value
	case address < 0x4000:

	case address < 0x4014:
	case address == 0x4014:
	case address == 0x4015:
	case address == 0x4016:

	case address == 0x4017:
	case address < 0x6000:
		// TODO: I/O registers
	case address >= 0x6000:
		mem.console.Mapper.Write(address, value)
	default:
		log.Fatalf("unhandled cpu memory write at address: 0x%04X", address)
	}
}
