package levelcache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"time"
	"ucapi/utils/str"
)

type Levelcache struct {
	db                *leveldb.DB
	defaultExpiration time.Duration
	stopSem           chan bool
}

type value struct {
	Val        []byte
	Expiration int64
}

func (v *value) isExpired() bool {
	if v.Expiration <= 0 {
		return true
	}
	return v.Expiration-time.Now().UnixNano() <= 0
}

//1 defaultExpiration:默认过期时间  2 cleanupInterval：设置过期时间
func NewLevecache(defaultExpiration, cleanupInterval time.Duration, dbPath string) *Levelcache {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err.Error())
	}
	c := &Levelcache{db: db, defaultExpiration: defaultExpiration}
	if cleanupInterval > 0 {
		go c.RunGc(cleanupInterval)
	}
	return c
}

func (this *Levelcache) Get(key string, val interface{}) bool {
	data, err := this.db.Get([]byte(key), nil)
	v := this.decodeData(data)
	if err != nil {
		return false
	}
	if v.isExpired() {
		this.db.Delete([]byte(key), nil)
		return false
	}

	if err := str.Deserialize([]byte(v.Val), val); err == nil {
		return true
	}
	return false
}

//val不能为指针类型
func (this *Levelcache) Set(key string, val interface{}, d time.Duration) {
	var t int64
	if d <= 0 {
		d = this.defaultExpiration
	} else {
		d = d * 1e9
	}
	t = time.Now().Add(d).UnixNano()
	data := this.encodeData(val, t)
	this.db.Put([]byte(key), data, nil)
}

func (this *Levelcache) Delete(key string) {
	this.db.Delete([]byte(key), nil)
}

func (this *Levelcache) DeleteExpired() {
	iter := this.db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		val := this.decodeData(iter.Value())
		if val == nil || val.isExpired() {
			this.Delete(string(key))
		}
	}
	iter.Release()
}

func (this *Levelcache) encodeData(val interface{}, t int64) []byte {
	enval, _ := str.Serialize(val)
	item := value{
		Val:        enval,
		Expiration: t,
	}
	data, _ := str.Serialize(item)
	return data
}

func (this *Levelcache) decodeData(data []byte) *value {
	v := new(value)
	if err := str.Deserialize(data, v); err == nil {
		return v
	}
	return nil
}

func (this *Levelcache) Close() {
	this.db.Close()
	this.stopSem <- true
}

func (this *Levelcache) RunGc(interval time.Duration) {
	this.stopSem = make(chan bool)
	tick := time.Tick(interval)
	for {
		select {
		case <-tick:
			this.DeleteExpired()
		case <-this.stopSem:
			return
		}
	}
}

func (this *Levelcache) IsExist(key string) bool {
	if _, err := this.db.Get([]byte(key), nil); err != nil {
		return false
	}
	return true
}
