// user_repository.go
package repository

import (
	"gormwrap/domain"
	"gormwrap/infrastructure/database"
	"reflect"
	"testing"
)

func TestUserRepositoryGorm_Create(t *testing.T) {
	db := database.TestContainerInit()
	db.DB.AutoMigrate(&domain.User{})

	type fields struct {
		db database.SQLHandler
	}
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			fields: fields{
				db: *db,
			},
			args: args{
				user: &domain.User{
					Name: "John",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepositoryGorm{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryGorm.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepositoryGorm_FindByID(t *testing.T) {
	db := database.TestContainerInit()
	db.DB.AutoMigrate(&domain.User{})
	user := domain.User{
		ID:   1,
		Name: "John",
	}
	db.DB.Create(&user)
	type fields struct {
		db database.SQLHandler
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			fields: fields{
				db: *db,
			},
			args: args{
				id: 1,
			},
			want: &domain.User{
				ID:   1,
				Name: "John",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepositoryGorm{
				db: tt.fields.db,
			}
			got, err := r.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryGorm.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepositoryGorm.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
