package main

func main() {
	e1 := ComplexEmployee{
		Name: "Akshay",
		Contact: []string{"988998899889","1234123412","112112","889997"},
		ID: 1,
		PrimaryAddress: Address{
			State: "MP",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450331,
		},
		SecondAddress: &Address{
			State: "MP",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450331,
		},
		Designation: "SWE",
	}
	GoTemplate(e1)
	conversionDirect(e1)
	GoRecursiveTemplate(10)
}

