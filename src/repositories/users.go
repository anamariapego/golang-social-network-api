package repositories

import (
	"database/sql"
	"fmt"
	"golang-social-network-api/src/models"
)

// repositories vai receber os dados e inserir, deletar ou atualizar no banco de dados

// Users representa um repositório de usuários
type users struct {
	db *sql.DB
}

// NewReposUsers interação com o banco de dados - comunicação com as tabelas do banco
func NewReposUsers(db *sql.DB) *users {
	return &users{db}
}

// Método para inserir usuário no banco de dados
func (repository users) Create(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare(`INSERT INTO users (name_user, nick, email, password_user) VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		return 0, fmt.Errorf("erro ao preparar a query de inserção: %w", err)
	}
	defer statement.Close()

	var lastIdInsert uint32

	err = statement.QueryRow(user.NameUser, user.Nick, user.Email, user.PasswordUser).Scan(&lastIdInsert)
	if err != nil {
		return 0, fmt.Errorf("erro ao executar a query de inserção: %w", err)
	}

	return uint64(lastIdInsert), nil
}

// Método para buscar os usuários no banco de dados
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick= fmt.Sprintf("%%%s%%", nameOrNick) // obter %nameOrKike%

	rows, err := repository.db.Query(
		"SELECT id, name_user, nick, email, created_at FROM users WHERE name_user LIKE $1 OR nick LIKE $2",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a query de seleção: %w", err)
	}
	defer rows.Close()

	// Interar nas linhas
	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.NameUser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Método para buscar um usuário por id no banco de dados
func (repository users) SearchId(Id uint64) (models.User, error) {

	rows, err := repository.db.Query("SELECT id, name_user, nick, email, created_at FROM users WHERE id = $1",Id)
	if err != nil {
		return models.User{}, fmt.Errorf("erro ao executar a query de seleção por id: %w", err)
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.NameUser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Método para atualizar informações de um usuário por Id no banco de dados
func (repository users) Update(Id uint64, user models.User) (error) {
	statement, err := repository.db.Prepare(`UPDATE users SET name_user = $1, nick = $2, email = $3 WHERE id = $4`)
	if err != nil {
		return fmt.Errorf("erro ao preparar a query de atualização: %w", err)
	}
	defer statement.Close()

	if _, err = statement.Exec(user.NameUser, user.Nick, user.Email, Id); err != nil {
		return fmt.Errorf("erro ao executar a query de atualização: %w", err)
	}
	return nil
}

// Método para deletar informações de um usuário por id no banco de dados
func (repository users) Delete(Id uint64) (error) {
	statement, err := repository.db.Prepare(`DELETE FROM users WHERE id = $1`)
	if err != nil {
		return fmt.Errorf("erro ao preparar consulta de deleção: %w", err)
	}
	defer statement.Close()

	if _, err = statement.Exec(Id); err != nil {
		return fmt.Errorf("erro ao executar a deleção do usuário com id %d: %w", Id, err)
	}
	return nil
}


// SearchEmail busca um usuário por email e retorna o seu id e senha hash
func (repository users) SearchEmail(email string) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, password_user FROM users WHERE email = $1",email)
	if err != nil {
		return models.User{}, fmt.Errorf("erro ao executar a query de seleção: %w", err)
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.PasswordUser); err != nil {
			return models.User{}, fmt.Errorf("erro ao escanear os dados do usuário: %w", err)
		}
	}

	return user, nil
}

// Follower permite que um usuário siga outro
func (repository users) Follower(userId, followerId uint64) error {
    statement, err := repository.db.Prepare(`
        INSERT INTO followers (user_id, follower_id)
        VALUES ($1, $2)
        ON CONFLICT DO NOTHING
    `)
    if err != nil {
        return nil
    }
    defer statement.Close()

    // Executa a instrução preparada
    if _, err = statement.Exec(userId, followerId); err != nil {
        return nil
    }
    return nil
}

// StopFollower permite que um usuário pare de seguir outro usuário
func (repository users) StopFollower(userId, followerId uint64) error {
    statement, err := repository.db.Prepare(`DELETE FROM followers WHERE user_id = $1 AND follower_id = $2`)
    if err != nil {
        return nil
    }
    defer statement.Close()

    // Executa a instrução preparada
    if _, err = statement.Exec(userId, followerId); err != nil {
        return nil
    }
    return nil
}

// SearchFollowers para buscar os seguidores
func (repository users) SearchFollowers(userId uint64) ([]models.User, error) {

	rows, err := repository.db.Query(`
		SELECT u.id, u.name_user, u.nick, u.email, u.created_at 
		FROM users u 
		INNER JOIN followers s ON u.id = s.follower_id 
		WHERE s.user_id = $1`, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Interar nas linhas
	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.NameUser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// SearchFollowers para buscar os usuários que o usuário segue
func (repository users) SearchFollowing(userId uint64) ([]models.User, error) {

	rows, err := repository.db.Query(`
		SELECT u.id, u.name_user, u.nick, u.email, u.created_at 
		FROM users u 
		INNER JOIN followers s ON u.id = s.user_id 
		WHERE s.follower_id = $1`, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Interar nas linhas
	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.NameUser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// SearchPassword traz a senha de um usuário por id
func (repository users) SearchPassword(userId uint64) (string, error) {

	rows, err := repository.db.Query(`SELECT password_user FROM users WHERE id = $1`, userId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Interar nas linhas
	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.PasswordUser,); err != nil {
			return "", err
		}
	}

	return user.PasswordUser, nil
}

// Updateassword atualizar senha de um usuário
func (repository users) UpdatePassword(userId uint64, password string) error {
    statement, err := repository.db.Prepare(`UPDATE users SET password_user = $1 WHERE id = $2`)
    if err != nil {
        return err
    }
    defer statement.Close()

    if _, err = statement.Exec(password, userId); err != nil {
        return err
    }

    return nil
}
