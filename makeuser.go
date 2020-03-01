package main

import(
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/websocket"

)


type client struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // userDataはユーザーに関する情報を保持します
    userData map[string]interface{}
    // 入力データのJSON
    inputdata *inputdata
}

type inputdata struct{
    forward chan *user
}

func newInputdata() *inputdata{
    return &inputdata{
        forward: make(chan *user),
    }
}

const(
    socketBufferSize = 1024
)

var upgrader = &websocket.Upgrader{
    ReadBufferSize:socketBufferSize,
    WriteBufferSize:socketBufferSize,
}

func (d *inputdata) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    client := &client{
        socket: socket,
        inputdata: d,
    }
    client.read()
}

func (c *client) read(){
    for{
        var usr *user
        if err := c.socket.ReadJSON(&usr); err == nil{
            usr.Inputuser()
            if err := c.socket.WriteJSON(&usr); err == nil{
                break
            }
        }else{
            break
        }
    }
    c.socket.Close()
}
