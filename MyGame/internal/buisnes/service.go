package buisnes

type Service interface {
	New(name string, cost int, pay int, grade string)
	GetBusinesByName(name string)
}

// type service struct {
// 	storage Storage
// }

func New(name string, cost int, pay int, grade string) *Buisnes {
	b := Buisnes{Name: name, Cost: cost, Pay: pay, Grade: grade}
	return &b
}

func GetBusinesByName(name string) *Buisnes {
	b := New("None", 0, 0, "None")
	if b != nil {
		return b
	}
	return nil
}
