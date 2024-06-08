// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type MenuTime string

const (
	MenuTimeBreakfast MenuTime = "breakfast"
	MenuTimeLunch     MenuTime = "lunch"
	MenuTimeDinner    MenuTime = "dinner"
	MenuTimeSnack     MenuTime = "snack"
)

func (e *MenuTime) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MenuTime(s)
	case string:
		*e = MenuTime(s)
	default:
		return fmt.Errorf("unsupported scan type for MenuTime: %T", src)
	}
	return nil
}

type NullMenuTime struct {
	MenuTime MenuTime `json:"menu_time"`
	Valid    bool     `json:"valid"` // Valid is true if MenuTime is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMenuTime) Scan(value interface{}) error {
	if value == nil {
		ns.MenuTime, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MenuTime.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMenuTime) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MenuTime), nil
}

type Status string

const (
	StatusPending  Status = "pending"
	StatusComplete Status = "complete"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type DiaryDay struct {
	ID        int64              `db:"id" json:"id"`
	Date      pgtype.Date        `db:"date" json:"date"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type DiaryDayView struct {
	DiaryDayID int64       `db:"diary_day_id" json:"diary_day_id"`
	DiaryDate  pgtype.Date `db:"diary_date" json:"diary_date"`
	DiaryMenus []DiaryMenu `db:"diary_menus" json:"diary_menus"`
}

type DiaryMenu struct {
	ID         int64              `db:"id" json:"id"`
	DiaryDayID *int64             `db:"diary_day_id" json:"diary_day_id"`
	Date       pgtype.Date        `db:"date" json:"date"`
	Img        *string            `db:"img" json:"img"`
	Summary    *string            `db:"summary" json:"summary"`
	TotalCal   *int32             `db:"total_cal" json:"total_cal"`
	Status     NullStatus         `db:"status" json:"status"`
	MenuTime   NullMenuTime       `db:"menu_time" json:"menu_time"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}
