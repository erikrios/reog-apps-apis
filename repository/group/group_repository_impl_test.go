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
			inputGroup:    entity.Group{},
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
			inputGroup:    entity.Group{},
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

func TestInsertAll(t *testing.T) {
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
		inputGroups   []entity.Group
		expectedError error
		mockBehaviour func()
	}{
		{
			name:          "it should return nil error, when successfully insert the data to database",
			inputGroups:   []entity.Group{{}},
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
				).WillReturnResult(sqlmock.NewResult(1, 0))
				mock.ExpectCommit()
			},
		},
		{
			name:          "it should return ErrDatabase, when database return an error",
			inputGroups:   []entity.Group{{}},
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour()

			gotError := repo.InsertAll(context.Background(), testCase.inputGroups)

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

func TestFindAll(t *testing.T) {
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
		name           string
		expectedGroups []entity.Group
		expectedError  error
		mockBehaviour  func()
	}{
		{
			name: "it should return valid groups, when database successfully return the data",
			expectedGroups: []entity.Group{
				{
					ID:     "g-xyz",
					Name:   "Paguyuban Reog",
					Leader: "Erik",
				},
			},
			expectedError: nil,
			mockBehaviour: func() {
				returnedRows := sqlmock.NewRows([]string{"id", "name", "leader", "created_at", "updated_at", "deleted_at"})
				returnedRows.AddRow(
					"g-xyz",
					"Paguyuban Reog",
					"Erik",
					nil,
					nil,
					nil,
				)
				mock.ExpectQuery(".*").WillReturnRows(returnedRows)
				mock.ExpectQuery(".*").
					WillReturnRows(sqlmock.NewRows([]string{"id", "address", "village_id", "villlage_name", "district_id", "district_name", "regency_id", "regency_name, province_id", "province_name", "created_at", "updated_at", "deleted_at"}))
				mock.ExpectQuery(".*").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "amount", "group_id", "created_at", "updated_at", "deleted_at"}))
			},
		},
		{
			name:           "it should return ErrDatabase, when database return an error",
			expectedGroups: []entity.Group{},
			expectedError:  repository.ErrDatabase,
			mockBehaviour: func() {
				mock.ExpectQuery(".*").WillReturnError(gorm.ErrInvalidDB)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour()

			gotEntity, gotError := repo.FindAll(context.Background())

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatal(err)
			}

			if testCase.expectedError != nil {
				assert.Equal(t, testCase.expectedError, gotError)
			} else {
				assert.NoError(t, gotError)
				assert.Equal(t, len(testCase.expectedGroups), len(gotEntity))
				for i, group := range testCase.expectedGroups {
					assert.Equal(t, group.ID, gotEntity[i].ID)
					assert.Equal(t, group.Name, gotEntity[i].Name)
					assert.Equal(t, group.Leader, gotEntity[i].Leader)
				}
			}
		})
	}
}
