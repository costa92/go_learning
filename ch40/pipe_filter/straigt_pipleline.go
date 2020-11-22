package pipe_filter

import "fmt"

func NewStraightPipeline(name string,filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:name,
		Filters:&filters,
	}

}

type StraightPipeline struct {
	Name string
	Filters *[]Filter
}


func (f *StraightPipeline) Process(data Request) (Response,error){
	var ret interface{}
	var err error
	for _,filter := range *f.Filters {
		ret ,err = filter.Process(data)
		fmt.Println(ret)
		if err != nil {
			return ret,err
		}
		// 前面一个 filter 处理的 重新赋值给data
		data = ret
	}

   return ret,nil
}