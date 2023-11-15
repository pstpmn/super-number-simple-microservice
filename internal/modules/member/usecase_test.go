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
		want    *MemberProfile
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
			want:    &MemberProfile{},
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
			want:    &MemberProfile{},
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
			want:    &MemberProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase{
				memberRepo: memberRepo,
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
