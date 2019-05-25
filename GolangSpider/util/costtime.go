package util

import (
	"fmt"
	"time"
)

//Cost
type Cost struct {
	Start time.Time
	End  	time.Time
}

func NewCost(start time.Time) *Cost {
	return &Cost{Start: start}
}

func (this *Cost) SetEndTime(end time.Time)  {
		this.End=end
}


//设置终止时间
func (this *Cost) CalcCost() time.Duration {
		//return this.
	duration := this.End.Sub(this.Start)
	//return fmt.Sprintf("%+v",duration)
	return duration
}
//设置终止时间
func (this *Cost) ShowCalcCostAsString() string {
		//return this.
	return fmt.Sprintf("%+v",this.CalcCost())
}

//不用设置终止时间
func (this *Cost) CostWithNow() time.Duration {
	//return this.
	duration := time.Since(this.Start)
	//return fmt.Sprintf("%+v",duration)
	return duration
}

func (this *Cost) CostWithNowAsString() string {
		//return this.
	return fmt.Sprintf("%+v",this.CostWithNow())
}


