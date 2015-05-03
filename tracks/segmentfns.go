package tracks

func isPossibleChainLift(val *Segment) bool {
	// XXX, check other pieces that can be chain lifts. in theory you can also
	// place chain lifts on straight pieces of track but this seems silly
	// (maybe a super in-future optimization will allow this)
	return val.InputDegree == TRACK_UP_25 &&
		(val.OutputDegree == TRACK_UP_25 || val.OutputDegree == TRACK_NONE) &&
		val.StartingBank == TRACK_BANK_NONE &&
		val.EndingBank == TRACK_BANK_NONE
}

// Possibilities computes all of the possible track pieces that can be built
func (s *Element) Possibilities() (o []Element) {
	for _, val := range TS_MAP {
		// XXX, need to check diagonal
		if val.InputDegree == s.Segment.OutputDegree && val.StartingBank == s.Segment.EndingBank {
			o = append(o, Element{Segment: val, ChainLift: false})
			if isPossibleChainLift(val) {
				o = append(o, Element{Segment: val, ChainLift: true})
			}
			// XXX, cache possibilities, since they're not going to change, or
			// just build this once
		}
	}
	return o
}

// Compatible returns true if the end of the first element is compatible with
// the beginning of the second element.
func Compatible(first Element, second Element) bool {
	return first.Segment.OutputDegree == second.Segment.InputDegree &&
		first.Segment.EndingBank == second.Segment.StartingBank
}
