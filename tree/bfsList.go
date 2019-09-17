/**
 * @Author: Pinelliaw
 * @Description:
 * @File:  bfsList
 * @Date: 2019/9/17 7:51 下午
 */

package tree

import (
    "container/list"
)

type ListQueue struct {
    List *list.List
}

func (q *ListQueue) enqueue(e Item) {
    q.List.PushBack(e)
}

func (q *ListQueue) dequeue() Item {
    if ele := q.List.Front(); ele != nil {
        q.List.Remove(ele)
        return ele.Value
    }
    return nil
}

func (t *Tree) BfsList(f func(item Item)) {
    queue := ListQueue{List: list.New()}
    queue.enqueue(t)

    for {
        if queue.List.Len() == 0 {
            break
        }

        node, ok := queue.dequeue().(*Tree)
        if ok {
            near := node.Children
            for i := 0; i < len(near); i++ {
                queue.enqueue(near[i])
            }
        } else {
            break
        }
        if f != nil {
            f(node.Value)
        }
    }
}
