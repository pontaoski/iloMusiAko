// Code generated by entc, DO NOT EDIT.

package user

import (
	"iloMusiAko/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DiscordID applies equality check predicate on the "DiscordID" field. It's identical to DiscordIDEQ.
func DiscordID(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordID), v))
	})
}

// Games applies equality check predicate on the "Games" field. It's identical to GamesEQ.
func Games(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGames), v))
	})
}

// Points applies equality check predicate on the "Points" field. It's identical to PointsEQ.
func Points(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// Rating applies equality check predicate on the "Rating" field. It's identical to RatingEQ.
func Rating(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// DiscordIDEQ applies the EQ predicate on the "DiscordID" field.
func DiscordIDEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordID), v))
	})
}

// DiscordIDNEQ applies the NEQ predicate on the "DiscordID" field.
func DiscordIDNEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscordID), v))
	})
}

// DiscordIDIn applies the In predicate on the "DiscordID" field.
func DiscordIDIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscordID), v...))
	})
}

// DiscordIDNotIn applies the NotIn predicate on the "DiscordID" field.
func DiscordIDNotIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscordID), v...))
	})
}

// DiscordIDGT applies the GT predicate on the "DiscordID" field.
func DiscordIDGT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscordID), v))
	})
}

// DiscordIDGTE applies the GTE predicate on the "DiscordID" field.
func DiscordIDGTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscordID), v))
	})
}

// DiscordIDLT applies the LT predicate on the "DiscordID" field.
func DiscordIDLT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscordID), v))
	})
}

// DiscordIDLTE applies the LTE predicate on the "DiscordID" field.
func DiscordIDLTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscordID), v))
	})
}

// GamesEQ applies the EQ predicate on the "Games" field.
func GamesEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGames), v))
	})
}

// GamesNEQ applies the NEQ predicate on the "Games" field.
func GamesNEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGames), v))
	})
}

// GamesIn applies the In predicate on the "Games" field.
func GamesIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGames), v...))
	})
}

// GamesNotIn applies the NotIn predicate on the "Games" field.
func GamesNotIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGames), v...))
	})
}

// GamesGT applies the GT predicate on the "Games" field.
func GamesGT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGames), v))
	})
}

// GamesGTE applies the GTE predicate on the "Games" field.
func GamesGTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGames), v))
	})
}

// GamesLT applies the LT predicate on the "Games" field.
func GamesLT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGames), v))
	})
}

// GamesLTE applies the LTE predicate on the "Games" field.
func GamesLTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGames), v))
	})
}

// PointsEQ applies the EQ predicate on the "Points" field.
func PointsEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// PointsNEQ applies the NEQ predicate on the "Points" field.
func PointsNEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoints), v))
	})
}

// PointsIn applies the In predicate on the "Points" field.
func PointsIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPoints), v...))
	})
}

// PointsNotIn applies the NotIn predicate on the "Points" field.
func PointsNotIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPoints), v...))
	})
}

// PointsGT applies the GT predicate on the "Points" field.
func PointsGT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPoints), v))
	})
}

// PointsGTE applies the GTE predicate on the "Points" field.
func PointsGTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPoints), v))
	})
}

// PointsLT applies the LT predicate on the "Points" field.
func PointsLT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPoints), v))
	})
}

// PointsLTE applies the LTE predicate on the "Points" field.
func PointsLTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPoints), v))
	})
}

// RatingEQ applies the EQ predicate on the "Rating" field.
func RatingEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// RatingNEQ applies the NEQ predicate on the "Rating" field.
func RatingNEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRating), v))
	})
}

// RatingIn applies the In predicate on the "Rating" field.
func RatingIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRating), v...))
	})
}

// RatingNotIn applies the NotIn predicate on the "Rating" field.
func RatingNotIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRating), v...))
	})
}

// RatingGT applies the GT predicate on the "Rating" field.
func RatingGT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRating), v))
	})
}

// RatingGTE applies the GTE predicate on the "Rating" field.
func RatingGTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRating), v))
	})
}

// RatingLT applies the LT predicate on the "Rating" field.
func RatingLT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRating), v))
	})
}

// RatingLTE applies the LTE predicate on the "Rating" field.
func RatingLTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRating), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
