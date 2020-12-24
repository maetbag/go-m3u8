package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// CueOutItem represents a #EXT-X-CUE-OUT tag
type CueOutItem struct {
	Duration float64
}

func NewCueOutItem(text string) (item *CueOutItem, err error) {
	attributes := ParseAttributes(text)
	value := text[len(CueOutTag)+1:]
	parts := strings.Split(value, ",")
	item = &CueOutItem{}
	item.Duration, _ = strconv.ParseFloat(parts[0], 64)
	v, err := parseFloat(attributes, DurationTag)
	if err != nil {
		return nil, err
	} else if v != nil {
		item.Duration = *v
	}
	return item, nil
}

func (coi *CueOutItem) String() string {
	return fmt.Sprintf("%s:%f", CueOutTag, coi.Duration)
}

// CueInItem represents a #EXT-X-CUE-IN tag
type CueInItem struct{}

func NewCueInItem(text string) (*CueInItem, error) {
	return &CueInItem{}, nil
}

func (coi *CueInItem) String() string {
	return fmt.Sprintf("%s", CueInTag)
}
