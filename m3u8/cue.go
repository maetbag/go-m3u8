package m3u8

import (
	"fmt"
	"strings"
)

// CueOutItem represents a #EXT-X-CUE-OUT tag
type CueOutItem struct {
	Duration float64
}

func NewCueOutItem(text string) (*CueOutItem, error) {
	attributes := ParseAttributes(text)
	duration, err := parseFloat(attributes, DurationTag)
	if err != nil {
		return nil, err
	}
	return &CueOutItem{
		Duration: *duration,
	}, nil
}

func (coi *CueOutItem) String() string {
	var slice []string
	slice = append(slice, fmt.Sprintf(formatString, DurationTag, coi.Duration))
	return fmt.Sprintf("%s:%s", CueOutTag, strings.Join(slice, ","))
}

// CueInItem represents a #EXT-X-CUE-IN tag
type CueInItem struct{}

func NewCueInItem(text string) (*CueInItem, error) {
	return &CueInItem{}, nil
}

func (coi *CueInItem) String() string {
	return fmt.Sprintf("%s", CueInTag)
}
