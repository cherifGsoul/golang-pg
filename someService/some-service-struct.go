package someService

type SomeServiceStruct struct {
	Greetings string
}

func NewSomeServiceStruct() SomeServiceStruct {
	return SomeServiceStruct{
		Greetings: "Hello, World!",
	}
}
