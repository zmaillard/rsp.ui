package main

import (
	"highway-sign-portal-builder/pkg/config"
	"highway-sign-portal-builder/pkg/database"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/services"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := database.InitializeDatabase(cfg)
	if err != nil {
		panic(err)
	}

	datastore := services.NewDatastore(db)
	placeService := datastore.GetPlaceService()
	highwayService := datastore.GetHighwayService()

	hwys, err := highwayService.GetAllHighways()
	if err != nil {
		panic(err)
	}

	for _, v := range hwys {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	highwayTypes, err := highwayService.GetAllHighwayTypes()
	if err != nil {
		panic(err)
	}
	for _, v := range highwayTypes {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	signService := datastore.GetSignService()

	signs, err := signService.GetAllSigns()
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs)
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs.GetHighQualityLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs.GetPlaceLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs.GetCountyLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs.GetStateLookup())
	if err != nil {
		panic(err)
	}

	err = generator.SaveLookup(cfg.HugoPath, signs.GetGeoJsonLookup())
	if err != nil {
		panic(err)
	}

	for _, v := range signs {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	featureService := datastore.GetFeatureService()

	features, err := featureService.GetAllFeatures()
	if err != nil {
		panic(err)
	}

	for _, v := range features {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}
	countries, err := placeService.GetAllCountries()
	if err != nil {
		panic(err)
	}

	for _, v := range countries {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	states, err := placeService.GetAllStates()
	if err != nil {
		panic(err)
	}
	for _, v := range states {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	places, err := placeService.GetAllPlaces()
	if err != nil {
		panic(err)
	}
	for _, v := range places {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	counties, err := placeService.GetAllCounties()
	if err != nil {
		panic(err)
	}
	for _, v := range counties {
		err = generator.SaveItem(cfg.HugoPath, v.ConvertToDto())
		if err != nil {
			panic(err)
		}
	}

	tags, err := signService.GetAllTags()
	if err != nil {
		panic(err)
	}
	err = generator.SaveLookup(cfg.HugoPath, tags)
	if err != nil {
		panic(err)
	}

}
