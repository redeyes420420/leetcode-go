type stack struct {
    backingArr []rune
}

func newStack () *stack {
    
    return new(stack)
}

func (s *stack) push(r rune) {
    
    s.backingArr = append(s.backingArr, r)
}

func (s *stack) pop() (r rune) {
    
    if len(s.backingArr) == 0 {
        r = ''
        return
    }
    r = s.backingArr[len(s.backingArr)-1]
    s.backingArr = s.backingArr[:len(s.backingArr)-1]
    return
}

func (s *stack) peek() rune {
    
    if len(s.backingArr) == 0 {
        return ''
    }
    return s.backingArr[len(s.backingArr)-1]
}

func (s *stack) len() int {
    
    return len(s.backingArr)
}

func isValid(s string) bool {
    
    pairs := map[rune]rune{
        '{': '}',
        '(': ')',
        '[': ']',
    }
    
    st := newStack()
    for _, r := range s {
        if r == '(' || r == '{' || r == '[' {
            st.push(r)
        } else {
            if pairs[st.peek()] == r {
                st.pop()
            } else {
                return false
            }
        }
    }
    
    return st.len() == 0
}
