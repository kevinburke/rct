// Dealing with entrances/exits from a ride
package td6

import (
	"bytes"
	"encoding/binary"

	"github.com/kevinburke/rct-rides/bits"
)

type Egress struct {
	Exit      bool
	Direction int
	XOffset   int16
	YOffset   int16
}

// Get list of entrances/exits for a ride from raw data
// Copied mostly from "Scenery Items" here:
// http://freerct.github.io/RCTTechDepot-Archive/TD6.html
func unmarshalEgress(buf []byte) []*Egress {
	var egrs []*Egress
	for i := 0; true; i += 6 {
		if buf[i] == 0xff {
			break
		}
		features := int(buf[i+1])
		egr := &Egress{
			Exit: bits.On(features, 7),
			// Direction set in lower two bits
			Direction: features & 3,
			XOffset:   int16(binary.LittleEndian.Uint16(buf[i+2 : i+4])),
			YOffset:   int16(binary.LittleEndian.Uint16(buf[i+4 : i+6])),
		}
		egrs = append(egrs, egr)
	}
	return egrs
}

func marshalEgress(egr Egress) ([]byte, error) {
	buf := make([]byte, 6)
	ftrBit := egr.Direction
	ftrBit = bits.SetCond(ftrBit, 7, egr.Exit)
	buf[1] = byte(ftrBit)

	b := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, egr.XOffset)
	if err != nil {
		return []byte{}, err
	}
	copy(buf[2:4], b.Bytes())
	b.Reset()
	err = binary.Write(b, binary.LittleEndian, egr.YOffset)
	if err != nil {
		return []byte{}, err
	}
	copy(buf[4:6], b.Bytes())
	return buf, nil
}

func marshalEgresses(egrs []*Egress) ([]byte, error) {
	buf := make([]byte, 6*len(egrs)+1)
	for i, egr := range egrs {
		bts, err := marshalEgress(*egr)
		if err != nil {
			return []byte{}, err
		}
		copy(buf[i*6:i*6+6], bts)
	}
	buf[len(buf)-1] = 0xff
	return buf, nil
}
