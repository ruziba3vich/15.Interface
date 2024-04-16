package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	m "github.com/ruziba3vich/files/models/user"
)

func main() {
	fileName := "user_data/sample.txt"
	f, err := os.Open(fileName)
	if err != nil {
		ErrorPrinter(err)
	} else {
		defer f.Close()
		var users []m.User
		scanner := bufio.NewScanner(f)

		if err := scanner.Err(); err != nil {
			ErrorPrinter(err)
		} else {
			var rawUsers []string
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) != 0 {
					rawUsers = append(rawUsers, line)
				}
				if len(rawUsers) == 3 {
					users = append(users, GetUser(rawUsers))
					rawUsers = [] string {}
				}
			}
		}
		for _, user := range users {
			fmt.Println("------------ user -------------")
			fmt.Println("\n{")
			RepresentUser(user)
			fmt.Println("}")
		}
	}
}

func RepresentUser(user m.User) {
	fmt.Printf("	\"name\":\"%s\",\n", user.Name)
	fmt.Printf("	\"age\":%d,\n", user.Age)
	fmt.Printf("	\"occupation\":\"%s\"\n", user.Occupation)
}

func ErrorPrinter(err error) {
	fmt.Println("Error :", err)
}

func GetUser(userData []string) (user m.User) {
	for _, line := range userData {
		data := strings.Split(line, ":")
		data[1] = strings.Trim(data[1], " ")
		if data[0] == "Name" {
			user.Name = data[1]
		} else if data[0] == "Age" {
			age, _ := strconv.Atoi(data[1])
			user.Age = int16(age)
		} else {
			user.Occupation = data[1]
		}
	}
	return user
}
