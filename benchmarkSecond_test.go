package main

import "testing"

func BenchmarkDirectConversion(b *testing.B) {
	initialiseTemplates()
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conversionDirect(e1)
	}
}


func BenchmarkGoTemplateNested_0Conversion(b *testing.B) {
	initialiseTemplates()
	e1 := ComplexEmployee{
		Name: "Akshay",
		Contact: []string{"988998899889","1234123412","112112","889997"},
		ID: 1,
		PrimaryAddress: Address{
			State: "MP",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450332,
		},
		SecondAddress: &Address{
			State: "MP",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450332,
		},
		Designation: "SWE",
	}
	for i := 0; i < b.N; i++ {
		GoTemplate(e1)
	}
}

func BenchmarkGoTemplateNested_1Conversion(b *testing.B) {
	initialiseTemplates()
	e1 := ComplexEmployee{
		Name: "Akshay",
		Contact: []string{"988998899889","1234123412","112112","889997"},
		ID: 1,
		PrimaryAddress: Address{
			State: "MH",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450331,
		},
		SecondAddress: &Address{
			State: "MH",
			City: "Indore",
			StreetName: "Lalbagh Road",
			Pin: 450331,
		},
		Designation: "SWE",
	}
	for i := 0; i < b.N; i++ {
		GoTemplate(e1)
	}
}

func BenchmarkGoTemplateNested_2Conversion(b *testing.B) {
	initialiseTemplates()
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
	for i := 0; i < b.N; i++ {
		GoTemplate(e1)
	}
}

func BenchmarkGoTemplateRecursive_1000Conversion_dynamic(b *testing.B) {
	initialiseTemplates()
	for i := 0; i < b.N; i++ {
		GoRecursiveTemplate(1000, false)
	}
}

func BenchmarkGoTemplateRecursive_1000Conversion_static(b *testing.B) {
	initialiseTemplates()
	for i := 0; i < b.N; i++ {
		GoRecursiveTemplate(1000, true)
	}
}

func BenchmarkGoPlugin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bindThroughPlugin(1)
	}
}