package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"liferich/Source/util/log"
	"net/http"
)

var mySigningKey = []byte("AllYourBase")

const (
	WebSite="liferich.com"
)

type LifeClaims struct {
	UserId int //用户ID
	UserEmail string//用户邮箱
	UserPhone string//用户手机
	jwt.StandardClaims
}

func loginHandler(w http.ResponseWriter, r *http.Request)  {
	cookie,err:=r.Cookie("Access-token")
	log.Info(err)
	tokenString:=cookie.Value
	parsetoken,err:=jwt.ParseWithClaims(tokenString,&LifeClaims{}, func(token *jwt.Token)(interface{},error) {
		return mySigningKey,nil
	})
	log.Info("start valid")
	fmt.Println(parsetoken.Valid)
	log.Info("end valid")
	if claims,ok:=parsetoken.Claims.(*LifeClaims);ok&&parsetoken.Valid{
		log.Info(claims)
	}else{
		fmt.Println(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request)  {

	claims := LifeClaims{
		20,
		"317699326@qq.com",
		"18867529162",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
			Issuer:    WebSite,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Info(err)
	cookie:=&http.Cookie{
		Name:"Access-token",
		Value:ss,
		MaxAge:int(time.Now().Add(time.Hour*24).Unix()),
		Path:"/",
		//Domain:"."+token.WebSite,
		Secure:false,
		HttpOnly:false,
	}
	http.SetCookie(w,cookie)
}

func main() {
	http.HandleFunc("/login/",loginHandler)
	http.HandleFunc("/register/",registerHandler)
	http.ListenAndServe(":8888", nil)
}
