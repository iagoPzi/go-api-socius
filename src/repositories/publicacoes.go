package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type publicacoes struct {
	db *sql.DB
}

func NewPublicacoesRepository(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

func (repository publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var ultimoId uint64
	err = statement.QueryRow(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID).Scan(&ultimoId)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	return ultimoId, nil

}

func (repository publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, err := repository.db.Query("select p.*, u.nick from publicacoes p inner join users u on u.id = p.autor_id where p.id = $1", publicacaoID)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linha.Close()

	var p models.Publicacao

	if linha.Next() {
		if err = linha.Scan(
			&p.ID, &p.Titulo, &p.Conteudo, &p.AutorID, &p.Curtidas, &p.CriadaEm, &p.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}
	return p, nil

}

func (repository publicacoes) Buscar(userID uint64) ([]models.Publicacao, error) {
	linhas, err := repository.db.Query("select distinct p.*, u.nick from publicacoes p inner join users u on u.id = p.autor_id join seguidores s on p.autor_id = s.user_id where u.id = $1 or s.seguidor_id = $1 order by 1 desc", userID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var publicacoes []models.Publicacao
	for linhas.Next() {
		var p models.Publicacao
		if err = linhas.Scan(
			&p.ID, &p.Titulo, &p.Conteudo, &p.AutorID, &p.Curtidas, &p.CriadaEm, &p.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, p)
	}
	return publicacoes, nil
}

func (repository publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, err := repository.db.Prepare("update publicacoes set titulo = $1, conteudo = $2 where id = $3")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}
	return nil
}

// Deletar exclui uma publicacao do DB
func (repository publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repository.db.Prepare("delete from publicacoes where id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}

func (repository publicacoes) BuscarPorUsuario(userID uint64) ([]models.Publicacao, error) {
	linhas, err := repository.db.Query("select p.*, u.nick from publicacoes p join users u on u.id = p.autor_id where p.autor_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var p models.Publicacao
		if err = linhas.Scan(
			&p.ID, &p.Titulo, &p.Conteudo, &p.AutorID, &p.Curtidas, &p.CriadaEm, &p.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, p)
	}
	return publicacoes, nil

}

func (repository publicacoes) Curtir(publicacaoID uint64) error {

	statement, err := repository.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

func (repository publicacoes) Descurtir(publicacaoID uint64) error {

	statement, err := repository.db.Prepare("update publicacoes set curtidas = case when curtidas > 0 then curtidas - 1 else curtidas end where id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
