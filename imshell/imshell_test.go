package imshell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRGBToHex(t *testing.T) {
	hex := string(RGBToHex(255, 0, 0))
	assert.Equal(t, hex, "ff0000", "Hex to RGB")
}

func TestRGBToShort(t *testing.T) {
	s := RGBToShort(255, 0, 0)
	assert.Equal(t, s, "09", "RGB to Short")
}

func TestGetText(t *testing.T) {
	imtext := NewImageText("ABC")
	assert.Equal(t, imtext.GetText(), "AB")
	assert.Equal(t, imtext.GetText(), "CA")
	assert.Equal(t, imtext.GetText(), "BC")
	assert.Equal(t, imtext.GetText(), "AB")

	imtext = NewImageText("あいう")
	imtext.index = 0
	assert.Equal(t, imtext.GetText(), "あ")
	assert.Equal(t, imtext.GetText(), "い")
	assert.Equal(t, imtext.GetText(), "う")
	assert.Equal(t, imtext.GetText(), "あ")

	imtext = NewImageText("あAいBCうえA")
	imtext.index = 0
	assert.Equal(t, imtext.GetText(), "あ")
	assert.Equal(t, imtext.GetText(), "AA")
	assert.Equal(t, imtext.GetText(), "い")
	assert.Equal(t, imtext.GetText(), "BC")
	assert.Equal(t, imtext.GetText(), "う")
	assert.Equal(t, imtext.GetText(), "え")
	assert.Equal(t, imtext.GetText(), "AA")
	assert.Equal(t, imtext.GetText(), "あ")
}
