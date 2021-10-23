/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/17 14:31
 * @version     v1.0
 * @filename    constant.go
 * @description
 ***************************************************************************/
package router

import (
	"sync"
)

var cst *Constant
var once sync.Once

func init() {
	once.Do(func() {
		cst = &Constant{}
		initInstance()
	})
}

func GetInstanceOfConstant() *Constant {
	return cst
}

type Constant struct {
	mux      sync.Mutex
	Website1 Website
}

type Website struct {
	Name    string
	Path    string
	Module1 Module
	Module2 Module
	Module4 Module
}

type Module struct {
	Name      string
	Path      string
	Function1 Function
	Function2 Function
}

type Function struct {
	Name string
	Path string
}

func initInstance() {
	cst := GetInstanceOfConstant()
	cst.Website1 = Website{
		Name: "Endangered Plants",
		Path: "/endangeredplants",
		Module1: Module{
			Name: "Plant Hub",
			Path: "/planthub",
			Function1: Function{
				Name: "Data Source",
				Path: "/datasource",
			},
		},
		Module2: Module{
			Name: "Stat Hub",
			Path: "/stathub",
			Function1: Function{
				Name: "Data Source",
				Path: "/datasource",
			},
		},
		Module4: Module{
			Name: "Blog",
			Path: "/blog",
			Function1: Function{
				Name: "Blog List",
				Path: "/bloglist",
			},
			Function2: Function{
				Name: "Article",
				Path: "/article",
			},
		},
	}
}
