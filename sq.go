// Created by Yuxi Luo; September 10, 2018
// CS3253 Parallel and Distributed System; homework 1

package sq
import ("log")

// [Start, End)
type StringQueue struct {
    Slice []string
    Start int
    End int
}

func NewQueue(name string) (*StringQueue) {
    return &StringQueue {
        Slice: []string{name},
        Start: 0,
        End: 1,
    }
}

func (a *StringQueue) IsEmpty() (bool) {
    return a.End == a.Start
}

func (a *StringQueue) Push(name string) {
    if (a.IsEmpty()) {
        if (a.Slice == nil) {
            // if not initialized
            *a = *NewQueue(name)
            // fmt.Println("Initialize queue", a)
        } else {
            // if emptied by pop()
            a.Start = 0
            a.End = 0
            a.Slice[a.End] = name
            a.End += 1
            // fmt.Println("Restruct empty queue", a)
        }
        return
    }

    if (a.End < len(a.Slice)) {
        a.Slice[a.End] = name
        a.End += 1
        // fmt.Println("Direct add", a)
    } else if (a.Start > (cap(a.Slice) / 2)) {
        // if more than half of the space in slice not used
        // to avoid multiple moves of every elements for Pop
        // move all arrays to [0]
        origSlice := a.Slice[a.Start:a.End]
        copy(a.Slice, origSlice)
        a.End = a.End - a.Start
        a.Start = 0

        a.Slice[a.End+1] = name
        a.End += 1
        // fmt.Println("Move and direct add", a, len(a.Slice), cap(a.Slice))
    } else {
        // expand the slice
        a.Slice = append(a.Slice, name)
        a.End += 1
        // fmt.Println("Append", a, len(a.Slice), cap(a.Slice))
    }
}

func (a *StringQueue) Pop() (name string) {
    if (a.IsEmpty()) {
        log.Fatal("Cannot pop from an empty StringQueue")
    }

    a.Start += 1
    // fmt.Println("Pop", a)
    return a.Slice[a.Start-1]
}
