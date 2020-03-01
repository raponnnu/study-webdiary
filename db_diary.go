package main

import(
    "database/sql"
    "log"
    "fmt"
    
    _ "github.com/lib/pq"
)

func (d *diary) inputDiary(){
    
    db, err := sql.Open("postgres", "user=zadrpmccdwzbjq dbname=d7uaktin79dcf5 password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=disable")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    
    row := db.QueryRow(`SELECT COUNT(name) FROM diary WHERE name = $1 AND day=$2`,d.Name,d.Day)
    var diarycnt int
    err = row.Scan(&diarycnt)
    if err != nil{
        log.Fatal(err)
    }
    
    if diarycnt > 0 {
        result, err := db.Query(`UPDATE diary SET hours=$1,minutes=$2,topic=$3,comment=$4,timestamp=CURRENT_TIMESTAMP WHERE name=$5 and day=$6`,d.Hour,d.Minute,d.Topic, d.Comment,d.Name,d.Day)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println(result)      
    }else{
        result, err := db.Query(`INSERT INTO diary(name, day,hours,minutes,topic,comment,timestamp) VALUES($1, $2,$3,$4,$5,$6,CURRENT_TIMESTAMP)`,d.Name,d.Day,d.Hour,d.Minute,d.Topic, d.Comment)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println(result)    
    }
}

func (d *diary) outputDiary(){
    
    db, err := sql.Open("postgres", "user=zadrpmccdwzbjq dbname=d7uaktin79dcf5 password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=disable")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    
    row := db.QueryRow(`SELECT COUNT(name) FROM diary WHERE name = $1 AND day=$2`,d.Name,d.Day)
    var diarycnt int
    err = row.Scan(&diarycnt)
    if err != nil{
        log.Fatal(err)
    }
    
    if diarycnt > 0 {
        row2 := db.QueryRow(`SELECT hours,minutes,topic,comment FROM diary WHERE name = $1 AND day=$2`,d.Name,d.Day)
        
        var hours int
        var minutes int
        var topic string
        var comment string
        err = row2.Scan(&hours,&minutes,&topic,&comment)
        d.Hour = hours
        d.Minute = minutes
        d.Topic = topic
        d.Comment = comment
        if err != nil{
            log.Fatal(err)
        }
    }else{
        return
    }
}

func (d *diarys) showDiary(){
    
    db, err := sql.Open("postgres", "user=zadrpmccdwzbjq dbname=d7uaktin79dcf5 password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=disable")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
    
    rows, err := db.Query(`SELECT day,hours,minutes,topic,comment FROM diary WHERE name = $1 AND date_part('year',day) = $2 AND date_part('month',day) = $3 `,d.Name,d.Day[:4],d.Day[5:7])
    
    
    defer rows.Close()
    for rows.Next() {
        var day string
        var hours int
        var minutes int
        var topic string
        var comment string
        err = rows.Scan(&day,&hours,&minutes,&topic,&comment)
        var diarydata Diarydata
        diarydata.Day = day
        diarydata.Hour = hours
        diarydata.Minute = minutes
        diarydata.Topic = topic
        diarydata.Comment = comment
        d.Diary = append(d.Diary, diarydata)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println(d.Diary)
    }
    return
}

func (d *diarys) showUser(){
    
    db, err := sql.Open("postgres", "user=zadrpmccdwzbjq dbname=d7uaktin79dcf5 password=b1f0d74eb636f9014724190aa77bca3548f482ff11fbd800251b1b48b3aaf1fc sslmode=disable")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()
        
    row := db.QueryRow(`SELECT COUNT(name) FROM diary WHERE name = $1`,d.Name)
    var diarycnt int
    err = row.Scan(&diarycnt)
    if err != nil{
        log.Fatal(err)
    }
    
    if diarycnt > 0 {
        row2 := db.QueryRow(`SELECT sex, years,comment FROM users WHERE name = $1`,d.Name)
        var sex int
        var year int
        var comment string
        err = row2.Scan(&sex,&year,&comment)
        d.Sex = sex
        d.Year = year
        d.UserComment = comment
        if err != nil{
            log.Fatal(err)
        }
    }else{
        d.UserCheck = false
    }
}
