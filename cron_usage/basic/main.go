package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		err      error
		exp      *cronexpr.Expression
		now      time.Time
		nextTime time.Time
	)
	//Field name     Mandatory?   Allowed values    Allowed special characters
	//----------     ----------   --------------    --------------------------
	//Seconds        No           0-59              * / , -
	//Minutes        Yes          0-59              * / , -
	//Hours          Yes          0-23              * / , -
	//Day of month   Yes          1-31              * / , - L W
	//Month          Yes          1-12 or JAN-DEC   * / , -
	//Day of week    Yes          0-6 or SUN-SAT    * / , - L #
	//Year           No           1970–2099         * / , -
	if exp, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	//当前时间
	now = time.Now()
	//下次调度时间
	nextTime = exp.Next(now)
	fmt.Println(now, nextTime)
	//等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("调度了...")
	})

	time.Sleep(10 * time.Second)
}
