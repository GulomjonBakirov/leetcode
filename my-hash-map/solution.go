package solution

type MyHashMap struct {
	hashMap []int
}

func Constructor() MyHashMap {
	myMap := make([]int, 1000001)

	return MyHashMap{myMap}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hashMap[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
	return this.hashMap[key] - 1
}

func (this *MyHashMap) Remove(key int) {
	this.hashMap[key] = 0
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
