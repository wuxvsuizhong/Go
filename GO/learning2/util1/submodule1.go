package util1

/*
普通的费大写字母开头的变量在package内可共享，也就是可以被同一个package内的其他文件访问
但是不能跨package被访问到
*/
var varInSubMod int = 100
