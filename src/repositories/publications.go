package repositories

import (
	"database/sql"
	"golang-social-network-api/src/models"
)

// Publications representa o repositório de publicações
type publications struct {
	db *sql.DB
}

// NewReposPublications cria um repositório de publicações
func NewReposPublications(db *sql.DB) *publications {
	return &publications{db}
}

// Método para inserir a publicação no banco de dados
func (repository publications) Create(publications models.Publications) (uint64, error) {
	statement, err := repository.db.Prepare(`INSERT INTO publications (title, text, author_id) VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	var lastIdInsert uint32

	err = statement.QueryRow(publications.Title, publications.Text, publications.AuthorId).Scan(&lastIdInsert)
	if err != nil {
		return 0, nil
	}

	return uint64(lastIdInsert), nil
}


// Método para buscar uma publicação
func (repository publications) SearchId(publicationId uint64) (models.Publications, error) {

	rows, err := repository.db.Query(`
		SELECT p.*, u.nick
		FROM publications p
		INNER JOIN users u ON u.id = p.author_id 
		WHERE p.id = $1`, publicationId,
	)
	if err != nil {
		return models.Publications{}, err
	}
	defer rows.Close()

	// Intera nas linhas
	var publication models.Publications

	if rows.Next() {
		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Text,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publications{}, err
		}
	}

	return publication, nil
}

// Busca as publicações do usuário e de seus seguidores
func (repository publications) Search(userId uint64) ([]models.Publications, error) {

	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publications p
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f on p.author_id = f.user_id
		WHERE u.id = $1 OR f.follower_id = $2
		ORDER BY 1 DESC`, userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Text,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

// Atualiza uma publicação DO USUÁRIO
func (repository publications) Update(publicationId uint64, publication models.Publications) error {
	statement, err := repository.db.Prepare(`UPDATE publications SET title = $1, text = $2 WHERE id = $3`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Text, publicationId); err != nil {
		return err
	}

	return nil
}

// DelEte exclui uma publicação do usuário no banco de dados
func (repository publications) Delete(publicationId uint64) error {
	statement, err := repository.db.Prepare(`DELETE FROM publications WHERE id = $1`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

// SearchUser traz todas as publicações de um usuário especÍfico
func (repository publications) SearchUser(userId uint64) ([]models.Publications, error ) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick
		FROM publications p
		INNER JOIN users u ON u.id = p.author_id 
		WHERE p.author_id = $1`, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Intera nas linhas
	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Text,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}


// Like para curtir uma publicação
func (repository publications) Like(publicationId uint64) error {

	statement, err := repository.db.Prepare(`UPDATE publications SET likes = likes + 1 WHERE id = $1`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

// DisLike para descurtir uma publicação
func (repository publications) DisLike(publicationId uint64) error {

	statement, err := repository.db.Prepare(`
	UPDATE publications SET likes = 
	CASE 
		WHEN likes > 0 THEN likes - 1
		ELSE 0 
	END
	WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}