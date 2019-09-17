/**
 * @Author: Pinelliaw
 * @Description: 广度优先搜索数组版 Breadth First Search
 * @File:  bfsArray
 * @Date: 2019/9/16 10:13 下午
 */

package tree

import "fmt"

type MyQueue struct {
    Items []Tree
}

func (q *MyQueue) New() *MyQueue {
    q.Items = []Tree{}
    return q
}

func (q *MyQueue) enqueue(t Tree) {
    q.Items = append(q.Items, t)
}

func (q *MyQueue) dequeue() *Tree {
    node := q.Items[0]
    q.Items = q.Items[1:len(q.Items)]
    return &node
}

func (q *MyQueue) IsEmpty() bool {
    return len(q.Items) == 0
}

func (t *Tree) BfsArray() {
    queue := MyQueue{}
    queue.New()
    queue.enqueue(*t)

    for {
        if queue.IsEmpty() {
            break
        }
        node := queue.dequeue()
        near := node.Children
        for i := 0; i < len(near); i++ {
            queue.enqueue(*near[i])
        }
        fmt.Printf("BFS visiting... %v\n", node.Value)
    }
}
