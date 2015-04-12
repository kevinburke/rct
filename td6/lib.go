// Parses TD6 files
package td6

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/kevinburke/rct-rides/bits"
	"github.com/kevinburke/rct-rides/rle"
	"github.com/kevinburke/rct-rides/tracks"
)

const DEBUG = false
const DEBUG_LENGTH = 20

func hasLoop(n int) bool {
	return bits.On(n, BIT_VERTICAL_LOOP)
}

type ControlFlags struct {
	UseMaximumTime             bool
	UseMinimumTime             bool
	SyncWithAdjacentStation    bool
	LeaveIfAnotherTrainArrives bool
	WaitForLoad                bool
	Load                       LoadType
}

func unmarshalControlFlags(n int) *ControlFlags {
	return &ControlFlags{
		UseMaximumTime:             n>>7 == 1,
		UseMinimumTime:             n>>6 == 1,
		SyncWithAdjacentStation:    n>>5 == 1,
		LeaveIfAnotherTrainArrives: n>>4 == 1,
		WaitForLoad:                n>>3 == 1,
		Load:                       LoadType(n & 7),
	}
}

func marshalControlFlags(c ControlFlags) (n int) {
	n = bits.SetCond(n, 7, c.UseMaximumTime)
	n = bits.SetCond(n, 6, c.UseMinimumTime)
	n = bits.SetCond(n, 5, c.SyncWithAdjacentStation)
	n = bits.SetCond(n, 4, c.LeaveIfAnotherTrainArrives)
	n = bits.SetCond(n, 3, c.WaitForLoad)
	n |= int(c.Load)
	return n
}

// Where to find various pieces of information in the decoded ride.
const (
	IDX_RIDE_TYPE        = 0x0
	IDX_VEHICLE_TYPE     = 0x1
	IDX_FEATURES_0       = 0x1
	IDX_FEATURES_1       = 0x1
	IDX_COST             = 0x2
	IDX_OPERATING_MODE   = 0x06
	IDX_COLOR_SCHEME     = 0x07
	IDX_CONTROL_FLAG     = 0x4b
	IDX_NUM_TRAINS       = 0x4c
	IDX_CARS_PER_TRAIN   = 0x4d
	IDX_MIN_WAIT_TIME    = 0x4e
	IDX_MAX_WAIT_TIME    = 0x4f
	IDX_MAX_SPEED        = 0x51
	IDX_AVERAGE_SPEED    = 0x52
	IDX_POSITIVE_G_FORCE = 0x55
	IDX_NEGATIVE_G_FORCE = 0x56
	IDX_LATERAL_G_FORCE  = 0x57
	IDX_NUM_INVERSIONS   = 0x58
	IDX_NUM_DROPS        = 0x59
	IDX_HIGHEST_DROP     = 0x5a

	IDX_EXCITEMENT = 0x5b
	IDX_INTENSITY  = 0x5c
	IDX_NAUSEA     = 0x5d

	IDX_VEHICLE_TYPE_STRING = 0x74

	IDX_TRACK_DATA = 0xa3
	IDX_X_SPACE    = 0x80
	IDX_Y_SPACE    = 0x81
)

// Technically this is track data that gets serialized to disk. A RCT2 Ride
// structure in memory has a different format.
type Ride struct {
	RideType    RideType
	VehicleType VehicleType

	XSpaceRequired int
	YSpaceRequired int

	OperatingMode OperatingMode
	ControlFlags  *ControlFlags
	NumTrains     uint8
	CarsPerTrain  uint8
	MinWaitTime   uint8
	MaxWaitTime   uint8
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

	NumInversions uint8
	MaxSpeed      uint8
	AverageSpeed  uint8

	// This is a little bit of a copout
	DatData []byte

	Egresses []*Egress

	Excitement int16
	Intensity  int16
	Nausea     int16
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
	r.SteepLiftChain = bits.On(featuresBit0, BIT_STEEP_LIFT_CHAIN)
	r.CurvedLiftChain = bits.On(featuresBit0, BIT_CURVED_LIFT_CHAIN)
	r.Banking = bits.On(featuresBit0, BIT_BANKING)
	r.HasLoop = bits.On(featuresBit0, BIT_VERTICAL_LOOP)

	featuresBit1 := int(buf[IDX_FEATURES_1])
	r.SteepSlope = bits.On(featuresBit1, BIT_STEEP_SLOPE)
	r.FlatToSteep = bits.On(featuresBit1, BIT_FLAT_TO_STEEP)
	r.SlopedCurves = bits.On(featuresBit1, BIT_SLOPED_CURVES)
	r.SteepTwist = bits.On(featuresBit1, BIT_STEEP_TWIST)
	r.SBends = bits.On(featuresBit1, BIT_S_BENDS)
	r.SmallRadiusCurves = bits.On(featuresBit1, BIT_SMALL_RADIUS_CURVES)
	r.SmallRadiusBanked = bits.On(featuresBit1, BIT_SMALL_RADIUS_BANKED)

	r.OperatingMode = OperatingMode(buf[IDX_OPERATING_MODE])
	r.ControlFlags = unmarshalControlFlags(int(buf[IDX_CONTROL_FLAG]))
	r.NumTrains = uint8(buf[IDX_NUM_TRAINS])
	r.CarsPerTrain = uint8(buf[IDX_CARS_PER_TRAIN])
	r.MinWaitTime = uint8(buf[IDX_MIN_WAIT_TIME])
	r.MaxWaitTime = uint8(buf[IDX_MAX_WAIT_TIME])
	r.NumInversions = uint8(buf[IDX_NUM_INVERSIONS])
	r.MaxSpeed = uint8(buf[IDX_MAX_SPEED])
	r.AverageSpeed = uint8(buf[IDX_AVERAGE_SPEED])

	d := new(tracks.Data)
	tracks.Unmarshal(buf[IDX_TRACK_DATA:], d)
	r.TrackData = *d

	r.VehicleType = VehicleType(string(buf[IDX_VEHICLE_TYPE_STRING : IDX_VEHICLE_TYPE_STRING+LENGTH_VEHICLE_TYPE]))

	r.XSpaceRequired = int(buf[IDX_X_SPACE])
	r.YSpaceRequired = int(buf[IDX_Y_SPACE])

	r.DatData = buf[0x70:0x80]

	// XXX, not sure if this fn should know about this, or the track data
	entranceExitIdx := IDX_TRACK_DATA + 2*len(r.TrackData.Elements) + 1
	r.Egresses = unmarshalEgress(buf[entranceExitIdx:])

	r.Excitement = int16(buf[IDX_EXCITEMENT])
	r.Intensity = int16(buf[IDX_INTENSITY])
	r.Nausea = int16(buf[IDX_NAUSEA])

	// Ignore scenery data for now.
	// sceneryIdx := entranceExitIdx + 6*len(r.Egresses) + 1

	return nil
}

// Deserialize the ride to a string of bytes.
//
// This is a little bit tricky, and requires implementing the format
// described in the tycoon technical depot, available here:
// https://github.com/UnknownShadow200/RCTTechDepot-Archive/blob/master/td4.html
func Marshal(r *Ride) ([]byte, error) {
	// at a minimum, rides have this much data
	rideb := make([]byte, 0xa3)

	rideb[IDX_RIDE_TYPE] = byte(r.RideType)

	copy(rideb[IDX_VEHICLE_TYPE:IDX_VEHICLE_TYPE+8], r.VehicleType)
	featureBit0 := 0
	if r.SteepLiftChain {
		bits.Set(featureBit0, BIT_STEEP_LIFT_CHAIN)
	}
	if r.CurvedLiftChain {
		bits.Set(featureBit0, BIT_CURVED_LIFT_CHAIN)
	}
	if r.Banking {
		bits.Set(featureBit0, BIT_BANKING)
	}
	if r.HasLoop {
		bits.Set(featureBit0, BIT_VERTICAL_LOOP)
	}
	rideb[IDX_FEATURES_0] = byte(featureBit0)

	// Features bit 1
	featureBit1 := 0
	if r.SteepSlope {
		bits.Set(featureBit1, BIT_STEEP_SLOPE)
	}
	if r.FlatToSteep {
		bits.Set(featureBit1, BIT_FLAT_TO_STEEP)
	}
	if r.SlopedCurves {
		bits.Set(featureBit1, BIT_SLOPED_CURVES)
	}
	if r.SteepTwist {
		bits.Set(featureBit1, BIT_STEEP_TWIST)
	}
	rideb[IDX_FEATURES_1] = byte(featureBit1)

	rideb[IDX_OPERATING_MODE] = byte(r.OperatingMode)
	rideb[IDX_CONTROL_FLAG] = byte(marshalControlFlags(*r.ControlFlags))
	rideb[IDX_NUM_TRAINS] = byte(r.NumTrains)
	rideb[IDX_CARS_PER_TRAIN] = byte(r.CarsPerTrain)
	rideb[IDX_MIN_WAIT_TIME] = byte(r.MinWaitTime)
	rideb[IDX_MAX_WAIT_TIME] = byte(r.MaxWaitTime)
	rideb[IDX_NUM_INVERSIONS] = byte(r.NumInversions)
	rideb[IDX_X_SPACE] = byte(r.XSpaceRequired)
	rideb[IDX_Y_SPACE] = byte(r.YSpaceRequired)
	rideb[IDX_MAX_SPEED] = byte(r.MaxSpeed)
	rideb[IDX_AVERAGE_SPEED] = byte(r.AverageSpeed)

	// XXX The below information hasn't been parsed into ride data yet - this
	// is just writing data to the file so it parses

	// Color scheme, 3rd bit is RCT2 flag
	rideb[7] = 1 << 3
	// Set a colorscheme (taken from Mischief)
	rideb[8] = 18
	// air time
	rideb[0x4a] = 0x13
	// max speed
	rideb[0x51] = 0x15
	// average speed
	rideb[0x52] = 9

	// ride length. also just copied from mischief
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint16(0x03b6))
	copy(rideb[0x53:0x53+2], buf.Bytes())

	// positive G force
	rideb[0x55] = 0x0a
	// negative g force
	rideb[0x56] = 0xfb
	// lateral G force
	rideb[0x57] = 0x07

	// XXX
	copy(rideb[0x70:0x80], r.DatData)

	// Top 3 bytes are the number of circuits
	rideb[0xa2] = 1 << 5
	// Low 5 bytes are the lift speed
	rideb[0xa2] |= 7

	trackBits, err := tracks.Marshal(&r.TrackData)
	if err != nil {
		return []byte{}, err
	}
	rideb = append(rideb, trackBits...)

	egresses, err := marshalEgresses(r.Egresses)
	rideb = append(rideb, egresses...)

	// append empty scenery data, XXX when we have scenery.
	rideb = append(rideb, 0xff)
	rideb = append(rideb, bytes.Repeat([]byte{0}, 21)...)
	return rideb, nil
}

func ReadRide(filename string) *Ride {
	encodedBits, err := ioutil.ReadFile(filename)
	bitsWithoutChecksum := encodedBits[:len(encodedBits)-4]

	if err != nil {
		panic(err)
	}
	z := rle.NewReader(bytes.NewReader(bitsWithoutChecksum))
	if err != nil {
		panic(err)
	}
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	decrypted := bitbuffer.Bytes()

	// r is a pointer
	r := new(Ride)
	Unmarshal(decrypted, r)

	if DEBUG {
		begin := 0xa2 + 2*len(r.TrackData.Elements) - 3
		for i := begin; i < begin+DEBUG_LENGTH; i++ {
			// encode the value of i as hex
			if i > len(decrypted) {
				continue
			}
			ds := hex.EncodeToString([]byte{byte(i)})
			bitValueInHex := hex.EncodeToString([]byte{decrypted[i]})
			fmt.Printf("%s: %s\n", ds, bitValueInHex)
		}
	}

	//bits, err := Marshal(r)
	//fmt.Println(bits)
	return r
}

func _pad(bits []byte, n int) []byte {
	return append(bits, bytes.Repeat([]byte{0}, n)...)
}

const RCT2_TD6_LENGTH = 24735

func Pad(bits []byte) []byte {
	return _pad(bits, RCT2_TD6_LENGTH-len(bits))
}
