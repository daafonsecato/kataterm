package models

func (store *SessionStore) GetMachineHost(machineID string) (string, error) {
	var host string
	query := `SELECT domain FROM machines WHERE aws_instance_id = $1`
	err := store.db.QueryRow(query, machineID).Scan(&host)
	if err != nil {
		return "", err
	}
	return host, nil
}
