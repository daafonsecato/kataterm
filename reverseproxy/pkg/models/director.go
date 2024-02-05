package models

import (
	"fmt"
)

func (store *SessionStore) GetServiceFromSessionID(SessionID string, service string) (string, error) {
	svc_host := fmt.Sprintf("%s-svc-%s", service, SessionID)
	return svc_host, nil
}
