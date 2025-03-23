package models

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Scan(input interface{}) error {
	hexStr := fmt.Sprint(input)
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return err
	}

	gt, err := ewkb.Unmarshal(bytes)
	if err != nil {
		return err
	}

	pt := gt.(*geom.Point)
	p.Y = pt.Y()
	p.X = pt.X()

	return nil
}

func (p Point) GormDataType() string {
	return "geography"
}

func (p Point) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", p.X, p.Y)},
	}
}

func (p Point) MarshalYAML() (interface{}, error) {
	outPoint := struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  p.Y,
		Longitude: p.X,
	}

	return outPoint, nil
}

func (p Point) MarshalJSON() ([]byte, error) {
	outPoint := struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  p.Y,
		Longitude: p.X,
	}

	return json.Marshal(outPoint)
}
