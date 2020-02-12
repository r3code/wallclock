package wallclock

import "time"

var (  	
	nowFunc = currentTimeUTC
)

func currentTimeUTC() time.Time {
	return time.Now().UTC()
}

// Now gets tim by calling assigned nowFunc. 
func Now() time.Time {
	return nowFunc()
}

// FakeTime changes wallclock.Now behaviour to return "2014-11-12T11:45:26.371Z".
// Returns cancelFunc which should be called to turn time back to realtime.
// Example:
//   cancelFake = wallclock.FakeTime()
//   // manual cancel
//   cancelFake() // or defer cancelFake()
// or
//   defer wallclock.FakeTime()()
func FakeTime() (cancelFunc func()) {
	nowFunc = fakeNow
	cancelFunc = func() {
		nowFunc = currentTimeUTC
	}
	return cancelFunc
}

// FakeMoment changes wallclock.Now behaviour to return the time specified by `timeMoment`.
// Returns cancelFunc which should be called to turn time back to realtime.
func FakeMoment(timeMoment time.Time) (cancelFunc func()) {
	nowFunc = func() time.Time {
		return timeMoment
	}
	cancelFunc = func() {
		nowFunc = currentTimeUTC
	}
	return cancelFunc
}

func fakeNow() time.Time {
	timeNow, _ := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")
	return timeNow
}
