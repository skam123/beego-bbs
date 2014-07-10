package post

import (
	"time"
)

type Post struct {
	Id int64
	Content string `orm:"type(text)"`
	Created time.Time
	Modified time.Time
}
