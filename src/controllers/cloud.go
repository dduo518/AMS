package controllers

import (
	"AMS/src/app"
	service "AMS/src/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)
func GetFederationToken(c *gin.Context)  {

	result,err:=service.GetFederationToken();

	if err!=nil{
		app.Response(c,err,400)
		return
	}
	app.Response(c,result,0)
}

func GetFileContentType(out *os.File) (string, error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func UploadFile(c *gin.Context){
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		app.Response(c,err,400)
		return
	}
	content, err := ioutil.ReadAll(file)
	filename := header.Filename
	filetype := http.DetectContentType(content)
	result,err:=service.UploadFile(content,filename,filetype);

	if err!=nil{
		app.Response(c,err,400)
		return
	}
	app.Response(c,result,0)
}