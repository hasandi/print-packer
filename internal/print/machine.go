package print

type Machine struct {
	length, width float64
}

func NewMachine(length, width float64) *Machine {
	return &Machine{length: length, width: width}
}

func (m Machine) Width() float64 {
	return m.width
}

func (m Machine) Length() float64 {
	return m.length
}
