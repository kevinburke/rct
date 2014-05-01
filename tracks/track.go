// A whole bunch of data relating to tracks.

// This menu table page is information about the RCT1 saved game files. Those
// addresses are for .SV4 files. In RCT2, the ride types are in a different
// order/arrangement. Some rides and shops use the same type as this is how
// custom objects work.

// I have attached the IDA database I have made so far.

// Current_possible_ride_configurations appears to be 32 bytes which map to
// ride_configuration_string_ids, those string ids map to the english text file
// I have sent you. current_possible_ride_configurations is the order of track
// configurations shown in the dropdown for the ride construction window when
// currently constructing a ride. In order to find the tables used to see what
// configurations are possible for a particular ride type, address 0x006C8681
// is the code which sets the current_possible_ride_configurations array using
// these other tables: track_config_table_1 and track_config_table_2. This is
// all I have been able to figure out so far.
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
	te.Segment = TS_MAP[SegmentType(rawElement[0])]
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
	for i := 0; i < len(buf); i += 2 {
		elem, err := parseElement(buf[i : i+2])
		if err != nil {
			if err == EndOfRide {
				break
			}
			return err
		}
		d.Elements = append(d.Elements, elem)
	}

	// XXX where is this data actually stored?
	d.Clearance = 2
	d.ClearanceDirection = CLEARANCE_ABOVE
	return nil
}
