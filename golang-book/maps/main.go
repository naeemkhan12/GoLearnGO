package main

import (
	"fmt"
)

func main() {
	customMap := Map{
		"lang":       {"en", "fr", "urd"},
		"department": {"Fashion", "Tech"},
	}
	fmt.Println(customMap["lang"])

}

type Map map[string][]string
