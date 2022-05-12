package group

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockLog struct{}

func (m *mockLog) Error(message string) {}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  "sqlmock_db_0",
		PreferSimpleProtocol: true,
		Conn:                 db,
	})
	mockDB, err := gorm.Open(dialector, &gorm.Config{})
	var repo GroupRepository = NewGroupRepositoryImpl(mockDB, &mockLog{})

	testCases := []struct {
		name          string
		inputGroup    entity.Group
		expectedError error
		mockBehaviour func()
	}{
		{
			name:          "it should return nil error, when successfully insert the data to database",
			expectedError: nil,
			mockBehaviour: func() {
				mock.ExpectBegin()
				mock.ExpectExec(".*").WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name:          "it should return ErrDatabase, when database return an error",
			expectedError: repository.ErrDatabase,
			mockBehaviour: func() {
				mock.ExpectBegin()
				mock.ExpectExec(".*").WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnError(gorm.ErrInvalidDB)
			},
		},
		// {
		// 	name:          "it should return ErrRecordAlreadyExists, when generated id is already exists in the database",
		// 	expectedError: repository.ErrRecordAlreadyExists,
		// 	mockBehaviour: func() {
		// 		mock.ExpectBegin()
		// 		mock.ExpectExec(".*").WithArgs(
		// 			sqlmock.AnyArg(),
		// 			sqlmock.AnyArg(),
		// 			sqlmock.AnyArg(),
		// 			sqlmock.AnyArg(),
		// 			sqlmock.AnyArg(),
		// 			sqlmock.AnyArg(),
		// 		).WillReturnError(&pgconn.PgError{Code: "23505"})
		// 	},
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour()

			gotError := repo.Insert(context.Background(), testCase.inputGroup)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			if testCase.expectedError != nil {
				assert.Equal(t, testCase.expectedError, gotError)
			} else {
				assert.NoError(t, gotError)
			}
		})
	}
}
