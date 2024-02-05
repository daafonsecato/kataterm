package models

func (store *SessionStore) StoreMachineAndSession(awsInstanceID string, ipAddress string, sessionID string) error {
	// Start a transaction
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}

	// Insert into machines table
	var pod_id int
	insertMachineQuery := `INSERT INTO pods (pod_name, pod_status, domain) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(insertMachineQuery, awsInstanceID, "pending", ipAddress).Scan(&pod_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into sessions table
	insertSessionQuery := `INSERT INTO sessions (session_id, pod_id) VALUES ($1, $2)`
	_, err = tx.Exec(insertSessionQuery, sessionID, pod_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	// Commit the transaction
	return nil
}
