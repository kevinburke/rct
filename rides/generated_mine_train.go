package rides

import "github.com/kevinburke/rct/td6"

var ride

func init() {
	ride = &td6.Ride{
		RideType:           17,
		VehicleColorScheme: 0,
		XSpaceRequired:     32,
		YSpaceRequired:     36,
		OperatingMode:      1,
		ControlFlags:       (*td6.ControlFlags)(0xc20802b0e0),
		NumTrains:          0x2,
		CarsPerTrain:       0x6,
		MinWaitTime:        0x0,
		MaxWaitTime:        0x0,
		HasLoop:            false,
		SteepLiftChain:     false,
		CurvedLiftChain:    false,
		Banking:            false,
		SteepSlope:         false,
		FlatToSteep:        false,
		SlopedCurves:       false,
		SteepTwist:         false,
		SBends:             false,
		SmallRadiusCurves:  false,
		SmallRadiusBanked:  false,
		NumInversions:      0x0,
		MaxSpeed:           0x0,
		AverageSpeed:       0x0,
		VehicleType:        "AMT1    ",
		DatData:            []uint8(nil),
		Egresses:           []*td6.Egress(nil),
		Excitement:         0,
		Intensity:          0,
		Nausea:             0}
}
