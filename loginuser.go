package main
import(
    "net/http"
    "fmt"
)

type passHandler struct{
    next http.Handler
    returnpass string
}

func (h *passHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    var usrdef user
    var usr *user = &(usrdef)
    
    if name, err := r.Cookie("name"); err == http.ErrNoCookie{
        //未認証
        w.Header().Set("Location", h.returnpass)
        w.WriteHeader(http.StatusTemporaryRedirect)
    }else if err !=nil{
        // 何らかの別のエラーが発生
        panic(err.Error())
    }else{
        usr.Name = name.Value
        fmt.Println(" 000010 " , usr.Name)
    }
    if password, err := r.Cookie("password"); err == http.ErrNoCookie{
        //未認証
        w.Header().Set("Location",h.returnpass)
        w.WriteHeader(http.StatusTemporaryRedirect)
    }else if err !=nil{
        // 何らかの別のエラーが発生
        panic(err.Error())
    }else{
        usr.Password = password.Value
        fmt.Println(" 000010 " , usr.Password)
    }
    usr.Access()
    
    if !usr.Check{
        //未認証
        w.Header().Set("Location", h.returnpass)
        w.WriteHeader(http.StatusTemporaryRedirect)
    }else{
        h.next.ServeHTTP(w,r)
    }
}

func MustPass(handler http.Handler,defpass string ) http.Handler{
    return &passHandler{next: handler, returnpass:defpass}
}


