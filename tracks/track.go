// A whole bunch of data relating to tracks.
package tracks

import (
	"encoding/hex"
	"errors"
)

type Data struct {
	Elements           []Element
	Clearance          int
	ClearanceDirection ClearanceDirection
}

type ClearanceDirection int

const (
	CLEARANCE_ABOVE = iota
	// For suspended coasters
	CLEARANCE_BELOW = iota
)

type Degree int

const (
	DEGREE_FLAT    = 0
	DEGREE_25_UP   = 25
	DEGREE_60_UP   = 60
	DEGREE_25_DOWN = -25
	DEGREE_60_DOWN = -60
)

type DirectionDelta int

const (
	// All of these expressed in positives to help with degree calculations
	DIRECTION_STRAIGHT     DirectionDelta = 0
	DIRECTION_45_DEG_RIGHT                = 45
	DIRECTION_90_DEG_RIGHT                = 90
	DIRECTION_180_DEG                     = 180
	DIRECTION_90_DEG_LEFT                 = 270
	DIRECTION_45_DEG_LEFT                 = 315
)

type Element struct {
	// XXX, add color schemes

	Segment       *Segment
	ChainLift     bool
	InvertedTrack bool
	Station       bool
	StationNumber int

	// bits 3, 2, 1 and 0 are a magnitude value (0..15) for brake and booster
	// track segments. The value obtained from these four bits is multiplied
	// by 7.6 km/hr = 4.5 mph.
	//
	// We store the value in km/h
	BoostMagnitude float32

	// For RCT2 "Multi Dimensional Coaster", these four bits specify the amount
	// of rotation.
	Rotation int
}

func (te Element) String() string {
	if te.Segment.Type > 0x17 {
		return hex.EncodeToString([]byte{byte(te.Segment.Type)})
	}
	return ""
}

var EndOfRide = errors.New("End of ride")

// Parse the serialized data into a Element struct
// When the end of ride is encountered a EndOfRide error is returned
func parseElement(rawElement []byte) (te Element, e error) {
	if len(rawElement) != 2 {
		return Element{}, errors.New("invalid length for element input")
	}
	te.Segment = &Segment{
		Type: SegmentType(rawElement[0]),
	}
	if te.Segment.Type == ELEM_END_OF_RIDE {
		return Element{}, EndOfRide
	}
	q := int(rawElement[1])
	te.ChainLift = q>>7 == 1
	te.InvertedTrack = q>>6 == 1
	te.Station = q>>3 == 1
	te.StationNumber = q & 3

	te.BoostMagnitude = float32(q&15) * 7.6
	te.Rotation = (q&15)*45 - 180
	return
}

// Turn RCT encoded track data into a Data object.
func Unmarshal(buf []byte, d *Data) error {
	td := new(Data)
	for i := 0; i < len(buf); i += 2 {
		elem, err := parseElement(buf[i : i+2])
		if err != nil {
			if err == EndOfRide {
				break
			}
			return err
		}
		td.Elements = append(td.Elements, elem)
	}

	// XXX where is this data actually stored?
	td.Clearance = 2
	td.ClearanceDirection = CLEARANCE_ABOVE
	return nil
}
