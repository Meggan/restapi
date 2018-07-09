package main

import (
	"database/sql"
)

//funcs init
type Store interface {
	createFood(food *Food) error
	getFood(food *Food) error
	getFoods() ([]*Food, error)
	deleteFood(food *Food) error
	updateFood(food *Food) error
}

//db connection
type dbStore struct {
	db *sql.DB
}

//funcs
func (store *dbStore) createFood(food *Food) error {
	_, err := store.db.Query("INSERT INTO cheesetest.test(id, foodgroup, subgroup, name, origincountry, origincity) VALUES ($1, $2, $3, $4, $5, $6)",
		food.ID, food.Group, food.Subgroup, food.Name, food.Origin.Country, food.Origin.City)
	return err
}

func (store *dbStore) getFood(food *Food) error {
	_, err := store.db.Query("SELECT * FROM cheesetest.test WHERE id = ($1)",
		food.ID)
	return err
}

func (store *dbStore) getFoods() ([]*Food, error) {
	rows, err := store.db.Query("SELECT * FROM cheesetest.test ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := []*Food{}
	for rows.Next() {
		food := &Food{}
		foods = append(foods, food)
	}
	return foods, nil
}

func (store *dbStore) deleteFood(food *Food) error {
	_, err := store.db.Query("DELETE * FROM cheesetest.test WHERE id = ($1)",
		food.ID)
	return err
}

var store Store

func InitStore(s Store) {
	store = s
}
