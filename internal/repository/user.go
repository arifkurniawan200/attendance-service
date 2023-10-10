package repository

import (
	"context"
	"database/sql"
	"fmt"
	"template/internal/model"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserHandler{db}
}

func (h UserHandler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func (h UserHandler) GetUserExcludeMe(userId int) ([]model.Member, error) {
	var (
		datas []model.Member
		err   error
	)
	query := fmt.Sprintf(baseGetMember, `WHERE id != ?`)
	rows, err := h.db.Query(query, userId)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.Member
		if err = rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Email, &data.Password,
			&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
}

func (h UserHandler) GetUserByID(userID int) (model.Member, error) {
	var (
		data model.Member
		err  error
	)
	query := fmt.Sprintf(baseGetMember, `WHERE id = ?`)
	rows, err := h.db.Query(query, userID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Email, &data.Password,
			&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (h UserHandler) RegisterUser(c model.MemberParam) error {
	_, err := h.db.Exec(insertNewMembers, c.FirstName, c.LastName, c.Email, c.Password)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) GetUserByEmail(email string) (model.Member, error) {
	var (
		data model.Member
		err  error
	)
	query := fmt.Sprintf(baseGetMember, `WHERE email = ?`)
	rows, err := h.db.Query(query, email)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Email, &data.Password,
			&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}
