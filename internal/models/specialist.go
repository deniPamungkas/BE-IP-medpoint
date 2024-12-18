package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Specialist struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	Specialist string `json:"specialist,omitempty" column:"name:specialist;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"specialist" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorSpecialists []*Doctor `json:"doctor_specialists,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:specialist_id"`
	Users []*Users `json:"users,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:specialist_id;targetPrimaryKey:id;targetForeign:specialist_id"`
}
