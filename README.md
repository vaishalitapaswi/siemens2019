package main

import "fmt"

func main() {
	e := &Emp{Salary: 1111.1, Empno: 10, Ename: "AAA"}
	fmt.Println(*e)
	e1 := Emp{Salary: 1111.1, Empno: 10, Ename: "AAA"}
	fmt.Println(&e1)

	empmgr := EmpMgr{}
	empmgr.add(*e)
	fmt.Println(empmgr)
	empmgr.add(e1)
	fmt.Println(empmgr)
}

type Emp struct {
	Empno  int
	Ename  string
	Salary float32
}
type EmpMgr struct {
	emparr [5]Emp
}

func (mgr EmpMgr) add(e Emp) {
	fmt.Println(" in add", e)
	mgr.emparr[0] = e
}
