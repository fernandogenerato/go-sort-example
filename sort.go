package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	now := time.Now()

	u := []User{
		{
			now.Add(time.Hour * 24 * 5),
		}, {
			now,
		}, {
			now.Add(time.Hour * 24 * 1),
		}, {
			now.Add(time.Hour * 24 * 3),
		},
		{
			now.Add(time.Hour * 24 * 12),
		},
	}
	fmt.Println("\nusers.date")
	for _, t := range u {
		fmt.Println(t.date.Format("2006-01-02"))
	}


	fmt.Println("\nsort by users.date")
	sort.Sort(Users(u))
	for _, t := range Users(u) {
		fmt.Println(t.date.Format("2006-01-02"))
	}

	fmt.Println("\nreverse by users.date")
	sort.Sort(sort.Reverse(Users(u)))
	for _, t := range Users(u) {
		fmt.Println(t.date.Format("2006-01-02"))
	}

}

type Users []User

type User struct {
	date time.Time
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].date.Before(u[j].date)
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

