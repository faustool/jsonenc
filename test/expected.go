package test

type ActualJson struct {
	StringField string `json:"stringField"`
	IntField    int `json:"intField"`
	StringArray []string `json:"stringArray"`
	IntArray    []int `json:"intArray"`
	ObjectArray []Object`json:"objectArray"`
}

type Object struct {
	Field string `json:"field"`
}