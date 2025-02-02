// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUser_Engine_UpdateUser(t *testing.T) {
	// setup types
	_user := testUser()
	_user.SetID(1)
	_user.SetName("foo")
	_user.SetToken("bar")
	_user.SetHash("baz")

	_postgres, _mock := testPostgres(t)
	defer func() { _sql, _ := _postgres.client.DB(); _sql.Close() }()

	// ensure the mock expects the query
	_mock.ExpectExec(`UPDATE "users"
SET "name"=$1,"refresh_token"=$2,"token"=$3,"hash"=$4,"favorites"=$5,"active"=$6,"admin"=$7
WHERE "id" = $8`).
		WithArgs("foo", AnyArgument{}, AnyArgument{}, AnyArgument{}, nil, false, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_sqlite := testSqlite(t)
	defer func() { _sql, _ := _sqlite.client.DB(); _sql.Close() }()

	err := _sqlite.CreateUser(_user)
	if err != nil {
		t.Errorf("unable to create test user for sqlite: %v", err)
	}

	// setup tests
	tests := []struct {
		failure  bool
		name     string
		database *engine
	}{
		{
			failure:  false,
			name:     "postgres",
			database: _postgres,
		},
		{
			failure:  false,
			name:     "sqlite3",
			database: _sqlite,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = test.database.UpdateUser(_user)

			if test.failure {
				if err == nil {
					t.Errorf("UpdateUser for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("UpdateUser for %s returned err: %v", test.name, err)
			}
		})
	}
}
