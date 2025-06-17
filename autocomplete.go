package autocomplete

import (
	"sync"
)

// AutoComplete provides a thread-safe API to Load `string` entries to later be used in an auto-complete.
type AutoComplete struct {
	root *autoCompleteNode
	mu   *sync.Mutex
}

// Create a new AutoComplete instance to load completion entries into.
func New() *AutoComplete {
	return &AutoComplete{
		root: newAutoCompleteNode(false),
		mu:   &sync.Mutex{},
	}
}

// Loads values into the AutoComplete so they can later be retrieved.
func (ac *AutoComplete) Load(values ...string) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	for _, value := range values {
		prefix := make([]rune, 0)
		for indx, letter := range value {
			isEnd := indx == len(value)-1
			ac.root.append(prefix, letter, isEnd)
			prefix = append(prefix, letter)
		}
	}
}

// Returns a slice of `string` that begin with `prefix`.
//
// If `prefix` is an empty string, all the entries will be returned.
func (ac *AutoComplete) Complete(prefix string) []string {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	return ac.root.searchByPrefix([]rune(prefix))
}

func (ac *AutoComplete) Count() int {
	return len(ac.Complete(""))
}
