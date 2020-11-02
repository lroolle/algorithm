package solution

import "sync"

func isAnagramBitMapGoroutine(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tableS := [26]int{}
	tableT := [26]int{}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		makeTable(&tableS, s)
		wg.Done()
	}()
	go func() {
		makeTable(&tableT, t)
		wg.Done()
	}()
	wg.Wait()
	return tableS == tableT

}

func makeTable(wordtable *[26]int, s string) {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		wordtable[index]++
	}
}

func isAnagramBitMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tableS := [26]int{}
	tableT := [26]int{}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		tableS[index]++
	}

	for i := 0; i < len(t); i++ {
		index := s[i] - 'a'
		tableT[index]++
	}

	return tableS == tableT
}
