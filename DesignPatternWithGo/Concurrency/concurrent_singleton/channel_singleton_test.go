package concurrent_singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T){
	singleTon := GetInstance()
	singleTon2 := GetInstance()

	n := 5000

	for i:=0;i<n;i++{
		go singleTon.AddOne()
		go singleTon2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleTon.GetCount())

	var val int

	for val != n*2{
		val = singleTon.GetCount()
		time.Sleep(10*time.Millisecond)
	}
	fmt.Printf("After loop, current count is %d\n", singleTon.GetCount())
	singleTon.Stop()
}
