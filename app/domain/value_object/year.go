package value_object

type Year struct {
	value string
}

func NewYear(brandString string) (Year, error) {
	return Year{value: brandString}, nil
}

func (b Year) String() string {
	return b.value
}

func (b Year) Equals(other Year) bool {
	return b.value == other.value
}
