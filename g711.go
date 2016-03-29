// Package g711 implements the ITU-T G.711 standard for audio companding.
package g711

import "errors"

var (
	errNilBuffer = errors.New("g711: nil buffer")
	errNilReader = errors.New("g711: nil reader")
)
