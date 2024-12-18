package repositories

import (
	"database/sql"
	"errors"
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

	var lastIdInsert uint64

	err = statement.QueryRow(publications.Title, publications.Text, publications.AuthorId).Scan(&lastIdInsert)
	if err != nil {
		return 0, nil
	}

	return lastIdInsert, nil
}


// Busca as publicações do usuário e de seus seguidores
func (repository publications) SearchPublications(userId uint64) ([]models.Publications, error) {

	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publications p
		INNER JOIN users u ON u.id = p.author_id 
		LEFT JOIN followers f ON p.author_id = f.user_id AND f.follower_id = $1
        WHERE p.author_id = $1 OR f.follower_id IS NOT NULL
        ORDER BY p.id DESC`, userId,
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


// Método para buscar uma publicação por id
func (repository publications) SearchPublicationsId(publicationId uint64) (models.Publications, error) {

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
			return models.Publications{}, errors.New("publicação não encontrada")
		}
	}

	return publication, nil
}


// Atualiza uma publicação do usuário autenticado
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

// Delete deleta uma publicação do usuário no banco de dados
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

// UserPublications verifica se a publicação existe
func (repository publications) ExistPublications(userId uint64) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM publications WHERE id = $1)"
    if err := repository.db.QueryRow(query, userId).Scan(&exists); err != nil {
        return false, err
    }
    return exists, nil
}