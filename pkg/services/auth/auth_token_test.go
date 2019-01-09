package auth

import (
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/services/sqlstore"
	. "github.com/smartystreets/goconvey/convey"
)

func createTestableUserAuthTokenService(t *testing.T) *UserAuthTokenService {
	t.Helper()

	sqlstore := sqlstore.InitTestDB(t)

	return &UserAuthTokenService{
		SQLStore: sqlstore,
		log:      log.New("test-logger"),
	}
}

func TestUserAuthToken(t *testing.T) {
	Convey("Test user auth token", t, func() {
		tokenService := createTestableUserAuthTokenService(t)

		t := time.Date(2018, 12, 13, 13, 45, 0, 0, time.UTC)
		now = func() time.Time {
			return t
		}

		Convey("Create a token", func() {
			clientIP := "192.168.10.11:1234"
			userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"

			t, err := tokenService.CreateToken(1, clientIP, userAgent)
			So(err, ShouldBeNil)
			So(t, ShouldNotBeNil)

			Convey("Lookup should return token", func() {
				userId, err := tokenService.LookupToken(t.unhashedToken)
				So(err, ShouldBeNil)
				So(userId, ShouldEqual, 1)
			})
		})

		// test cases: https://github.com/discourse/discourse/blob/5af9a69a3b9df4b08ebc40c87c4cb9236860fe16/spec/models/user_auth_token_spec.rb

		Convey("Lookup token", func() {
			Convey("When looking up a token that was rotated after expiry time should return error", func() {

			})

			Convey("When looking up a token that was rotated before expiry time should return user id", func() {
				Convey("Should update auth_token_seen and seen_at", func() {

				})
			})

			Convey("When looking up a previous token that was rotated after expiry time should return error", func() {

			})

			Convey("When looking up a previous token that was rotated before expiry time should return user id", func() {
				Convey("Should update auth_token_seen", func() {

				})
			})
		})

		Convey("Refresh token", func() {

		})

		Reset(func() {
			now = time.Now
		})
	})
}

func TestParseIPAddress(t *testing.T) {
	Convey("Test parse ip address", t, func() {
		So(parseIPAddress("192.168.0.140:456"), ShouldEqual, "192.168.0.140")
		So(parseIPAddress("[::1:456]"), ShouldEqual, "127.0.0.1")
	})
}
