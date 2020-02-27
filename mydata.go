package main

import(
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
    "strconv"
)


type clientdata struct{
    //socket はこのクライアントのためのWebSocketです
    socket *websocket.Conn
    // 入力データのJSON
    mydata *mydata
}

type mydata struct{
    forward chan *diarys
}



func newMydata() (*mydata){
    return &mydata{
        forward: make(chan *diarys),
    }
}


func (d *mydata) ServeHTTP(w http.ResponseWriter, r *http.Request){
    socket, err := upgrader.Upgrade(w,r,nil)
    if err != nil{
        log.Fatal("ServeHTTP:",err)
        return
    }
    clientdata:= &clientdata{
        socket: socket,
        mydata: d,
    }
    clientdata.read(r)
}

func (c *clientdata) read(r *http.Request){
    for{
        var diarys *diarys
        if err := c.socket.ReadJSON(&diarys); err == nil{
            if name, err := r.Cookie("name"); err != nil{
                panic(err.Error())
            }else{
                diarys.Name = name.Value
            }
            fmt.Println("koko:",diarys)
            diarys.showDiary()
            var diarydata Diarydata
            diarydata.Day = "00000000000"
            diarydata.Hour = 0
            diarydata.Minute = 0
            diarydata.Topic = ""
            diarydata.Comment = ""
            diarys.Diary = append(diarys.Diary, diarydata)
            diarys.Check = true
            if err := c.socket.WriteJSON(&diarys); err == nil{
                breaker := false
                j := 0;
                for{
                    if err := c.socket.ReadJSON(&diarys); err == nil{
                fmt.Println("soko:",diarys.Day[0:8])
                        for i:=1;i<=diarys.Days;i++{
                            im, _ := strconv.Atoi(diarys.Diary[j].Day[8:10])
                            fmt.Println("im:",diarys.Day[0:8])
                            if(im == i){
                                diarys.Day = diarys.Day[0:8] + diarys.Diary[j].Day[8:10]
                                diarys.Hour = diarys.Diary[j].Hour
                                diarys.Minute = diarys.Diary[j].Minute
                                diarys.Topic = diarys.Diary[j].Topic
                                diarys.Comment = diarys.Diary[j].Comment
                                diarys.Check = false
                                if err := c.socket.WriteJSON(&diarys); err == nil{
                                    j++
                                }else{
                                    break
                                }
                            }else{
                                diarys.Day = diarys.Day[0:8] + fmt.Sprintf("%02d", i)
                                diarys.Hour = 0
                                diarys.Minute = 0
                                diarys.Topic = ""
                                diarys.Comment = ""
                                diarys.Check = false
                                if err := c.socket.WriteJSON(&diarys); err != nil{
                                    break
                                }
                            }
                            breaker = true
                        }
                    }
                if breaker{
                    break
                }
            }
        }else{
            break
        }
        }
    }
    
    fmt.Println("owarimypage")
    c.socket.Close()
}
