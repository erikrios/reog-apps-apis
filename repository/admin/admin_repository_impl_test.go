package admin

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

func TestFindByUsername(t *testing.T) {
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
	var repo AdminRepository = NewAdminRepositoryImpl(mockDB, &mockLog{})

	testCases := []struct {
		name          string
		inputUsername string
		expectedAdmin entity.Admin
		expectedError error
		mockBehaviour func()
	}{
		{
			name:          "it should return valid admin, when database successfully return the data",
			inputUsername: "admin",
			expectedAdmin: entity.Admin{
				ID:       "a-xyz",
				Username: "admin",
				Name:     "Administrator",
				Password: "secret",
			},
			expectedError: nil,
			mockBehaviour: func() {
				returnedRows := sqlmock.NewRows([]string{"id", "username", "name", "password", "created_at", "updated_at", "deleted_at"})
				returnedRows.AddRow(
					"a-xyz",
					"admin",
					"Administrator",
					"secret",
					nil,
					nil,
					nil,
				)
				mock.ExpectQuery(".*").WithArgs(sqlmock.AnyArg()).WillReturnRows(returnedRows)
			},
		},
		{
			name:          "it should return ErrDatabase, when database return an error",
			inputUsername: "admin",
			expectedAdmin: entity.Admin{},
			expectedError: repository.ErrDatabase,
			mockBehaviour: func() {
				mock.ExpectQuery(".*").WithArgs(sqlmock.AnyArg()).WillReturnError(gorm.ErrInvalidDB)
			},
		},
		{
			name:          "it should return ErrRecordNotFound, when given username not found in the database",
			inputUsername: "admin",
			expectedAdmin: entity.Admin{},
			expectedError: repository.ErrRecordNotFound,
			mockBehaviour: func() {
				mock.ExpectQuery(".*").WithArgs(sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour()

			gotEntity, gotError := repo.FindByUsername(context.Background(), testCase.inputUsername)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			if testCase.expectedError != nil {
				assert.Equal(t, testCase.expectedError, gotError)
			} else {
				assert.NoError(t, gotError)
				assert.Equal(t, testCase.expectedAdmin.ID, gotEntity.ID)
				assert.Equal(t, testCase.expectedAdmin.Username, gotEntity.Username)
				assert.Equal(t, testCase.expectedAdmin.Name, gotEntity.Name)
				assert.Equal(t, testCase.expectedAdmin.Password, gotEntity.Password)
			}
		})
	}
}
