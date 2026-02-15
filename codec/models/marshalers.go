package models

import (
	"encoding/json"
)

// MarshalJSON implements custom JSON marshaling for AvlDataPacketTCP
func (a *AvlDataPacketTCP) MarshalJSON() ([]byte, error) {
	type Alias AvlDataPacketTCP
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	})
}

// MarshalIndent marshals AvlDataPacketTCP to indented JSON format
func (a *AvlDataPacketTCP) MarshalIndent(prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(a, prefix, indent)
}

// MarshalJSON implements custom JSON marshaling for AvlDataPacketUDP
func (a *AvlDataPacketUDP) MarshalJSON() ([]byte, error) {
	type Alias AvlDataPacketUDP
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	})
}

// MarshalJSON implements custom JSON marshaling for AvlDataPacketFlat
func (a *AvlDataPacketFlat) MarshalJSON() ([]byte, error) {
	type Alias AvlDataPacketFlat
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	})
}

// MarshalIndent marshals AvlDataPacketFlat to indented JSON format
func (a *AvlDataPacketFlat) MarshalIndent(prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(a, prefix, indent)
}

// MarshalIndent marshals AvlDataPacketUDP to indented JSON format
func (a *AvlDataPacketUDP) MarshalIndent(prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(a, prefix, indent)
}

func (r *ResponseMessage) MarshalJSON() ([]byte, error) {
	type Alias ResponseMessage

	base := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	baseMap := make(map[string]any)

	b, err := json.Marshal(base)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &baseMap); err != nil {
		return nil, err
	}

	if r.CodecSpecificParam != nil {
		pb, err := json.Marshal(r.CodecSpecificParam)
		if err != nil {
			return nil, err
		}

		var payloadMap map[string]any
		if err := json.Unmarshal(pb, &payloadMap); err != nil {
			return nil, err
		}

		for k, v := range payloadMap {
			baseMap[k] = v
		}
	}

	return json.Marshal(baseMap)
}

// MarshalIndent marshals ResponseMessage to indented JSON format
func (r *ResponseMessage) MarshalIndent(prefix, indent string) ([]byte, error) {
	type Alias ResponseMessage

	base := struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	baseMap := make(map[string]any)

	b, err := json.Marshal(base)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &baseMap); err != nil {
		return nil, err
	}

	if r.CodecSpecificParam != nil {
		pb, err := json.Marshal(r.CodecSpecificParam)
		if err != nil {
			return nil, err
		}

		var payloadMap map[string]any
		if err := json.Unmarshal(pb, &payloadMap); err != nil {
			return nil, err
		}

		for k, v := range payloadMap {
			baseMap[k] = v
		}
	}

	return json.MarshalIndent(baseMap, prefix, indent)
}
