package base_type

type DefinedStringType string

func (t DefinedStringType) String() string {
	return string(t)
}
