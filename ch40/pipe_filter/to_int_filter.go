package pipe_filter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")
type ToIntFiler struct {

}

func NewToIntFilter() *ToIntFiler{
	return  &ToIntFiler{}
}

func (tif *ToIntFiler) Process(data Request) (Response,error)  {
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
	parts,ok := data.([]string)
	if !ok {
		return nil,ToIntFilterWrongFormatError
	}
	var ret []int
	for _,part := range parts{
		s,err := strconv.Atoi(part)
		if err != nil {
			return nil,err
		}
		ret = append(ret,s)
	}

	return  ret,nil
}