package main


type diary struct{
    Name string
    Day string
    Hour int
    Minute int
    Topic string
    Comment string
    Input bool
}

type diarys struct{
    Name string
    Sex int
    Year int
    UserComment string
    UserCheck bool
    
    Day string
    Days int
    Hour int
    Minute int
    Topic string
    Comment string
    Check bool
    Diary []Diarydata
}

type Diarydata struct{
    Day string
    Hour int
    Minute int
    Topic string
    Comment string
}
