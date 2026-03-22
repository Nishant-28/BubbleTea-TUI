package store

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store interface {
	GetNotes() ([]Note, error)
	SaveNote(note Note) error
}
