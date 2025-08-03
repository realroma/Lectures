package buisnes

type Buisnes struct {
	Name  string //`json: "name"`
	Pay   int    //`json: "pay"`
	Cost  int    //`json: "cost"`
	Grade string //`json: "grade"`
}

func (b *Buisnes) ChangeGrade(grade string) {
	b.Grade = grade
}

func (b *Buisnes) Upgrade() {

}

func (b *Buisnes) Busines() Buisnes {
	return *b
}
