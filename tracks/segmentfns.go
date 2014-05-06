package tracks

// Compute all of the possible track pieces that can be built
func (s *Segment) Possibilities() (o []Segment) {
	for _, val := range TS_MAP {
		// XXX, need to check diagonal
		if val.InputDegree == s.OutputDegree && val.StartingBank == s.EndingBank {
			o = append(o, *val)
		}
	}
	return o
}
