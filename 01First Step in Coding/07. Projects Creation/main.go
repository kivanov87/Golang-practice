package main

import "fmt"

func main() {
	var architect string
	var projects int

	fmt.Scan(&architect)
	fmt.Scan(&projects)
	work := projects * 3
	fmt.Printf("The architect %v will need %v hours to complete %v project/s.", architect, work, projects)
}
