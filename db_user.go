package main


import(
    "database/sql"
    "log"
    "fmt"
    
    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"
)

func (u *user) Inputuser(){
    
    db, err := sql.Open("postgres", "dbname=d7uaktin79dcf5 host=ec2-184-72-236-57.compute-1.amazonaws.com port=5432 user=zadrpmccdwzbjq password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=require")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    
    row := db.QueryRow(`SELECT COUNT(name) FROM users WHERE name = $1`,u.Name)
    var namecnt int
    err = row.Scan(&namecnt)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(namecnt)
    
    if namecnt > 0 {
        u.Check = false
	    fmt.Println("そのユーザーIDは既に存在します")
    }else{
        u.Check = true
        hash, err := passwordHash(u.Password)
        if err != nil {
            log.Fatal(err)
        }
        result, err := db.Query(`INSERT INTO users(name, password,years,sex,comment,timestamp) VALUES($1, $2,$3,$4,$5,CURRENT_TIMESTAMP)`,u.Name,hash,u.Year,u.Sex," ")
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println(result)    
    }
}

func (u *user) Access(){
    db, err := sql.Open("postgres", "dbname=d7uaktin79dcf5 host=ec2-184-72-236-57.compute-1.amazonaws.com port=5432 user=zadrpmccdwzbjq password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=require")
    fmt.Println(u)
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    row := db.QueryRow(`SELECT COUNT(name) FROM users WHERE name = $1`,u.Name)
    fmt.Println(u)
    var namecnt int
    err = row.Scan(&namecnt)
    fmt.Println(namecnt)
    if err != nil{
        log.Fatal(err)
    }
    
    if namecnt == 0{
        u.Check = false
        return
    }
    
    row2 := db.QueryRow(`SELECT name, password FROM users WHERE name = $1`,u.Name)
    var name string
    var password string
    err = row2.Scan(&name,&password)
    if err != nil{
        log.Fatal(err)
    }
    
    
	err = passwordVerify(password, u.Password)
	if err != nil {
        fmt.Println(err)
        u.Check = false
        return
	}

	println("認証しました！")
    u.Check = true
}

func (u *user) Delete(){
    db, err := sql.Open("postgres", "dbname=d7uaktin79dcf5 host=ec2-184-72-236-57.compute-1.amazonaws.com port=5432 user=zadrpmccdwzbjq password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=require")
    fmt.Println(u)
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    

    _, err = db.Exec (`delete from users WHERE name = $1`, u.Name )

    if err != nil {
        fmt.Println(err)
        }

    fmt.Println ("*** 終了 ***")
}

func (u *user) Showuser(){
    db, err := sql.Open("postgres", "dbname=d7uaktin79dcf5 host=ec2-184-72-236-57.compute-1.amazonaws.com port=5432 user=zadrpmccdwzbjq password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=require")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    row2 := db.QueryRow(`SELECT sex, years,comment FROM users WHERE name = $1`,u.Name)
    var sex int
    var year int
    var comment string
    err = row2.Scan(&sex,&year,&comment)
    u.Sex = sex
    u.Year = year
    u.Comment = comment
    if err != nil{
        log.Fatal(err)
    }
}

func (u *user)Updateuser(){
    
    db, err := sql.Open("postgres", "dbname=d7uaktin79dcf5 host=ec2-184-72-236-57.compute-1.amazonaws.com port=5432 user=zadrpmccdwzbjq password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=require")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    fmt.Println(u.Name + "      c 7")
    
    hash, err := passwordHash(u.Password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(u)
    result, err := db.Query(`UPDATE users SET password=$1,years=$2,sex=$3,comment=$4,timestamp=CURRENT_TIMESTAMP WHERE name=$5`,hash,u.Year,u.Sex,u.Comment,u.Name)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(result)      
}
// パスワードハッシュを作る
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// パスワードがハッシュにマッチするかどうかを調べる
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
