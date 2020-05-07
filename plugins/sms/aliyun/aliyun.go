package aliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var client *dysmsapi.Client

func init() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accessKeyId>", "<accessSecret>")
	if err != nil {
		// Handle exceptions
		panic(err)
	}
}

func AddSMSSign(name, source, remark string, signf map[string]string) {
	request := dysmsapi.CreateAddSmsSignRequest()
	request.Scheme = "https"
	request.SignName = name
	request.SignSource = source
	request.Remark = remark
	request.SignFileList = &[]dysmsapi.AddSmsSignSignFileList{
		dysmsapi.AddSmsSignSignFileList{
			FileSuffix:   signf["suffix"],
			FileContents: signf["contents"],
		},
	}
	response, err := client.AddSmsSign(request)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	// Handle response
	fmt.Printf("response is %#v\n", response)
}

func QuerySMSSign(name string) {
	request := dysmsapi.CreateQuerySmsSignRequest()
	request.Scheme = "https"

	request.SignName = name

	response, err := client.QuerySmsSign(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func DeleteSMSSign(name string) {
	request := dysmsapi.CreateDeleteSmsSignRequest()
	request.Scheme = "https"
	request.SignName = name
	response, err := client.DeleteSmsSign(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func ModifySMSSign(name, source, remark string, signf map[string]string) {
	request := dysmsapi.CreateModifySmsSignRequest()
	request.Scheme = "https"

	request.SignName = name
	request.SignSource = source
	request.Remark = remark
	request.SignFileList = &[]dysmsapi.AddSmsSignSignFileList{
		dysmsapi.AddSmsSignSignFileList{
			FileSuffix:   signf["suffix"],
			FileContents: signf["contents"],
		},
	}

	response, err := client.ModifySmsSign(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func AddSMSTemplate(ttype, name, content, remark string) {
	request := dysmsapi.CreateAddSmsTemplateRequest()
	request.Scheme = "https"
	request.TemplateType = ttype
	request.TemplateName = name
	request.TemplateContent = content
	request.Remark = remark
	response, err := client.AddSmsTemplate(request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is %#v\n", response)
}

func QuerySMSTemplate(code string) {
	request := dysmsapi.CreateQuerySmsTemplateRequest()
	request.Scheme = "https"

	request.TemplateCode = code

	response, err := client.QuerySmsTemplate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func DeleteSMSTemplate(name string) {
	request := dysmsapi.CreateDeleteSmsTemplateRequest()
	request.Scheme = "https"
	request.SignName = name
	response, err := client.DeleteSmsTemplate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func ModifySMSTemplate(ttype, name, content, remark string) {
	request := dysmsapi.CreateModifySmsTemplateRequest()
	request.Scheme = "https"

	request.TemplateType = ttype
	request.TemplateName = name
	request.TemplateContent = content
	request.Remark = remark

	response, err := client.ModifySmsTemplate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func QuerySendDetails(phone, sdate, bizid string, pagesize, page int) {
	request := dysmsapi.CreateQuerySendDetailsRequest()
	request.Scheme = "https"

	request.PhoneNumber = phone
	request.SendDate = sdate
	request.PageSize = pagesize
	request.CurrentPage = page
	request.BizId = bizid

	response, err := client.QuerySendDetails(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func SendSMS(phone, name, code string) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "1234134"
	request.SignName = "asd"
	request.TemplateCode = "1242"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func SendBatchSMS(phones, names, code string) {
	request := dysmsapi.CreateSendBatchSmsRequest()
	request.Scheme = "https"

	request.PhoneNumberJson = phones
	request.SignNameJson = names
	request.TemplateCode = code

	response, err := client.SendBatchSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
