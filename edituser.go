package main

import(
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
)


type editclient struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // userDataはユーザーに関する情報を保持します
    userData map[string]interface{}
    // 入力データのJSON
    editdata *editdata
}

type editdata struct{
    forward chan *user
}

func newEditdata() *editdata{
    return &editdata{
        forward: make(chan *user),
    }
}


func (d *editdata) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    editclient := &editclient{
        socket: socket,
        editdata: d,
    }
    editclient.read(r)
}

func (c *editclient) read(r *http.Request){
    for{
        var usr *user
        if err := c.socket.ReadJSON(&usr); err == nil{
            if name, err := r.Cookie("name"); err != nil{
                panic(err.Error())
            }else{
                usr.Name = name.Value
            }
            if usr.PageOpen{
                usr.Showuser()
                if err := c.socket.WriteJSON(&usr); err != nil{
                    break
                }
            }else{
                if password, err := r.Cookie("password"); err != nil{
                    panic(err.Error())
                }else{
                    if(usr.OldPassword == password.Value){
                        usr.Check = true;
                        usr.Updateuser()
                    }else{
                        usr.Check = false;
                    }
                    
                    if err := c.socket.WriteJSON(&usr); err != nil{
                        break
                    }
                }
            }
        }else{
            break
        }
    }
    fmt.Println("owari")
    c.socket.Close()
}
