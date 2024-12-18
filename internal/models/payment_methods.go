package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type PaymentMethods struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:auth.uid()"`
	Method string `json:"method,omitempty" column:"name:method;type:text;nullable:false"`
	Name string `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	AccountNumber float64 `json:"account_number,omitempty" column:"name:account_number;type:numeric;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payment_methods" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Payments []*Payments `json:"payments,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:payment_method_id"`
	Reservations []*Reservations `json:"reservations,omitempty" join:"joinType:manyToMany;through:payments;sourcePrimaryKey:id;sourceForeignKey:payment_method_id;targetPrimaryKey:id;targetForeign:payment_method_id"`
}
