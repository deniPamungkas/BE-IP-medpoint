package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctor struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UserId uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false;default:gen_random_uuid()"`
	Name string `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Gender string `json:"gender,omitempty" column:"name:gender;type:text;nullable:false"`
	SpecialistId uuid.UUID `json:"specialist_id,omitempty" column:"name:specialist_id;type:uuid;nullable:false;default:gen_random_uuid()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Patients []*Patients `json:"patients,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	Payments []*Payments `json:"payments,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	DoctorScheduleDoctors []*DoctorSchedule `json:"doctor_schedule_doctors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	ReservationDoctors []*Reservations `json:"reservation_doctors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	Specialist *Specialist `json:"specialist,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:specialist_id"`
}
