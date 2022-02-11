package main

 
const Size = 10
type Hashtable struct{
     array [Size]*bucket
}
//Insert will take in a key and add it to the hash table array
func (h *Hashtable)Insert(key string)bool{
    index := hash(key)
    h.array[index].Insert(key)
    return true
}
//Search will take in key and return true if key in hash table
func (h *Hashtable)search(key string)bool{
    index := hash(key)
    return h.array[index].search(key)
}
//Delete will take in a key and delete it from the hash table
func (h *Hashtable)delete(key string)bool{
    index := hash(key)
    return h.array[index].delete(key)
}
 
type bucket struct{
    head *bucketNode
}
func (b *bucket)Insert(keyVal string){
    newNode  := &bucketNode{key: keyVal}
    newNode.next = b.head
    b.head = newNode
} 
func (b *bucket)search(key string)bool{
     cm := b.head 
     for cm != nil{
        if cm.key == key{
             return true
        }
       cm = cm.next
     }
    return false
}
func (b *bucket)delete(key string)bool{
    if b.head.key == key{
        b.head = b.head.next
        return true
    }
    prevNode := b.head

    for prevNode.next != nil{
            if prevNode.next.key == key{
                prevNode.next = prevNode.next.next
                return true
            }
     prevNode = prevNode.next       
    }
    return false
}

type bucketNode struct{
     key string
     next *bucketNode
}

func InitHashtable()*Hashtable{
    h := &Hashtable{}
    for i := range h.array{
        h.array[i] = &bucket{}
    }
    return h
}
// hash func
func hash(key string)int{
    var sum int

    for _,v := range key{
        sum += int(v)
    }
    return  sum % Size
}

 