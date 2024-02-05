package models

func (store *SessionStore) GetPodNameBySessionId(sessionID string) (string, error) {
	var awsInstanceID string
	query := `SELECT pod_name FROM pods JOIN sessions ON pods.id = sessions.pod_id WHERE sessions.session_id = $1`
	err := store.db.QueryRow(query, sessionID).Scan(&awsInstanceID)
	if err != nil {
		return "", err
	}

	return awsInstanceID, nil
}
