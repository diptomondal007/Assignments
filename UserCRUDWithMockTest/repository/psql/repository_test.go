package psql

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log"
	"regexp"
	"testing"
	"usercrudwithmocktest/crud"
	"usercrudwithmocktest/crud/models"
)

type Suite struct {
	suite.Suite
	DB         *sql.DB
	mock       sqlmock.Sqlmock
	repository crud.Repository
	user       *models.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.DB = db
	require.NoError(s.T(), err)
	s.repository, _ = NewPsqlRepository(db)
}
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestPsqlRepository_Get() {
	var (
		limit = 3
		page  = 1
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * from users limit $1 offset $2")).
		WithArgs(limit, limit*(page-1)).
		WillReturnRows(sqlmock.NewRows([]string{"username", "name"}).
			AddRow("dipto", "Dipto").
			AddRow("dipto01", "Dipto").
			AddRow("dipto0001", "Dipto Mondal"))
	res, err := s.repository.Get(page, limit)
	require.NoError(s.T(), err)
	result := []models.User{
		{
			Username: "dipto",
			Name:     "Dipto",
		}, {
			Username: "dipto01",
			Name:     "Dipto",
		},
		{
			Username: "dipto0001",
			Name:     "Dipto Mondal",
		},
	}
	require.Nil(s.T(), deep.Equal(result, res))

}

func (s *Suite) TestPsqlRepository_GetByUsername() {
	var (
		username = "dipto"
		name     = "Dipto"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"select * from users where username=$1")).
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "name"}).
			AddRow(username, name))

	res, err := s.repository.GetByUsername(username)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.User{
		Username: "dipto",
		Name:     "Dipto",
	}, res))
}

func (s *Suite) TestPsqlRepository_Create() {
	var (
		username = "dipto"
		name     = "test-name"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO users VALUES ($1,$2)`)).
		WithArgs(username, name).
		WillReturnRows(
			sqlmock.NewRows([]string{"username", "name"}).AddRow(username, name))
	user := &models.User{
		Username: username,
		Name:     name,
	}
	err := s.repository.Create(user)

	require.NoError(s.T(), err)
}

func (s *Suite) TestPsqlRepository_Delete() {
	var (
		username = "dipto"
		//name = "test-name"
	)
	s.mock.NewRows([]string{"username", "name"}).
		AddRow("dipto", "Dipto")
	s.mock.ExpectExec(regexp.QuoteMeta(
		`delete from users where username=$1`)).
		WithArgs(username).WillReturnResult(driver.RowsAffected(1))
	rowsAffected, err := s.repository.Delete(username)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(rowsAffected, int64(1)))
}

func (s *Suite) TestRowExists() {

}

func (s *Suite) TestPsqlRepository_Update() {
	var (
		username             = "dipto"
		usernameAfterUpdated = "dipto07"
		name                 = "test-name"
	)
	s.mock.NewRows([]string{"username", "name"}).
		AddRow(username, name)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT exists (select username from users where username=$1)")).
		WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{"exists"}).
		AddRow(true))
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"select username from users where username=$1")).
		WithArgs(usernameAfterUpdated).WillReturnRows(sqlmock.NewRows([]string{"username"}))
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"update users set username=$1, name =$2 where username=$3")).
		WithArgs(usernameAfterUpdated, name, username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "name"}).
			AddRow(username, name))
	user := &models.User{
		Username: usernameAfterUpdated,
		Name:     name,
	}
	err := s.repository.Update(username, user)
	log.Println(err)
	require.NoError(s.T(), err)
}
