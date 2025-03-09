package main

import "fmt"

//creating a single linked listi

type element[T any] stuct{
    val T
    next *element[T]
}

type List[T any] struct{
    head, tail *element[T]
}

func (1st *List[T]) Push(v T){

    if 1st.tail == nil {
        1st.head = &element[T]{val: v}
        1st.tail = 1st.head
    } else {
        1st.tail.next = &element[T]{val: v}
        1st.tail = 1st.tail.next
    }
}


