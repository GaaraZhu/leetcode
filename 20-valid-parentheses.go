func isValid(s string) bool {
    stack := NewStringStack()
    for i:=0; i<len(s); i++ {
        current := s[i:i+1]
        if current == "(" || current == "{" || current == "[" {
            stack.Push(current)
            continue
        }

        if current == ")" {
            if stack.Peek() != "(" {
                return false
            } else {
                stack.Pop()
            }
        } else if current == "}" {
            if stack.Peek() != "{" {
                return false
            } else {
                stack.Pop()
            }
        } else if current == "]" {
            if stack.Peek() != "[" {
                return false
            } else {
                stack.Pop()
            }
        }
    }
    
    if stack.Len() != 0 {
        return false
    }
    
    return true
}


type StringStack struct {
    items []string
}
func NewStringStack() *StringStack {
	s := &StringStack{}
	s.items = []string{}
	return s
}
func (s *StringStack) Len() int {
    return len(s.items)
}
func (s *StringStack) Peek() string {
    if(len(s.items)==0){
	return ""
    }
    return s.items[len(s.items)-1]
}
func (s *StringStack) Pop() string {
    result := s.items[len(s.items)-1]
    s.items = s.items[0:len(s.items)-1]
    return result
}
func (s *StringStack) Push(value string) {
    s.items = append(s.items, value)
}