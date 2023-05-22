package someService

func (s *SomeServiceStruct) Greet() string {
	return s.Greetings
}

func (s *SomeServiceStruct) GreetingSize() int {
	return len(s.Greetings)
}
