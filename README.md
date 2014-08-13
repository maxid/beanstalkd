beanstalkd
==========

beanstalkd golang client


## Implemented Commands

Producer commands:

* use
* put

Worker commands:

* reserve
* delete
* release
* bury
* touch
* watch
* ignore

Other commands:

* peek
* peek-ready
* peek-delayed
* peek-buried
* kick
* kick-job
* stats-job
* stats-tube
* stats
* list-tubes
* list-tube-used
* list-tubes-watched
* quit
* pause-tube


# Release Notes
Latest release is v1.0 that contains API changes, see release notes [here](https://github.com/maxid/beanstalkd/blob/master/ReleaseNotes.txt)

## Author

* [Maxid Tseng](http://uscan.cn)
