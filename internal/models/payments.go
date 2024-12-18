package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Payments struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:auth.uid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	PaymentStatus string `json:"payment_status,omitempty" column:"name:payment_status;type:text;nullable:false"`
	ReservationId uuid.UUID `json:"reservation_id,omitempty" column:"name:reservation_id;type:uuid;nullable:false;default:gen_random_uuid()"`
	PaymentMethodId uuid.UUID `json:"payment_method_id,omitempty" column:"name:payment_method_id;type:uuid;nullable:false;default:gen_random_uuid()"`
	PatmentDate time.Time `json:"patment_date,omitempty" column:"name:patment_date;type:timestampz;nullable:false"`
	Amount float64 `json:"amount,omitempty" column:"name:amount;type:numeric;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payments" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Patients []*Patients `json:"patients,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:payment_id;targetPrimaryKey:id;targetForeign:payment_id"`
	Doctors []*Doctor `json:"doctors,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:payment_id;targetPrimaryKey:id;targetForeign:payment_id"`
	PaymentMethodPayment *PaymentMethods `json:"payment_method_payment,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:payment_method_id"`
	Reservation *Reservations `json:"reservation,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
	ReservationPayments []*Reservations `json:"reservation_payments,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:payment_id"`
}
