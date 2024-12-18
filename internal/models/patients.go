package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Patients struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:auth.uid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UserId uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false;default:gen_random_uuid()"`
	Fullname string `json:"fullname,omitempty" column:"name:fullname;type:text;nullable:false"`
	Address string `json:"address,omitempty" column:"name:address;type:text;nullable:false"`
	Age int16 `json:"age,omitempty" column:"name:age;type:smallint;nullable:false"`
	Gender string `json:"gender,omitempty" column:"name:gender;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"patients" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctors []*Doctor `json:"doctors,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:patient_id;targetPrimaryKey:id;targetForeign:patient_id"`
	Payments []*Payments `json:"payments,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:patient_id;targetPrimaryKey:id;targetForeign:patient_id"`
	ReservationPatients []*Reservations `json:"reservation_patients,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:patient_id"`
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
