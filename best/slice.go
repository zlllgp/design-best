package best

func Slice() {
	//Good
	list := make([]*int, 100)
	for i := 0; i < 100; i++ {
		//a := i
		list[i] = &i
	}

	//Bad
	newList := make([]*int, 0, 100)
	for _, v := range list {
		newList = append(newList, v)
	}
}
