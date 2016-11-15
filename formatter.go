package formatter

import "time"

func Jst(t time.Time) string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	n := t.In(jst)
	return n.Format("2006/01/02 15:04:05")
}
