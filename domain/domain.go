package domain

type Page struct {
	Title string
	Body  []byte
}

type WikiStore interface {
	Save(*Page) error
	LoadPage(string) (*Page, error)
}
