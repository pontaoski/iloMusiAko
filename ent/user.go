// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"iloMusiAko/ent/user"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DiscordID holds the value of the "DiscordID" field.
	DiscordID uint64 `json:"DiscordID,omitempty"`
	// Games holds the value of the "Games" field.
	Games uint64 `json:"Games,omitempty"`
	// Points holds the value of the "Points" field.
	Points uint64 `json:"Points,omitempty"`
	// Rating holds the value of the "Rating" field.
	Rating uint64 `json:"Rating,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // DiscordID
		&sql.NullInt64{}, // Games
		&sql.NullInt64{}, // Points
		&sql.NullInt64{}, // Rating
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field DiscordID", values[0])
	} else if value.Valid {
		u.DiscordID = uint64(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Games", values[1])
	} else if value.Valid {
		u.Games = uint64(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Points", values[2])
	} else if value.Valid {
		u.Points = uint64(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Rating", values[3])
	} else if value.Valid {
		u.Rating = uint64(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", DiscordID=")
	builder.WriteString(fmt.Sprintf("%v", u.DiscordID))
	builder.WriteString(", Games=")
	builder.WriteString(fmt.Sprintf("%v", u.Games))
	builder.WriteString(", Points=")
	builder.WriteString(fmt.Sprintf("%v", u.Points))
	builder.WriteString(", Rating=")
	builder.WriteString(fmt.Sprintf("%v", u.Rating))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
