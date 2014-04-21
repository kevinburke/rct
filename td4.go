// td4 ride format parser

package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
)

func hasLoop(n int) bool {
	return n>>7 == 1
}

// Set the loop bit.
func marshalLoop(hasLoop bool, n int) {
	if hasLoop {
		n |= (1 << 7)
	} else {
		n &= ^(1 << 7)
	}
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

type StationType int

const (
	STATION_ONE   StationType = 0
	STATION_TWO               = 1
	STATION_THREE             = 2
	STATION_FOUR              = 3
)

// computes sines in degrees
func sindeg(deg DirectionDelta) int {
	for ; deg >= 360; deg -= 360 {
	}
	if deg%180 == 0 {
		return 0
	} else if deg == 90 {
		return 1
	} else if deg == 270 {
		return -1
	} else {
		return int(math.Sin(float64(deg) * math.Pi / 180))
	}
}

// computes sines in degrees
func cosdeg(deg DirectionDelta) int {
	for ; deg >= 360; deg -= 360 {
	}
	if deg == 0 {
		return 1
	} else if deg == 90 || deg == 270 {
		return 0
	} else if deg == 180 {
		return -1
	} else {
		return int(math.Sin(float64(deg) * math.Pi / 180))
	}
}

//func Decrypt() {

// Take a compressed byte stream representing a ride and turn it into a Ride
// struct. Returns an error if the byte array is too short.
func Unmarshal(buf []byte, r *Ride) error {
	if len(buf) < IDX_TRACK_DATA {
		return errors.New("buffer too short to be a ride")
	}
	r.RideType = RideType(buf[IDX_RIDE_TYPE])
	r.VehicleType = VehicleType(buf[IDX_VEHICLE_TYPE])
	r.HasLoop = hasLoop(int(buf[IDX_HAS_LOOP]))
	r.OperatingMode = OperatingMode(buf[IDX_OPERATING_MODE])
	r.ControlFlags = parseControlFlags(int(buf[IDX_CONTROL_FLAG]))
	r.NumTrains = int(buf[IDX_NUM_TRAINS])
	r.CarsPerTrain = int(buf[IDX_CARS_PER_TRAIN])
	r.MinWaitTime = int(buf[IDX_MIN_WAIT_TIME])
	r.MaxWaitTime = int(buf[IDX_MAX_WAIT_TIME])
	r.TrackData = parseTrackData(buf[IDX_TRACK_DATA:])

	fmt.Println(r.VehicleType)
	return nil
}

func Marshal(r *Ride) ([]byte, error) {
	// at a minimum, rides have this much data
	bits := make([]byte, 0xc4)

	// This is a little bit tricky, and requires implementing the format
	// described in the tycoon technical depot, available here:
	// https://github.com/UnknownShadow200/RCTTechDepot-Archive/blob/master/td4.html
	bits[IDX_RIDE_TYPE] = byte(r.RideType)
	bits[IDX_VEHICLE_TYPE] = byte(r.VehicleType)
	//bits[IDX_HAS_LOOP] = byte(r.HasLoop)

	return bits, nil
}

var EndOfRide = errors.New("End of ride")

// Parse the serialized data into a TrackElement struct
// When the end of ride is encountered a EndOfRide error is returned
func parseElement(rawElement []byte) (te TrackElement, e error) {
	if len(rawElement) != 2 {
		return TrackElement{}, errors.New("invalid length for element input")
	}
	te.Segment = &TrackSegment{
		Type: SegmentType(rawElement[0]),
	}
	if te.Segment.Type == ELEM_END_OF_RIDE {
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

	// XXX where is this data actually stored?
	td.Clearance = 2
	td.ClearanceDirection = CLEARANCE_ABOVE
	return *td
}

// Advance all of the values by one track segment.
func advanceTrack(ts *TrackSegment, ΔE int, ΔForward int, ΔSideways int,
	direction DirectionDelta) (int, int, int, DirectionDelta) {
	ΔE += ts.ElevationGain

	ΔForward += cosdeg(direction) * ts.ForwardDelta
	ΔForward += sindeg(direction) * ts.SidewaysDelta

	ΔSideways += sindeg(direction) * ts.ForwardDelta
	ΔSideways += cosdeg(direction) * ts.SidewaysDelta

	direction += ts.DirectionDelta
	for ; direction >= 360; direction -= 360 {
	}

	return ΔE, ΔForward, ΔSideways, direction
}

// Test whether the ride's track forms a continuous circuit. Does not test
// whether the ride collides with itself.
func IsCircuit(t *TrackData) bool {
	// X and Y don't really make sense as variable names, easier to think about
	// relative changes
	eΔ, forwardΔ, sidewaysΔ := 0, 0, 0
	direction := DIRECTION_STRAIGHT
	if len(t.Elements) == 0 {
		return false
	}
	for i := range t.Elements {
		ts := t.Elements[i].Segment
		eΔ, forwardΔ, sidewaysΔ, direction = advanceTrack(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
	}
	return forwardΔ == 0 && sidewaysΔ == 0 && eΔ == 0
}

// Detect whether the track collides with itself.
func HasCollision(t *TrackData) bool {
	matrix := make([][][]bool, 100)
	for i := range matrix {
		matrix[i] = make([][]bool, 100)
		for j := range matrix[i] {
			matrix[i][j] = make([]bool, 100)
		}
	}
	eΔ, forwardΔ, sidewaysΔ := 0, 0, 0
	direction := DIRECTION_STRAIGHT
	for i := range t.Elements {
		ts := t.Elements[i].Segment
		eΔ, forwardΔ, sidewaysΔ, direction = advanceTrack(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
		// if there already exists a piece there, we can't build.
		if matrix[forwardΔ][sidewaysΔ][eΔ] {
			return true
		}
		matrix[forwardΔ][sidewaysΔ][eΔ] = true
	}
	return false
}

func main() {
	fmt.Println(readRaptor())
}

func readRaptor() Ride {
	encodedBits, err := ioutil.ReadFile("rides/raptor.td4")
	if err != nil {
		panic(err)
	}
	z := NewReader(bytes.NewReader(encodedBits))
	if err != nil {
		panic(err)
	}
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	decrypted := bitbuffer.Bytes()
	for i := 0; i < 500; i++ {
		// encode the value of i as hex
		ds := hex.EncodeToString([]byte{byte(i)})
		bitValueInHex := hex.EncodeToString([]byte{decrypted[i]})
		fmt.Printf("%s: %s\n", ds, bitValueInHex)
	}

	// r is a pointer
	r := new(Ride)
	Unmarshal(decrypted, r)
	fmt.Println(r.VehicleType)
	return *r
}
