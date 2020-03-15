package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var (
	// ErrStateTooBig occurs when an instument does not contain note state
	ErrStateTooBig = errors.New("the state is too big")
)

const (
	sampleRate = 44100
)

// Note is a struct to represent a note
type Note struct {
	Buffer *beep.Buffer
}

// Instrument is a struct to represent an instrument
type Instrument struct {
	Name  string
	Notes []*Note
}

// NewInstrument creates a new instrument
func NewInstrument(name string) (Instrument, error) {
	notes, err := loadNotes(name)
	if err != nil {
		return Instrument{}, err
	}

	return Instrument{
		Name:  name,
		Notes: notes,
	}, nil
}

func loadNotes(instrumentName string) ([]*Note, error) {
	notes := []*Note{nil}

	for i := 1; i < 5; i++ {
		file, err := os.Open(fmt.Sprintf("sounds/%s/%d.wav", instrumentName, i))
		if err != nil {
			return nil, err
		}

		streamer, format, err := wav.Decode(file)
		if err != nil {
			return nil, err
		}

		resampledStreamer := beep.Resample(4, format.SampleRate, sampleRate, streamer)

		buffer := beep.NewBuffer(format)
		buffer.Append(resampledStreamer)
		err = streamer.Close()
		if err != nil {
			return nil, err
		}

		notes = append(notes, &Note{
			Buffer: buffer,
		})
	}

	return notes, nil
}

func (instrument *Instrument) playNote(state int) error {
	if state == 0 {
		return nil
	}

	if state >= len(instrument.Notes) {
		return fmt.Errorf("playing note with state %d: %w", state, ErrStateTooBig)
	}

	note := instrument.Notes[state]
	noteSeek := note.Buffer.Streamer(0, note.Buffer.Len())
	speaker.Play(noteSeek)

	return nil
}
