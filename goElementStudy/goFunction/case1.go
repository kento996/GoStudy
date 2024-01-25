package main

func main1() {
	c := 2
	dosometing1(&c)
	println(&c)
	dosomething2(c)

}

// 这里函数接收一个类型为 *int的参数，即int类型的指针
// 在函数内部，变量b被赋值为a，这意味着b也是一个int类型的指针，指向与a一样的内存地址，也就是c所在的地址
func dosometing1(a *int) {
	b := a
	println(*b)
}

// 这里函数接收一个类型为int的参数
// 在函数内部，变量b被赋值为&a，这是a的地址，所以b是一个指向a的指针
// 重要的是，这里的a是通过值传递的，这意味着a是原始数据的副本。因此，b指向的是这个副本的地址
// 而不是原始数据的地址
// 因此，函数外对原始数据的任何修改不会影响在函数内部创建的a的副本
func dosomething2(a int) {
	b := &a
	println(*b)
}

//总结：
//1. dosomething1通过指针传递，允许函数直接修改原始数据
//2. dosomething2通过值传递，创建原始数据的副本，并且函数内部的修改不回影响原始数据
