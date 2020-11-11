package scbstudents

import (
	"encoding/base64"
	"fmt"
	tools2 "go-admin/apis/tools"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
	"go-admin/tools/config"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nfnt/resize"
)

func GetScbStudentsList(c *gin.Context) {
	var data models.ScbStudents
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Id, _ = tools.StringToInt(c.Request.FormValue("id"))
	data.Name = c.Request.FormValue("name")
	data.Number = c.Request.FormValue("number")
	data.ClassId, _ = tools.StringToInt(c.Request.FormValue("classId"))
	data.LineId, _ = tools.StringToInt(c.Request.FormValue("lineId"))
	data.SiteIdUp, _ = tools.StringToInt(c.Request.FormValue("siteIdUp"))
	data.SiteIdDown, _ = tools.StringToInt(c.Request.FormValue("siteIdDown"))
	data.IsPickUp, _ = tools.StringToInt(c.Request.FormValue("isPickUp"))
	data.CarId, _ = tools.StringToInt(c.Request.FormValue("carId"))
	data.ParentPhone = c.Request.FormValue("parentPhone")
	data.Picture = c.Request.FormValue("picture")
	data.IsDeleted, _ = tools.StringToInt(c.Request.FormValue("isDeleted"))

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	var studentsData []models.ScbStudents
	for _, student := range result {
		//线路
		lineModel := models.ScbLines{}
		lineModel.Id = student.LineId
		lineData, err := lineModel.Get()
		if nil == err {
			student.LineName = lineData.Name
		}

		//下车站点
		siteModel := models.SchSites{}
		siteModel.Id = student.SiteIdDown
		siteDownData, err := siteModel.Get()
		if nil == err {
			student.SiteDownName = siteDownData.Name
		}

		//上车站点
		siteModel.Id = student.SiteIdUp
		siteUpData, err := siteModel.Get()
		if nil == err {
			student.SiteUpName = siteUpData.Name
		}
		//班级信息
		deptModel := models.ScbDept{}
		deptModel.DeptId = student.ClassId
		classData, err := deptModel.Get()
		if nil == err {
			student.ClassName = classData.DeptName
		}

		//车牌
		carModel := models.ScbCars{}
		carModel.Id = student.CarId
		carData, err := carModel.Get()
		if nil == err {
			student.CarNo = carData.CarNo
		}
		studentsData = append(studentsData, student)
	}

	app.PageOK(c, studentsData, count, pageIndex, pageSize, "")
}

func GetScbStudents(c *gin.Context) {
	var data models.ScbStudents
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	result.HeadImg = config.ApplicationConfig.ImageUrl + result.HeadImg
	result.HeadImgSmall = strings.Replace(result.HeadImg, ".", "_small.", 1)

	app.OK(c, result, "")
}

func InsertScbStudents(c *gin.Context) {
	var data models.ScbStudents
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)

	picture := data.Picture
	data.Picture = ""
	//新增学生信息
	result, err := data.Create()
	tools.HasError(err, "", -1)

	if "" != picture {
		api := &tools2.BdApi{}
		studentIdStr := strconv.Itoa(result.Id)
		faceToken := api.FacesetAdd(studentIdStr, picture)

		imageArr := strings.Split(picture, ";base64,")
		ext := "png"
		if len(imageArr) > 1 {
			imageArrTmp := strings.Split(imageArr[0], "/")
			if len(imageArrTmp) > 1 {
				ext = imageArrTmp[1]
			}
			ddd, _ := base64.StdEncoding.DecodeString(imageArr[1])          //成图片文件并把文件写入到buffer
			ioutil.WriteFile("./images/face_"+faceToken+"."+ext, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			saveSmall("./images/face_"+faceToken+"."+ext, ext, "./images/face_"+faceToken+"_small."+ext)
		}

		if "" != faceToken {
			var updateData models.ScbStudents
			updateData.FaceToken = faceToken
			updateData.HeadImg = "face_" + faceToken + "." + ext
			//更新人脸
			_, err := data.Update(result.Id)
			tools.HasError(err, "更换相片失败", -1)
		}
	}

	app.OK(c, nil, "")
}

// 更新学生信息
func UpdateScbStudents(c *gin.Context) {
	var data models.ScbStudents
	err := c.MustBindWith(&data, binding.JSON)
	fmt.Println("err...", err)
	tools.HasError(err, "数据解析失败", -1)

	api := &tools2.BdApi{}
	studentIdStr := strconv.Itoa(data.Id)
	faceToken := api.FacesetAdd(studentIdStr, data.Picture)

	imageArr := strings.Split(data.Picture, ";base64,")
	ext := "png"
	fmt.Println("imageArr. len......", len(imageArr))
	if len(imageArr) > 1 {
		imageArrTmp := strings.Split(imageArr[0], "/")
		if len(imageArrTmp) > 1 {
			ext = imageArrTmp[1]
		}
		ddd, _ := base64.StdEncoding.DecodeString(imageArr[1])          //成图片文件并把文件写入到buffer
		ioutil.WriteFile("./images/face_"+faceToken+"."+ext, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		saveSmall("./images/face_"+faceToken+"."+ext, ext, "./images/face_"+faceToken+"_small."+ext)
	}

	fmt.Println("faceToken...", faceToken)
	if "" != faceToken {
		data.FaceToken = faceToken
		data.HeadImg = "face_" + faceToken + "." + ext
	}

	data.Picture = ""
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbStudents(c *gin.Context) {
	var data models.ScbStudents

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

//保存小图
func saveSmall(filepath, ext, smallPath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("saveSmall Open err......", err)
		return
	}
	var img image.Image
	switch strings.ToLower(ext) {
	case "png":
		img, err = png.Decode(file)
	case "jpg":
		img, err = jpeg.Decode(file)
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	default:
		return
	}
	// decode jpeg into image.Image
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio

	//newWidth := 200
	kuan := img.Bounds().Dx() //
	gao := img.Bounds().Dy()  //
	newHeight := uint(kuan * 200 / gao)
	fmt.Println("newHeight....", newHeight, kuan, gao, 200)

	//m := resize.Resize(200, 200, img, resize.Bilinear)
	//m := resize.Resize(1000, 0, img, resize.Lanczos3)
	m := resize.Thumbnail(200, newHeight, img, resize.Lanczos3)

	out, err := os.Create(smallPath)
	if err != nil {
		fmt.Println("saveSmall Create err......", err)
		return
	}
	defer out.Close()
	fmt.Println("saveSmall strings.ToLower(ext)......", strings.ToLower(ext))

	// write new image to file
	var sErr error
	switch strings.ToLower(ext) {
	case "png":
		sErr = png.Encode(out, m)
	case "jpg":
		sErr = jpeg.Encode(out, m, &jpeg.Options{Quality: 100})
	case "jpeg":
		sErr = jpeg.Encode(out, m, &jpeg.Options{Quality: 100})
	case "gif":
		sErr = gif.Encode(out, m, &gif.Options{})
	default:
		return
	}
	fmt.Println("saveSmall sErr......", sErr)
}
