package test

import (
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestCueOutItem_Parse(t *testing.T) {
	// #EXT-X-CUE-OUT:<DURATION>
	line := `#EXT-X-CUE-OUT:177.100,SignalID=
`
	cue, err := m3u8.NewCueOutItem(line)
	assert.Nil(t, err)
	assert.Equal(t, 177.100, cue.Duration)
	// #EXT-X-CUE-OUT:DURATION=<time>
	line = `#EXT-X-CUE-OUT:DURATION=177.100
`
	cue, err = m3u8.NewCueOutItem(line)
	assert.Nil(t, err)
	assert.Equal(t, 177.100, cue.Duration)
}

func TestCueOutItem_String(t *testing.T) {
	item := &m3u8.CueOutItem{
		Duration: 177.100,
	}
	assert.Equal(t, item.String(), "#EXT-X-CUE-OUT:177.100")
}
