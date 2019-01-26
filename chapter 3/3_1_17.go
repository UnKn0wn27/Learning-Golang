package main

import (
	"fmt"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Dob					time.Time
	Email, Location		string
}

//A person method
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location)
}

//A person method with pointer receiver
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

type Admin struct {
	Person //type embedding for composition
	Roles []string
}

//overrides PrintDetails
func (a Admin) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type Member struct {
	Person //type embedding for composition
	Skills []string
}

//overrides PrintDetails
func (m Member) PrintDetails(){
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

//Combine with interface
type Team struct {
	Name, Description string
	Users 			  []User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

func main() {
	var p Person
	p.FirstName="Otto"
	p.LastName="Mount"
	p.Dob=time.Date(1996, time.April, 27, 0,0,0,0, time.UTC)
	p.Email="unkn0wn27md@gmail.com"
	p.Location="Kochi"

	fmt.Println("Method 1")
	fmt.Println(p)

	pe := Person{
		FirstName: "Unknown",
		LastName: "Void",
		Dob: time.Date(1996, time.April, 27, 0,0,0,0, time.UTC),
		Email: "unkn0wn27md@gmail.com",
		Location: "Kochi",
	}

	fmt.Println("Method 2")
	fmt.Println(pe)

	per := Person{
		"Unknown",
		"Void",
		time.Date(1996, time.April, 27, 0,0,0,0, time.UTC),
		"unkn0wn27md@gmail.com",
		"Kochi",
	}

	fmt.Println("Method 3")
	fmt.Println(per)

	pers := Person{FirstName: "Unknown", LastName: "Void"}
	pers.Dob=time.Date(1996, time.April, 27, 0,0,0,0, time.UTC)
	pers.Email="unkn0wn27md@gmail.com"
	pers.Location="Kochi"

	fmt.Println("Method 4")
	fmt.Println(pers)

	fmt.Println("Person Methods!")
	pers.PrintName()
	pers.PrintDetails()


	perso := &Person{
		"Unknown",
		"Void",
		time.Date(1996, time.April, 27, 0,0,0,0, time.UTC),
		"unkn0wn27md@gmail.com",
		"Kochi",
	}

	fmt.Println("Pointer Receiver!")
	perso.ChangeLocation("Moldova")
	perso.PrintName()
	perso.PrintDetails()

	fmt.Println("Embeding Type Struct!")

	alex := Admin{
		Person{"Alex",
		"Void",
		time.Date(1994, time.February, 7, 0,0,0,0, time.UTC),
		"unkn0wn27md@gmail.com",
		"Kochi"},
		[]string{"Manage Team", "Manage Tasks"},
	}

	octa := Member{
		Person{"Octavian",
		"Muntean",
		time.Date(1996, time.April, 27, 0,0,0,0, time.UTC),
		"unkn0wn27md@gmail.com",
		"Chisinau"},
		[]string{"Python", "Golang", "Django"},
	}

	//call method for alex
	alex.PrintName()
	alex.PrintDetails()
	//call method for octa
	octa.PrintName()
	octa.PrintDetails()

	fmt.Println("Overriding Methods of Embedded type")

	fmt.Println("Using Interface!")
	users := []User{alex, octa}

	for _, v := range users {
		v.PrintName()
		v.PrintDetails()
	}

	fmt.Println("Team!")
	team := Team{
		"Go",
		"Golang and Python",
		[]User{pers, octa},
	}
	//get details of Team
	team.GetTeamDetails()
}