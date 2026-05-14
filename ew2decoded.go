package ew2decoded

import (
	"mime"
)

func DecoderNew(decoder *mime.WordDecoder) func(string) (string, error) {
	return func(original string) (string, error) {
		return decoder.Decode(original)
	}
}

func DecodeNewDefault() func(string) (string, error) {
	var dec mime.WordDecoder
	return DecoderNew(&dec)
}
