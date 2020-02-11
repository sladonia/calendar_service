package models

import "testing"

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
