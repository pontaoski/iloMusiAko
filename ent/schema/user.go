package schema

import (
	"context"
	lent "iloMusiAko/ent"
	"iloMusiAko/ent/hook"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("DiscordID").Unique(),
		field.Uint64("Games").Default(0),
		field.Uint64("Points").Default(0),
		field.Uint64("Rating").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(c context.Context, m *lent.UserMutation) (ent.Value, error) {
					id, ok := m.ID()
					if !ok {
						goto pini
					}
					{
						user := m.Client().User.GetX(c, id)

						m.SetRating(1000 * user.Points / (user.Games + 5))
					}
				pini:
					return next.Mutate(c, m)
				})
			},
			ent.OpUpdateOne,
		),
	}
}
