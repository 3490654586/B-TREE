package main

import (
	"Btree/DataReacrd"
	"fmt"
)

func main()  {
	Data:=  new(DataReacrd.DataRecord)
     head := &DataReacrd.DataRecord{}

     head1 := &DataReacrd.DataRecord{
		 Key:      1,
		 Data:     "张",
	 }

	head2 := &DataReacrd.DataRecord{
		Key:      2,
		Data:     "冲",
	}
	head3 := &DataReacrd.DataRecord{
		Key:      2,
		Data:     "冲",
	}
	head4 := &DataReacrd.DataRecord{
		Key:      2,
		Data:     "冲",
	}
	head5 := &DataReacrd.DataRecord{
		Key:      2,
		Data:     "冲",
	}
	Data.InsertNode(head,head1)
	Data.InsertNode(head,head2)
	Data.InsertNode(head,head3)
	Data.InsertNode(head,head4)
	Data.InsertNode(head,head5)
  fmt.Println(Data.Lenth(head))

}