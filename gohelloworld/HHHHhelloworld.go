package gohelloworld //定义包名 必须在源文件的非注释的第一行致命这个文件属于那个
//package main 代表这个是一个可独立执行的程序,每个go应用程序都包含一个名为
//main 的包

/**
import 告诉go 编译器这个程序需要使用fmt包(里的函数,或者其他元素)
fmt包实现了格式化io的函数
 */
import "fmt"

/**
程序开始执行的函数,main函数是每一个可执行程序必须包含的,启动顺序:init() 然后是main()
如果没有init,首先执行main
 */

//func init()  {
//	fmt.Printf("this is init\n")
//}

func main()  {
 	strtest()
	fmt.Printf("hello world\n")
}

/**
标识符(包括常量,变量,类型,函数名,结构字段)以一个大写字母开头,对外部可见
如果以小写字母开头,则对保外不可见,但在整个包的内部是可见并且可用的(类似protected)

必须以字母或者是下划线开头,不能是数字
 */

 /**
 关键字
下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr
程序一般由关键字、常量、变量、运算符、类型和函数组成。
程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
程序中可能会使用到这些标点符号：.、,、;、: 和 …。
  */

  /**
  go 语言中的变量声明必须用空格隔开如:var age int;
   */
   var x int


   /**
   数据类型:
   1.布尔型
    */

    var t bool = true
    var f bool = false
    //2.数字类型

    var u uint;//unsigned int
    var uu uint8//无符号8位整型 0-255
    var nnn uint16 //16位整形
    var uuuu uint32
    var uuuuu uint64

    var i int8//有符号整型 8位,下面同理
    var ii int16
    var iii int32
    var iiii int64

    var n2 float32 = 1.0 //IEEE754 32位浮点数
    var n3 float64 = 2.0

    var com complex64 = 1+2i;//这是32位实数,还有64位实数

    //还有其他数字类型,比如:byte,rune(类似int32),uint(32位或64位)
    //int(和uint一样大小),uintptr(无符号整型,用于存放一个指针)
func numtest()  {

}
    //3.字符串类型(用utf8编码表示的unicode文本)
var string = "abcdefg"

func strtest()  {
	 print(string)
}

	//4.派生类型:
	/**
	包括:a,指针类型pointer
		b,数组类型
		c,结构化类型(struct)
		d,channel类型
		e,函数类型
		f,切片类型
		g接口类型(interface)
		h,map类型
	 */



	 //变量声明
	 var v int = 10
	 var vv = 10

	 var xxx,yyy int;
	 var (
	 	bbb bool
	 )

	 var ccc,ddd int = 1,2
	 var eee,fff  = 123, "asdf"

func vartest()  {

}

/**
值引用和引用类型:
所有相int,bool这种基本类型都属于值类型,使用这些类型的变量直接
指向存在内存中的值
在用=将一个变量赋值给另一个变量时,如:j=i 在内存中的i 的值进行了拷贝
可以通过&来获取内存地址
一个引用类型的变量r1存储的是r1的值的所在的内存地址,或内存地址中第一个字所在的位置

 */

/**
array(栈 值类型)======slice 不限长数组(引用类型)
struct 值类型
string 值类型
值类型 go强类型
 */

func givetest()  {

	var  xxx = ""
	fmt.Printf(xxx)


	//vvvva string = ""
}


//uintptr 保存指针的32位和64位