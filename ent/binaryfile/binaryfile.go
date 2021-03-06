// Code generated by entc, DO NOT EDIT.

package binaryfile

import (
	"time"
)

const (
	// Label holds the string label denoting the binaryfile type in the database.
	Label = "binary_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the binaryfile in the database.
	Table = "binary_files"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "binary_files"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_owned"
)

// Columns holds all SQL columns for binaryfile fields.
var Columns = []string{
	FieldID,
	FieldFilename,
	FieldBody,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "binary_files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_owned",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	FilenameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
