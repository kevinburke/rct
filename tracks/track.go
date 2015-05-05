// A whole bunch of data relating to tracks.

// This menu table page is information about the RCT1 saved game files. Those
// addresses are for .SV4 files. In RCT2, the ride types are in a different
// order/arrangement. Some rides and shops use the same type as this is how
// custom objects work.

package tracks

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/kevinburke/rct/bits"
)

// An Element containts a single track segment and some metadata about the
// track segment - whether it is a station piece, or has a chain lift.
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

type Data struct {
	Elements           []Element
	Clearance          int
	ClearanceDirection ClearanceDirection
}

// The amount of space each track piece needs above or below. XXX should this
// go on the ride object?
type ClearanceDirection int

const (
	CLEARANCE_ABOVE = iota
	// For suspended coasters
	CLEARANCE_BELOW = iota
)

func (te Element) String() string {
	return hex.EncodeToString([]byte{byte(te.Segment.Type)})
}

var EndOfRide = errors.New("End of ride")

// Parse the serialized data into a Element struct
// When the end of ride is encountered a EndOfRide error is returned
// Documentation from
// http://freerct.github.io/RCTTechDepot-Archive/trackQualifier.html
func unmarshalElement(rawElement []byte) (te Element, e error) {
	if len(rawElement) != 2 {
		return Element{}, errors.New("invalid length for element input")
	}
	// XXX hack for brakes
	if rawElement[0] == 0xd8 {
		rawElement[0] = ELEM_BRAKES
	}
	te.Segment = TS_MAP[SegmentType(rawElement[0])]
	if te.Segment.Type == ELEM_END_OF_RIDE {
		return Element{}, EndOfRide
	}
	q := int(rawElement[1])
	te.ChainLift = bits.On(q, 7)
	te.InvertedTrack = bits.On(q, 6)
	te.Station = bits.On(q, 3)
	if te.Station {
		fmt.Println("found station piece")
		fmt.Println(te.Segment)
	}
	te.StationNumber = q & 3

	te.BoostMagnitude = float32(q&15) * 7.6
	te.Rotation = (q&15)*45 - 180
	return
}

// Turn track elements back into bytes (opposite of above function)
//
// Taken from http://freerct.github.io/RCTTechDepot-Archive/trackQualifier.html
func marshalElement(e Element) ([]byte, error) {
	buf := make([]byte, 2)
	buf[0] = byte(e.Segment.Type)
	featureBit := 0
	featureBit = bits.SetCond(featureBit, 7, e.ChainLift)
	featureBit = bits.SetCond(featureBit, 6, e.InvertedTrack)
	featureBit = bits.SetCond(featureBit, 3, e.Station)
	if e.Segment.Type == ELEM_END_STATION || e.Segment.Type == ELEM_BEGIN_STATION || e.Segment.Type == ELEM_MIDDLE_STATION {
		// Set the station bit
		featureBit |= e.StationNumber
	}

	// XXX booster?
	if e.Segment.Type == ELEM_BRAKES || e.Segment.Type == ELEM_BLOCK_BRAKES {
		bm := e.BoostMagnitude / 7.6
		featureBit |= int(bm)
	} else {
		// 2nd bit seems to be set in most cases.
		featureBit |= 1 << 2
	}
	// XXX, rotation for multi dimensional coasters
	buf[1] = byte(featureBit)
	return buf, nil
}

// Turn RCT encoded track data into a Data object.
func Unmarshal(buf []byte, d *Data) error {
	for i := 0; i < len(buf); i += 2 {
		elem, err := unmarshalElement(buf[i : i+2])
		if err != nil {
			if err == EndOfRide {
				break
			}
			return err
		}
		fmt.Println(elem.Segment)
		d.Elements = append(d.Elements, elem)
	}

	// XXX where is this data actually stored?
	d.Clearance = 2
	d.ClearanceDirection = CLEARANCE_ABOVE
	return nil
}

// Turn track data into a series of bytes
func Marshal(d *Data) ([]byte, error) {
	buf := make([]byte, 2*len(d.Elements)+1)
	for i := range d.Elements {
		bts, err := marshalElement(d.Elements[i])
		if err != nil {
			return []byte{}, err
		}
		copy(buf[i*2:i*2+2], bts)
	}
	buf[len(buf)-1] = 0xff
	return buf, nil
}
