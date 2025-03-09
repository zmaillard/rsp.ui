// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/twpayne/go-geom"
	"highway-sign-portal-builder/pkg/types"
)

type SignAdminAreaCountry struct {
	ID              int32
	Name            pgtype.Text
	SubdivisionName pgtype.Text
	Slug            pgtype.Text
	FeaturedSignID  pgtype.Int4
	ImageCount      pgtype.Int4
}

type SignAdminAreaCounty struct {
	ID               int32
	Name             pgtype.Text
	Slug             pgtype.Text
	AdminAreaStateid pgtype.Int4
	ImageCount       pgtype.Int4
}

type SignAdminAreaPlace struct {
	ID               int32
	Name             pgtype.Text
	Slug             pgtype.Text
	AdminAreaStateid pgtype.Int4
	ImageCount       pgtype.Int4
}

type SignAdminAreaState struct {
	ID                 int32
	Name               pgtype.Text
	SubdivisionName    pgtype.Text
	Slug               pgtype.Text
	AdminareaCountryID pgtype.Int4
	FeaturedSignID     pgtype.Int4
	ImageCount         pgtype.Int4
}

type SignFeature struct {
	ID                 int32
	Point              interface{}
	Name               pgtype.Text
	AdminAreaCountryID pgtype.Int4
	AdminAreaStateID   pgtype.Int4
	Featured           pgtype.Bool
	FeatureTypeID      pgtype.Int4
}

type SignFeatureLink struct {
	ID              int32
	FromFeature     pgtype.Int4
	ToFeature       pgtype.Int4
	RoadName        pgtype.Text
	TempPlaceholder pgtype.Text
	Link            interface{}
}

type SignFeatureLinkAlias struct {
	ID                 int32
	Name               string
	FeatureAliasTypeID pgtype.Int4
	FeatureLinkID      pgtype.Int4
}

type SignFeatureLinkAliasType struct {
	ID   int32
	Name pgtype.Text
}

type SignFeatureLinkHighway struct {
	ID            int32
	HighwayID     pgtype.Int4
	FeatureLinkID pgtype.Int4
	IsDescending  pgtype.Bool
}

type SignFeatureType struct {
	ID   int32
	Name string
}

type SignFlickrSet struct {
	ID              int32
	HighwayID       pgtype.Int4
	FlickrSetID     pgtype.Text
	PrimaryFlickrID pgtype.Text
	LastUpdated     pgtype.Date
}

type SignHighway struct {
	ID                 int32
	HighwayName        pgtype.Text
	ScopeID            pgtype.Int4
	Slug               pgtype.Text
	HighwayTypeID      pgtype.Int4
	ImageName          pgtype.Text
	DateAdded          pgtype.Date
	SortNumber         pgtype.Int4
	AdminAreaCountryID pgtype.Int4
	AdminAreaStateID   pgtype.Int4
}

type SignHighwayScope struct {
	ID    int32
	Scope pgtype.Text
}

type SignHighwayType struct {
	ID                 int32
	HighwayTypeName    pgtype.Text
	Sort               pgtype.Int4
	Slug               pgtype.Text
	DisplayImageID     pgtype.Int4
	ImageCount         pgtype.Int4
	AdminAreaCountryID pgtype.Int4
}

type SignHighwaysign struct {
	ID                 int32
	Flickrid           pgtype.Text
	DateTaken          pgtype.Timestamp
	DateAdded          pgtype.Timestamp
	Title              pgtype.Text
	SignDescription    pgtype.Text
	ImageWidth         pgtype.Int4
	ImageHeight        pgtype.Int4
	Point              interface{}
	Imageid            pgtype.Int8
	Lastsyncwithflickr pgtype.Date
	LastUpdate         pgtype.Timestamp
	CroppedImageID     pgtype.Int8
	LastIndexed        pgtype.Timestamp
	Archived           pgtype.Bool
	FeatureID          pgtype.Int4
	AdminAreaCountryID pgtype.Int4
	AdminAreaStateID   pgtype.Int4
	AdminAreaCountyID  pgtype.Int4
	AdminAreaPlaceID   pgtype.Int4
	Quality            int32
}

type SignHighwaysignHighway struct {
	ID            int32
	HighwayID     pgtype.Int4
	HighwaysignID pgtype.Int4
	IsTo          pgtype.Bool
}

type SignHighwaysignPendingChange struct {
	ID            int32
	HighwaysignID pgtype.Int4
	ChangedOn     pgtype.Timestamp
}

type SignHighwaysignStaging struct {
	ID          pgtype.Int4
	ImageWidth  pgtype.Int4
	ImageHeight pgtype.Int4
	DateTaken   pgtype.Timestamp
	Imageid     pgtype.Int8
	Latitude    pgtype.Float8
	Longitude   pgtype.Float8
}

type SignTag struct {
	ID              int32
	Name            pgtype.Text
	Slug            pgtype.Text
	FlickrOnly      pgtype.Bool
	CategoryDetails pgtype.Text
	IsCategory      bool
}

type SignTagHighwaysign struct {
	ID            int32
	TagID         pgtype.Int4
	HighwaysignID pgtype.Int4
}

type SignVwHighwayCountByState struct {
	AdminAreaStateID pgtype.Int4
	HighwayID        pgtype.Int4
	ImageCount       int64
}

type SignVwUnprocessedLink struct {
	ID              int32
	FromFeature     pgtype.Int4
	ToFeature       pgtype.Int4
	RoadName        pgtype.Text
	TempPlaceholder pgtype.Text
	Link            interface{}
}

type SignVwfeaturelinktile struct {
	ID              int32
	FromFeature     pgtype.Int4
	ToFeature       pgtype.Int4
	RoadName        pgtype.Text
	TempPlaceholder pgtype.Text
	Link            geom.Point
}

type SignVwfeaturelinkwithhighway struct {
	ID              int32
	Link            interface{}
	Highways        string
	TempPlaceholder pgtype.Text
}

type SignVwfeaturetile struct {
	ID                 int32
	Point              geom.Point
	Name               pgtype.Text
	AdminAreaCountryID pgtype.Int4
	AdminAreaStateID   pgtype.Int4
}

type SignVwhighwaysigntile struct {
	ID                 int32
	Flickrid           pgtype.Text
	DateTaken          pgtype.Timestamp
	DateAdded          pgtype.Timestamp
	Title              pgtype.Text
	SignDescription    pgtype.Text
	ImageWidth         pgtype.Int4
	ImageHeight        pgtype.Int4
	Point              geom.Point
	Imageid            pgtype.Int8
	Lastsyncwithflickr pgtype.Date
	LastUpdate         pgtype.Timestamp
	CroppedImageID     pgtype.Int8
	LastIndexed        pgtype.Timestamp
	Archived           pgtype.Bool
	FeatureID          pgtype.Int4
	AdminAreaCountryID pgtype.Int4
	AdminAreaStateID   pgtype.Int4
	AdminAreaCountyID  pgtype.Int4
	AdminAreaPlaceID   pgtype.Int4
}

type SignVwhugocountry struct {
	ID              int32
	CountryName     pgtype.Text
	CountrySlug     pgtype.Text
	SubdivisionName pgtype.Text
	ImageCount      int64
	States          []types.AdminArea
	Featured        *types.ImageID
	HighwayTypes    []types.AdminArea
}

type SignVwhugocounty struct {
	ID         int32
	CountySlug pgtype.Text
	CountyName pgtype.Text
	StateSlug  pgtype.Text
	StateName  pgtype.Text
	ImageCount int64
}

type SignVwhugofeature struct {
	ID          int32
	Point       interface{}
	Name        pgtype.Text
	Signs       interface{}
	StateName   pgtype.Text
	StateSlug   pgtype.Text
	CountryName pgtype.Text
	CountrySlug pgtype.Text
}

type SignVwhugofeaturelink struct {
	ID          int32
	FromFeature pgtype.Int4
	ToFeature   pgtype.Int4
	RoadName    pgtype.Text
	Highways    []string
	ToPoint     geom.Point
	FromPoint   geom.Point
}

type SignVwhugohighway struct {
	ID               int32
	HighwayName      pgtype.Text
	Slug             pgtype.Text
	SortNumber       pgtype.Int4
	ImageName        pgtype.Text
	HighwayTypeSlug  pgtype.Text
	HighwayTypeName  pgtype.Text
	States           interface{}
	Counties         interface{}
	Places           interface{}
	PreviousFeatures interface{}
	NextFeatures     interface{}
}

type SignVwhugohighwaysign struct {
	ID              int32
	Title           pgtype.Text
	SignDescription pgtype.Text
	FeatureID       pgtype.Int4
	DateTaken       pgtype.Timestamp
	Imageid         types.ImageID
	Flickrid        string
	Point           geom.Point
	CountrySlug     string
	StateSlug       string
	PlaceSlug       string
	CountySlug      string
	Tags            []string
	Categories      []string
	Highways        []string
	IsTo            []string
	ImageHeight     pgtype.Int4
	ImageWidth      pgtype.Int4
	Quality         int32
}

type SignVwhugohighwaytype struct {
	ID              int32
	HighwayTypeName pgtype.Text
	HighwayTypeSlug pgtype.Text
	Sort            pgtype.Int4
	Imagecount      int64
	Imageid         *types.ImageID
	Highways        interface{}
	Country         pgtype.Text
}

type SignVwhugoplace struct {
	ID         int32
	PlaceSlug  pgtype.Text
	PlaceName  pgtype.Text
	StateSlug  pgtype.Text
	StateName  pgtype.Text
	ImageCount int64
}

type SignVwhugostate struct {
	ID              int32
	StateName       pgtype.Text
	StateSlug       pgtype.Text
	SubdivisionName pgtype.Text
	ImageCount      int64
	Highways        []string
	Places          []types.AdminArea
	Counties        []types.AdminArea
	Featured        *types.ImageID
	CountrySlug     pgtype.Text
	Categories      []string
}

type SignVwindexsign struct {
	Imageid         pgtype.Int8
	Title           pgtype.Text
	SignDescription pgtype.Text
	DateTaken       pgtype.Timestamp
	CountrySlug     pgtype.Text
	CountryName     pgtype.Text
	StateSlug       pgtype.Text
	StateName       pgtype.Text
	CountyName      pgtype.Text
	CountySlug      pgtype.Text
	PlaceName       pgtype.Text
	PlaceSlug       pgtype.Text
	Tagitems        interface{}
	Hwys            []byte
	Point           interface{}
	LastIndexed     pgtype.Timestamp
	LastUpdate      pgtype.Timestamp
	Quality         int32
}
