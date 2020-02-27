package main
import(
    "net/http"
    "fmt"
    "github.com/gorilla/websocket"
    "log"
)
type deleteclient struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // userDataはユーザーに関する情報を保持します
    userData map[string]interface{}
    // 入力データのJSON
    deletedata *deletedata
}

type deletedata struct{
    forward chan *user
}

func newDeletedata() *deletedata{
    return &deletedata{
        forward: make(chan *user),
    }
}

func (d *deletedata) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    deleteclient := &deleteclient{
        socket: socket,
        deletedata: d,
    }
    deleteclient.deleteuser(w,r)
}

func (c *deleteclient) deleteuser(w http.ResponseWriter,r *http.Request){
    for{
        var usr *user
        if err := c.socket.ReadJSON(&usr); err == nil{
    
            if name, err := r.Cookie("name"); err == http.ErrNoCookie{
                //未認証
                w.Header().Set("Location", "/start")
                w.WriteHeader(http.StatusTemporaryRedirect)
            }else if err !=nil{
                // 何らかの別のエラーが発生
                panic(err.Error())
            }else{
                usr.Name = name.Value
                fmt.Println(" 000010 " , usr.Name)
            }
            usr.Access()
            if usr.Check{
                usr.Delete()
            }
            if err := c.socket.WriteJSON(&usr); err == nil{
                break
            }
        }else{
            break
        }
    }
    c.socket.Close()
}

