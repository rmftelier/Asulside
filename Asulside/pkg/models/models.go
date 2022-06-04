package models

import(
    "errors"
    "time"
)


var ErrNoRecord = errors.New("models: no matching record Found")


type Blog struct{
    ID int
    Title string
    Article string 
    BannerImage string
    PublishedAt time.Time
}