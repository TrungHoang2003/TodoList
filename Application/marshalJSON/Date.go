package marshalJSON

import "time"

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2 Jan")
	return []byte(`"` + formatted + `"`), nil
}
