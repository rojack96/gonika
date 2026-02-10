package parsers

import (
	"encoding/binary"
	"testing"

	"github.com/rojack96/gonika/codec/models"
)

func TestTimestamp(t *testing.T) {
	// TODO
	data := []byte{0x00, 0x00, 0x01, 0x6B, 0x40, 0xD8, 0xEA, 0x30}
	timestamp, bytesRead := Timestamp(0, data)
	expectedTimestamp := 1560161086000

	if timestamp != models.Timestamp(expectedTimestamp) {
		t.Errorf("Expected timestamp %d, got %d", expectedTimestamp, timestamp)
	}

	if bytesRead != 8 {
		t.Errorf("Expected bytes read to be 8, got %d", bytesRead)
	}
}

func TestDecodeCoordinate_Positive(t *testing.T) {
	pos := []byte{0x10, 0xC6, 0xB3, 0x7A} // 281,000,826
	got := decodeCoordinate(pos)
	want := 28.145753

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDecodeCoordinate_Negative(t *testing.T) {
	raw := int32(-124964000)

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(raw))

	got := decodeCoordinate(buf)
	want := -12.4964

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
