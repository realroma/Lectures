package main

import (
	"fmt"
	"project/Lectures/Lesson5/pkg/filer"
	"project/Lectures/Lesson5/pkg/index"
	"project/Lectures/Lesson5/pkg/lecture_flag"
	"project/Lectures/Lesson5/pkg/mapper"
	"project/Lectures/Lesson5/pkg/sorter"
)

func main() {
	//Делаем флаг для поиска по словам в консоли.
	word := lecture_flag.ParseFlag()

	m := mapper.Scan()

	//Объединяем ссылки в единое.
	name := "Link.txt"
	filer.New(name)

	mapper.MapWrite(m, name)

	ma := mapper.MapRead(name)
	fmt.Println(ma)

	itd := index.Indexer(m)

	sorter.Sorter(m, word)
	if word != "" {
		index.Read(word, itd)
	}
}
