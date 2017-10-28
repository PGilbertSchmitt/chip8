package cpu

type CPU struct {
	/*  */
	// Entire memory storage
	memory [4096]byte

	// V Registers from 0 to F
	vRegister [16]uint16

	// Register for storing memory adresses
	iRegister uint16

	// Count down at 60Hz
	delay byte
	buzz byte

	// Program counter
	programCounter uint16

	// Stack counter
	stackPointer uint16
	stack [16]uint16
}

type CPUOption func(c* CPU)

func NewCPU(options ...CPUOption) (*CPU, error) {
	cpu := &CPU{}
	
	return cpu, nil
}