package main


import(
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
)


type editclientdiary struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // 入力データのJSON
    editdiary *editdiary
}

type editdiary struct{
    forward chan *diary
}

func newEditdiary() *editdiary{
    return &editdiary{
        forward: make(chan *diary),
    }
}


func (d *editdiary) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    editclientdiary := &editclientdiary{
        socket: socket,
        editdiary: d,
    }
    editclientdiary.read(r)
}

func (c *editclientdiary) read(r *http.Request){
    for{
        var diary *diary
        if err := c.socket.ReadJSON(&diary); err == nil{
            if name, err := r.Cookie("name"); err != nil{
                panic(err.Error())
            }else{
                diary.Name = name.Value
            }
            fmt.Println(diary)
            if(diary.Input){
                diary.inputDiary()
            }else{
                diary.outputDiary()
            }
            if err := c.socket.WriteJSON(&diary); err != nil{
                break
            }
        }else{
            break
        }
    }
    fmt.Println("owaridiary")
    c.socket.Close()
}
