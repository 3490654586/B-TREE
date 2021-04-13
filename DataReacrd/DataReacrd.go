package DataReacrd

import (
	"Btree/DataPage"
	InterfaceStuct "Btree/Interface"
	"fmt"
)


var (
	DataPageNum = 0  //数据页的个数
	Stor = 0  //页目录
)

//自身数据,类似与SQL的一条记录 这里是单向链表
type DataRecord struct {
	Key  int
	Next *DataRecord
	Data interface{}
	Datapage *DataPage.DataPage
	pace InterfaceStuct.DatapageInster
}


//页目录(页面槽) 用于存放数据页的上一级几点
type PageContents []int


//插入数据
func (data *DataRecord)InsertNode(head *DataRecord ,newDataRecord *DataRecord){
	temp := head
	flag := true

	sint := data.Lenth(temp)
	if sint > 3 {
		AllReacrd :=data.QuerNode(temp)
		 data.Pr(&AllReacrd)
	    	data.InsterDatall()
	}

	for {
		//链表到最后
		if temp.Next == nil{
			break
		}else if temp.Next.Key >= newDataRecord.Key {
			//由于这里是模仿SQL主键进行插入所以是有序链表,普通链表一般没有这个判断
			break
		}else if temp.Next.Key == newDataRecord.Key {
			flag = false
			break
		}
		temp = temp.Next
	}

	if !flag{
		fmt.Println("该数据已经存在")
		return
	}else {
		//是为了key不是顺序插入而准备
		newDataRecord.Next =temp.Next

		temp.Next = newDataRecord
	}
}

//查询链表长度
func (data *DataRecord) Lenth(head *DataRecord)int{
	  temp := head
	  count := 0
	  for {
	  	if temp.Next == nil {
	  		break
		}
		  count++
		  temp = temp.Next
	  }

	  return count
}

//查询链表 四条记录为一个数据页
func (data *DataRecord)QuerNode(head *DataRecord)[]DataRecord{
      temp := head
      recroc  := make([]DataRecord,4)

      for {
		  recroc = append(recroc,*temp.Next)
		  //每四条记录为一组再次生成一个数据页
		  if len(recroc) >4 {
			  recroc = make([]DataRecord,4)
		  }

		  if temp.Next ==nil{
		  	fmt.Println("切片分组完成")
		  	break
		  }
	  }

	  return recroc
}


func (page *DataRecord)Pr(dataPage *[]DataRecord){
	page.pace = page.Datapage
	page.pace.Getrecords(dataPage)
}

func (page *DataRecord)InsterDatall(){
	page.pace = page.Datapage
	page.pace.InsterDatall()
}