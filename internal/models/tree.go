package models

type Tree struct {
	GisID          int64   `db:"gis_id" json:"gis_id,omitempty" csv:"Номер дерева (ГИС/Прибит на ствол),omitempty"`          // Номер дерева (ГИС/Прибит на ствол)
	PcdName        string  `db:"pcd_name" json:"pcd_name,omitempty" csv:"Номер .pcd (для проверки),omitempty"`               // Номер .pcd (для проверки)
	XCoordinate    float64 `db:"x_coordinate" json:"x_coordinate,omitempty" csv:"Координата Х,omitempty"`                    // Координата Х
	YCoordinate    float64 `db:"y_coordinate" json:"y_coordinate,omitempty" csv:"Координата Y,omitempty"`                    // Координата Y
	GISHeightMitro float64 `db:"gis_height_mitro" json:"gis_height_mitro,omitempty" csv:"Высота ГИС (Митрофанов),omitempty"` // Высота ГИС (Митрофанов)
	GISHeightIl    float64 `db:"gis_height_il" json:"gis_height_il,omitempty" csv:"Высота (Илья),omitempty"`                 // Высота (Илья)
	OrderNumber    int64   `db:"order_number" json:"order_number,omitempty" csv:"№,omitempty"`                               // №
	TreeType       string  `db:"tree_type" json:"tree_type,omitempty" csv:"порода,omitempty"`                                // порода
	Circle         float64 `db:"circle" json:"circle,omitempty" csv:"Окр.,omitempty"`                                        // Окр.
	DiameterMitro  float64 `db:"diameter_mitro" json:"diameter_mitro,omitempty" csv:"Диаметр (Митрофанов),omitempty"`        // Диаметр, см (Митрофанов)
	DiameterIl     float64 `db:"diameter_il" json:"diameter_il,omitempty" csv:"Диаметр (Илья),omitempty"`                    // Диаметр, м (Илья)
}
