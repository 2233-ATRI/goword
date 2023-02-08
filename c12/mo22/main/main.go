package main

//双向链表
import "fmt"

type heronode struct {
	no       int
	name     string
	nickname string
	next     *heronode
}

func inserhero(x *heronode, newheronode *heronode) {
	heee := x
	for {
		if heee == nil {
			break
		}
		x = x.next
	}
	x.next = newheronode
}

func Modifyhero(head *heronode, id int) {
	heee := head
	flag := false
	for {
		if heee.next == nil {
			break
		} else if heee.next.no == id {
			flag = true
			break
		}
		heee = heee.next
	}
	var a string
	var b string
	if flag {
		fmt.Scanln(&a)
		heee.name = a
		fmt.Scanln(&b)
		heee.nickname = b
	}
}

// 插入可以根据no大小进行适当的插入
func inserhero2(x *heronode, newheronode *heronode) {
	heee := x
	flag := true
	for {
		if heee.next == nil {
			break
		} else if heee.next.no > newheronode.no {
			break
		} else if heee.next.no == newheronode.no {
			fmt.Println("错误")
			flag = false
			break
		}
		x = x.next
	}

	if !flag {
		fmt.Println("已经存在", newheronode.no)
		return
	} else {
		newheronode = x.next
		x.next = newheronode
	}

}

func Delhernode(head *heronode, id int) {
	heee := head
	flag := false
	for {
		if heee.next == nil {
			break
		} else if heee.next.no == id {
			flag = true
			break
		}
		heee = heee.next
	}
	if flag {
		heee.next = heee.next.next
	} else {
		fmt.Println("sorry dont id")
	}
}

func listheronode(x *heronode) {
	temp := x
	if temp.next == nil {
		fmt.Println("nothing")
		return
	}
	for {
		fmt.Printf("[%d ,%s %s]==>", temp.next.no, temp.next.name,
			temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}

}

func main() {
	herd := &heronode{}
	herd1 := &heronode{
		no:       1,
		name:     "aaww",
		nickname: "wwww",
	}

	inserhero(herd, herd1)
	//inserhero2(herd,herd1)
	//Delhernode(herd,2)
	listheronode(herd)
}
