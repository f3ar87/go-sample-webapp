package viewmodel

//Home is a test comment
type Home struct {
	Title  string
	Active string
}

//NewHome is a test comment
func NewHome() Home {
	return Home{
		Title:  "Lemonade Stand Supply - Home",
		Active: "home",
	}
}
