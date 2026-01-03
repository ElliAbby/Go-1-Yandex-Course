package main

import (
	"fmt"
	"time"
)

type User struct {
	ID int
	Name string
	Email string
	Age int
}

type Report struct {
	User
	ReportID int
	Date string
}

func CreateReport(user User, reportDate string) Report {
	reportID := int(time.Now().Unix())
	return Report{user, reportID, reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	result := make([]Report, len(users))
	for i, user := range users {
		result[i] = CreateReport(user, reportDate)
	}
	return result
}

func main() {
	user1 := User{1, "test1", "email1", 34}
	fmt.Println(CreateReport(user1, "31-03-2006"))

	user2 := User{2, "test2", "email2", 35}
	user3 := User{3, "test3", "email3", 36}
	user4 := User{4, "test4", "email4", 37}
	user5 := User{5, "test5", "email5", 38}

	users := []User{user1, user2, user3, user4, user5}
	fmt.Println(GenerateUserReports(users, "31-02-2023"))
}
