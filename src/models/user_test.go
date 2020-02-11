package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		Base      Base
		FirstName string
		LastName  string
		Email     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid_user",
			fields: fields{
				FirstName: "Jorge",
				LastName:  "TheGreat",
				Email:     "gorge@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "invalid_user",
			fields: fields{
				FirstName: "Jorge",
				LastName:  "TheGreat",
				Email:     "gorge@gmailcom",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Base:      tt.fields.Base,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	defer DropAllData(db)
	type fields struct {
		Base      Base
		FirstName string
		LastName  string
		Email     string
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
			name: "success_creation",
			fields: fields{
				FirstName: "Jorge",
				LastName:  "TheGreat",
				Email:     "gorge@gmail.com",
			},
			args:    args{db: db},
			wantErr: false,
		},
		{
			name: "error_user_first_last_name_exists",
			fields: fields{
				FirstName: "Jorge",
				LastName:  "TheGreat",
				Email:     "new_gorge@gmail.com",
			},
			args:    args{db: db},
			wantErr: true,
		},
		{
			name: "error_user_first_last_name_exists",
			fields: fields{
				FirstName: "Gorgio",
				LastName:  "Morod",
				Email:     "gorge@gmail.com",
			},
			args:    args{db: db},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Base:      tt.fields.Base,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if err := u.Create(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	defer DropAllData(db)
	user := &User{
		FirstName: "Jorge",
		LastName:  "TheGreat",
		Email:     "gorge@gmail.com",
	}
	err := user.Delete(db)
	assert.NotNil(t, err)
	user.Create(db)
	err = user.Delete(db)
	assert.Nil(t, err)
}

func TestUser_Update(t *testing.T) {
	defer DropAllData(db)

	user := &User{
		FirstName: "Jorge",
		LastName:  "TheGreat",
		Email:     "gorge@gmail.com",
	}
	err := user.Update(db)
	assert.NotNil(t, err)

	err = user.Create(db)
	assert.Nil(t, err)

	user.FirstName = "Kotlin"
	err = user.Update(db)
	assert.Nil(t, err)

	var user2 User
	db.Find(&user2, "id = ?", user.ID)
	assert.Equal(t, user.FirstName, user2.FirstName)

	user.Email = "adf"
	err = user.Update(db)
	assert.NotNil(t, err)
}

func TestUser_Read(t *testing.T) {
	defer DropAllData(db)

	user := &User{}
	err := user.Read(db)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, EmptyIdError))

	user = &User{
		FirstName: "Jorge",
		LastName:  "TheGreat",
		Email:     "gorge@gmail.com",
	}
	user.Create(db)

	user2 := new(User)
	user2.ID = user.ID
	err = user2.Read(db)
	assert.Nil(t, err)
	assert.Equal(t, user.FirstName, user2.FirstName)
}
