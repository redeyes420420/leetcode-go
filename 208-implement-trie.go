type TrieNode struct {
    children map[rune]*TrieNode
    end bool
}

type Trie struct {
    root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() Trie {
    
    root := TrieNode{
        children: make(map[rune]*TrieNode, 0),
        end: false,
    }
    
    trie := Trie{
        root: &root,
    }
    return trie
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    
    cur := this.root
    for _, r := range word {
        node, f := cur.children[r]
        if !f {
            node = &TrieNode{
                children: make(map[rune]*TrieNode, 0),
                end: false,
            }
            cur.children[r] = node
        }
        cur = node
    }
    
    cur.end = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    
    cur := this.root
    for _, r := range word {
        node, f := cur.children[r]
        if !f {
            return false
        }
        cur = node
    }
    
    return cur.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    
    cur := this.root
    for _, r := range prefix {
        node, f := cur.children[r]
        if !f {
            return false
        }
        cur = node
    }
    
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
