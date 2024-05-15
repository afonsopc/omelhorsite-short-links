package utils

import (
	"database/sql"
	"errors"
)

type Link struct {
	ID         string `json:"id"`
	ForwardUrl string `json:"forwardUrl"`
	UserID     string `json:"userId"`
	CreatedAt  string `json:"createdAt"`
}

func GetLinkInfo(linkID string) (Link, error) {
	db := GetDatabaseConnection()

	defer db.Close()

	var link Link

	err := db.QueryRow(`
		SELECT * FROM links WHERE id = $1 LIMIT 1
	`, linkID).Scan(
		&link.ID,
		&link.ForwardUrl,
		&link.UserID,
		&link.CreatedAt,
	)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func GetAllLinks(userID string) ([]Link, error) {
	db := GetDatabaseConnection()

	defer db.Close()

	var query string

	if userID == "" {
		query = `SELECT * FROM links`
	} else {
		query = `SELECT * FROM links WHERE user_id = $1`
	}

	rows, err := db.Query(query, userID)

	if err != nil {
		return nil, err
	}

	var links []Link

	for rows.Next() {
		var link Link

		err := rows.Scan(
			&link.ID,
			&link.ForwardUrl,
			&link.UserID,
			&link.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func CreateLink(
	forwardUrl string,
	userID string,
) (Link, error) {
	db := GetDatabaseConnection()

	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO links (id, forward_url, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, forward_url, user_id, created_at
	`)

	if err != nil {
		return Link{}, err
	}

	id, err := getRandomUnusedLinkID()

	if err != nil {
		return Link{}, err
	}

	var link Link

	err = stmt.QueryRow(
		id,
		forwardUrl,
		userID,
	).Scan(
		&link.ID,
		&link.ForwardUrl,
		&link.UserID,
		&link.CreatedAt,
	)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func getRandomUnusedLinkID() (string, error) {
	db := GetDatabaseConnection()

	defer db.Close()

	counter := 1
	maxTries := 10
	linkIDLength := 12

	var id string
	for {
		id = random_string(linkIDLength)
		err := db.QueryRow("SELECT id FROM links WHERE id = $1", id).Scan(&id)
		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return "", err
		}

		if counter > maxTries {
			return "", errors.New("error generating unique id for link")
		}

		counter++
	}

	return id, nil
}

func DeleteLink(linkID string) error {
	db := GetDatabaseConnection()

	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM links WHERE id = $1
	`, linkID,
	)

	return err
}
