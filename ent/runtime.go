// Code generated by entc, DO NOT EDIT.

package ent

import (
	"sample/ent/binaryfile"
	"sample/ent/schema"
	"sample/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	binaryfileFields := schema.BinaryFile{}.Fields()
	_ = binaryfileFields
	// binaryfileDescFilename is the schema descriptor for filename field.
	binaryfileDescFilename := binaryfileFields[0].Descriptor()
	// binaryfile.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	binaryfile.FilenameValidator = binaryfileDescFilename.Validators[0].(func(string) error)
	// binaryfileDescCreatedAt is the schema descriptor for created_at field.
	binaryfileDescCreatedAt := binaryfileFields[2].Descriptor()
	// binaryfile.DefaultCreatedAt holds the default value on creation for the created_at field.
	binaryfile.DefaultCreatedAt = binaryfileDescCreatedAt.Default.(func() time.Time)
	// binaryfileDescUpdatedAt is the schema descriptor for updated_at field.
	binaryfileDescUpdatedAt := binaryfileFields[3].Descriptor()
	// binaryfile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	binaryfile.DefaultUpdatedAt = binaryfileDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
