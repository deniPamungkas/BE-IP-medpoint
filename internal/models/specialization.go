package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Specialization struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	Specialization string `json:"specialization,omitempty" column:"name:specialization;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"specialization" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
