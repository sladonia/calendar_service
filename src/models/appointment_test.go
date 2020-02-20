package models

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAppointment_Validate(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)
	type fields struct {
		Base        Base
		Subject     string
		Description string
		CalendarId  string
		Start       time.Time
		End         time.Time
		WholeDay    bool
		Users       []*User
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "failure no time",
			fields: fields{
				Subject:     "Meet friends",
				Description: "just have fun",
				CalendarId:  KnownCalendarId,
				WholeDay:    false,
			},
			wantErr: true,
		},
		{
			name: "failure no end time",
			fields: fields{
				Subject:     "Meet friends",
				Description: "just have fun",
				CalendarId:  KnownCalendarId,
				Start:       time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:    false,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Subject:    "to do job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:   true,
			},
			wantErr: false,
		},
		{
			name: "success2",
			fields: fields{
				Subject:    "to do another job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				End:        time.Date(2020, 1, 17, 21, 0, 0, 0, time.UTC),
				WholeDay:   false,
			},
			wantErr: false,
		},
		{
			name: "failure time overlaps",
			fields: fields{
				Subject:    "to do another job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 22, 0, 0, 0, time.UTC),
				End:        time.Date(2020, 1, 17, 21, 0, 0, 0, time.UTC),
				WholeDay:   false,
			},
			wantErr: true,
		},
		{
			name: "failure both whole_day=true and end time provided",
			fields: fields{
				Subject:    "to do another job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				End:        time.Date(2020, 1, 17, 21, 0, 0, 0, time.UTC),
				WholeDay:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Appointment{
				Base:        tt.fields.Base,
				Subject:     tt.fields.Subject,
				Description: tt.fields.Description,
				CalendarId:  tt.fields.CalendarId,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
				WholeDay:    tt.fields.WholeDay,
				Attendees:   tt.fields.Users,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppointment_Create(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	type fields struct {
		Base        Base
		Subject     string
		Description string
		CalendarId  string
		Start       time.Time
		End         time.Time
		WholeDay    bool
		Users       []*User
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
			name: "success",
			fields: fields{
				Subject:    "to do job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:   true,
			},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name: "failure unique index violation",
			fields: fields{
				Subject:    "to do job",
				CalendarId: KnownCalendarId,
				Start:      time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:   true,
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Appointment{
				Base:        tt.fields.Base,
				Subject:     tt.fields.Subject,
				Description: tt.fields.Description,
				CalendarId:  tt.fields.CalendarId,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
				WholeDay:    tt.fields.WholeDay,
				Attendees:   tt.fields.Users,
			}
			if err := a.Create(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppointment_Delete(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	type fields struct {
		Base        Base
		Subject     string
		Description string
		CalendarId  string
		Start       time.Time
		End         time.Time
		WholeDay    bool
		Users       []*User
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
			name: "success",
			fields: fields{
				Base: Base{ID: AppointmentFixedTimeId},
			},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name:    "failure no id",
			fields:  fields{},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name:    "failure unexisting id",
			fields:  fields{Base: Base{ID: UnexistingId}},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Appointment{
				Base:        tt.fields.Base,
				Subject:     tt.fields.Subject,
				Description: tt.fields.Description,
				CalendarId:  tt.fields.CalendarId,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
				WholeDay:    tt.fields.WholeDay,
				Attendees:   tt.fields.Users,
			}
			if err := a.Delete(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppointment_Update(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	type fields struct {
		Base        Base
		Subject     string
		Description string
		CalendarId  string
		Start       time.Time
		End         time.Time
		WholeDay    bool
		Users       []*User
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
			name: "success",
			fields: fields{
				Base:        Base{ID: AppointmentFixedTimeId},
				Subject:     "Meet friends",
				Description: "just have fun",
				CalendarId:  KnownCalendarId,
				Start:       time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:    true,
			},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name: "failure violate index",
			fields: fields{
				Base:        Base{ID: AppointmentFixedTimeId},
				Subject:     "take a rest",
				Description: "just have fun",
				CalendarId:  KnownCalendarId,
				Start:       time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
				WholeDay:    true,
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Appointment{
				Base:        tt.fields.Base,
				Subject:     tt.fields.Subject,
				Description: tt.fields.Description,
				CalendarId:  tt.fields.CalendarId,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
				WholeDay:    tt.fields.WholeDay,
				Attendees:   tt.fields.Users,
			}
			if err := a.Update(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppointment_Read(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	type fields struct {
		Base        Base
		Subject     string
		Description string
		CalendarId  string
		Start       time.Time
		End         time.Time
		WholeDay    bool
		Users       []*User
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
			name:    "success",
			fields:  fields{Base: Base{ID: AppointmentWholeDayId}},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name:    "failure unexisting id",
			fields:  fields{Base: Base{ID: UnexistingId}},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Appointment{
				Base:        tt.fields.Base,
				Subject:     tt.fields.Subject,
				Description: tt.fields.Description,
				CalendarId:  tt.fields.CalendarId,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
				WholeDay:    tt.fields.WholeDay,
				Attendees:   tt.fields.Users,
			}
			if err := a.Read(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppointment_AddAttendees(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	t.Run("success", func(tt *testing.T) {
		userIds := []string{KnownUserId, SecondKnownUserId}
		apt := &Appointment{Base: Base{ID: AppointmentWholeDayId}}
		err := apt.AddAttendees(userIds, db)
		assert.Nil(tt, err)
		assert.Len(tt, apt.Attendees, 2)
	})

	t.Run("fail unexisting appointment id", func(tt *testing.T) {
		userIds := []string{KnownUserId, SecondKnownUserId}
		apt := &Appointment{Base: Base{ID: UnexistingId}}
		err := apt.AddAttendees(userIds, db)
		assert.NotNil(tt, err)
	})

	t.Run("fail unexisting user ids", func(tt *testing.T) {
		userIds := []string{UnexistingId}
		apt := &Appointment{Base: Base{ID: UnexistingId}}
		err := apt.AddAttendees(userIds, db)
		assert.NotNil(tt, err)
	})

	t.Run("fail invalid user ids", func(tt *testing.T) {
		userIds := []string{"ha-ha-ha"}
		apt := &Appointment{Base: Base{ID: UnexistingId}}
		err := apt.AddAttendees(userIds, db)
		assert.NotNil(tt, err)
	})
}

func TestAppointment_RemoveAttendees(t *testing.T) {
	err := MockDbData(db)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer DropAllData(db)

	t.Run("success", func(tt *testing.T) {
		appt := &Appointment{Base: Base{ID: AppointmentWholeDayId}}
		err := appt.RemoveAttendees([]string{ThirdKnownUserId}, db)
		assert.Nil(tt, err)
		assert.Len(tt, appt.Attendees, 0)
	})

	t.Run("success remove unexisting user id", func(tt *testing.T) {
		appt := &Appointment{Base: Base{ID: AppointmentWholeDayId}}
		err := appt.RemoveAttendees([]string{UnexistingId}, db)
		assert.Nil(tt, err)
		assert.Len(tt, appt.Attendees, 0)
	})
}
