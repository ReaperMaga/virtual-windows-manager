package vw

type VirtualWindows struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	OS   string `bson:"os"`
}
