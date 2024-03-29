package vw

type VirtualWindows struct {
	Id      string `bson:"_id" json:"id"`
	Name    string `bson:"name" json:"name"`
	OS      string `bson:"os" json:"os"`
	Port    int    `bson:"port" json:"port"`
	Running bool   `bson:"-"  json:"running"`
}
