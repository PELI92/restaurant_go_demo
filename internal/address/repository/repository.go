package repository

import (
	"context"
	"database/sql"
	"demo/config"
	"demo/internal/address/repository/entity"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const (
	insertQ = `INSERT INTO addresses (
        id,
        street,
        city,
        state,
        postal_code,
        country
    ) VALUES (?, ?, ?, ?, ?, ?)`

	selectQ = `SELECT
        id,
        street,
        city,
        state,
        postal_code,
        country,
        created_at,
        updated_at
    FROM addresses
    WHERE id = ?`

	selectByRestaurantIdQ = `SELECT
        a.id,
        a.street,
        a.city,
        a.state,
        a.postal_code,
        a.country,
        a.created_at,
        a.updated_at
    FROM addresses a
    INNER JOIN restaurants r ON a.id = r.address_id
    WHERE r.id = ?`

	updateQ = `UPDATE addresses
    SET
        street = ?,
        city = ?,
        state = ?,
        postal_code = ?,
        country = ?
    WHERE id = ?`

	deleteQ = `DELETE FROM addresses WHERE id = ?`
)

func CreateAddress(ctx context.Context, a entity.Address) (*uuid.UUID, error) {
	fmt.Printf("CreateAddress: %+v\n", a)
	newID := uuid.New()
	if _, err := config.DB.ExecContext(
		ctx,
		insertQ,
		newID[:],
		a.Street,
		a.City,
		a.State,
		a.PostalCode,
		a.Country,
	); err != nil {
		return nil, err
	}
	return &newID, nil
}

func GetAddressByRestaurantId(ctx context.Context, restaurantID uuid.UUID) (*entity.Address, error) {
	row := config.DB.QueryRowContext(ctx, selectByRestaurantIdQ, restaurantID[:])

	var addr entity.Address
	if err := row.Scan(
		&addr.ID,
		&addr.Street,
		&addr.City,
		&addr.State,
		&addr.PostalCode,
		&addr.Country,
		&addr.CreatedAt,
		&addr.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}
	return &addr, nil
}

func GetAddressByID(ctx context.Context, id uuid.UUID) (*entity.Address, error) {
	row := config.DB.QueryRowContext(ctx, selectQ, id[:])

	var addr entity.Address
	if err := row.Scan(
		&addr.ID,
		&addr.Street,
		&addr.City,
		&addr.State,
		&addr.PostalCode,
		&addr.Country,
		&addr.CreatedAt,
		&addr.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}
	return &addr, nil
}

func UpdateAddress(ctx context.Context, id uuid.UUID, a entity.Address) (*entity.Address, error) {
	if _, err := config.DB.ExecContext(
		ctx,
		updateQ,
		a.Street,
		a.City,
		a.State,
		a.PostalCode,
		a.Country,
		id[:],
	); err != nil {
		return nil, err
	}
	return GetAddressByID(ctx, id)
}

func DeleteAddress(ctx context.Context, id uuid.UUID) error {
	_, err := config.DB.ExecContext(ctx, deleteQ, id[:])
	return err
}
