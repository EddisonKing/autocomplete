package autocomplete

type autoCompleteNode struct {
	isEnd bool
	data  map[rune]*autoCompleteNode
}

func newAutoCompleteNode(isEnd bool) *autoCompleteNode {
	return &autoCompleteNode{
		isEnd: isEnd,
		data:  make(map[rune]*autoCompleteNode),
	}
}

func (acn *autoCompleteNode) append(prefix []rune, letter rune, isEnd bool) {
	if len(prefix) == 0 {
		if _, exists := acn.data[letter]; !exists {
			acn.data[letter] = newAutoCompleteNode(isEnd)
		}
		return
	}

	acn.data[prefix[0]].append(prefix[1:], letter, isEnd)
}

func (acn *autoCompleteNode) rootByPrefix(prefix []rune) *autoCompleteNode {
	if len(prefix) == 0 {
		return acn
	}

	if n, exists := acn.data[prefix[0]]; exists {
		return n.rootByPrefix(prefix[1:])
	}

	return nil
}

func (acn *autoCompleteNode) searchByPrefix(prefix []rune) []string {
	searchResults := make([]string, 0)

	root := acn.rootByPrefix(prefix)
	if root == nil {
		return searchResults
	}

	root.search(&searchResults, prefix, []rune{})

	return searchResults
}

func (acn *autoCompleteNode) search(searchResults *[]string, prefix []rune, r []rune) {
	if acn.isEnd {
		word := string(append(prefix, r...))
		*searchResults = append(*searchResults, word)
	}

	for letter, node := range acn.data {
		node.search(searchResults, prefix, append(r, letter))
	}
}
