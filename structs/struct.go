package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Contact ContactInfo
}

type ContactInfo struct {
	Phone string
	Mail  string
}

func (p Person) Print() {
	fmt.Println(p.Name, p.Age)
	fmt.Println(p.Contact.Phone, p.Contact.Mail)
}

func main() {
	var p1 Person
	p1 = Person{"John", 72, ContactInfo{"0102030405", "test@mail.com"}}

	p2 := Person{"Harry", 21, ContactInfo{"1111111", "test@email.com"}}

	p3 := Person{Name: "Will", //recommmended  kind of Static Initialisation
		Age: 32,
		Contact: ContactInfo{
			Phone: "12131415",
			Mail:  "test@mail.com",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2.Name, p2.Age, p2.Contact.Mail, p2.Contact.Phone)
	fmt.Printf("%+v\n", p3)

	p3.Print()

}
