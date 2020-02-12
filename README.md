# `wallclock`

[![Build Status](https://travis-ci.org/r3code/wallclock.svg?branch=master)](https://travis-ci.org/r3code/wallclock)
[![Coverage Status](https://coveralls.io/repos/github/r3code/wallclock/badge.svg?branch=master)](https://coveralls.io/github/r3code/wallclock?branch=master)
[![GoDoc](https://godoc.org/github.com/r3code/wallclock?status.svg)](https://godoc.org/github.com/r3code/wallclock)

`wallclock` Golang is a package to replace go `time` package allowing you to mock the time during test.

## Usage 

Just replace your `time.Now()` by a `wallclock.Now()` call, etc.

## Mocking 

Use `wallclock.FakeMoment(timeMoment)` or `wallclock.FakeFixedMoment()` to replace the current `wallclock.Now()` time reading function, and use the function it returns to restore the realtime. 


## Notice                                                                                                   

Package changing its internal global variable so you can't fake the time from several tests at the same time. 


### Usage

```go
func TestWithFakeTime(t *testing.T) {
	testMoment, _ := time.Parse(time.RFC3339, "1999-01-02T12:34:56.001Z")	
	t1 := wallclock.Now()
	cancelFunc := wallclock.FakeMoment(testMoment) 
  defer cancelFunc()
	fmt.Printf("%+v") // Outputs: 1999-01-02T12:34:56.001Z	
}
```

## Drawbacks

### Performance

The indirection of the call makes `wallclock.Now` actually little slower.   
It's not easy to get a stable result because the difference is small.

```
> go test -run=NONE -benchmem -benchtime=5s -bench=. .
goos: windows
goarch: 386
pkg: github.com/r3code/wallclock
BenchmarkTimeNow-4                      234912168               25.5 ns/op             0 B/op
0 allocs/op
BenchmarkWallclockNow-4                 142611286               42.1 ns/op             0 B/op
0 allocs/op
BenchmarkWallclockFakeMoment-4          447401670               13.4 ns/op             0 B/op
0 allocs/op
BenchmarkWallclockFakeFixedMoment-4     468723194               12.8 ns/op             0 B/op
0 allocs/op
PASS
ok      github.com/r3code/wallclock     33.543s
```      

## Faking time practice

Actually 
> Changing the system time while making tests (or in general) is a bad idea. 
> You don't know what depends on the system time while executing tests and you don't want to find out the hard way by spending days of debugging into that. Just don't do it." 
by [nemo](https://stackoverflow.com/a/18970352/469898)

*But sometimes we need to do it*.  
