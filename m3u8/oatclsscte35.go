package m3u8

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type OATCLSScte35 struct {
	Data []byte
}

// Ad marker tag with a SCTE-35 payload in base64-encoded binary. The decoded binary must provide a SCTE-35
// splice_info_section containing the cue-out marker 0x34, for provider placement opportunity start, and the cue-in
// marker 0x35, for provider placement opportunity end.
func NewOATCLSScte35(text string) (*OATCLSScte35, error) {
	line := strings.Replace(text, OATCLSScte35Tag+":", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	data, err := base64.StdEncoding.DecodeString(line)
	if err != nil {
		return nil, err
	}
	return &OATCLSScte35{Data: data}, nil
}

func (sp *OATCLSScte35) String() string {
	return fmt.Sprintf("%s:%s", OATCLSScte35Tag, base64.StdEncoding.EncodeToString(sp.Data))
}
