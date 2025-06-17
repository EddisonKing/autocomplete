# AutoComplete
A simple implementation of a prefix based auto-completer that is thread safe using a trie data structure.

## Installation and Import
```bash
go get github.com/EddisonKing/autocomplete
```

```go
import "github.com/EddisonKing/autocomplete"
```

## Usage

### Getting Started
```go
ac := autocomplete.New()
```

This will create a new `AutoComplete` that you can then load `string` entries into for later use in auto-completion.

### Loading Entries
```go
// Single string
ac.Load("apples")
// Or multiple
ac.Load([]string{"apples", "bananas", "ewwwww_fruit"}...)
```

This will load entries into the `AutoComplete` allowing them to be returned as entries during auto-completion. 

Duplicate values will not change the underlying state, but the algorithm can't know an entry is a duplicate without walking the trie, so effectively, you will be wasting compute on duplicates. 

### Completions
```go
results := ac.Complete("a") // Would return []string{"apples"} if the above Load data was used
```

Call `Complete` and pass in a prefix string to get a slice containing entries that satisfy the auto-completion.

Calling `Complete` with an empty prefix, `""`, will return all entries that were stored. Note that this requires a traversal of the full trie so the effort to return all entries is the same effort as what it took to store all the entries originally.

## Future Features
- Ranking entries on Load, so higher ranked entries come back first upon Complete
- Lazy yield of results over time by passing back a `chan string` instead

## Notes On Performance
The underlying data structure is known as a "trie" which is a portion of the word "retrieval", the reason why the structure was initially designed.

Loading entries into the tree requires a traversal of the tree at a depth of however long the input is. Consider the performance is likely:
```
O(k) where `k` is the length of the input
```

Complete takes a traversal of the prefix, then a full traversal of the remainder of the trie past the prefix. Likely:
```
O(p + c) where `p` is the length of the prefix and `c` is how many entries there remaining in that sub-trie
```

A trie is fairly space efficient since overlapping words in the data (ex. "apple" and "app") are also stored utilising this overlap.
