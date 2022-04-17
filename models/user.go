package models

type User struct {
	Name     string `json:"name" bson:"user_name"`
	Age      int    `json:"age" bson:"user_age"`
	Birthday Birthday
	Address  Address
}

type Birthday struct {
	Day   int    `json:"day" bson:"day"`
	Month string `json:"month" bson:"month"`
	Year  int    `json:"year" bson:"year"`
}

type Address struct {
	Province string `json:"province" bson:"province"`
	County   string `json:"county" bson:"county"`
	Alley    string `json:"alley" bson:"alley"`
}
