// client_test.go
package beanstalkd

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// dial & put & reserve & delete & quit
func TestPut(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go consumer(0, t)
	go producer(1, t)
	time.Sleep(1 * time.Second)
}

// put & reserve-with-timeout & delete & quit
func TestReserve(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go consumer(10, t)
	go producer(100, t)
	time.Sleep(15 * time.Second)
}

// consumer
func consumer(timeout int, t *testing.T) {
	queue, err := Dial("127.0.0.1:11300")
	if err != nil {
		t.Fatal(err)
		return
	}
	for {
		job, err := queue.Reserve(timeout)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("reserve", job.Id, string(job.Data))
		err = queue.Delete(job.Id)
		if err != nil {
			t.Fatal(err)
			break
		
	}
	err = queue.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

// producer
func producer(size int, t *testing.T) {
	queue, err := Dial("127.0.0.1:11300")
	if err != nil {
		t.Fatal(err)
		return
	}
	for i := 0; i < size; i++ {
		_, err := queue.Put(1, 0*time.Second, 5*time.Second, []byte("test "+strconv.Itoa(i)))
		if err != nil {
			t.Fatal(err)
			break
		}
		fmt.Println("put test " + strconv.Itoa(i))
	}
	err = queue.Quit()
	if err != nil {
		t.Fatal(err)
	}
}
