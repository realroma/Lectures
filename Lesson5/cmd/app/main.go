package main

import (
	"project/Lectures/Lesson5/pkg/filer"
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
	m := make(map[string]string)

	if name != "" {
		m = mapper.MapRead(name)
	} else {
		name := "Link.txt"
		filer.New(name)
		m = mapper.Scan()
	}

	mapper.MapWrite(m, name)

	itd := index.Indexer(m)

	sorter.Sorter(m, word)
	if word != "" {
		index.Read(word, itd)
	}
}
