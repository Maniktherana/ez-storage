package model

import "time"

type ObjAttrs struct {
	Name    string
	Size    int64
	Created time.Time
	Updatd  time.Time
}
