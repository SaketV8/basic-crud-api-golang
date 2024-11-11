package sqlite

import (
	"database/sql"

	"github.com/saketV8/basic-crud-api-golang/models"
)

// for query data using sql in sqlite
type UserAccountsDbModel struct {
	DB *sql.DB
}

// this function is method for UserAccountsModel

// to get data from the table by <user_name>
func (UserAccDb_Model *UserAccountsDbModel) GetByUserName(user_name string) (UserAcc models.UserAccount, err error) {
	statement := `SELECT id, user_name, first_name, last_name, phone_number, email FROM user_accounts WHERE user_name = ?`

	err = UserAccDb_Model.DB.QueryRow(statement, user_name).Scan(&UserAcc.ID, &UserAcc.UserName, &UserAcc.FirstName, &UserAcc.LastName, &UserAcc.PhoneNumber, &UserAcc.Email)
	if err != nil {
		return UserAcc, err
	}
	return UserAcc, nil
}

// to get all data from the table
func (UserAccDb_Model *UserAccountsDbModel) All() ([]models.UserAccount, error) {
	statement := `SELECT id, user_name, first_name, last_name, phone_number, email FROM user_accounts ORDER BY id DESC;`

	rows, err := UserAccDb_Model.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	//initializing the empty slice <UserAccounts> of data type <models.UserAccount>
	UserAccounts := []models.UserAccount{}
	for rows.Next() {
		// initializing the empty variable <UserAccounts> of data type <models.UserAccount>
		UserAcc := models.UserAccount{}
		//extracting data from rows and setting in UserAcc Variable
		err := rows.Scan(&UserAcc.ID, &UserAcc.UserName, &UserAcc.FirstName, &UserAcc.LastName, &UserAcc.PhoneNumber, &UserAcc.Email)
		if err != nil {
			return nil, err
		}

		// finally adding that single user account data
		// as element in slice UserAccounts
		UserAccounts = append(UserAccounts, UserAcc)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// final return
	return UserAccounts, nil
}

// Insert Data in table
func (UserAccDb_Model *UserAccountsDbModel) Insert(user_name, first_name, last_name, phone_number, email string) (rowAffected int64, err error) {
	statement := `INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email) VALUES (?, ?, ?, ?, ?);`

	// use exec when there is no return
	result, err := UserAccDb_Model.DB.Exec(statement, user_name, first_name, last_name, phone_number, email)
	if err != nil {
		return rowAffected, err
	}

	// getting row affected from the result
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, err
	}
	return rowAffected, nil
}

// Update Data in table
func (UserAccDb_Model *UserAccountsDbModel) Update(user_name, first_name, last_name, phone_number, email string) (rowAffected int64, err error) {
	statement := `UPDATE user_accounts SET first_name = ?, last_name = ?, phone_number = ?, email = ? WHERE user_name = ?`

	// use exec when there is no return
	result, err := UserAccDb_Model.DB.Exec(statement, first_name, last_name, phone_number, email, user_name)
	if err != nil {
		return rowAffected, err
	}

	// getting row affected from the result
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, err
	}
	return rowAffected, nil
}

// Delete Data in table
func (UserAccDb_Model *UserAccountsDbModel) Delete(user_name string) (rowAffected int64, err error) {
	statement := `DELETE FROM user_accounts WHERE user_name = ?;`

	// use exec when there is no return
	result, err := UserAccDb_Model.DB.Exec(statement, user_name)
	if err != nil {
		return rowAffected, err
	}

	// getting row affected from the result
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, err
	}
	return rowAffected, nil
}
