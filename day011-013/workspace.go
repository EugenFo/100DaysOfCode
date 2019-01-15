package main

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

func main() {
	newUUID, err := uuid.NewV4()
	fmt.Println("The new UUID is:", newUUID)
	if err != nil {
		fmt.Println(err)
	}
}
