package DataPage

import (
	"fmt"
)

//数据页 双向链表
type DataPage struct {
	PageNumber int  //数据页编号
	Head *DataPage //下一个节点
	Tail *DataPage //上一个节点
	DataAll []interface{} //数据页里面的所有SQL记录
}

var datapage *DataPage
//获取记录
func (page *DataPage)Getrecords(s interface{}) {
    datapage = new(DataPage)
	datapage.DataAll = append(datapage.DataAll,s)
}

//四条记录为一组
func (page *DataPage)InsterDatall(){
	  head := &DataPage{}
	  datapage.PageNumber=1
	  page.GetList(head,datapage)
	  page.forxun(head)
}

func (page *DataPage)forxun(heap *DataPage)  {
	  temp := heap

	 for {
	 	fmt.Printf("[%d,%s]",temp.Head.PageNumber,temp.Head.DataAll)
	 	break
	 }
}

//创建双向链表
func (page *DataPage)GetList(head *DataPage,newHead *DataPage){
	temp := head
	flag := true

	for {
		//链表到最后
		if temp.Head == nil{
			break
		}else if temp.Head.PageNumber >= newHead.PageNumber {
			break
		}else if temp.Head.PageNumber == newHead.PageNumber {
			flag = false
			break
		}
		temp = temp.Head
	}

	if !flag{
		fmt.Println("该数据已经存在")
		return
	}else {
		//是为了key不是顺序插入而准备
		newHead.Head =temp.Head
		newHead.Tail = temp
		if temp.Head != nil{
			temp.Head.Tail = newHead
		}
		temp.Head = newHead
	}
}
