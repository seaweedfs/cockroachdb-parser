// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package geomfn

import (
	"github.com/seaweedfs/cockroachdb-parser/pkg/geo"
	"github.com/seaweedfs/cockroachdb-parser/pkg/geo/geos"
)

// Snap returns the input geometry with the vertices snapped to the target
// geometry. Tolerance is used to control where snapping is performed.
// If no snapping occurs then the input geometry is returned unchanged.
func Snap(input, target geo.Geometry, tolerance float64) (geo.Geometry, error) {
	snappedEWKB, err := geos.Snap(input.EWKB(), target.EWKB(), tolerance)
	if err != nil {
		return geo.Geometry{}, err
	}
	return geo.ParseGeometryFromEWKB(snappedEWKB)
}
