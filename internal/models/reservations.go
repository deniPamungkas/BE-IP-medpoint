package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Reservations struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:auth.uid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	PatientId uuid.UUID `json:"patient_id,omitempty" column:"name:patient_id;type:uuid;nullable:false"`
	DoctorId uuid.UUID `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	PaymentId uuid.UUID `json:"payment_id,omitempty" column:"name:payment_id;type:uuid;nullable:false"`
	Schedule time.Time `json:"schedule,omitempty" column:"name:schedule;type:timestampz;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	PaymentReservations []*Payments `json:"payment_reservations,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	Patient *Patients `json:"patient,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:patient_id"`
	Doctor *Doctor `json:"doctor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Payment *Payments `json:"payment,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:payment_id"`
	PaymentMethods []*PaymentMethods `json:"payment_methods,omitempty" join:"joinType:manyToMany;through:payments;sourcePrimaryKey:id;sourceForeignKey:reservation_id;targetPrimaryKey:id;targetForeign:reservation_id"`
}
