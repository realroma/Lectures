package main

import (
	"fmt"
	"project/Lectures/Lesson5/pkg/index"
	"project/Lectures/Lesson5/pkg/lecture_flag"
	"project/Lectures/Lesson5/pkg/mapper"
	"project/Lectures/Lesson5/pkg/sorter"
)

func main() {
	//Делаем флаг для поиска по словам в консоли.
	word := lecture_flag.ParseFlag()

	//Объединяем ссылки в единое.
	name := "Link.txt"
	fmt.Println("Check")
	mapper.ChangeDir()
	m := mapper.ChekDataFile(name)
	fmt.Println(m)

	mapper.MapWrite(m, name)

	itd := index.Indexer(m)

	sorter.Sorter(m, word)
	if word != "" {
		index.Read(word, itd)
	}
}
