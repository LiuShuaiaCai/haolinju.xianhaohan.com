package core

import (
	"haolinju.xianhaohan.com/internal/app/services"
)

type Job struct {
	Svc *services.Service
}

// TODO::备用
func NewJob(svc *services.Service) (job *Job, closeFunc func(), err error) {
	job = &Job{
		Svc: svc,
	}
	//err = s.AdvertSyncReceive()
	//if err != nil {
	//	panic(err)
	//}

	return
}
