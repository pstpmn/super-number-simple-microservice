package member

import (
	"context"
	"super-number-simple-microservice/configs"
	"super-number-simple-microservice/pkg"
	"testing"
)

func Test_usecase_CreateMember(t *testing.T) {
	// connect mongo
	mongo := pkg.NewMongo()
	uri := "mongodb://root:root@0.0.0.0:27017"
	cfg := configs.NewMongo(uri)
	conn := mongo.Conn(context.Background(), &cfg)

	// warpper repository
	hash := pkg.NewHash()
	memberRepo := NewRepository(conn)

	type args struct {
		name  string
		user  string
		pass  string
		email string
	}
	tests := []struct {
		name    string
		u       usecase
		args    args
		want    *Profile
		wantErr bool
	}{
		{
			name: "create member should be success",
			u: usecase{
				memberRepo: memberRepo,
			},
			args: args{
				name:  "test",
				user:  "test",
				pass:  "test",
				email: "test@test.com",
			},
			want:    &Profile{},
			wantErr: false,
		},
		{
			name: "create member should erorr duplicate email member",
			u: usecase{
				memberRepo: memberRepo,
			},
			args: args{
				name:  "test",
				user:  "notDuplicate",
				pass:  "test",
				email: "test@test.com",
			},
			want:    &Profile{},
			wantErr: true,
		},
		{
			name: "create member should erorr duplicate username member",
			u: usecase{
				memberRepo: memberRepo,
			},
			args: args{
				name:  "test",
				user:  "test",
				pass:  "test",
				email: "notDuplicate@test.com",
			},
			want:    &Profile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase{
				memberRepo: memberRepo,
				jwt:        nil,
				hash:       hash,
			}
			_, err := u.CreateMember(tt.args.name, tt.args.user, tt.args.pass, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CreateMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("usecase.CreateMember() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_usecase_Authentication(t *testing.T) {
	// connect mongo
	mongo := pkg.NewMongo()
	uri := "mongodb://root:root@0.0.0.0:27017"
	cfg := configs.NewMongo(uri)
	conn := mongo.Conn(context.Background(), &cfg)

	// warpper repository
	jwt := pkg.NewJwt([]byte("test"), 123)
	hash := pkg.NewHash()
	memberRepo := NewRepository(conn)

	type fields struct {
		memberRepo IMemberRepository
		jwt        pkg.IJwt
		hash       pkg.IHash
	}
	type args struct {
		user string
		pass string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialCombindProfile
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should be success",
			fields: fields{
				memberRepo: memberRepo,
			},
			args: args{
				user: "test",
				pass: "test",
			},
			// want: &CredentialCombindProfile{
			// 	Profile:    &Profile{},
			// 	Credential: &Credential{},
			// },
			wantErr: false,
		},
		{
			name: "should be error invalid username or password",
			fields: fields{
				memberRepo: memberRepo,
				jwt:        jwt,
			},
			args: args{
				user: "test",
				pass: "test1",
			},
			// want:    &CredentialCombindProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase{
				memberRepo: tt.fields.memberRepo,
				jwt:        jwt,
				hash:       hash,
			}
			_, err := u.Authentication(tt.args.user, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Authentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("usecase.Authentication() = %v, want %v", got, tt.want)
			// }
		})
	}
}
