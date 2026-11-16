type Node struct {
	val, key int
	prev *Node
	next *Node
}

type LRUCache struct {
    cp int
	mp map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	h := &Node{val: -1}
	t := &Node{val: -1}

	h.next= t
	t.prev= h

	return LRUCache{
		cp: capacity,
		mp: map[int]*Node{},
		head: h,
		tail: t,
	}
    
}

func (this *LRUCache) addNode(node *Node){
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev= node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node){
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) movetoHead(node *Node){
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) Get(key int) int {
	node, ok:= this.mp[key]

	if !ok {
		return -1
	}
	this.movetoHead(node)
	return node.val
    
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok:= this.mp[key] ; ok{
		node.val=value
		this.movetoHead(node)
		return
	}

	node := &Node{
		key: key,
		val: value,
	}

	this.mp[key]= node
	this.addNode(node)

	if len(this.mp) > this.cp {
        lru := this.tail.prev

        this.removeNode(lru)
        delete(this.mp, lru.key)
	}
}
