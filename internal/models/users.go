package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Users struct {
	db.ModelBase
	Id uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	Username string `json:"username,omitempty" column:"name:username;type:text;nullable:false"`
	Email string `json:"email,omitempty" column:"name:email;type:text;nullable:false"`
	Password string `json:"password,omitempty" column:"name:password;type:text;nullable:false"`
	Role string `json:"role,omitempty" column:"name:role;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Specialists []*Specialist `json:"specialists,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	DoctorUsers []*Doctor `json:"doctor_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	PatientUsers []*Patients `json:"patient_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
}
