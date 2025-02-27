// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnaplockLog snaplock log
//
// swagger:model snaplock_log
type SnaplockLog struct {

	// links
	Links *SnaplockLogLinks `json:"_links,omitempty"`

	// log archive
	LogArchive *SnaplockLogLogArchive `json:"log_archive,omitempty"`

	// log files
	LogFiles []*SnaplockLogFile `json:"log_files,omitempty"`

	// log volume
	LogVolume *SnaplockLogVolume `json:"log_volume,omitempty"`

	// svm
	Svm *SnaplockLogSvm `json:"svm,omitempty"`
}

// Validate validates this snaplock log
func (m *SnaplockLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLog) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) validateLogArchive(formats strfmt.Registry) error {
	if swag.IsZero(m.LogArchive) { // not required
		return nil
	}

	if m.LogArchive != nil {
		if err := m.LogArchive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) validateLogFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.LogFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.LogFiles); i++ {
		if swag.IsZero(m.LogFiles[i]) { // not required
			continue
		}

		if m.LogFiles[i] != nil {
			if err := m.LogFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnaplockLog) validateLogVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.LogVolume) { // not required
		return nil
	}

	if m.LogVolume != nil {
		if err := m.LogVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_volume")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log based on the context it is used
func (m *SnaplockLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogArchive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLog) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) contextValidateLogArchive(ctx context.Context, formats strfmt.Registry) error {

	if m.LogArchive != nil {
		if err := m.LogArchive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) contextValidateLogFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LogFiles); i++ {

		if m.LogFiles[i] != nil {
			if err := m.LogFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnaplockLog) contextValidateLogVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.LogVolume != nil {
		if err := m.LogVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_volume")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLog) UnmarshalBinary(b []byte) error {
	var res SnaplockLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogLinks snaplock log links
//
// swagger:model SnaplockLogLinks
type SnaplockLogLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snaplock log links
func (m *SnaplockLogLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log links based on the context it is used
func (m *SnaplockLogLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogLogArchive snaplock log log archive
//
// swagger:model SnaplockLogLogArchive
type SnaplockLogLogArchive struct {

	// links
	Links *SnaplockLogLogArchiveLinks `json:"_links,omitempty"`

	// Archive the specified SnapLock log file for the given base_name, and create a new log file. If base_name is not mentioned, archive all log files.
	Archive *bool `json:"archive,omitempty"`

	// Base name of log file
	// Enum: [legal_hold privileged_delete system]
	BaseName string `json:"base_name,omitempty"`

	// Expiry time of the log file in date-time format. Value '9999-12-31T00:00:00Z' indicates infinite expiry time.
	// Example: 2058-06-04T19:00:00Z
	// Read Only: true
	// Format: date-time
	ExpiryTime *strfmt.DateTime `json:"expiry_time,omitempty"`

	// Absolute path of the log file in the volume
	// Example: /snaplock_log/system_logs/20180822_005947_GMT-present
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Size of the log file in bytes
	// Example: 20000
	// Read Only: true
	Size int64 `json:"size,omitempty"`
}

// Validate validates this snaplock log log archive
func (m *SnaplockLogLogArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLogArchive) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var snaplockLogLogArchiveTypeBaseNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legal_hold","privileged_delete","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snaplockLogLogArchiveTypeBaseNamePropEnum = append(snaplockLogLogArchiveTypeBaseNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SnaplockLogLogArchive
	// SnaplockLogLogArchive
	// base_name
	// BaseName
	// legal_hold
	// END DEBUGGING
	// SnaplockLogLogArchiveBaseNameLegalHold captures enum value "legal_hold"
	SnaplockLogLogArchiveBaseNameLegalHold string = "legal_hold"

	// BEGIN DEBUGGING
	// SnaplockLogLogArchive
	// SnaplockLogLogArchive
	// base_name
	// BaseName
	// privileged_delete
	// END DEBUGGING
	// SnaplockLogLogArchiveBaseNamePrivilegedDelete captures enum value "privileged_delete"
	SnaplockLogLogArchiveBaseNamePrivilegedDelete string = "privileged_delete"

	// BEGIN DEBUGGING
	// SnaplockLogLogArchive
	// SnaplockLogLogArchive
	// base_name
	// BaseName
	// system
	// END DEBUGGING
	// SnaplockLogLogArchiveBaseNameSystem captures enum value "system"
	SnaplockLogLogArchiveBaseNameSystem string = "system"
)

// prop value enum
func (m *SnaplockLogLogArchive) validateBaseNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snaplockLogLogArchiveTypeBaseNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnaplockLogLogArchive) validateBaseName(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseName) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseNameEnum("log_archive"+"."+"base_name", "body", m.BaseName); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogLogArchive) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("log_archive"+"."+"expiry_time", "body", "date-time", m.ExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snaplock log log archive based on the context it is used
func (m *SnaplockLogLogArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiryTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLogArchive) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLogLogArchive) contextValidateExpiryTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_archive"+"."+"expiry_time", "body", m.ExpiryTime); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogLogArchive) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_archive"+"."+"path", "body", string(m.Path)); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogLogArchive) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_archive"+"."+"size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogLogArchive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogLogArchive) UnmarshalBinary(b []byte) error {
	var res SnaplockLogLogArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogLogArchiveLinks snaplock log log archive links
//
// swagger:model SnaplockLogLogArchiveLinks
type SnaplockLogLogArchiveLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snaplock log log archive links
func (m *SnaplockLogLogArchiveLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLogArchiveLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log log archive links based on the context it is used
func (m *SnaplockLogLogArchiveLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogLogArchiveLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogLogArchiveLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogLogArchiveLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogLogArchiveLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogSvm snaplock log svm
//
// swagger:model SnaplockLogSvm
type SnaplockLogSvm struct {

	// links
	Links *SnaplockLogSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this snaplock log svm
func (m *SnaplockLogSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log svm based on the context it is used
func (m *SnaplockLogSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogSvm) UnmarshalBinary(b []byte) error {
	var res SnaplockLogSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogSvmLinks snaplock log svm links
//
// swagger:model SnaplockLogSvmLinks
type SnaplockLogSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snaplock log svm links
func (m *SnaplockLogSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log svm links based on the context it is used
func (m *SnaplockLogSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogSvmLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
