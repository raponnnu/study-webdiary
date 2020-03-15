package main

import(
    "net/http"
    "log"
    //"fmt"
    "github.com/gorilla/websocket"
)

//var wm http.ResponseWriter

type loginclient struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // userDataはユーザーに関する情報を保持します
    userData map[string]interface{}
    // 入力データのJSON
    inputeddata *inputeddata
}

type inputeddata struct{
    forward chan *user
}

func newInputeddata() *inputeddata{
    return &inputeddata{
        forward: make(chan *user),
    }
}

func (d *inputeddata) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    loginclient := &loginclient{
        socket: socket,
        inputeddata: d,
    }
    loginclient.accessuser(w,r)
}

func (c *loginclient) accessuser(w http.ResponseWriter,r *http.Request){
    for{
        var usr *user
        if err := c.socket.ReadJSON(&usr); err == nil{
            usr.Access()/*
            if usr.Check {
                fmt.Println("cookie入力")
                
                cookie := &http.Cookie{
                    Name: "name",
                    Value: usr.Name,
                }
                //fmt.Println(cookie.Name)
                
                http.SetCookie(w,cookie)
                http.SetCookie(w,&http.Cookie{
                    Name: "password",
                    Value: usr.Password,
                    Path:"/"})
                if name, err := r.Cookie("name"); err == http.ErrNoCookie{
                    //未認証
                    fmt.Println("うんち!",err)
                    w.Header().Set("Location", "/login")
                    w.WriteHeader(http.StatusTemporaryRedirect)
                }else{
                    fmt.Println("name:",name)
                }
            }*/
            if err := c.socket.WriteJSON(&usr); err == nil{
                break
            }
        }else{
            break
        }
    }
    c.socket.Close()
}
/*
func setCookies(w http.ResponseWriter, r *http.Request) {
    for{
        usr := <- loguser
        if usr.Name != ""{
                http.SetCookie(w,&http.Cookie{
                    Name: "name",
                    Value: usr.Name,
                    Path:"/"})
                http.SetCookie(w,&http.Cookie{
                    Name: "password",
                    Value: usr.Password,
                    Path:"/"})
        }
    }
}**/