package test

import (
	"encoding/base64"
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestOATCLSScte35_Parse(t *testing.T) {
	line := `#EXT-OATCLS-SCTE35:/DAnAAAAAAAAAP/wBQb+AA27oAARAg9DVUVJAAAAAX+HCQA0AAE0xUZn
`
	sp, err := m3u8.NewOATCLSScte35(line)

	assert.Nil(t, err)
	assert.Equal(t, "/DAnAAAAAAAAAP/wBQb+AA27oAARAg9DVUVJAAAAAX+HCQA0AAE0xUZn", base64.StdEncoding.EncodeToString(sp.Data))
}
