// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeParents holds the string denoting the parents edge name in mutations.
	EdgeParents = "parents"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeFriendships holds the string denoting the friendships edge name in mutations.
	EdgeFriendships = "friendships"
	// EdgeParentHood holds the string denoting the parent_hood edge name in mutations.
	EdgeParentHood = "parent_hood"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "owner_id"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friendships"
	// ParentsTable is the table that holds the parents relation/edge. The primary key declared below.
	ParentsTable = "parents"
	// ChildrenTable is the table that holds the children relation/edge. The primary key declared below.
	ChildrenTable = "parents"
	// FriendshipsTable is the table that holds the friendships relation/edge.
	FriendshipsTable = "friendships"
	// FriendshipsInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsInverseTable = "friendships"
	// FriendshipsColumn is the table column denoting the friendships relation/edge.
	FriendshipsColumn = "user_id"
	// ParentHoodTable is the table that holds the parent_hood relation/edge.
	ParentHoodTable = "parents"
	// ParentHoodInverseTable is the table name for the Parent entity.
	// It exists in this package in order to avoid circular dependency with the "parent" package.
	ParentHoodInverseTable = "parents"
	// ParentHoodColumn is the table column denoting the parent_hood relation/edge.
	ParentHoodColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// ParentsPrimaryKey and ParentsColumn2 are the table columns denoting the
	// primary key for the parents relation (M2M).
	ParentsPrimaryKey = []string{"user_id", "parent_id"}
	// ChildrenPrimaryKey and ChildrenColumn2 are the table columns denoting the
	// primary key for the children relation (M2M).
	ChildrenPrimaryKey = []string{"user_id", "parent_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPetsCount orders the results by pets count.
func ByPetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetsStep(), opts...)
	}
}

// ByPets orders the results by pets terms.
func ByPets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFriendsCount orders the results by friends count.
func ByFriendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendsStep(), opts...)
	}
}

// ByFriends orders the results by friends terms.
func ByFriends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentsCount orders the results by parents count.
func ByParentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParentsStep(), opts...)
	}
}

// ByParents orders the results by parents terms.
func ByParents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFriendshipsCount orders the results by friendships count.
func ByFriendshipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendshipsStep(), opts...)
	}
}

// ByFriendships orders the results by friendships terms.
func ByFriendships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendshipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentHoodCount orders the results by parent_hood count.
func ByParentHoodCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParentHoodStep(), opts...)
	}
}

// ByParentHood orders the results by parent_hood terms.
func ByParentHood(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentHoodStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PetsTable, PetsColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
	)
}
func newFriendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
	)
}
func newParentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ParentsTable, ParentsPrimaryKey...),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ChildrenTable, ChildrenPrimaryKey...),
	)
}
func newFriendshipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendshipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, FriendshipsTable, FriendshipsColumn),
	)
}
func newParentHoodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParentHoodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ParentHoodTable, ParentHoodColumn),
	)
}
