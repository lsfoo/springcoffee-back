package photo

import (
	"bufio"
	"cafe.lsfoo.com/auth"
	"cafe.lsfoo.com/controllers/user"
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang/freetype"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	_ "strings"
	"time"
)

var engine = db.Orm

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "FZMHJW.TTF", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 50, "font size in points")
	spacing  = flag.Float64("spacing", 1.2, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

//用户图片
func FromUserHandler(w http.ResponseWriter, r *http.Request) {
	u, err := auth.FromUserAuth(r)
	if err != nil {
		fmt.Print(err)
	}
	var userPhotos []model.UserPhoto
	engine.Desc("create_time").Where("user_id = ?", u.UserId).Find(&userPhotos)
	if err != nil {
		panic(err)
	}
	jsons, _ := json.Marshal(userPhotos)
	fmt.Fprint(w, string(jsons))
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var fonts = []string{"方正美黑简.TTF", "方正清刻本悦宋简体.TTF", "方正宋刻本秀楷简体.TTF", "方正瘦金书简体.TTF", "hdqc.ttf"}
	u, err := auth.FromUserAuth(r)
	if err != nil {
		panic(err)
	}
	fmt.Print(u.UserId)
	uFile, _, err := r.FormFile("file")

	if err != nil {
		panic(err)
		return
	}
	defer uFile.Close()

	flag.Parse()
	//处理图片开始

	uf, _, err := image.Decode(uFile)

	if err != nil {
		panic(err)
	}

	var canvasWidth int
	x := uf.Bounds().Size().X
	y := uf.Bounds().Size().Y

	col := 23
	if x > y {
		canvasWidth = 1600
		col = 31
	} else {
		canvasWidth = 1200
	}

	resizeImage := resize.Resize(uint(canvasWidth), 0, uf, resize.Lanczos3)

	imgH := resizeImage.Bounds().Size().Y

	bg := image.NewRGBA(image.Rect(0, 0, canvasWidth, imgH))
	for x := 0; x < canvasWidth; x++ {
		for y := 0; y < imgH; y++ {
			bg.Set(x, y, color.NRGBA{255, 255, 255, 255})
		}
	}

	draw.Draw(bg, bg.Bounds(), resizeImage, resizeImage.Bounds().Min, draw.Over)

	queryWords := r.FormValue("words")
	if len(queryWords) > 0 {

		text := []rune("　　" + queryWords)
		//文字
		textLen := len(text)

		textRow := textLen/col + 1

		var textArr []string
		for i := 0; i < textRow; i++ {
			if textRow-i > 1 {
				if i == 0 {
					textArr = append(textArr, string(text[col*i:col*(i+1)]))
				} else {
					textArr = append(textArr, string(text[col*i:col*(i+1)]))
				}
			} else {
				textArr = append(textArr, string(text[col*i:textLen]))
			}

		}

		i := rand.Intn(4)
		fontBytes, err := ioutil.ReadFile(fonts[i])
		if err != nil {
			//	log.Println(err)
			return
		}
		f, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return
		}

		tfg, tbg := image.White, image.NewUniform(color.RGBA{0, 0, 0, 80})
		if *wonb {
			tfg, tbg = image.White, image.Black
		}
		//行高
		rgbaTextH := len(textArr) * 70
		rgbaText := image.NewRGBA(image.Rect(0, 0, canvasWidth, rgbaTextH))
		draw.Draw(rgbaText, rgbaText.Bounds(), tbg, image.ZP, draw.Src)
		c := freetype.NewContext()
		c.SetDPI(*dpi)
		c.SetFont(f)
		c.SetFontSize(*size)
		c.SetClip(rgbaText.Bounds())
		c.SetDst(rgbaText)
		c.SetSrc(tfg)
		switch *hinting {
		default:
			c.SetHinting(font.HintingNone)
		case "full":
			c.SetHinting(font.HintingFull)
		}

		// 画文本.
		pt := freetype.Pt(30, 10+int(c.PointToFixed(*size)>>6))
		for _, s := range textArr {
			_, err = c.DrawString(s, pt)
			if err != nil {
				//log.Println(err)
				return
			}
			pt.Y += c.PointToFixed(*size * *spacing)
		}

		textH := rgbaText.Bounds().Size().Y

		draw.Draw(bg, bg.Bounds(), rgbaText, rgbaText.Bounds().Min.Sub(image.Pt(0, imgH-textH-50)), draw.Over)
	}

	// Save that rgbaText image to disk.
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".jpeg"
	filePath := "./public/upload/" + fileName

	outFile, err := os.Create(filePath)
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}

	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	err = jpeg.Encode(b, bg, nil)
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}
	//fmt.Print(http.FileServer(http.Dir("./public")))
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAICmLY9c6bDHAX", "qUbMQOIx0qRRV11zSRHKR7E2rjuIUP")
	if err != nil {
		panic(err)

	}

	bucket, err := client.Bucket("lfo")
	if err != nil {
		panic(err)
	}

	err = bucket.PutObjectFromFile(fileName, filePath)
	if err != nil {
		panic(err)

	}

	//保存用户图片数据
	var up model.UserPhoto
	up.UserId = u.UserId
	up.Src = fileName
	up.Words = queryWords
	up.CreateTime = time.Now()

	//存库
	affected, err := engine.Insert(&up)
	if err != nil {
		fmt.Print(up)
	}
	if affected == 0 {
		fmt.Println("图片添加失败")
		return
	}

	isPrint := r.FormValue("isPrint")
	//添加打印任务
	if isPrint == "1" {
		printShopId, _ := strconv.Atoi(r.FormValue("printShopId"))
		var print model.PhotoPrintQueue
		print.UserId = u.UserId
		print.ShopId = printShopId
		print.UserPhotoId = up.UserPhotoId
		print.CreateTime = time.Now()
		_, err := engine.Insert(&print)
		if err != nil {
			panic(err)
		}

	}
	fmt.Fprintf(w, fileName)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	u, err := user.FromAuthorization(r)
	if err != nil {
		return
	}
	photo := &model.UserPhoto{UserPhotoId: id, UserId: u.UserId}
	affected, _ := engine.Delete(photo)

	fmt.Fprint(w, affected)
	//	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAICmLY9c6bDHAX", "qUbMQOIx0qRRV11zSRHKR7E2rjuIUP")
	//	if err != nil {
	//		panic(err)
	//
	//	}
	//
	//	bucket, err := client.Bucket("lfo")
	//	if err != nil {
	//		panic(err)
	//
	//	}
	//
	//	err = bucket.DeleteObject(photo.Src)
	//	if err != nil {
	//		panic(err)
	//	}

}
