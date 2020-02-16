package models

import (
	"github.com/jinzhu/gorm"
	"testing"
)

func TestCalendar_Create(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)
	type fields struct {
		Base
		Name   string
		UserId string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "fail no user_id",
			fields: fields{
				Name: "bad calendar",
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Name:   "bad calendar",
				UserId: KnownUserId,
			},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name: "fail user_id foreign key constraint",
			fields: fields{
				Name:   "bad calendar",
				UserId: UnexistingId,
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Name:   tt.fields.Name,
				UserId: tt.fields.UserId,
			}
			if err := c.Create(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalendar_Delete(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)
	type fields struct {
		Base   Base
		Name   string
		UserId string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "fail empty id",
			fields: fields{
				Name:   "dfd",
				UserId: UnexistingId,
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "fail unexisting id",
			fields: fields{
				Base:   Base{ID: UnexistingId},
				Name:   "dfd",
				UserId: UnexistingId,
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Base:   Base{ID: KnownCalendarId},
				Name:   "dfd",
				UserId: KnownUserId,
			},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Base:   tt.fields.Base,
				Name:   tt.fields.Name,
				UserId: tt.fields.UserId,
			}
			if err := c.Delete(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalendar_Update(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)
	type fields struct {
		Base   Base
		Name   string
		UserId string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "fail no calendar id",
			fields: fields{
				Name:   "cavabunga",
				UserId: KnownUserId,
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "fail unexisting calendar id",
			fields: fields{
				Base:   Base{ID: UnexistingId},
				Name:   "cavabunga",
				UserId: KnownUserId,
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "fail unexisting user_id",
			fields: fields{
				Base:   Base{ID: KnownCalendarId},
				Name:   "cavabunga",
				UserId: UnexistingId,
			},
			args:    args{db: db},
			wantErr: true,
		},

		{
			name: "success",
			fields: fields{
				Base:   Base{ID: KnownCalendarId},
				Name:   "cavabunga",
				UserId: KnownUserId,
			},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Base:   tt.fields.Base,
				Name:   tt.fields.Name,
				UserId: tt.fields.UserId,
			}
			if err := c.Update(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalendar_Read(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)
	type fields struct {
		Base   Base
		Name   string
		UserId string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "fail no such calendar",
			fields: fields{
				Base: Base{ID: UnexistingId},
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Base: Base{ID: KnownCalendarId},
			},
			args:    args{db: db},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Base:   tt.fields.Base,
				Name:   tt.fields.Name,
				UserId: tt.fields.UserId,
			}
			if err := c.Read(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
