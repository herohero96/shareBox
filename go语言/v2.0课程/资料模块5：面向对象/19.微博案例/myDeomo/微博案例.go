package main

func main()  {



}


type Blogger struct {
	Base
	WeiBos []*PostContent
	Comments map[int][]*PostContent
	Fans []FansInterfacer
}

func NewBlogger(name string) *Blogger  {
	blg := new(Blogger)
	blg.Name = name
	blog.Comments = make(map[int][]*PostContent)
	blg.WeiBos = make([]*PostContent, 0)
	return blg

}

