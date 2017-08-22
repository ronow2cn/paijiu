package route

import (
	"comm"
	"comm/config"
	"comm/dbmgr"
	"comm/sched/asyncop"
	"encoding/json"
	"fmt"
	"net/http"
	"proto/errorcode"
	"proto/macrocode"
	"time"
)

// ============================================================================
//weixin server res
type weixinTokenRet struct {
	Token     string `json:"access_token"`
	ExpireIn  int64  `json:"expires_in"`
	RefrToken string `json:"refresh_token"`
	OpenId    string `json:"openid"`
	Scope     string `json:"scope"`
	UnionId   string `json:"unionid"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

//weixin get user info
type weixinUserInfo struct {
	OpenId     string `json:"openid"`
	NickName   string `json:"nickname"`
	Sex        string `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege  string `json:"privilege"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// ============================================================================

func HandlerWeiXinAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method != "POST" {
		return
	}

	//werver token
	serverToken := r.PostFormValue("servertoken")
	if !checkServerToken(serverToken) {
		makeAuthRes(w, &AuthRes{Result: Err.Failed})
		return
	}

	//auth token
	token := r.PostFormValue("token")
	if len(token) == 0 {
		makeAuthRes(w, &AuthRes{Result: Err.Failed})
		return
	}

	//auth type: default token auth
	err := Err.OK
	authType := r.PostFormValue("authtype")
	if authType != comm.I32toa(macrocode.LoginType_WeiXinCode) { //token auth
		openid := r.PostFormValue("uid")
		err = processWeiXinTokenAuth(w, openid, token)
	} else { //code auth
		err = processWeiXinCodeAuth(w, token)
	}

	if err != Err.OK {
		makeAuthRes(w, &AuthRes{Result: int32(err)})
	}
}

func processWeiXinTokenAuth(w http.ResponseWriter, openid string, token string) int {
	UserInfo := dbmgr.CenterGetUserInfo(macrocode.ChannelType_WeiXin, openid)
	if UserInfo == nil {
		log.Error("processWeiXinTokenAuth get UserInfo failed")
		return Err.Failed
	}

	if token != UserInfo.Token {
		log.Error("processWeiXinTokenAuth token error")
		return Err.Failed
	}

	if time.Now().After(UserInfo.ExpireT) {
		err := updateWeinXinTokenByRefreshToken(w, UserInfo.RefreshToken)
		return err
	} else {
		makeAuthRes(w, &AuthRes{
			Result: Err.OK,
			OpenId: openid,
			Token:  token,
			Expire: UserInfo.ExpireT.Unix(),
		})
	}

	return Err.OK
}

func updateWeinXinTokenByRefreshToken(w http.ResponseWriter, refrtoken string) int {
	url := fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s",
		config.Auth.WeiXin.RefrUrl, config.Auth.WeiXin.AppId, refrtoken)

	ret, err := comm.HttpGetT(url, HttpTimeOutSecond)
	if err != nil {
		log.Error("updateWeinXinTokenByRefreshToken HttpGetT error", err)
		return Err.Failed
	}

	var jret weixinTokenRet
	err = json.Unmarshal([]byte(ret), &jret)
	if err != nil {
		log.Error("updateWeinXinTokenByRefreshToken Unmarshal error", ret, err)
		return Err.Failed
	}

	if jret.ErrCode > 0 {
		log.Error("updateWeinXinTokenByRefreshToken ErrCode", jret.ErrCode, jret.ErrMsg)
		return Err.Failed
	}

	makeAuthRes(w, &AuthRes{
		Result: Err.OK,
		OpenId: jret.OpenId,
		Token:  jret.Token,
		Expire: jret.ExpireIn + time.Now().Unix(),
	})

	//flush
	asyncop.Push(func() {
		dbmgr.CenterUpdateUserToken(macrocode.ChannelType_WeiXin, jret.OpenId, jret.Token, jret.RefrToken, jret.ExpireIn)
	}, nil)

	//update name and headimg
	UpdateWeiXinUserInfo(jret.OpenId, jret.Token)

	return Err.OK
}

func processWeiXinCodeAuth(w http.ResponseWriter, code string) int {
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		config.Auth.WeiXin.TokenUrl, config.Auth.WeiXin.AppId, config.Auth.WeiXin.AppKey, code)

	ret, err := comm.HttpGetT(url, HttpTimeOutSecond)
	if err != nil {
		log.Error("processWeiXinCodeAuth HttpGetT error", err)
		return Err.Failed
	}

	var jret weixinTokenRet
	err = json.Unmarshal([]byte(ret), &jret)
	if err != nil {
		log.Error("processWeiXinCodeAuth Unmarshal error", ret, err)
		return Err.Failed
	}

	if jret.ErrCode > 0 {
		log.Error("processWeiXinCodeAuth ErrCode", jret.ErrCode, jret.ErrMsg)
		return Err.Failed
	}

	makeAuthRes(w, &AuthRes{
		Result: Err.OK,
		OpenId: jret.OpenId,
		Token:  jret.Token,
		Expire: jret.ExpireIn + time.Now().Unix(),
	})

	//flush
	asyncop.Push(func() {
		dbmgr.CenterUpdateUserToken(macrocode.ChannelType_WeiXin, jret.OpenId, jret.Token, jret.RefrToken, jret.ExpireIn)
	}, nil)

	//update name and headimg
	UpdateWeiXinUserInfo(jret.OpenId, jret.Token)

	return Err.OK
}

func UpdateWeiXinUserInfo(token, openid string) {
	url := fmt.Sprintf("%s?access_token=%s&openid=%s", config.Auth.WeiXin.UserInfoUrl, token, openid)

	ret, err := comm.HttpGetT(url, HttpTimeOutSecond)
	if err != nil {
		log.Error("UpdateWeiXinUserInfo HttpGetT error", err)
		return
	}

	var jret weixinUserInfo
	err = json.Unmarshal([]byte(ret), &jret)
	if err != nil {
		log.Error("UpdateWeiXinUserInfo Unmarshal error", ret, err)
		return
	}

	if jret.ErrCode > 0 {
		log.Error("UpdateWeiXinUserInfo ErrCode", jret.ErrCode, jret.ErrMsg)
		return
	}

	//flush
	asyncop.Push(func() {
		dbmgr.CenterUpdateUserNameHead(macrocode.ChannelType_WeiXin, openid, jret.NickName, jret.Headimgurl)
	}, nil)

}

func makeAuthRes(w http.ResponseWriter, ret *AuthRes) {
	jres, err := json.Marshal(ret)
	if err != nil {
		log.Error("makeAuthRes Marshal failed")
		return
	}

	fmt.Fprint(w, string(jres))
}
