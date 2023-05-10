package models

import (
	"database/sql"
	"time"
)

type GrowthTree struct {
	GisID    int64    `db:"gis_id" json:"gis_id,omitempty" csv:"gis_id,omitempty"`       // Номер дерева (ГИС/Прибит на ствол)
	TS       DateTime `db:"ts" json:"ts,omitempty" csv:"ts,omitempty"`                   // Координата Х
	Age      int64    `db:"age" json:"age,omitempty" csv:"age,omitempty"`                // Координата Y
	Diameter float64  `db:"diameter" json:"diameter,omitempty" csv:"diameter,omitempty"` // Высота ГИС (Митрофанов)
	Height   float64  `db:"height" json:"height,omitempty" csv:"height,omitempty"`       // Высота (Илья)
	IsAlive  bool     `db:"is_alive" json:"is_alive,omitempty" csv:"is_alive,omitempty"` // №
}

type DateTime struct {
	sql.NullTime
}

// MarshalCSV convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

// You could also use the standard Stringer interface
func (date *DateTime) String() string {
	return date.Time.Format("2006-01-02 15:04:05") // Redundant, just for example
}

// UnmarshalCSV convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}
