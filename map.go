package main

import (
	"fmt"
	"strconv"
)

type ADDRESS struct {
	area    string
	country string
}
type Slice1 struct {
	id      int
	name    string
	address ADDRESS
}
type TECHDETS struct {
	tech string
	exp  float64
}
type Slice2 struct {
	id       int
	techdets []TECHDETS
}
type ContactDETS struct {
	email string
	phone string
}
type Slice3 struct {
	id          int
	contactDets ContactDETS
}
type ALL struct {
	//id int
	name     string
	address  ADDRESS
	techdets []TECHDETS
	email    string
	phone    string
}

func main() {
	s1 := make([]Slice1, 2)
	s1[0] = Slice1{id: 1, name: "radha", address: ADDRESS{area: "bopal", country: "IND"}}
	s1[1] = Slice1{id: 2, name: "aniket", address: ADDRESS{area: "maninagar", country: "UK"}}

	s2 := make([]Slice2, 2)
	s2[0] = Slice2{id: 1, techdets: []TECHDETS{{tech: "react", exp: 1}, {tech: "go", exp: 1.5}}}
	s2[1] = Slice2{id: 2, techdets: []TECHDETS{{tech: "react", exp: 0.9}, {tech: "go", exp: 1.5}}}

	s3 := make([]Slice3, 2)
	s3[0] = Slice3{id: 1, contactDets: ContactDETS{email: "abc@gmail.com", phone: "123455778"}}
	s3[1] = Slice3{id: 2, contactDets: ContactDETS{email: "xyz@gmail.com", phone: "123455778"}}

	MAP := make(map[string]string)
	MAP["IND"] = "+91"
	MAP["UK"] = "+41"

	answer := make(map[string]ALL, 2)

	for i := range s1 {
		var temp ALL
		id := s1[i].id
		temp.name = s1[i].name
		temp.address = s1[i].address
		country := s1[i].address.country
		for j := 0; j < len(s2); j++ {
			if s2[j].id == id {
				temp.techdets = s2[j].techdets
			}
		}
		for j := 0; j < len(s3); j++ {
			if s3[j].id == id {
				temp.email = s3[j].contactDets.email
				temp.phone = s3[j].contactDets.phone
			}
		}
		for j := range MAP {
			if j == country {
				temp.phone = MAP[j] + temp.phone
			}
		}
		answer[strconv.Itoa(id)] = temp
	}

	fmt.Println(answer)
}
