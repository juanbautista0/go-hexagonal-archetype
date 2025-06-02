package value_object

type Brand struct {
	value string
}

func NewBrand(brandString string) (Brand, error) {
	return Brand{value: brandString}, nil
}

func (b Brand) String() string {
	return b.value
}

func (b Brand) Equals(other Brand) bool {
	return b.value == other.value
}
