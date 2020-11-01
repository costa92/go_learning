package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓存可重用对象
}

func NewObjPool(numOfObj int) * ObjPool  {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj,numOfObj)
	for i:=0; i< numOfObj ; i++  {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}


func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj ,error)  { // 如果放任何对象 ReusableObj 改成 obj interface{}
	select {
	case ret := <-p.bufChan:
		return ret,nil
	case <- time.After(timeout): // 超时控制
		return nil,errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error  {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}