package encodestruct

import (
	"fmt"
	"time"
)

type ComplexStruct struct {
	Name     string
	Age      uint
	Birthday time.Time
}

func ExampleEncode() {
	t, _ := time.ParseInLocation(time.RFC3339, "1990-06-04T03:15:47.000Z", time.UTC)
	v := ComplexStruct{
		Name:     "jesse michael",
		Age:      29,
		Birthday: t,
	}

	encoded := Encode(&v)
	fmt.Println(encoded)
	// Output:
	// Ov+BAwEBDUNvbXBsZXhTdHJ1Y3QB/4IAAQMBBE5hbWUBDAABA0FnZQEGAAEIQmlydGhkYXkB/4QAAAAQ/4MFAQEEVGltZQH/hAAAACX/ggENamVzc2UgbWljaGFlbAEdAQ8BAAAADp37yWMAAAAA//8A
}

func ExampleDecode() {
	encoded := "Ov+BAwEBDUNvbXBsZXhTdHJ1Y3QB/4IAAQMBBE5hbWUBDAABA0FnZQEGAAEIQmlydGhkYXkB/4QAAAAQ/4MFAQEEVGltZQH/hAAAACX/ggENamVzc2UgbWljaGFlbAEdAQ8BAAAADp37yWMAAAAA//8A"
	var v ComplexStruct
	Decode(encoded, &v)
	fmt.Printf("%v", v)
	// Output:
	// {jesse michael 29 1990-06-04 03:15:47 +0000 UTC}
}
