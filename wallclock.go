package wallclock

import "time"

var (  	
	nowFunc = currentTime
)

// Now gets tim by calling assigned nowFunc. 
func Now() time.Time {
	return nowFunc()
}

// FakeMoment changes wallclock.Now behaviour to return the time specified by `timeMoment`.
// Returns cancelFunc which should be called to turn time back to realtime.
// Example:
//   cancelFake = wallclock.FakeMoment()
//   // manual cancel
//   cancelFake() // or defer cancelFake()
// or
//   defer wallclock.FakeMoment()()
func FakeMoment(timeMoment time.Time) (cancelFunc func()) {
	nowFunc = func() time.Time {
		return timeMoment
	}
	cancelFunc = func() {
		nowFunc = currentTime
	}
	return cancelFunc
}

// FakeFixedMoment changes wallclock.Now behaviour to return "2014-11-12T11:45:26.371Z".
// Returns cancelFunc which should be called to turn time back to realtime.
func FakeFixedMoment() (cancelFunc func()) {
  timeFrozen, _ := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z") 	
	return FakeMoment(timeFrozen)
}

// return realtime clock 
func currentTime() time.Time {
	return time.Now()
}
