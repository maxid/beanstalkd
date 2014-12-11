package beanstalkd

import (
	"testing"
	"time"
)

func getPool() *BeanstalkdPool {
	pool := NewBeanstalkdPool(func(tube string) (Client, error) {
		queue, err := Dial("127.0.0.1:11300")
		if err != nil {
			return nil, err
		}
		queue.Use(tube)
		_, err = queue.Watch(tube)
		if err != nil {
			return nil, err
		}
		queue.Ignore("default")
		return queue, nil
	}, 30)
	return pool
}

func TestPool(t *testing.T) {
	pool := getPool()
	client := pool.Get("TEST")
	t.Logf("%#v", client)
	client.Put(1, 0*time.Second, 5*time.Second, []byte("hello world!"))
	job, _ := client.Reserve(60)
	t.Log(string(job.Data))
	client.Quit()
	client = pool.Get("TEST")
	t.Logf("%#v", client)
	client = pool.Get("TEST")
	t.Logf("%#v", client)
}
