package basics

// Find type of the variable using type switch
import "fmt"

func Type() {
	checkType("raed")
	checkType(1)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
