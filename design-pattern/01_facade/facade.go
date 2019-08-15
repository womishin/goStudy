package facade

import "fmt"

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (b *bModuleImpl) TestB() string {
	return "B module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
