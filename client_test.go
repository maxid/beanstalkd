// client_test.go
package beanstalkd

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// dial & list-tubes & quit
func TestListTubes(t *testing.T) {
	queue := dial(t)
	tubes, err := queue.ListTubes()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%#v", tubes))
	queue.Quit()
}

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
	queue := dial(t)
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
	}
	err := queue.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

// producer
func producer(size int, t *testing.T) {
	queue := dial(t)
	for i := 0; i < size; i++ {
		_, err := queue.Put(1, 0*time.Second, 5*time.Second, []byte("hello world "+strconv.Itoa(i)))
		if err != nil {
			t.Fatal(err)
			break
		}
		fmt.Println("put hello world " + strconv.Itoa(i))
	}
	err := queue.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

// dial
func dial(t *testing.T) *BeanstalkdClient {
	queue, err := Dial("127.0.0.1:11300")
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return queue
}
