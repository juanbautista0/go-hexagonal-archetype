package value_object

type Model struct {
	value string
}

func NewModel(brandString string) (Model, error) {
	return Model{value: brandString}, nil
}

func (b Model) String() string {
	return b.value
}

func (b Model) Equals(other Model) bool {
	return b.value == other.value
}
