package viewmodel

//StandLocator is a test comment
type StandLocator struct {
	Title  string
	Active string
}

//NewStandLocator is a test comment
func NewStandLocator() StandLocator {
	return StandLocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "standlocator",
	}
}
