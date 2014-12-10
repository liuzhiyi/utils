package captcha

import (
	"crypto/rand"
	r "math/rand"
	"strconv"
	"time"
	"ucapi/utils/image"
	"ucapi/utils/levelcache"
)

type Captcha struct {
	cache  *levelcache.Levelcache
	sessid string
	width  int
	height int
	length int
}

func NewCaptcha(e time.Duration) *Captcha {
	lc := levelcache.NewLevecache(e, 10, "./db/captcha")
	cpt := new(Captcha)
	cpt.cache = lc
	return cpt
}

func (this *Captcha) SetWidth(w int) *Captcha {
	this.width = w
	if this.width <= 0 {
		this.width = 240
	}
	return this
}

func (this *Captcha) Setheight(h int) *Captcha {
	this.height = h
	if this.height <= 0 {
		this.height = 80
	}
	return this
}

func (this *Captcha) SetLength(l int) *Captcha {
	this.length = l
	return this
}

func (this *Captcha) SetSessid(sessid string) *Captcha {
	this.sessid = sessid
	return this
}

func (this *Captcha) CheckID(v string) bool {
	var val string
	defer this.cache.Delete(this.sessid)
	if this.cache.Get(this.sessid, &val) && val == v {
		return true
	}
	return false

}
func (this *Captcha) Colse() {
	this.cache.Close()
}
func (this *Captcha) CreateImageCaptcha() []byte {
	val := this.GetRandChars(this.length)
	str := ""
	for i := 0; i < len(val); i++ {
		str += strconv.Itoa(int(val[i]))
	}
	this.cache.Set(this.sessid, str, 0)
	img := image.NewImage(val, this.width, this.height)
	return img.EncodedPNG()
}

func (this *Captcha) GetRandChars(len int) []byte {
	return RandomCreateBytes(len, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
}

func RandomCreateBytes(n int, alphabets ...byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[r.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[r.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}
