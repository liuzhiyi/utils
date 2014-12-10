package captcha

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	c "github.com/astaxie/beego/utils/captcha"
	"testing"
	//"ucapi/utils/image"
)

var (
	defaultChars = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func Test_Action(t *testing.T) {
	cpt := NewCaptcha(24 * 60 * 60)
	cpt.SetSessid("test")
	cpt.SetWidth(10)
	cpt.Setheight(20)
	cpt.SetLength(4)
	fmt.Println(utils.RandomCreateBytes(6, defaultChars...))
	//image.NewImage(utils.RandomCreateBytes(6, defaultChars...), 10, 20)
	c.NewImage(utils.RandomCreateBytes(6, defaultChars...), 150, 100)
	//fmt.Println(img.EncodedPNG())

}
