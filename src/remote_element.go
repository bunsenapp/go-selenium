package goselenium

func newSeleniumElement(i string, w *seleniumWebDriver) *seleniumElement {
	return &seleniumElement{
		id: i,
		wd: w,
	}
}

type seleniumElement struct {
	id string
	wd *seleniumWebDriver
}

func (s *seleniumElement) ID() string {
	return s.id
}
