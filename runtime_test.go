package faaslib_go_handler

import (
	"log"
	"testing"
	"time"
)

//func TestAbs(t *testing.T) {
//	got := math.Abs(-1)
//	if got != 1 {
//		t.Errorf("Abs(-1) = %03.6f; want 1", got)
//	}
//}
//
//func BenchmarkRandInt(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		rand.Int()
//	}
//}

//func BenchmarkLogging(b *testing.B) {
//	defer Inject(time.Now(), "BOB")
//	x := shortuuid.New()
//	println(x)
//}
//
//func BenchmarkLogging2(b *testing.B) {
//	defer Inject(time.Now(), "BOB")
//	x := xid.New()
//	println(x.String())
//}

//func TestLogging2(t *testing.T) {
//	defer Inject(time.Now(), "BOB")
//	log.Print("Here")
//	CallHomeStart()
//	log.Print("There")
//}
//
func TestLogging3(t *testing.T) {
	defer InjectCallback("2345", time.Now(), CallHomeStart("2345"))
	log.Print("here")
	log.Print("now")
}
