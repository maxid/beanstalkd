// tube.go
package beanstalkd

type BeanstalkdTube struct {
	clients   []*BeanstalkdClient //
	maxActive int                 //
	useTime   int                 //
	idleTime  int                 //
	addr      string              //
	tube      string              //
}

var (
	DEFAULT_MAX_ACTIVE = 30             // default pool size
	DEFAULT_USE_TIME   = 20 * 60 * 1000 // default checkout time is 20 minutes
	DEFAULT_IDEL_TIME  = 20 * 60 * 1000 // connection will be removed after not use
)

func NewBeanstalkdTube1(addr, tube string) *BeanstalkdTube {
	return &BeanstalkdTube{
		clients:   make([]*BeanstalkdClient, DEFAULT_MAX_ACTIVE),
		maxActive: DEFAULT_MAX_ACTIVE,
		useTime:   DEFAULT_USE_TIME,
		idleTime:  DEFAULT_IDEL_TIME,
		addr:      addr,
		tube:      tube,
	}
}

func NewBeanstalkdTube2(addr, tube string, poolSize, useTime, idleTime int) *BeanstalkdTube {
	return &BeanstalkdTube{
		clients:   make([]*BeanstalkdClient, poolSize),
		maxActive: poolSize,
		useTime:   useTime,
		idleTime:  idleTime,
		addr:      addr,
		tube:      tube,
	}
}
