package repository

import (
	"database/sql"
	"fmt"
	"time"
)

func GetAll(db *sql.DB) (CakeCollection, error) {
	resDb, err := db.Query("SELECT * FROM cakes ORDER BY id DESC")
	if err != nil {
		return CakeCollection{}, err
	}

	result := CakeCollection{}

	for resDb.Next() {
		temp := Cake{}
		err := resDb.Scan(&temp.ID, &temp.Name, &temp.Description,
			&temp.Rating, &temp.Image, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return CakeCollection{}, err
		}
		result.Cakes = append(result.Cakes, temp)
	}

	return result, nil
}

func GetOne(db *sql.DB, id int) (Cake, error) {
	resDb, err := db.Query("SELECT * FROM cakes WHERE id=?", id)
	if err != nil {
		return Cake{}, err
	}

	result := Cake{}

	for resDb.Next() {
		temp := Cake{}
		err := resDb.Scan(&temp.ID, &temp.Name, &temp.Description,
			&temp.Rating, &temp.Image, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return Cake{}, err
		}
		result = temp
	}

	return result, nil
}

func Create(db *sql.DB, cake Cake) error {

	insForm, err := db.Prepare("INSERT INTO cakes(name, description,rating, image ) VALUES(?,?,?,?)")

	if err != nil {
		return err
	}

	insForm.Exec(cake.Name, cake.Description, cake.Rating, cake.Image)

	return nil

}

func Update(db *sql.DB, cake Cake) error {

	insForm, err := db.Prepare("UPDATE cakes set name = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?")

	if err != nil {
		return err
	}

	insForm.Exec(cake.Name, cake.Description, cake.Rating, cake.Image, time.Now(), cake.ID)

	return nil

}

func Delete(db *sql.DB, id int) error {

	insForm, err := db.Prepare("DELETE FROM cakes WHERE id = ?")

	if err != nil {
		return err
	}

	insForm.Exec(id)

	return nil

}
