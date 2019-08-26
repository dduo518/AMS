package services

import (
	"AMS/config"
	"context"
	"encoding/json"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getSignature(params map[string]interface{},key string,method string) string{

	return ""
}

func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

func GetFederationToken()(str string,err error)  {

	var uri = "https://sts.tencentcloudapi.com"
	var Action = "GetFederationToken"
	var Version = "2018-08-13"
	var Region = "GetFederationToken"
	var Name = ""
	var DurationSeconds = 1800;
	var method = "POST"
	var SecretId= "SecretId"
	var SecretKey= "SecretKey"

	var policy map[string]interface{};
	//var policy = {
	//	'version': '2.0',
	//	'statement': [{
	//		'action': config.allowActions,
	//			'effect': 'allow',
	//		'principal': {'qcs': ['*']},
	//		'resource': [
	//		'qcs::cos:' + config.region + ':uid/' + appId + ':prefix//' + appId + '/' + shortBucketName + '/' + config.allowPrefix,
	//		],
	//	}],
	//};

	policy["version"] = Version

	 policyStr ,_:= json.Marshal(policy)

	var params map[string]interface{}
	params["SecretId"] =  SecretId
	params["Timestamp"] = time.Now().Unix()
	params["Nonce"]= rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	params["Policy"]= encodeURIComponent(string(policyStr))
	params["Region"]= Region
	params["Name"]= Name
	params["DurationSeconds"]= DurationSeconds
	params["Version"]= Version
	params["Action"]= Action

	params["Signature"] = getSignature(params, SecretKey, method)
	resp, err := http.PostForm(uri,
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body),nil
}

func UploadFile(content []byte, fileName string, fileType string)( string, error){
	var uri= "http://"+config.Conf.Cos.Bucket+"-"+config.Conf.Cos.AppId+"."+config.Conf.Cos.Name+"."+config.Conf.Cos.Region+"."+config.Conf.Cos.Host
	u,_:=url.Parse(uri)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: config.Conf.Cos.SecretID,
			SecretKey: config.Conf.Cos.SecretKey,
		},
	})
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: fileType,
			ContentLength:len(content),
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "public-read",
		},
	}
	resp, err := client.Object.Put(context.Background(), fileName, strings.NewReader(string(content)), opt)
	defer resp.Body.Close()
	if err != nil {
		return "",err
	}
	return uri+"/"+fileName,nil
}