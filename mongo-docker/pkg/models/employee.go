package models

// bson for mongoDB unsderstand, json for JSON understand
// omitempty require not null field
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name" bson:"name"`
	Salary float64 `json:"salary" bson:"salary"`
	Age    float64 `json:"age" bson:"age"`
}
