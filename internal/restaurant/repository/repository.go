package repository

import (
	"context"
	"demo/config"
	"demo/internal/restaurant/repository/entity"
	"fmt"
	"github.com/google/uuid"
)

const (
	insertQ          = `INSERT INTO restaurants (id,name, phone, email, website, address_id) VALUES (?, ?, ?, ?, ?, ?)`
	selectQ          = `SELECT id, name, phone, email, website, status, created_at, updated_at FROM restaurants WHERE id = ?`
	selectPaginatedQ = `SELECT id, name, phone, email, website, status, created_at, updated_at FROM restaurants ORDER BY created_at DESC LIMIT ? OFFSET ?`
	updateQ          = `UPDATE restaurants SET name = ?, phone = ?, email = ?, website = ? WHERE id = ?`
	deleteQ          = `DELETE FROM restaurants WHERE id = ?`
)

func CreateRestaurant(ctx context.Context, r entity.Restaurant) (*uuid.UUID, error) {
	newID := uuid.New()
	if _, err := config.DB.ExecContext(ctx, insertQ, newID[:], r.Name, r.Phone, r.Email, r.Website, r.AddressID); err != nil {
		fmt.Printf("RestaurantRepositry - CreateRestaurant error: %+v\n", err)
		return nil, err
	}
	return &newID, nil
}

func GetRestaurantsPaginated(ctx context.Context, limit, offset int) ([]entity.Restaurant, error) {
	rows, err := config.DB.QueryContext(ctx, selectPaginatedQ, limit, offset)
	if err != nil {
		fmt.Printf("RestaurantRepositry - GetRestaurantsPaginated error: %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	var list []entity.Restaurant
	for rows.Next() {
		var r entity.Restaurant
		if err := rows.Scan(
			&r.ID, &r.Name, &r.Phone, &r.Email,
			&r.Website, &r.Status, &r.CreatedAt, &r.UpdatedAt,
		); err != nil {
			fmt.Printf("RestaurantRepositry - GetRestaurantsPaginated error: %+v\n", err)
			return nil, err
		}
		list = append(list, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func GetRestaurantByID(ctx context.Context, id uuid.UUID) (*entity.Restaurant, error) {
	res := config.DB.QueryRowContext(ctx, selectQ, id[:])
	var restaurant entity.Restaurant
	if err := res.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Phone, &restaurant.Email, &restaurant.Website, &restaurant.Status, &restaurant.CreatedAt, &restaurant.UpdatedAt); err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func UpdateRestaurant(ctx context.Context, id uuid.UUID, r entity.Restaurant) (*entity.Restaurant, error) {
	_, err := config.DB.ExecContext(ctx,
		updateQ,
		r.Name,
		r.Phone,
		r.Email,
		r.Website,
		id[:],
	)
	if err != nil {
		fmt.Printf("RestaurantRepositry - UpdateRestaurant error: %+v\n", err)
		return nil, err
	}

	updated, err := GetRestaurantByID(ctx, id)
	if err != nil {
		fmt.Printf("RestaurantRepositry - UpdateRestaurant(get) error: %+v\n", err)
		return nil, err
	}
	return updated, nil
}

func DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	_, err := config.DB.ExecContext(ctx, deleteQ, id[:])
	if err != nil {
		fmt.Printf("RestaurantRepositry - DeleteRestaurant error: %+v\n", err)
		return err
	}
	return nil
}
