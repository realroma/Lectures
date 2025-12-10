package index

import (
	"fmt"
	"project/Lectures/Lesson5/pkg/sorter"
)

type Index struct {
	ID    int
	URL   string
	Title string
}

func SerializeIndex(Url string, Title string, id int) Index {
	//Если список проиндексированных ссылок не пустой, добавляем в конец.
	id++
	//Задаём структуру добавляляемых в конец структур.
	c := Index{
		ID:    id,
		URL:   Url,
		Title: Title,
	}
	fmt.Println(c)

	//Файл постоянно перезаписывается. Индекс всегда нулевой. По итогу возвращается ноль.
	return c
}

// Проходит по всем структурам массива вынимая название и разбирая его на слова, возвращая map слов и их Id.
func add(mapWord map[string][]int, dictIndex []Index) map[string][]int {
	for _, v := range dictIndex {
		arr := sorter.Sort(v.Title)
		for _, a := range arr {
			if len(a) > 1 {
				var ids []int
				ids = append(ids, v.ID)
				mapWord[a] = ids
			}
		}
	}
	return mapWord
}

func Read(w string, mW map[string][]int) {
	dict := make([]int, 0)

	arr := sorter.Sort(w)
	for _, a := range arr {
		fmt.Println("A, in read: ", a)
		v, ok := mW[a]
		fmt.Println(ok)
		if ok {
			fmt.Println(cap(dict))
			if cap(dict) == 0 {
				dict = v
			}
			if cap(dict) <= len(v) {
				dict = sorter.Elements(dict, mW[a])
				fmt.Println("Dict: ", dict)
			}
		}
	}
}

func Indexer(m map[string]string) map[string][]int {
	mapWord := make(map[string][]int)
	dictIndex := make([]Index, 0)

	//Проходим по всем ссылкам и разбираем их на переменные.
	for Url, Title := range m {
		dictIndex = append(dictIndex, SerializeIndex(Url, Title, len(dictIndex)))
		//Добавляем слова из ссылок в Map для индексации.
		mapWord = add(mapWord, dictIndex)
	}
	// for i, v := range mapWord {
	// 	fmt.Println("Id cm", i)
	// 	fmt.Println("Result Indexer:", v)
	// }
	// fmt.Println(dictIndex)
	return mapWord
}
