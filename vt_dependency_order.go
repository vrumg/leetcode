package leetcode

/*
Печать зависимостей в порядке импорта

условие:
Есть мапа библиотек и их зависимостей. Надо распечатать библиотеки в порядке правильного импорта.
Неправильный порядок - если печатаешь библиотеку, а ее зависимость еще не напечатана.
Все остальные порядки правильные
NB: Циклических зависимостей во входных данных нет

пример:
deps = {
    "tensorflow": ["nvcc", "gpu", "linux"],
    "nvcc": ["linux"],
    "linux": ["core"],
    "mylib": ["tensorflow"],
    "mylib2": ["requests"]
}

core linux nvcc gpu tensorflow mylib requests mylib2
*/

type Node struct {
	prev *Node
	next *Node
	key  string
}

func dependencyOrder(deps map[string][]string) string {
	var head *Node

	// fill first
	for k, depsInner := range deps {
		insertNodes(head, k, depsInner)
	}

	return ""
}

func insertNodes(head *Node, key string, deps []string) {
	var prevDep *Node

	for _, dep := range deps {
		if head == nil {
			head = &Node{
				prev: nil,
				next: nil,
				key:  dep,
			}
			prevDep = head
		}

		newDep := &Node{
			prev: prevDep,
			next: nil,
			key:  dep,
		}

		prevDep.next = newDep
		prevDep = newDep
	}

	newDep := &Node{
		prev: prevDep,
		next: nil,
		key:  key,
	}
	prevDep.next = newDep
}
