package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Test struct {
	db.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	Name *string `json:"name,omitempty" column:"name:name;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"test" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
