package tests

import "github.com/revel/revel"

type LanguageTest struct {
	revel.TestSuite
}

func (t *LanguageTest) Before() {
	println("Set up")
}

func (t LanguageTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *LanguageTest) After() {
	println("Tear down")
}
