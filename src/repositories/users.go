package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.Usuario) (uint64, error) {

	statement, erro := repository.db.Prepare("insert into users (nome, nick, email, senha) values($1, $2, $3, $4) RETURNING id")
	if erro != nil {
		fmt.Println(erro)
		return 0, nil
	}
	defer statement.Close()

	var ultimoId uint64
	erro = statement.QueryRow(user.Nome, user.Nick, user.Email, user.Senha).Scan(&ultimoId)
	if erro != nil {
		fmt.Println(erro)
		return 0, nil
	}

	return ultimoId, nil

}

func (repository users) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repository.db.Query(
		"SELECT id, nome, nick, email, createdat FROM users WHERE nome LIKE '%' || $1 || '%' OR nick LIKE '%' || $1 || '%'", nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []models.Usuario

	for linhas.Next() {
		var user models.Usuario

		if err := linhas.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, erro
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository users) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repository.db.Query(
		"select id, nome, nick, email, createdat from users where id = $1", ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var user models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return user, nil
}

func (repository users) Atualizar(ID uint64, user models.Usuario) error {
	statement, erro := repository.db.Prepare("update users set nome = $1, nick = $2, email = $3 where id = $4")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Nome, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repository users) Deletar(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repository users) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repository.db.Query("select id, senha from users where email = $1", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var user models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&user.ID, &user.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return user, nil
}

func (repository users) Seguir(userID, seguidorID uint64) error {
	statement, erro := repository.db.Prepare("insert into seguidores (user_id, seguidor_id) values ($1, $2) on conflict (user_id, seguidor_id) do nothing")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(userID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (repository users) UnFollow(userID, seguidorID uint64) error {
	statement, erro := repository.db.Prepare("delete from seguidores where user_id = $1 and seguidor_id = $2")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(userID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (repository users) BuscarSeguidores(userID uint64) ([]models.Usuario, error) {
	linha, erro := repository.db.Query("select u.id, u.nome, u.nick, u.email, u.createdat from users u inner join seguidores s on u.id = s.seguidor_id where s.user_id = $1", userID)
	if erro != nil {
		return nil, erro
	}
	defer linha.Close()

	var users []models.Usuario
	for linha.Next() {
		var u models.Usuario

		if err := linha.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CreatedAt); err != nil {
			return nil, erro

		}
		users = append(users, u)
	}
	return users, nil

}

func (repository users) BuscarSeguindo(userID uint64) ([]models.Usuario, error) {
	linha, erro := repository.db.Query("select u.id, u.nome, u.nick, u.email, u.createdat from users u inner join seguidores s on u.id = s.user_id where s.seguidor_id = $1", userID)
	if erro != nil {
		return nil, erro
	}
	defer linha.Close()

	var users []models.Usuario
	for linha.Next() {
		var u models.Usuario

		if err := linha.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CreatedAt); err != nil {
			return nil, erro

		}
		users = append(users, u)
	}
	return users, nil

}

func (repository users) BuscarSenha(userID uint64) (string, error) {
	linha, err := repository.db.Query("select senha from users where id = $1", userID)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var user models.Usuario

	if linha.Next() {
		if err = linha.Scan(&user.Senha); err != nil {
			return "", err
		}
	}
	return user.Senha, nil
}

func (repository users) AtualizarSenha(userID uint64, senha string) error {
	statement, erro := repository.db.Prepare("update users set senha = $1 where id = $2")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senha, userID); erro != nil {
		return erro
	}
	return nil
}
