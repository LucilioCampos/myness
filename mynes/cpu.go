package mynes

const CPUFrequency = 1789773

// CPU has instructions

// Opcodes - 256 - 151 Realy used

// Registers
// Accumulator(A) - The accumulator can read and write to memory. It is used to store arithmetic and logic results such as addition and subtraction.
// X Index(X) - The x index is can read and write to memory. It is used primarily as a counter in loops, or for addressing memory, but can also temporarily store data like the accumulator.
// Y Index(Y) - Much like the x index, however they are not completely interchangeable. Some operations are only available for each register.
// Flag(P) - The register holds value of 7 different flags which can only have a value of 0 or 1 and hence can be represented in a single register. The bits represent the status of the processor.
// Stack Pointer(SP) - The stack pointer hold the address to the current location on the Stack. The stack is a way to store data by pushing or popping data to and from a section of memory.
// Program Counter(PC) - This is a 16-bit register unlike other registers which are only 8-bit in length, it indicates where the processor is in the program sequence.
type Operation struct {
	addr uint16
	PC   uint16
	mode byte
}
type Cpu struct {
	Memory
	status  byte
	A       byte
	X       byte
	Y       byte
	flags   Flag
	PC      uint16
	SP      byte
	operate [256]func(*Operation)
	cycles  uint16
}

func NewCPU(console *Console) *Cpu {
	cpu := Cpu{Memory: NewCPUMemory(console)}
	return &cpu
}

type Flag struct {
	C byte
	Z byte
	I byte
	D byte
	B byte
	U byte
	V byte
	S byte
}

func (cpu *Cpu) GetFlags() {
	cpu.flags.C |= (1 << 0)
	cpu.flags.Z |= (1 << 1)
	cpu.flags.I |= (1 << 2)
	cpu.flags.D |= (1 << 3)
	cpu.flags.B |= (1 << 4)
	cpu.flags.U |= (1 << 5)
	cpu.flags.V |= (1 << 6)
	cpu.flags.S |= (1 << 7)
}

func (cpu *Cpu) GetFlag(flag byte) byte {
	if (cpu.status & flag) > 0 {
		return 1
	}
	return 0
}

func (cpu *Cpu) SetFlag(flag byte, v bool) {
	if v {
		cpu.status |= flag
	} else {
		cpu.status &= ^flag
	}
}

func (cpu *Cpu) Step() int {
	if cpu.cycles > 0 {
		cpu.cycles--
		return 1
	}

	cycles := cpu.cycles

	return int(cpu.cycles - cycles)
}
