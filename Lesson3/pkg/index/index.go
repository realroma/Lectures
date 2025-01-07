package index

import (
	"fmt"
	"project/Lectures/Lesson3/pkg/sorter"
)

type Index struct {
	ID    int
	URL   string
	Title string
}

func Const(Url string, Title string, dict []Index) []Index {
	var id int = 1
	if len(dict) != 0 {
		id = len(dict) + 1
	}
	c := Index{
		ID:    id,
		URL:   Url,
		Title: Title,
	}
	id++
	dict = append(dict, c)
	return dict
}

// Проходит по всем структурам массива вынимая название и разбирая его на слова, возвращая map слов и их Id.
func add(mapWord map[string][]int, dict []Index) map[string][]int {
	for _, v := range dict {
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

//  func read(w string) {
//  	var dict []Index
// 	var mapWord map[string][]int
// 	mapWord = add(mapWord, dict)

//  }

func Indexer(m map[string]string) map[string][]int {
	mapWord := make(map[string][]int)
	var dict []Index
	for Url, Title := range m {
		dict = Const(Url, Title, dict)
		mapWord = add(mapWord, dict)
	}
	// for _, v := range dict {
	// 	fmt.Println("Id: ", v.ID)
	// 	fmt.Println("Title: ", v.Title)
	// 	fmt.Println("Url: ", v.URL)
	// }
	for i, v := range mapWord {
		fmt.Println("Id cm", i)
		fmt.Println("Result Indexer:", v)
	}
	return mapWord
}
