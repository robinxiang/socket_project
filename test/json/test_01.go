package main

import (
	"encoding/json"
	"fmt"
)

// test the json marshal and unmarshal

// define the struck
// the first letter must be uppercase later in struct
type employee struct {
	Name     string
	Age      int
	Partment string
	Gender   string
}

func main() {
	// create the instance
	var (
		emp_01, emp_02 employee
	)

	emp_01 = employee{
		Name:     "zhangsan",
		Age:      42,
		Partment: "big data",
		Gender:   "male",
	}

	// create the json string to Unmarshal
	str_json_emp := `{"Name":"lisi","Age":42,"Partment":"cyberspaceSecurity","Gender":"male"}`

	fmt.Println("the struct is:", emp_01)

	// use Marshal to encode the struct to json string
	// the first letter must be uppercase later in struct
	// the json marshal result is :
	// {"Name":"zhangsan","Age":42,"Partment":"big data","Gender":"male"}
	str_json_encode, err := json.Marshal(&emp_01)
	if err != nil {
		fmt.Println("the marshal is error:", err)
	}
	// print the marshal result
	fmt.Printf("json marshal result is :\n%s\n", string(str_json_encode))

	// print the json string to Unmarshal
	fmt.Printf("the json string wait for unmarshal:%s\n", str_json_emp)
	// use Unmarshal to decode the json string to struct
	if err := json.Unmarshal([]byte(str_json_emp), &emp_02); err != nil {
		fmt.Println("the unmarshal is error:", err)
	}

	fmt.Printf("after the json unmarshal result is :%v\n", emp_02)

}
