// Parses TD6 files
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/kevinburke/rct-rides/tracks"
)

const DEBUG = false
const DEBUG_LENGTH = 20

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
	n = setConditionalBit(n, 7, c.UseMaximumTime)
	n = setConditionalBit(n, 6, c.UseMinimumTime)
	n = setConditionalBit(n, 5, c.SyncWithAdjacentStation)
	n = setConditionalBit(n, 4, c.LeaveIfAnotherTrainArrives)
	n = setConditionalBit(n, 3, c.WaitForLoad)
	n |= int(c.Load)
	return n
}

// Where to find various pieces of information in the decoded ride.
const (
	IDX_RIDE_TYPE = 0x0
	// Indications this isn't stored here in td6
	IDX_FEATURES_0       = 0x2
	IDX_FEATURES_1       = 0x3
	IDX_OPERATING_MODE   = 0x06
	IDX_COLOR_SCHEME     = 0x07
	IDX_CONTROL_FLAG     = 0x4b
	IDX_NUM_TRAINS       = 0x4c
	IDX_CARS_PER_TRAIN   = 0x4d
	IDX_MIN_WAIT_TIME    = 0x4e
	IDX_MAX_WAIT_TIME    = 0x4f
	IDX_POSITIVE_G_FORCE = 0x55
	IDX_NEGATIVE_G_FORCE = 0x56
	IDX_LATERAL_G_FORCE  = 0x57
	IDX_NUM_INVERSIONS   = 0x58
	IDX_NUM_DROPS        = 0x59
	IDX_HIGHEST_DROP     = 0x5a
	IDX_EXCITEMENT       = 0x5b
	IDX_INTENSITY        = 0x5c
	IDX_NAUSEA           = 0x5d

	IDX_VEHICLE_TYPE = 0x74

	IDX_TRACK_DATA = 0xa3
	IDX_X_SPACE    = 0x80
	IDX_Y_SPACE    = 0x81
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

	NumInversions int

	// This is a little bit of a copout
	DatData []byte

	Egresses []*Egress
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
	r.ControlFlags = unmarshalControlFlags(int(buf[IDX_CONTROL_FLAG]))
	r.NumTrains = int(buf[IDX_NUM_TRAINS])
	r.CarsPerTrain = int(buf[IDX_CARS_PER_TRAIN])
	r.MinWaitTime = int(buf[IDX_MIN_WAIT_TIME])
	r.MaxWaitTime = int(buf[IDX_MAX_WAIT_TIME])
	r.NumInversions = int(buf[IDX_NUM_INVERSIONS])

	d := new(tracks.Data)
	tracks.Unmarshal(buf[IDX_TRACK_DATA:], d)
	r.TrackData = *d

	r.VehicleType = VehicleType(string(buf[IDX_VEHICLE_TYPE : IDX_VEHICLE_TYPE+LENGTH_VEHICLE_TYPE]))

	r.XSpaceRequired = int(buf[IDX_X_SPACE])
	r.YSpaceRequired = int(buf[IDX_Y_SPACE])

	r.DatData = buf[0x70:0x80]

	// XXX, not sure if this fn should know about this, or the track data
	entranceExitIdx := IDX_TRACK_DATA + 2*len(r.TrackData.Elements) + 1
	r.Egresses = unmarshalEgress(buf[entranceExitIdx:])

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
	bits := make([]byte, 0xa3)

	fmt.Println(r.RideType)
	bits[IDX_RIDE_TYPE] = byte(r.RideType)

	copy(bits[IDX_VEHICLE_TYPE:IDX_VEHICLE_TYPE+8], r.VehicleType)
	featureBit0 := 0
	if r.SteepLiftChain {
		setBit(featureBit0, BIT_STEEP_LIFT_CHAIN)
	}
	if r.CurvedLiftChain {
		setBit(featureBit0, BIT_CURVED_LIFT_CHAIN)
	}
	if r.Banking {
		setBit(featureBit0, BIT_BANKING)
	}
	if r.HasLoop {
		setBit(featureBit0, BIT_VERTICAL_LOOP)
	}
	bits[IDX_FEATURES_0] = byte(featureBit0)

	// Features bit 1
	featureBit1 := 0
	if r.SteepSlope {
		setBit(featureBit1, BIT_STEEP_SLOPE)
	}
	if r.FlatToSteep {
		setBit(featureBit1, BIT_FLAT_TO_STEEP)
	}
	if r.SlopedCurves {
		setBit(featureBit1, BIT_SLOPED_CURVES)
	}
	if r.SteepTwist {
		setBit(featureBit1, BIT_STEEP_TWIST)
	}
	bits[IDX_FEATURES_1] = byte(featureBit1)

	bits[IDX_OPERATING_MODE] = byte(r.OperatingMode)
	bits[IDX_CONTROL_FLAG] = byte(marshalControlFlags(*r.ControlFlags))
	bits[IDX_NUM_TRAINS] = byte(r.NumTrains)
	bits[IDX_CARS_PER_TRAIN] = byte(r.CarsPerTrain)
	bits[IDX_MIN_WAIT_TIME] = byte(r.MinWaitTime)
	bits[IDX_MAX_WAIT_TIME] = byte(r.MaxWaitTime)
	bits[IDX_NUM_INVERSIONS] = byte(r.NumInversions)
	bits[IDX_X_SPACE] = byte(r.XSpaceRequired)
	bits[IDX_Y_SPACE] = byte(r.YSpaceRequired)

	// XXX The below information hasn't been parsed into ride data yet - this
	// is just writing data to the file so it parses

	// Color scheme, 3rd bit is RCT2 flag
	bits[7] = 1 << 3
	// Set a colorscheme (taken from Mischief)
	bits[8] = 18
	// air time
	bits[0x4a] = 0x13
	// max speed
	bits[0x51] = 0x15
	// average speed
	bits[0x52] = 9

	// ride length. also just copied from mischief
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint16(0x03b6))
	copy(bits[0x53:0x53+2], buf.Bytes())

	// positive G force
	bits[0x55] = 0x0a
	// negative g force
	bits[0x56] = 0xfb
	// lateral G force
	bits[0x57] = 0x07

	// XXX
	copy(bits[0x70:0x80], r.DatData)

	// Top 3 bytes are the number of circuits
	bits[0xa2] = 1 << 5
	// Low 5 bytes are the lift speed
	bits[0xa2] |= 7

	trackBits, err := tracks.Marshal(&r.TrackData)
	if err != nil {
		return []byte{}, err
	}
	bits = append(bits, trackBits...)

	egresses, err := marshalEgresses(r.Egresses)
	bits = append(bits, egresses...)

	// append empty scenery data, XXX when we have scenery.
	bits = append(bits, 0xff)
	bits = append(bits, bytes.Repeat([]byte{0}, 21)...)
	return bits, nil
}

func ReadRide(filename string) *Ride {
	encodedBits, err := ioutil.ReadFile(filename)
	bitsWithoutChecksum := encodedBits[:len(encodedBits)-4]

	if err != nil {
		panic(err)
	}
	z := NewReader(bytes.NewReader(bitsWithoutChecksum))
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

func pad(bits []byte) []byte {
	return _pad(bits, RCT2_TD6_LENGTH-len(bits))
}

func main() {
	encodedBits, err := ioutil.ReadFile("rides/mischief.td6")
	fmt.Println(hex.EncodeToString(encodedBits[len(encodedBits)-4 : len(encodedBits)]))
	bitsWithoutChecksum := encodedBits[:len(encodedBits)-4]

	if err != nil {
		panic(err)
	}
	z := NewReader(bytes.NewReader(bitsWithoutChecksum))
	if err != nil {
		panic(err)
	}

	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	decrypted := bitbuffer.Bytes()

	// r is a pointer
	r := new(Ride)
	Unmarshal(decrypted, r)

	bits, err := Marshal(r)
	if err != nil {
		panic(err)
	}

	for i := range bits {
		if bits[i] != decrypted[i] {
			fmt.Printf("%d: ", i)
			fmt.Printf("byte %x differs, in my mischief it is %d, in orig it is %d\n", i, bits[i], decrypted[i])
		}
	}

	if DEBUG {
		begin := 0xa2 + 2*len(r.TrackData.Elements) - 3
		for i := begin; i < begin+DEBUG_LENGTH; i++ {
			// encode the value of i as hex
			ds := hex.EncodeToString([]byte{byte(i)})
			bitValueInHex := hex.EncodeToString([]byte{bits[i]})
			fmt.Printf("%s: %s\n", ds, bitValueInHex)
		}
	}

	paddedBits := pad(bits)

	fmt.Println(paddedBits)
	fmt.Println(decrypted)

	var buf bytes.Buffer
	w := NewWriter(&buf)
	w.Write(paddedBits)
	ioutil.WriteFile("/Users/kevin/Applications/Wineskin/rct2.app/Contents/Resources/drive_c/GOG Games/RollerCoaster Tycoon 2 Triple Thrill Pack/Tracks/mymischief.TD6", buf.Bytes(), 0644)
	fmt.Println("Wrote rides/mischief.td6.out.")
}
