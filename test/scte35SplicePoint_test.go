package test

import (
	"encoding/base64"
	"testing"

	"github.com/quangngotan95/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestScte35SplicePointItem_Parse(t *testing.T) {
	line := `#EXT-X-SPLICEPOINT-SCTE35:/DA9AAAAAAAAAP/wBQb+uYbZqwAnAiVDVUVJAAAKqX//AAEjW4AMEU1EU05CMDAxMTMyMjE5M19ONAAAmXz5JA==
`
	sp, err := m3u8.NewScte35SplicePoint(line)

	assert.Nil(t, err)
	assert.Equal(t, "/DA9AAAAAAAAAP/wBQb+uYbZqwAnAiVDVUVJAAAKqX//AAEjW4AMEU1EU05CMDAxMTMyMjE5M19ONAAAmXz5JA==", base64.StdEncoding.EncodeToString(sp.Data))
}
