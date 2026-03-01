package models

import "time"

type Clients struct {
	Id             int32     `pg:"id,pk"`
	ClientName     string    `pg:"company_name"`
	PrimaryEmail   string    `pg:"primary_email"`
	SecondaryEmail string    `pg:"secondary_email"`
	Address        string    `pg:"address"`
	State          string    `pg:"state"`
	Country        string    `pg:"country"`
	CreatedAt      time.Time `pg:"created_at,default:now()"`
}
type User struct {
	Id        int32     `pg:"id,pk"`
	Name      string    `pg:"name"`
	Email     string    `pg:"email,unique"`
	Age       int       `pg:"age"`
	Picture   string    `pg:"picture"`
	Iat       int32     `pg:"iat"`
	Exp       int32     `pg:"exp"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
}
type UserOtherDetails struct {
	Id int32
}
