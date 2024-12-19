package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorSchedule struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	DoctorId uuid.UUID `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	AvailableDay string `json:"available_day,omitempty" column:"name:available_day;type:varchar;nullable:false"`
	StartTime string `json:"start_time,omitempty" column:"name:start_time;type:text;nullable:false"`
	EndTime string `json:"end_time,omitempty" column:"name:end_time;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor_schedule" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor *Doctor `json:"doctor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
}
