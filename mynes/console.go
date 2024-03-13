package mynes

type Console struct {
	CPU       *Cpu
	Cartridge *Cartridge
	Mapper    Mapper
	RAM       []byte
}

func NewConsole(path string) (*Console, error) {
	cartridge, err := LoadNESFile(path)
	if err != nil {
		return nil, err
	}
	ram := make([]byte, 2048)
	console := Console{nil, cartridge, nil, ram}
	mapper, err := NewMapper(&console)
	if err != nil {
		return nil, err
	}
	console.Mapper = mapper
	console.CPU = NewCPU(&console)
	return &console, nil
}

func (console *Console) StepSeconds(seconds float64) {
	cycles := int(CPUFrequency * seconds)
	for cycles > 0 {
		cycles -= console.Step()
	}
}

func (console *Console) Step() int {
	cpuCycles := console.CPU.Step()
	ppuCycles := cpuCycles * 3
	for i := 0; i < ppuCycles; i++ {
		console.Mapper.Step()
	}
	return cpuCycles
}
