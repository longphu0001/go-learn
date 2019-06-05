package main

// User is user
type User struct {
	ID     int
	Name   string
	Active bool
}

// Data is list user
type Data struct {
	User  []User
	Count int
}

func dataTest() map[string]interface{} {

	data := Data{
		User: []User{
			User{ID: 1, Name: "User 1", Active: true},
			User{ID: 2, Name: "User 2", Active: false},
		},
		Count: 1,
	}

	dataMapArray := make(map[string]interface{})
	dataMapArray["text"] = "ascsaca"
	dataMapArray["Num"] = 123456
	dataMapArray["Flag"] = true
	dataMapArray["data"] = data

	return dataMapArray
}
