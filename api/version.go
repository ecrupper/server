// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-vela/server/version"
)

// swagger:operation GET /version base Version
//
// Get the version of the Vela API
//
// ---
// produces:
// - application/json
// parameters:
// responses:
//   '200':
//     description: Successfully retrieved the Vela API version
//     schema:
//       "$ref": "#/definitions/Version"

// Version represents the API handler to
// report the version number for Vela.
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, version.New())
}
