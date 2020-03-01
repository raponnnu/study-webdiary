package main

import(
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
    "flag"
    "os"
)

type templateHandler struct{
    once sync.Once
    filename string
    templ *template.Template
}

//ServerHTTPはHTTPリクエストを処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    t.once.Do(func(){
        t.templ = 
        template.Must(template.ParseFiles(filepath.Join("./templates",t.filename)))
    })
    data := map[string]interface{}{
        "Host":r.Host,
    }
    if nameCookie, err := r.Cookie("name"); err == nil{
        if nameCookie.Value != ""{
            data["Name"] = nameCookie.Value
            var usrdef user
        var user *user = &(usrdef)
        user.Name = nameCookie.Value
        user.Showuser()
        if user.Sex==1{
            data["Sex"] = "男性"
        }else if user.Sex ==2{
            data["Sex"] = "女性"
        }else{
            data["Sex"] = "設定なし"
        }
        if user.Year==0{
            data["Year"] = "設定なし"
        }else{
            data["Year"] = user.Year
        }
        data["Comment"] = user.Comment
        }
    }
    t.templ.Execute(w, data)
}

func main(){
    var addr = flag.String("addr",os.Getenv("PORT"),"アプリケーションのアドレス")
    flag.Parse() //フラグを解釈します
    
    makedata := newInputdata()
    logindata := newInputeddata()
    deletedata := newDeletedata()
    editdata := newEditdata()
    editdiary := newEditdiary()
    mydata := newMydata()
    viewdata := newViewdata()
    http.Handle("/start",&templateHandler{filename: "start.html"})
    http.Handle("/makeuser",&templateHandler{filename: "makeuser.html"})
    http.Handle("/makeduser",&templateHandler{filename: "makeduser.html"})
    http.Handle("/login",&templateHandler{filename: "login.html"})
    http.Handle("/mypage",MustPass(&templateHandler{filename: "mypage.html"},"/login"))
    http.Handle("/logout",MustPass(&templateHandler{filename: "logout.html"},"/start"))
    http.Handle("/delete",MustPass(&templateHandler{filename: "delete.html"},"/deleted"))
    http.Handle("/commit",MustPass(&templateHandler{filename: "commit.html"},"/login"))
    http.Handle("/edituser",MustPass(&templateHandler{filename: "edituser.html"},"/login"))
    http.Handle("/editdiary",MustPass(&templateHandler{filename: "editdiary.html"},"/login"))
    http.Handle("/deleted",&templateHandler{filename: "deleted.html"})
    http.Handle("/view",&templateHandler{filename: "view.html"})
    http.Handle("/style.css",&templateHandler{filename: "style.css"})
    http.Handle("/makeruser", makedata)
    http.Handle("/loginuser", logindata)
    http.Handle("/deleteuser", deletedata)
    http.Handle("/editeduser", editdata)
    http.Handle("/editeddiary", editdiary)
    http.Handle("/mydata",mydata)
    http.Handle("/viewdata",viewdata)
    
    //Webサーバーを開始します
    log.Println("Webサーバーを開始します。ポート： ", *addr)
    if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil{
        log.Fatal("ListenAndServe",err)
    }
}
