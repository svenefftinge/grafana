package auth

import (
	"net"
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/registry"
	"github.com/grafana/grafana/pkg/services/sqlstore"
)

func init() {
	registry.RegisterService(&UserAuthTokenService{})
}

var now = time.Now

// UserAuthTokenService are used for generating and validating user auth tokens
type UserAuthTokenService struct {
	SQLStore *sqlstore.SqlStore `inject:""`
	log      log.Logger
}

// Init this service
func (s *UserAuthTokenService) Init() error {
	s.log = log.New("auth")
	return nil
}

func (s *UserAuthTokenService) CreateToken(userId int64, clientIP, userAgent string) (token *userAuthToken, err error) {
	// generate token
	// hash token
	// store in db
	// return unhashed token

	return nil, nil
}

func (s *UserAuthTokenService) LookupToken(unhashedToken string) (token *userAuthToken, err error) {
	// lookup token in db
	// return error if not found/expired

	return nil, nil
}

func (s *UserAuthTokenService) RefreshToken(userId int64, token, clientIP, userAgent string) (unhashedToken string, err error) {
	// lookup token in db
	// refresh token if needed

	return "", nil
}

func parseIPAddress(input string) string {
	var s string
	lastIndex := strings.LastIndex(input, ":")

	if lastIndex != -1 {
		s = input[:lastIndex]
	}

	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)

	ip := net.ParseIP(s)

	if ip.IsLoopback() {
		return "127.0.0.1"
	}

	return ip.String()
}
