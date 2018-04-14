package database

import (
	"time"
)

// Birthday struc
type Birthday struct {
	Name string    //`db:"name"`
	Born time.Time //`db:"born"`
}

//AllBirthdays tu vieja
func AllBirthdays() ([]*Birthday, error) {
	/*rows, err := db.Query("SELECT Name, Born FROM birthday")
	if err != nil {
		return nil, err
	}

	log.Println("Successfully loaded birthdays from database")
	birthdays := make([]*Birthday, 0)
	for rows.Next() {
		birthday := new(Birthday)
		err := rows.Scan(&birthday.Name, &birthday.Born)
		if err != nil {
			return nil, err
		}
		birthdays = append(birthdays, birthday)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return birthdays, nil*/
	return nil, nil
}

//InsertBirthday inserta registro en la tabla birdthday
func InsertBirthday(birthday *Birthday) error {
	_, err := db.Exec("INSERT INTO birthday (Name, Born) VALUES (?,?)", birthday.Name, birthday.Born)
	// db.NewRecord(birthday)
	/*
	    id, _ := result.LastInsertId()

	   	contact.Id = int(id)
	*/

	return err
}
