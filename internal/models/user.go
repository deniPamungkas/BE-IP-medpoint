package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type User struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	Username string `json:"username,omitempty" column:"name:username;type:text;nullable:false"`
	Email string `json:"email,omitempty" column:"name:email;type:text;nullable:false;unique"`
	Password string `json:"password,omitempty" column:"name:password;type:text;nullable:false"`
	Role string `json:"role,omitempty" column:"name:role;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"user" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
