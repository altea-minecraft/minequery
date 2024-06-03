package minequery

import (
	"errors"
)

// ErrInvalidStatus wraps errors occurred during ping status deserialization.
// Some errors may be ignored if UseStrict is not set to true.
var ErrInvalidStatus = errors.New("invalid status")

// ErrPacketIO indicates there was an IO error when sending or receiving a packet.
var ErrPacketIO = errors.New("packet IO error")

// ErrNoSRV indicates there were no SRV records found for hostname.
var ErrNoSRV = errors.New("no SRV records found")
