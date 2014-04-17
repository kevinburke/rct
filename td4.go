// td4 ride format parser

package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func hasLoop(n int) bool {
	return n>>7 == 1
}

type ControlFlags struct {
	UseMaximumTime             bool
	UseMinimumTime             bool
	SyncWithAdjacentStation    bool
	LeaveIfAnotherTrainArrives bool
	WaitForLoad                bool
	Load                       LoadType
}

func parseControlFlags(n int) *ControlFlags {
	return &ControlFlags{
		UseMaximumTime:             n>>7 == 1,
		UseMinimumTime:             n>>6 == 1,
		SyncWithAdjacentStation:    n>>5 == 1,
		LeaveIfAnotherTrainArrives: n>>4 == 1,
		WaitForLoad:                n>>3 == 1,
		Load:                       LoadType(n & 7),
	}
}

type RideType int
type VehicleType int
type LoadType int
type OperatingMode int
type SegmentType int

const (
	WOODEN_RIDE    RideType = iota
	STAND_UP_STEEL          = iota
	SUSPENDED               = iota
)

const (
	STEEL              VehicleType = iota
	STEEL_BACKWARDS                = iota
	WOODEN_VEHICLE                 = iota
	INVERTED                       = iota
	SUSPENDED_SWINGING             = iota
	LADYBIRD                       = iota
	STAND_UP                       = iota
)

const (
	MODE_NORMAL                  OperatingMode = 0x00
	MODE_CONTINUOUS_CIRCUIT                    = 0x01
	MODE_REVERSE_INCLINE_SHUTTLE               = 0x02
	MODE_POWERED_LAUNCH                        = 0x03
	MODE_SHUTTLE                               = 0x04
	MODE_BOAT_HIRE                             = 0x05
	MODE_UPWARD_LAUNCH                         = 0x06
	MODE_ROTATING_LIFT                         = 0x07
	MODE_STATION_TO_STATION                    = 0x08
)

const (
	LOAD_ONE_QUARTER    = 0
	LOAD_HALF           = 1
	LOAD_THREE_QUARTERS = 2
	LOAD_FULL           = 3
	LOAD_ANY            = 4
)

const (
	IDX_RIDE_TYPE      = 0
	IDX_VEHICLE_TYPE   = 1
	IDX_HAS_LOOP       = 2
	IDX_OPERATING_MODE = 0x06
	IDX_CONTROL_FLAG   = 0x23
	IDX_NUM_TRAINS     = 0x24
	IDX_CARS_PER_TRAIN = 0x25
	IDX_MIN_WAIT_TIME  = 0x26
	IDX_MAX_WAIT_TIME  = 0x27
	IDX_TRACK_DATA     = 0xc4
)

type Ride struct {
	RideType      RideType
	VehicleType   VehicleType
	HasLoop       bool
	OperatingMode OperatingMode
	ControlFlags  *ControlFlags
	NumTrains     int
	CarsPerTrain  int
	MinWaitTime   int
	MaxWaitTime   int
	TrackData     TrackData
}

const (
	ELEM_FLAT                              SegmentType = 0
	ELEM_END_STATION                                   = 0x1
	ELEM_BEGIN_STATION                                 = 0x2
	ELEM_MIDDLE_STATION                                = 0x3
	ELEM_25_DEG_UP                                     = 0x4
	ELEM_60_DEG_UP                                     = 0x5
	ELEM_FLAT_TO_25_DEG_UP                             = 0x6
	ELEM_25_DEG_TO_60_DEG_UP                           = 0x7
	ELEM_60_DEG_TO_25_DEG_UP                           = 0x8
	ELEM_25_DEG_UP_TO_FLAT                             = 0x9
	ELEM_25_DEG_DOWN                                   = 0x0a
	ELEM_60_DEG_DOWN                                   = 0x0b
	ELEM_FLAT_TO_25_DOWN                               = 0x0c
	ELEM_25_DEG_TO_60_DEG_DOWN                         = 0x0d
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN                    = 0x0e
	ELEM_25_DEG_DOWN_TO_FLAT                           = 0x0f
	ELEM_LEFT_QUARTER_TURN_5_TILES                     = 0x10
	ELEM_RIGHT_QUARTER_TURN_5_TILES                    = 0x11
	ELEM_FLAT_TO_LEFT_BANK                             = 0x12
	ELEM_FLAT_TO_RIGHT_BANK                            = 0x13
	ELEM_LEFT_BANK_TO_FLAT                             = 0x14
	ELEM_RIGHT_BANK_TO_FLAT                            = 0x15
	ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES              = 0x16
	ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES             = 0x16
	ELEM_BRAKES                                        = 0x63
	ELEM_BOOSTER                                       = 0x64
	ELEM_WATERFALL                                     = 0x70
	ELEM_ON_RIDE_PHOTO                                 = 0x72
	ELEM_END_OF_RIDE                                   = 0xff
)

type TrackElement struct {
	// XXX, add color schemes

	Segment       SegmentType
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

const (
	STATION_ONE   = 0
	STATION_TWO   = 1
	STATION_THREE = 2
	STATION_FOUR  = 3
)

type TrackData struct {
	Elements []TrackElement
}

// Take a compressed byte stream representing a ride and turn it into a Ride
// struct.
func Deserialize(buf io.Reader) *Ride {
	z, err := NewReader(buf)
	if err != nil {
		panic(err)
	}
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	out := bitbuffer.Bytes()
	for i := 0; i < 500; i++ {
		ds := hex.EncodeToString([]byte{byte(i)})
		fmt.Printf("%s: %s\n", ds, hex.EncodeToString([]byte{out[i]}))
	}
	r := Ride{
		RideType:      RideType(out[IDX_RIDE_TYPE]),
		VehicleType:   VehicleType(out[IDX_VEHICLE_TYPE]),
		HasLoop:       hasLoop(int(out[IDX_HAS_LOOP])),
		OperatingMode: OperatingMode(out[IDX_OPERATING_MODE]),
		ControlFlags:  parseControlFlags(int(out[IDX_CONTROL_FLAG])),
		NumTrains:     int(out[IDX_NUM_TRAINS]),
		CarsPerTrain:  int(out[IDX_CARS_PER_TRAIN]),
		MinWaitTime:   int(out[IDX_MIN_WAIT_TIME]),
		MaxWaitTime:   int(out[IDX_MAX_WAIT_TIME]),
		TrackData:     parseTrackData(out[IDX_TRACK_DATA:]),
	}
	return &r
}

var EndOfRide = errors.New("End of ride")

// Parse the serialized data into a TrackElement struct
// When the end of ride is encountered a EndOfRide error is returned
func parseElement(rawElement []byte) (te TrackElement, e error) {
	if len(rawElement) != 2 {
		return TrackElement{}, errors.New("invalid length for element input")
	}
	te.Segment = SegmentType(rawElement[0])
	if te.Segment == ELEM_END_OF_RIDE {
		return TrackElement{}, EndOfRide
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

func parseTrackData(trackData []byte) TrackData {
	td := new(TrackData)
	for i := 0; i < len(trackData); i += 2 {
		elem, err := parseElement(trackData[i : i+2])
		if err != nil {
			if err == EndOfRide {
				break
			}
			panic(err)
		}
		td.Elements = append(td.Elements, elem)
	}
	return *td
}

func main() {
	fmt.Println(readRaptor())
}

func readRaptor() *Ride {
	f, err := os.Open("rides/raptor.td4")
	if err != nil {
		panic(err)
	}
	return Deserialize(f)
}
