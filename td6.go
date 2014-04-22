// Parses TD6 files
package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/kevinburke/rct-rides/tracks"
	"io/ioutil"
)

func hasLoop(n int) bool {
	return hasBit(n, BIT_VERTICAL_LOOP)
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

// Where to find various pieces of information in the decoded ride.
const (
	IDX_RIDE_TYPE = 0x1
	// Indications this isn't stored here in td6
	IDX_FEATURES_0     = 0x2
	IDX_FEATURES_1     = 0x3
	IDX_OPERATING_MODE = 0x06
	IDX_COLOR_SCHEME   = 0x07
	IDX_CONTROL_FLAG   = 0x4b
	IDX_NUM_TRAINS     = 0x4c
	IDX_CARS_PER_TRAIN = 0x4d
	IDX_MIN_WAIT_TIME  = 0x4e
	IDX_MAX_WAIT_TIME  = 0x4f
	IDX_VEHICLE_TYPE   = 0x74
	IDX_TRACK_DATA     = 0xa3
	IDX_X_SPACE        = 0x80
	IDX_Y_SPACE        = 0x81
)

type Ride struct {
	RideType    RideType
	VehicleType VehicleType

	XSpaceRequired int
	YSpaceRequired int

	OperatingMode OperatingMode
	ControlFlags  *ControlFlags
	NumTrains     int
	CarsPerTrain  int
	MinWaitTime   int
	MaxWaitTime   int
	TrackData     tracks.Data

	// set in bit 0 of the ride features list
	HasLoop         bool
	SteepLiftChain  bool
	CurvedLiftChain bool
	Banking         bool

	// set in bit 1
	SteepSlope        bool
	FlatToSteep       bool
	SlopedCurves      bool
	SteepTwist        bool
	SBends            bool
	SmallRadiusCurves bool
	SmallRadiusBanked bool
}

type RideType int

const (
	RIDE_SPIRAL     RideType = 0x00
	RIDE_STAND_UP            = 0x01
	RIDE_SUSPENDED           = 0x02
	RIDE_INVERTED            = 0x03
	RIDE_STEEL_MINI          = 0x04
	RIDE_WOODEN              = 0x34
)

type VehicleType string
type LoadType int
type OperatingMode int

const (
	VEHICLE_SPIRAL                  VehicleType = "SPDRCR"
	VEHICLE_STAND_UP                            = "TOGST"
	VEHICLE_WOODEN_ARTICULATED                  = "MFT"
	VEHICLE_WOODEN_4SEATER                      = "PTCT1"
	VEHICLE_WOODEN_6SEATER                      = "PTCT2"
	VEHICLE_WOODEN_6SEATER_REVERSED             = "PTCT2R"
)

const (
	LENGTH_VEHICLE_TYPE = 8

	BIT_LIFT_CHAIN        = 3
	BIT_STEEP_LIFT_CHAIN  = 4
	BIT_CURVED_LIFT_CHAIN = 5
	BIT_BANKING           = 6
	BIT_VERTICAL_LOOP     = 7

	BIT_STEEP_SLOPE         = 1
	BIT_FLAT_TO_STEEP       = 2
	BIT_SLOPED_CURVES       = 3
	BIT_STEEP_TWIST         = 4
	BIT_S_BENDS             = 5
	BIT_SMALL_RADIUS_CURVES = 6
	BIT_SMALL_RADIUS_BANKED = 7
)

// Take a compressed byte stream representing a ride and turn it into a Ride
// struct. Returns an error if the byte array is too short.
func Unmarshal(buf []byte, r *Ride) error {
	if len(buf) < IDX_TRACK_DATA {
		return errors.New("buffer too short to be a ride")
	}
	r.RideType = RideType(buf[IDX_RIDE_TYPE])

	featuresBit0 := int(buf[IDX_FEATURES_0])
	r.SteepLiftChain = hasBit(featuresBit0, BIT_STEEP_LIFT_CHAIN)
	r.CurvedLiftChain = hasBit(featuresBit0, BIT_CURVED_LIFT_CHAIN)
	r.Banking = hasBit(featuresBit0, BIT_BANKING)
	r.HasLoop = hasBit(featuresBit0, BIT_VERTICAL_LOOP)

	featuresBit1 := int(buf[IDX_FEATURES_1])
	r.SteepSlope = hasBit(featuresBit1, BIT_STEEP_SLOPE)
	r.FlatToSteep = hasBit(featuresBit1, BIT_FLAT_TO_STEEP)
	r.SlopedCurves = hasBit(featuresBit1, BIT_SLOPED_CURVES)
	r.SteepTwist = hasBit(featuresBit1, BIT_STEEP_TWIST)
	r.SBends = hasBit(featuresBit1, BIT_S_BENDS)
	r.SmallRadiusCurves = hasBit(featuresBit1, BIT_SMALL_RADIUS_CURVES)
	r.SmallRadiusBanked = hasBit(featuresBit1, BIT_SMALL_RADIUS_BANKED)

	r.OperatingMode = OperatingMode(buf[IDX_OPERATING_MODE])
	r.ControlFlags = parseControlFlags(int(buf[IDX_CONTROL_FLAG]))
	r.NumTrains = int(buf[IDX_NUM_TRAINS])
	r.CarsPerTrain = int(buf[IDX_CARS_PER_TRAIN])
	r.MinWaitTime = int(buf[IDX_MIN_WAIT_TIME])
	r.MaxWaitTime = int(buf[IDX_MAX_WAIT_TIME])

	d := new(tracks.Data)
	tracks.Unmarshal(buf[IDX_TRACK_DATA:], d)
	r.TrackData = *d

	r.VehicleType = VehicleType(string(buf[IDX_VEHICLE_TYPE : IDX_VEHICLE_TYPE+LENGTH_VEHICLE_TYPE]))

	r.XSpaceRequired = int(buf[IDX_X_SPACE])
	r.YSpaceRequired = int(buf[IDX_Y_SPACE])

	// Not sure where the source of truth for this is yet
	// r.VehicleType = VehicleType(buf[IDX_VEHICLE_TYPE])

	return nil
}
func readRide(filename string) Ride {
	encodedBits, err := ioutil.ReadFile(filename)

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
	for i := 0; i < 200; i++ {
		// encode the value of i as hex
		ds := hex.EncodeToString([]byte{byte(i)})
		bitValueInHex := hex.EncodeToString([]byte{decrypted[i]})
		fmt.Printf("%s: %s\n", ds, bitValueInHex)
	}

	// r is a pointer
	r := new(Ride)
	Unmarshal(decrypted, r)
	//bits, err := Marshal(r)
	//fmt.Println(bits)
	return *r
}

func main() {
	fmt.Println(readRide("rides/mischief.td6"))
}
