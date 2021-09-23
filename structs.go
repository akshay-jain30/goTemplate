package main


type SimpleEmployee struct {
	FirstName string
	FamilyName string
	ID int
	Designation string
	Address Address
	Contact []string
	Metadata Metadata
}

type ComplexEmployee struct {
	Name string
	ID int
	Designation string
	PrimaryAddress Address
	SecondAddress *Address
	Contact []string
}

type Address struct {
	StreetName string
	City string
	State string
	Pin int
	FormattedAddress string
}

type Metadata struct {
	metadata_1 string
	metadata_2 string
	metadata_3 string
	metadata_4 string
	metadata_5 string
	metadata_6 string
	metadata_7 string
	metadata_8 string
	metadata_9 string
	metadata_10 string
}

type RecursiveStruct struct {
	Depth int
	Property_1 string
	Property_2 string
	Property_3 string
	Property_4 string
	Property_5 string
	Property_6 string
	Property_7 string
	Property_8 string
	Property_9 string
	Property_10 string
	Next *RecursiveStruct
	DynamicTemplateChecker string
}

func (r *RecursiveStruct) Decrement() int{
	r.Depth--
	return r.Depth
}

func (r *RecursiveStruct) GetProperty() string{
	if r.Depth%2==1{
		return "odd"
	}else{
		return "even"
	}
}

