package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	//Take naem as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name")
	name, _ = reader.ReadString('\n')

	//Take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa Center between 1 and 5 :")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back End
	//.Printf("%v %v", name, mynumRating)
	fmt.Printf("Hello %v, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v ", name, mynumRating, time.Now().Format(time.Kitchen))

	if mynumRating == 5 {
		fmt.Println("Bonus for team for 5 star service")

	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("We are always improving")
	} else if mynumRating < 3 {
		fmt.Println("Need Serious work on our side ")
	}

}
