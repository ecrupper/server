// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package dml

const (
	// ListHooks represents a query to
	// list all webhooks in the database.
	ListHooks = `
SELECT *
FROM hooks;
`

	// ListRepoHooks represents a query to list
	// all webhooks for a repo_id in the database.
	ListRepoHooks = `
SELECT *
FROM hooks
WHERE repo_id = ?
ORDER BY id DESC
LIMIT ?
OFFSET ?;
`

	// SelectRepoHookCount represents a query to select
	// the count of webhooks for a repo_id in the database.
	SelectRepoHookCount = `
SELECT count(*) as count
FROM hooks
WHERE repo_id = ?;
`

	// SelectRepoHook represents a query to select
	// a webhook for a repo_id in the database.
	SelectRepoHook = `
SELECT *
FROM hooks
WHERE repo_id = ?
AND number = ?
LIMIT 1;
`

	// SelectLastRepoHook represents a query to select
	// the last hook for a repo_id in the database.
	SelectLastRepoHook = `
SELECT *
FROM hooks
WHERE repo_id = ?
ORDER BY number DESC
LIMIT 1;
`

	// DeleteHook represents a query to
	// remove a webhook from the database.
	DeleteHook = `
DELETE
FROM hooks
WHERE id = ?;
`
)
