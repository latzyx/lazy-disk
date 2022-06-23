package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"lazy/ReturnMessage"
	"lazy/config"
	"net/http"
)

func Gettoken(c *config.Client) *ReturnMessage.ApifoxModal {
	r := new(ReturnMessage.ApifoxModal)
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + c.Corpid + "&corpsecret=" + c.CorPsecret
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	marshal, err := r.Marshal(body)
	err = json.Unmarshal(marshal, &r)
	return r
}
