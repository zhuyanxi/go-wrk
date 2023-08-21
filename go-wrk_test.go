package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestReq(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			url := "http://10.196.8.103:7001/api/v1/spaces/9d875fd3-b411-4f23-8841-027f51e4e7ae/wave/metriccollectors"
			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Ubi-AppId", "c7c727e9-c210-472d-8340-eacd1e14f19f")
			req.Header.Add("Authorization", "ubi_v1 t=ewogICJ2ZXIiOiAiMSIsCiAgImFpZCI6ICJjN2M3MjdlOS1jMjEwLTQ3MmQtODM0MC1lYWNkMWUxNGYxOWYiLAogICJlbnYiOiAiVUFUIiwKICAic2lkIjogImExYTc3YmMyLTcxZDUtNGVjNy1iMjNjLWNiZWY2YjQwNTBmNSIsCiAgInR5cCI6ICJKV0UiLAogICJlbmMiOiAiQTEyOENCQyIsCiAgIml2IjogImZUUG1qNmdJdDVLQkt0TUZ0bnEyWEEiLAogICJpbnQiOiAiSFMyNTYiCn0.3m-bHFO6LTAzt2vndqN8ewkFwPyo5Z-EGcCBKO2VYXuX2G5mX5sz7GrWXxkPXx_1MAw93ERvsXAOAHL7w_UIaLX4Bd1j6wBFuEnu3jkwpL4ifWdcs1y3iATXBrx4OAzoPEwPLXQ5anB5gl8vYL-ulR0kzJ3WKGnt62Z5GBiRejS6F7bjtLNkdhBiwMfR99Pl5teNFrFyYWhy2HXqpZG9DtFAS9cA1TL-1fYkLejL8HKG8duvl4EJzfbix5PqsEGR8LsH9_6o-uTgmWvPFAKbSTqhGsLJlkG8Mq02JRKES66x3O9_0z9h6gxH_4eklBVKpbPcSCBrxtoumnAx_vA2bFEF2mDbfOvr1gyFentKeb4cPKxK2eH_MemHzbyMH48QCsCePT2iY5r2FeB5z7QbF2_vZlv_p8vEgmGRQ4uuRSHzCqAVuNCWVNWXzPWrlllg5qyno6_odu6nwDKlg61WkesSNzuVoJebtF5aA-BIIPp9OdzYUJPZ8HVZ39cX2gvp-bJoo5pSixCIJ6SD878QwaeLMTwi_RbJmeoRDH7BtFPm_-lifiG9GqGkRnNY9gTLsZD5AImgX_-ZQtyn_7eNyn4thrff6lPolqINvcSFGQl25EaKzT5dsanfcACs7oXBaoNXcJEwcNM4AU3ZMRHBsId9lqJY2X5u2j7xTSZ36VsZ4hDSDssQVxIVFhVlE0fsHKVa_Zr3ACkcFLrVdvXA63ZHwrf7qwWu4qEerafj3SWGPOfLuFGcdmS3wwacuPrl8pn-ON_S4H9UZfJrwXnCNlVgb1tRmo_ocMb7JgJ7Rb_ygoMgWaCU3XJPevZHpNRI7Omuq8DHbfFQwSa-SzNujlBCbtxtSyYbU_P6caW_PH4IO7VslKv6z7VWKmpgicVXeQb1Uhf78Kvb6CQrWwEp6t56wC6J4X4cDYgEhBLIJSB5_4W2ZIM2UxuoI5m3xyQH8A2NAlQ_E_lYO_3997UafAN2LgA3KDhQ8jjcp4PZY5s.GE59G6TkbMobSEgcbKSF4m3SPVIIhGa4tgmU_F9hxRs")
			req.Header.Add("Ubi-SessionId", "bbd0d53f-72b9-49bd-aa38-a631c07c7c16")
			// req.Header.Add("User-Agent", "go-wrk")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println(string(body))
			log.Println(string(body))
		}()
	}
	time.Sleep(30 * time.Second)
	fmt.Println("Test Finish")
}
