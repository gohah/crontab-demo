package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

//代表一个任务
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		now           time.Time
		exp           *cronexpr.Expression
		scheduleTable map[string]*CronJob
		cronJob       *CronJob
	)

	scheduleTable = make(map[string]*CronJob)
	//当前时间
	now = time.Now()
	//我们定义2个cronJob
	exp = cronexpr.MustParse("*/5 * * * * * *")

	cronJob = &CronJob{
		expr:     exp,
		nextTime: exp.Next(now),
	}

	//任务注册到调度表
	scheduleTable["job1"] = cronJob

	exp = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     exp,
		nextTime: exp.Next(now),
	}

	//任务注册到调度表
	scheduleTable["job2"] = cronJob

	//启动一个调度协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		//定时检查一下调度表
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				//判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					//启动一个协程执行任务
					go func(jobName string) {
						fmt.Println(jobName)
					}(jobName)

					//计算下次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次调度时间", cronJob.nextTime)
				}
			}

			//睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: //将在100毫秒刻度，返回
			}
			//睡眠100毫秒
			//time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(100 * time.Second)
}
