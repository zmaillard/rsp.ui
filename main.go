package main

import (
	"context"
	"highway-sign-portal-builder/pkg/config"
	"highway-sign-portal-builder/pkg/converter"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/generator"
)

func main() {

	ctx := context.Background()
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dbConn, err := db.NewDatabase(&cfg)
	if err != nil {
		panic(err)
	}

	sqlMgr := db.NewSqlManager(dbConn)

	hwys, err := converter.NewHighwayConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}

	for v := range hwys.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	highwayTypes, err := converter.NewHighwayTypeConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	for v := range highwayTypes.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	signConverter, err := converter.NewHighwaySignConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}

	signLookup := signConverter.(*converter.SignConverter)
	err = generator.SaveLookup(cfg.HugoPath, signLookup)
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signLookup.GetHighQualityLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signLookup.GetPlaceLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signLookup.GetCountyLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signLookup.GetStateLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signLookup.GetGeoJsonLookup())
	if err != nil {
		panic(err)
	}

	for v := range signConverter.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	featureConverter, err := converter.NewFeatureConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}

	for v := range featureConverter.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	countryConverter, err := converter.NewCountryConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	for v := range countryConverter.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	stateConverter, err := converter.NewStateConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	for v := range stateConverter.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	placeConverter, err := converter.NewPlaceConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	for v := range placeConverter.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	counties, err := converter.NewCountyConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	for v := range counties.Convert() {
		err = generator.SaveItem(cfg.HugoPath, v)
		if err != nil {
			panic(err)
		}
	}

	tags, err := converter.NewTagConverter(ctx, sqlMgr)
	if err != nil {
		panic(err)
	}
	err = generator.SaveLookup(cfg.HugoPath, tags)
	if err != nil {
		panic(err)
	}

}
