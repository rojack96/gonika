package parsers

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/rojack96/gonika/codec/models"
)

func TestPreamble(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0x00, 0x00}
	got := bp.Preamble(data)
	want := models.Preamble(0)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDataFieldLength(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0x00, 0x9A}
	got := bp.DataFieldLength(data)
	want := models.DataFieldLength(154)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDataSize(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0x00, 0x9A}
	got := bp.DataSize(data)
	want := models.DataSize(154)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestResponseSize(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0x00, 0x9A}
	got := bp.ResponseSize(data)
	want := models.ResponseSize(154)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCrc16(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0xBF, 0x30}
	got := bp.Crc16(data)
	want := models.Crc16(48944)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTimestamp(t *testing.T) {
	bp := NewBaseParser()
	data := []byte{0x00, 0x00, 0x01, 0x6B, 0x40, 0xD8, 0xEA, 0x30}
	timestamp, bytesRead := bp.Timestamp(0, data)
	expectedTimestamp := time.UnixMilli(1560161086000).UTC()

	if timestamp != models.Timestamp(expectedTimestamp) {
		t.Errorf("Expected timestamp %v, got %v", expectedTimestamp, timestamp)
	}

	if bytesRead != 8 {
		t.Errorf("Expected bytes read to be 8, got %d", bytesRead)
	}
}

func TestGpsElement(t *testing.T) {
	// TODO
	bp := NewBaseParser()
	data := []byte{0x0F, 0x0D, 0xC3, 0x9B, 0x20, 0x95, 0x96, 0x4A, 0x00, 0xAC, 0x00, 0xF8, 0x0B, 0x00, 0x00} // Type byte for GPS element
	got, _ := bp.GpsElement(0, data)
	want := models.GpsElement{
		Latitude:   54.667425,
		Longitude:  25.2560283,
		Altitude:   172,
		Angle:      248,
		Satellites: 11,
		Speed:      0,
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
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
