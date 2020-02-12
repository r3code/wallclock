package wallclock_test

import (
	"testing"
	"time"

	"github.com/r3code/wallclock"
)

func TestFakeTime(t *testing.T) {
	var t1, t2, ftA, ftB time.Time
	t1 = wallclock.Now()
	cancelFunc := wallclock.FakeTime()
	ftA = wallclock.Now()
	ftB = wallclock.Now()
	if t1.Equal(t2) {
		t.Errorf("FakeTime() failed")
	}
	if !ftA.Equal(ftB) {
		t.Errorf("FakeTime() fake times not equal %v != %v", ftA.String(), ftB.String())
	}
	cancelFunc()
	t2 = wallclock.Now()
	if t2.Equal(ftA) {
		t.Errorf("FakeTime() cancelFunc failed to return normal time")
	}
}

func TestFakeTimeDeferred(t *testing.T) {
	var t1, t2, ftA, ftB time.Time
	wrapDeferFunc := func() {
		t1 = wallclock.Now()
		defer wallclock.FakeTime()
		ftA = wallclock.Now()
		ftB = wallclock.Now()
	}
	wrapDeferFunc()
	t2 = wallclock.Now()

	if t1.Equal(t2) {
		t.Errorf("FakeTime() failed after defer called")
	}
	if !ftA.Equal(ftB) {
		t.Errorf("FakeTime() fake times not equal %v != %v", ftA.String(), ftB.String())
	}
}

func TestFakeMoment(t *testing.T) {
	testMoment, _ := time.Parse(time.RFC3339, "1999-01-02T12:34:56.001Z")
	var t1, t2, ftA time.Time
	t1 = wallclock.Now()
	cancelFunc := wallclock.FakeMoment(testMoment)
	if !t1.After(testMoment) {
		t.Errorf("FakeMoment() failed to fake time")
	}

	ftA = wallclock.Now()
	if !testMoment.Equal(ftA) {
		t.Errorf("FakeMoment() failed, returned other time then set")
	}
	cancelFunc()
	t2 = wallclock.Now()
	if t2.Equal(testMoment) {
		t.Errorf("FakeMoment() cancelFunc failed to return normal time")
	}
}

func TestFakeMomentDeferred(t *testing.T) {
	testMoment, _ := time.Parse(time.RFC3339, "1999-01-02T12:34:56.001Z")
	var t1, t2, ftA, ftB time.Time
	wrapDeferFunc := func() {
		t1 = wallclock.Now()
		defer wallclock.FakeMoment(testMoment)
		ftA = wallclock.Now()
		ftB = wallclock.Now()
	}
	wrapDeferFunc()
	t2 = wallclock.Now()

	if t1.Equal(t2) {
		t.Errorf("FakeMoment() failed after defer called")
	}
	if !ftA.Equal(ftB) {
		t.Errorf("FakeMoment() fake times not equal %v != %v", ftA.String(), ftB.String())
	}
}
