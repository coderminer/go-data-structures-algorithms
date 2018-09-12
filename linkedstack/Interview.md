### 关于栈相关的问题-Go语言实现

#### 判断给定字符串中的括号是否匹配

```
func (s *Stack) IsBalancedParenthesis(exp string) bool {
	for _, ch := range exp {
		switch ch {
		case '{', '[', '(':
			s.Push(ch)
		case '}':
			val := s.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := s.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := s.Pop()
			if val != '(' {
				return false
			}
		}

	}

	return s.IsEmpty()
}
```