package dbuser

import (
	"flag"
	"testing"

	persistancev1 "github.com/infobloxopen/db-controller/api/v1"
)

// The following gingo struct and associted init() is required to run go test with ginkgo related flags
// Since this test is not using ginkgo, this is a hack to get around the issue of go test complaining about
// unknown flags.
var ginkgo struct {
	dry_run      string
	label_filter string
}

func init() {
	flag.StringVar(&ginkgo.dry_run, "ginkgo.dry-run", "", "Ignore this flag")
	flag.StringVar(&ginkgo.label_filter, "ginkgo.label-filter", "", "Ignore this flag")
}

func TestDBUser_IsUserChanged(t *testing.T) {
	type mockDBUser struct {
		role  string
		userA string
		userB string
	}
	type args struct {
		dbClaim *persistancev1.DatabaseClaim
	}
	tests := []struct {
		name   string
		dbuser mockDBUser
		args   args
		want   bool
	}{
		{
			"User unchanged",
			mockDBUser{
				role: "oldUser",
			},
			args{
				dbClaim: &persistancev1.DatabaseClaim{
					Status: persistancev1.DatabaseClaimStatus{
						ActiveDB: persistancev1.Status{ConnectionInfo: &persistancev1.DatabaseClaimConnectionInfo{
							Username: "oldUser",
						},
						},
					},
				},
			},
			false,
		},
		{
			"User unchanged",
			mockDBUser{
				role: "oldUser",
			},
			args{
				dbClaim: &persistancev1.DatabaseClaim{
					Status: persistancev1.DatabaseClaimStatus{
						ActiveDB: persistancev1.Status{ConnectionInfo: &persistancev1.DatabaseClaimConnectionInfo{
							Username: "",
						},
						},
					},
				},
			},
			false,
		},
		{
			"User changed",
			mockDBUser{
				role: "newUser",
			},
			args{
				dbClaim: &persistancev1.DatabaseClaim{
					Spec: persistancev1.DatabaseClaimSpec{
						Username: "newUser",
					},
					Status: persistancev1.DatabaseClaimStatus{
						ActiveDB: persistancev1.Status{ConnectionInfo: &persistancev1.DatabaseClaimConnectionInfo{
							Username: "oldUser",
						},
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DBUser{
				rolename: tt.dbuser.role,
			}
			if got := d.IsUserChanged(tt.args.dbClaim.Status.ActiveDB); got != tt.want {
				t.Errorf("isUserChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}
