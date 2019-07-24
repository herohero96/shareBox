package main

import "fmt"

type Stroager interface {
	Read()
	Writer()
}


type Mdisk struct {

}

func (m *Mdisk) Read()  {
	fmt.Println("移动硬盘读取数据")
}
func (m *Mdisk) Writer()  {
	fmt.Println("移动硬盘写入数据")
}

type UDisk struct {

}

func (u *UDisk) Read()  {
	fmt.Println("U盘读取数据")
}

func (u *UDisk) Writer()  {
	fmt.Println("U盘写入数据")
}

func Computer(c Stroager)  {
	c.Read()
	c.Writer()
}

func main()  {
	var udisk UDisk
	var mdisk Mdisk
	Computer(&udisk)
	Computer(&mdisk)
}