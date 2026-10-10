package main

import "fmt"

func greet(name string) string {

	return fmt.Sprintf("hello %s", name)
}

func main() {
	names := []string{"Hilary", "Lucky", "omondi"}
	for _, name := range names {
		fmt.Println(greet(name))

	}

}
