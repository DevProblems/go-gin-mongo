package models

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Address Address `json:"address" bson:"address"`
}
