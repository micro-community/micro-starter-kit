// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: srv/account/proto/account/account.proto

/*
Package go_micro_srv_account is a generated protocol buffer package.

It is generated from these files:
	srv/account/proto/account/account.proto

It has these top-level messages:
	User
	UserRequest
	UserResponse
	UserExistResponse
	UserListQuery
	UserListResponse
	Profile
	ProfileRequest
	ProfileResponse
	ProfileListQuery
	ProfileListResponse
*/
package go_micro_srv_account

import context "context"
import strings "strings"
import time "time"

import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import ptypes1 "github.com/golang/protobuf/ptypes"

import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type UserORM struct {
	CreatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
	Email     string      `gorm:"not null"`
	FirstName string      `gorm:"size:255;not null"`
	Id        string      `gorm:"type:uuid;primary_key"`
	LastName  string      `gorm:"not null"`
	Profile   *ProfileORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UpdatedAt *time.Time  `gorm:"not null"`
	Username  string      `gorm:"size:100;not null"`
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	to.Username = m.Username
	to.FirstName = m.FirstName
	to.LastName = m.LastName
	to.Email = m.Email
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	to.Username = m.Username
	to.FirstName = m.FirstName
	to.LastName = m.LastName
	to.Email = m.Email
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

type ProfileORM struct {
	Avatar    string
	Birthday  *time.Time `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
	Gender    string
	Id        string `gorm:"type:uuid;primary_key"`
	Tz        string
	UpdatedAt *time.Time `gorm:"not null"`
	UserId    *string
}

// TableName overrides the default tablename generated by GORM
func (ProfileORM) TableName() string {
	return "profiles"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Profile) ToORM(ctx context.Context) (ProfileORM, error) {
	to := ProfileORM{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	to.Tz = m.Tz
	to.Avatar = m.Avatar
	to.Gender = m.Gender
	if m.Birthday != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.Birthday); err != nil {
			return to, err
		}
		to.Birthday = &t
	}
	if posthook, ok := interface{}(m).(ProfileWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ProfileORM) ToPB(ctx context.Context) (Profile, error) {
	to := Profile{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	to.Tz = m.Tz
	to.Avatar = m.Avatar
	to.Gender = m.Gender
	if m.Birthday != nil {
		if to.Birthday, err = ptypes1.TimestampProto(*m.Birthday); err != nil {
			return to, err
		}
	}
	if posthook, ok := interface{}(m).(ProfileWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Profile the arg will be the target, the caller the one being converted from

// ProfileBeforeToORM called before default ToORM code
type ProfileWithBeforeToORM interface {
	BeforeToORM(context.Context, *ProfileORM) error
}

// ProfileAfterToORM called after default ToORM code
type ProfileWithAfterToORM interface {
	AfterToORM(context.Context, *ProfileORM) error
}

// ProfileBeforeToPB called before default ToPB code
type ProfileWithBeforeToPB interface {
	BeforeToPB(context.Context, *Profile) error
}

// ProfileAfterToPB called after default ToPB code
type ProfileWithAfterToPB interface {
	AfterToPB(context.Context, *Profile) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadUser executes a basic gorm read call
func DefaultReadUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm1.DB) error
}

// DefaultStrictUpdateUser clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	filterProfile := ProfileORM{}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	filterProfile.UserId = new(string)
	*filterProfile.UserId = ormObj.Id
	if err = db.Where(filterProfile).Delete(ProfileORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	var updatedProfile bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"DeletedAt" {
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"FirstName" {
			patchee.FirstName = patcher.FirstName
			continue
		}
		if f == prefix+"LastName" {
			patchee.LastName = patcher.LastName
			continue
		}
		if f == prefix+"Email" {
			patchee.Email = patcher.Email
			continue
		}
		if !updatedProfile && strings.HasPrefix(f, prefix+"Profile.") {
			updatedProfile = true
			if patcher.Profile == nil {
				patchee.Profile = nil
				continue
			}
			if patchee.Profile == nil {
				patchee.Profile = &Profile{}
			}
			if o, err := DefaultApplyFieldMaskProfile(ctx, patchee.Profile, patcher.Profile, &field_mask1.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"Profile.", db); err != nil {
				return nil, err
			} else {
				patchee.Profile = o
			}
			continue
		}
		if f == prefix+"Profile" {
			updatedProfile = true
			patchee.Profile = patcher.Profile
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm1.DB) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]UserORM) error
}

// DefaultCreateProfile executes a basic gorm create call
func DefaultCreateProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ProfileORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadProfile executes a basic gorm read call
func DefaultReadProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &ProfileORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ProfileORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ProfileORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ProfileORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteProfile(ctx context.Context, in *Profile, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ProfileORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ProfileORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteProfileSet(ctx context.Context, in []*Profile, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ProfileORM{})).(ProfileORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ProfileORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ProfileORM{})).(ProfileORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ProfileORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Profile, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Profile, *gorm1.DB) error
}

// DefaultStrictUpdateProfile clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateProfile")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ProfileORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type ProfileORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchProfile executes a basic gorm update call with patch behavior
func DefaultPatchProfile(ctx context.Context, in *Profile, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Profile
	var err error
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadProfile(ctx, &Profile{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskProfile(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateProfile(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ProfileWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ProfileWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultApplyFieldMaskProfile patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskProfile(ctx context.Context, patchee *Profile, patcher *Profile, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Profile, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"DeletedAt" {
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if f == prefix+"Tz" {
			patchee.Tz = patcher.Tz
			continue
		}
		if f == prefix+"Avatar" {
			patchee.Avatar = patcher.Avatar
			continue
		}
		if f == prefix+"Gender" {
			patchee.Gender = patcher.Gender
			continue
		}
		if f == prefix+"Birthday" {
			patchee.Birthday = patcher.Birthday
			continue
		}
		if f == prefix+"Age" {
			patchee.Age = patcher.Age
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListProfile executes a gorm list call
func DefaultListProfile(ctx context.Context, db *gorm1.DB) ([]*Profile, error) {
	in := Profile{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &ProfileORM{}, &Profile{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ProfileORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Profile{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ProfileORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]ProfileORM) error
}
