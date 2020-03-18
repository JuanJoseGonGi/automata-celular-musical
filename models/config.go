package models

// Config is a struct to handle configuration
type Config struct {
	Speed       float64  `json:"speed"`
	Instruments []string `json:"instruments"`
	NotesAmount int      `json:"notes_amount"`
}
