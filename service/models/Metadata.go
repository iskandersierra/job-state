package models

import "encoding/json"

var emptyMetadata = "{}"

func SerializeMetadata(metadata map[string]string) (string, error) {
	if metadata == nil {
		return emptyMetadata, nil
	}

	bytes, err := json.Marshal(metadata)
	if err != nil {
		return emptyMetadata, err
	}

	return string(bytes), nil
}

func SerializeOptionalMetadata(metadata map[string]string) (*string, error) {
	if len(metadata) == 0 {
		return nil, nil
	}

	text, err := SerializeMetadata(metadata)
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func DeserializeMetadata(text string) (map[string]string, error) {
	var metadata map[string]string
	err := json.Unmarshal([]byte(text), &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func DeserializeOptionalMetadata(text *string) (map[string]string, error) {
	if text == nil || *text == "" {
		return nil, nil
	}

	return DeserializeMetadata(*text)
}
