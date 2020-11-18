package model

import "time"

// History model
type History struct {
	ID                         uint8                   `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Measure                    float32                 `gorm:"column:measure;type:float unsigned"`
	MeasureUnit                string                  `gorm:"column:measure_unit;type:enum('s', 'kg', 'ms')"`
	MeasureType                string                  `gorm:"column:measure_type;type:enum('weight', 'exerciseMetric')"`
	Date                       time.Time               `gorm:"column:date;type:date"`
	RoutineSpecificExercisesID uint8                   `gorm:"column:routine_specific_exercises_id;type:mediumint unsigned"`
	RoutineSpecificExercise    RoutineSpecificExercise `gorm:"foreignKey:RoutineSpecificExercisesID;references:ID"`
}

// TableName Function to change the name of a table. In this case of History model
func (h *History) TableName() string {
	return "history"
}

// CreateHistory Create a new history point
func (h *History) CreateHistory() string {
	return "CreateHistory"
}

// UpdateHistory Updates an existing history point
func (h *History) UpdateHistory() string {
	return "UpdateHistory"
}