package main

import "fmt"


type Person struct {
	name string
	phone map[string]string
}

// 通讯录

var PersonList []Person

func main()  {
	ScanNum()
}

//增加通讯录
func AddPersonList()  {
	var person Person
	var name string
	phone := make(map[string]string, 10)
	var phoneType string
	var phoneNumber string
	var isQ string
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	for {
		fmt.Println("请输入电话类型")
		fmt.Scan(&phoneType)
		fmt.Println("请输入电话")
		fmt.Scan(&phoneNumber)
		phone[phoneType] = phoneNumber
		fmt.Println("结束输入请输入Q")
		fmt.Scan(&isQ)
		if (isQ == "Q") {
			break
		}
	}
	person = Person{name,phone }
	PersonList = append(PersonList, person)
	ShowPersonList()
	ScanNum()
}

// 删除通讯录
func DeletePersonList()  {
	var name string

	if len(PersonList) == 0 {
		fmt.Println("通讯录为空，没有可以删除的名单")
	} else {
		fmt.Println("请输入要删除的通讯录姓名")
		fmt.Scan(&name)
		for k, v := range PersonList{
			if v.name == name {
				//if len(PersonList) == 1 {
				//	PersonList = []Person{}
				//} else {
				//	if k == 0 {
				//		PersonList = PersonList[k+1:]
				//	} else if k == len(PersonList) - 1 {
				//		PersonList = PersonList[:k]
				//	} else {
				//		PersonList = append(PersonList[:k], PersonList[(k+1)])
				//	}
				//}
				PersonList = append(PersonList[0:k], PersonList[k+1:]...) // append函数第二个参数如果是切片后面要给三个点。

			}
		}
	}

	ShowPersonList()
	ScanNum()
}

// 编辑通讯录
func EditPersonList()  {
	var name string
	var phoneType string
	var phoneNumber string
	var isQ string
	fmt.Println("请输入要编辑的通讯录的姓名")
	fmt.Scan(&name)
	if len(PersonList) == 0 {
		fmt.Println("通讯录为空，没有可以删除的名单")
	}else {
		for i := 0; i <len(PersonList) ; i ++ {
			if name == PersonList[i].name {
				PersonList[i].phone = map[string]string{}
				for {
					fmt.Println("请输入电话类型")
					fmt.Scan(&phoneType)
					fmt.Println("请输入电话")
					fmt.Scan(&phoneNumber)
					PersonList[i].phone[phoneType] = phoneNumber
					fmt.Println("结束输入请输入Q")
					fmt.Scan(&isQ)
					if (isQ == "Q") {
						break
					}
				}

			}
		}

	}

	ShowPersonList()
	ScanNum()

}

// 查询通讯录

func QueryPeronsList()  {
	var name string
	if len(PersonList) == 0 {
		fmt.Println("通讯录为空")
		return
	}
	fmt.Println("请输入要查询的通讯录的姓名")
	fmt.Scan(&name)
	for _, v := range PersonList {
		if name == v.name {
			fmt.Printf("\n姓名：%s\n", v.name)
			for k, val := range v.phone {
				fmt.Printf("电话类型%s:%s", k, val)
			}
		}
	}

	ScanNum()

}

// 展示通讯录
func ShowPersonList()  {
	fmt.Println( "PersonList：" ,PersonList)
	if len(PersonList) == 0 {
		fmt.Println("通讯录为空")
		return
	}
	for _, v := range PersonList {
		fmt.Printf("\n姓名：%s\n", v.name)
		for k, val := range v.phone {
			fmt.Println("电话类型：", k)
			fmt.Println("电话：", val)
		}
	}
}


//第一次输入数字
func ScanNum() {
	fmt.Println("添加通讯录，请输入1")
	fmt.Println("删除通讯录，请输入2")
	fmt.Println("编辑通讯录，请输入3")
	fmt.Println("查询通讯录，请输入4")
	var n int
	fmt.Scan(&n)
	Action(n)
}

// 根据输入的数字进行操作
func Action( n  int )  {
	switch n {
		case 1: AddPersonList()
		case 2: DeletePersonList()
		case 3: EditPersonList()
		case 4: QueryPeronsList()
	}
}



