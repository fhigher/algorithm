package main

import (
	"fmt"
	"strings"
)

type PersonNode struct {
	Name string
	Age int
	Next *PersonNode
}

func NewPersonNode(name string, age int) *PersonNode {
	return &PersonNode{
		Name: name,
		Age: age,
	}
}

type SingleList struct {
	Head *PersonNode
	Tail *PersonNode
	Length int
}

// 无序尾部插入
func (list *SingleList) DisorderTailInsert(node *PersonNode) {
	if list.Head == nil {
		list.Head = node	// 第一个结点 既是头结点，也是尾结点
		list.Tail = node
	} else {
		list.Tail.Next = node 	// 之后的结点往尾结点上加
		// list.Tail = node
		list.Tail = list.Tail.Next 	// 更新尾结点
	}

	list.Length++
}

func (list *SingleList) Find(name string) *PersonNode {
	if list.Head == nil {
		return nil
	}

	temp := list.Head
	for {
		if 0 == strings.Compare(temp.Name, name) {
			return temp
		}
		temp = temp.Next
		if temp == nil {
			break
		}
	}

	return nil
}

func (list *SingleList) Delete(name string) (popNode *PersonNode) {
	if list.Head == nil {
		return nil
	}

	if 0 == strings.Compare(list.Head.Name, name) {
		// 头结点就是要删除的结点
		popNode = list.Head
		list.Head = list.Head.Next
		list.Length--
		return
	}

	// 删除其他结点
	temp := &PersonNode{}	// temp 指向一个空结点
	temp.Next = list.Head	// temp.Next 指向头结点
	for {
		if 0 != strings.Compare(temp.Next.Name, name) {	// temp始终是正在比较结点的前一个结点
			temp = temp.Next	// temp后移
		} else {
			// 找到要删除的结点，先保存
			popNode = temp.Next
			// 删除
			temp.Next = temp.Next.Next
			list.Length--
			return
		}

		if temp.Next == nil {
			break
		}
	}

	return nil
}

func (list *SingleList) ShowList() {
	if list.Head == nil {
		fmt.Println("empty list")
		return
	}
	fmt.Println("链表长度为：", list.Length)
	temp := list.Head
	for {
		fmt.Printf("[%s(%d)] --> ", temp.Name, temp.Age)
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	fmt.Println()
	fmt.Println()
}

func NewSingleList() *SingleList {
	return &SingleList{}
}

func main() {

	var singleList *SingleList

	singleList = NewSingleList()

	fmt.Println("添加结点：")
	singleList.DisorderTailInsert(NewPersonNode("张三", 22))
	singleList.DisorderTailInsert(NewPersonNode("李四", 19))
	singleList.DisorderTailInsert(NewPersonNode("王天", 33))
	singleList.DisorderTailInsert(NewPersonNode("玉玉", 10))
	singleList.DisorderTailInsert(NewPersonNode("蛋蛋", 20))

	singleList.ShowList()

	fmt.Println("查找结点：")
	person := singleList.Find("李四")
	unknown := singleList.Find("张飒")
	fmt.Println(person, unknown)
	fmt.Println()
	fmt.Println("删除结点：")
	// 删除第一个结点
	firstPerson := singleList.Delete("张三")
	fmt.Println("被删除的头结点", firstPerson)
	singleList.ShowList()

	// 删除尾结点
	tailPerson := singleList.Delete("蛋蛋")
	fmt.Println("被删除的尾结点", tailPerson)
	singleList.ShowList()

	// 删除中间的某个结点
	midPerson := singleList.Delete("王天")
	fmt.Println("被删除的中间某结点", midPerson)
	singleList.ShowList()

	// 删除不存在的结点
	unknownPerson := singleList.Delete("无名")
	fmt.Println("删除不存在的结点", unknownPerson)
	singleList.ShowList()
}


