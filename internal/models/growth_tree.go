package models

import "time"

type GrowthTree struct {
	GisID    int64     `db:"gis_id" json:"gis_id,omitempty" csv:"Номер дерева (ГИС/Прибит на ствол),omitempty"` // Номер дерева (ГИС/Прибит на ствол)
	TS       time.Time `db:"ts" json:"ts,omitempty" csv:"ts,omitempty"`                                         // Координата Х
	Age      int64     `db:"age" json:"age,omitempty" csv:"age,omitempty"`                                      // Координата Y
	Diameter float64   `db:"diameter" json:"diameter,omitempty" csv:"diameter,omitempty"`                       // Высота ГИС (Митрофанов)
	Height   float64   `db:"height" json:"height,omitempty" csv:"height,omitempty"`                             // Высота (Илья)
	IsAlive  bool      `db:"is_alive" json:"is_alive,omitempty" csv:"is_alive,omitempty"`                       // №
}
