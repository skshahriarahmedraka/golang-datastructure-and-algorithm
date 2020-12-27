package main

import (
	"fmt"
)

func main(){
	var (
		t int
		n int
		singly_list Node
	)

	fmt.Printf("how many variable : ")
	fmt.Scanf("%d",&t)

	for i:=0;i<t;i++{
		fmt.Printf("give value : ")
		fmt.Scanf("%d",&n)
		Add_Node1(&singly_list,n)
	}
	singly_list=*singly_list.Address
	fmt.Println("singly list :- ")
	Value_List(&singly_list)

}

type Node struct{

		Value int
		Address *Node
}

func Add_Node1(n1 *Node,v int){
	if n1 ==nil{
		n1=new(Node)
		n1.Value=v
		// fmt.Println("first insert")
	} 
	
	for n1 !=nil{
		if n1 !=nil && n1.Address ==nil{
			n1.Address=new(Node)
			n1.Address.Value=v
			// fmt.Println("value list : ",n1.Address.Value," address : ",n1.Address.Address)

			return
		}
		// fmt.Println("value list : ",n1.Value," address : ",n1.Address)
		n1=n1.Address
	}
}


func Value_List(n1 *Node){
	if n1 ==nil{
		return
	}
	fmt.Println("Value : ",n1.Value," next address : ",n1.Address)
	Value_List(n1.Address)
}




























func  Add_Node(n1 *Node,v int) {
	fmt.Println(n1.Value," ",n1.Address)
	if n1.Value ==0 {
		n1.Value=v
		fmt.Println("first insert ",n1.Value," ",n1.Address)
		
		return
	}else if n1.Address==nil {
		n1.Address=new(Node)
		n1.Address.Value=v
		fmt.Println("first insert ",n1.Value," ",n1.Address)

		fmt.Println("2nd insert ",n1.Address.Value," ",n1.Address.Address)
		
		return
	}
	p1:=n1
	for i:=1; ; i++ {

		if p1.Value!=0 && p1.Address==nil{
			p1.Address=new(Node)
			p1.Address.Value=v
			fmt.Println(i," insert complete : ",p1.Address.Value," ",p1.Address.Address)
			

			break
		}else{
			fmt.Println(p1.Value," ",p1.Address)
			fmt.Println(i," insert was : ",p1.Value," ",p1.Address)


			p1=p1.Address
		}

	}
}