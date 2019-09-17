/**
 * @Author: Pinelliaw
 * @Description:
 * @File:  dfsList.go
 * @Date: 2019/9/17 8:31 下午
 */

package tree

import "container/list"

type ListStack struct {
    List *list.List
}

func (s *ListStack) push(item Item) {
    s.List.PushBack(item)
}

func (s *ListStack) pop() Item {
    if ele := s.List.Back(); ele != nil {
        s.List.Remove(ele)
        return ele.Value
    }
    return nil
}

func (t *Tree) DfsList(f func(item Item)) {
    stack := ListStack{List: list.New()}
    stack.push(t)
    for {
        if stack.List.Len() == 0 {
            break
        }

        node, ok := stack.pop().(*Tree)
        if ok {
            near := node.Children
            for i := len(near) - 1; i >= 0; i-- {
                stack.push(near[i])
            }
        } else {
            break
        }

        if f != nil {
            f(node.Value)
        }
    }
}
