package _interface

import (
	"AMS/config"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Ext string `json:"ext"`
	Extend string `json:"extend"`
	Params []string `json:"params"`
	Sig string `json:"sig"`
	Sign string `json:"sign"`
	Tel map[string]string `json:"tel"`
	Time string `json:"time"`
	TplId string `json:"tpl_id"`
	Random string
}

type Resp struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// send api to message platform
func (params *Message)Send()(interface{},error)  {
	jsonBytes, err := json.Marshal(params)
	url:=config.Conf.Letter.Api["SINGLE"]+"?sdkappid="+config.Conf.Letter.AppId+"&random="+params.Random

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body,nil
}


func (params *Message)VerifierMessage(num ,code string) (interface{},error)  {
	params.SetParams(code)
	params.Tel= map[string]string{
		"mobile":num,
		"nationcode":"86",
	}
	params.SetSha256()
	resp,err:=params.Send()
	return resp,err
}

func (params *Message)SetParams(code string)  {
	params.Params = []string{code,"2"}
	params.Sign = config.Conf.Letter.Sign
	params.TplId = config.Conf.Letter.TplId
	timeStamp:= time.Now().Unix()
	timeStr :=strconv.FormatInt(timeStamp,10)
	params.Time= timeStr
}

func (params *Message)SetSha256()  {
	sig := sha256.New()
	r := rand.New(rand.NewSource(99))
	params.Random= string(r.Int31())
	strSig:="appkey="+config.Conf.Letter.AppKey+"&random="+params.Random+"&time="+params.Time+"&mobile=13423805473"
	sig.Write([]byte(strSig))
	params.Sig = hex.EncodeToString(sig.Sum([]byte("")))
}
