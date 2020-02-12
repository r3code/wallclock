package wallclock_test

import (
	"testing"
	"time"

	"github.com/r3code/wallclock"
)

var dontOptimizeMePlease time.Time

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		dontOptimizeMePlease = time.Now()
	}
}

func BenchmarkWallclockNow(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		dontOptimizeMePlease = wallclock.Now()
	}
}    


func BenchmarkWallclockFakeMoment(b *testing.B) {             
  timeMoment, _ := time.Parse(time.RFC3339, "1999-01-02T12:34:56.001Z")	   
  cancelFunc := wallclock.FakeMoment(timeMoment) 
  defer cancelFunc()
  for i := 0; i <= b.N; i++ {                                           
		dontOptimizeMePlease = wallclock.Now()            
	}      
  
} 

func BenchmarkWallclockFakeFixedMoment(b *testing.B) {	
  cancelFunc := wallclock.FakeFixedMoment()
  defer cancelFunc()
  for i := 0; i <= b.N; i++ {           
		dontOptimizeMePlease = wallclock.Now()      
	}      
}  

