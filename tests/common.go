package tests

type TestData struct {
	Name          string
	Requester     func(s string) ([]byte, error)
	ExpectedError error
}

type JsonValidationTest struct {
	Name     string
	Expected any
	Actual   any
}
