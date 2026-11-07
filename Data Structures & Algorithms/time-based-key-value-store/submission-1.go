type keyvalue struct {
	val string
	ts int
}

type TimeMap struct {
	hmap map[string][]keyvalue
}

func Constructor() TimeMap {
	return TimeMap {
		hmap : make(map[string][]keyvalue),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if values, ok := this.hmap[key]; ok{
		this.hmap[key]=append(values, keyvalue{val: value, ts:timestamp})
	} else {
		this.hmap[key]=[]keyvalue{{val: value, ts:timestamp}}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {

	if values,ok := this.hmap[key]; ok{
		if values[len(values)-1].ts<timestamp {
			return values[len(values)-1].val
		} 

		if values[0].ts>timestamp {
			return ""
		}

		l, r := 0, len(values)

		for l+1<r {
			mid := l+(r-l)/2

			if values[mid].ts>timestamp{
				r = mid
			} else {
				l = mid
			}
		}
		return values[l].val
	}
	return ""
}
