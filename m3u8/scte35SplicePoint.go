package m3u8

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Scte35SplicePoint struct {
	Data string
}

// Ad marker tag with a SCTE-35 payload in base64-encoded binary. The decoded binary must provide a SCTE-35
// splice_info_section containing the cue-out marker 0x34, for provider placement opportunity start, and the cue-in
// marker 0x35, for provider placement opportunity end.
func NewScte35SplicePoint(text string) (*Scte35SplicePoint, error) {
	line := strings.Replace(text, Scte35SplicePointTag+":", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	if _, err := base64.StdEncoding.DecodeString(line); err != nil {
		return nil, err
	}
	return &Scte35SplicePoint{Data: line}, nil
}

func (sp *Scte35SplicePoint) String() string {
	return fmt.Sprintf("%s:%s", Scte35SplicePointTag, sp.Data)
}
