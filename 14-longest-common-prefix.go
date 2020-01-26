type trieNode struct {
    children map[rune]*trieNode
    isEnd bool
}

func newTrieNode() *trieNode {
    
    tn := trieNode{make(map[rune]*trieNode), false}
    return &tn
}

type trie struct {
    root *trieNode   
    size int
}

func newTrie() *trie{
    
    tr := trie{newTrieNode(), 0}
    return &tr
}

func (tr *trie) insert(word string) {
    
    if word == "" {
        word = " "
    }
    cur := tr.root
    for _, r := range word {
        node := cur.children[r]
        if node == nil {
            node = newTrieNode()
            cur.children[r] = node
        }
        cur = node
    }
    tr.size++
    cur.isEnd = true
}

func (tr *trie) commonPrefix() string {
    
    var prefix string
    
    cur := tr.root
    if len(cur.children) != 1 {
        return prefix
    }
    var r rune
    for len(cur.children) == 1 && !cur.isEnd{
        for r = range cur.children {
            prefix += string(r)
        }
        cur = cur.children[r]
    }
    
    if prefix == " " {
        prefix = ""
    }
    return prefix
}

func longestCommonPrefix(strs []string) string {
    
    tr := newTrie()
    for _, str := range strs {
        tr.insert(str)
    }
    
    return tr.commonPrefix()
}
