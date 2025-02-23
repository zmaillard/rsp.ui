// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddHighwayToFeatureLink(ctx context.Context, arg AddHighwayToFeatureLinkParams) (int32, error)
	AddTagToSign(ctx context.Context, arg AddTagToSignParams) (SignTagHighwaysign, error)
	AdminAreaCountryCountByName(ctx context.Context, name pgtype.Text) (int64, error)
	AdminAreaCountyCountByName(ctx context.Context, arg AdminAreaCountyCountByNameParams) (int64, error)
	AdminAreaPlaceCountByName(ctx context.Context, arg AdminAreaPlaceCountByNameParams) (int64, error)
	AdminAreaStateCountByName(ctx context.Context, arg AdminAreaStateCountByNameParams) (int64, error)
	CreateFeature(ctx context.Context, arg CreateFeatureParams) (int32, error)
	CreateFeatureLink(ctx context.Context, arg CreateFeatureLinkParams) (int32, error)
	CreateHighway(ctx context.Context, arg CreateHighwayParams) (SignHighway, error)
	CreateHighwayType(ctx context.Context, arg CreateHighwayTypeParams) (SignHighwayType, error)
	CreateSign(ctx context.Context, arg CreateSignParams) (int32, error)
	CreateTag(ctx context.Context, arg CreateTagParams) (SignTag, error)
	DeleteFeature(ctx context.Context, id int32) error
	DeleteFeatureLink(ctx context.Context, id int32) error
	DeleteFeatureLinkHighway(ctx context.Context, featureLinkID pgtype.Int4) error
	DeleteHighwaysOnSign(ctx context.Context, highwaysignID pgtype.Int4) error
	DeletePendingChange(ctx context.Context, id int32) error
	DeleteStaging(ctx context.Context, id pgtype.Int4) error
	DeleteTagsOnSign(ctx context.Context, highwaysignID pgtype.Int4) error
	GetAdminAreaCountryByName(ctx context.Context, name pgtype.Text) (SignAdminAreaCountry, error)
	GetAdminAreaCountyByName(ctx context.Context, arg GetAdminAreaCountyByNameParams) (SignAdminAreaCounty, error)
	GetAdminAreaPlaceByName(ctx context.Context, arg GetAdminAreaPlaceByNameParams) (SignAdminAreaPlace, error)
	GetAdminAreaStateByName(ctx context.Context, arg GetAdminAreaStateByNameParams) (SignAdminAreaState, error)
	GetAllCountries(ctx context.Context) ([]SignAdminAreaCountry, error)
	GetAllFeaturedFeatures(ctx context.Context, featured pgtype.Bool) ([]GetAllFeaturedFeaturesRow, error)
	GetAllFeatures(ctx context.Context) ([]GetAllFeaturesRow, error)
	GetAllHighwayTypes(ctx context.Context) ([]SignHighwayType, error)
	GetAllHighways(ctx context.Context) ([]SignHighway, error)
	GetAllSigns(ctx context.Context) ([]GetAllSignsRow, error)
	GetAllSignsByCountrySearch(ctx context.Context, arg GetAllSignsByCountrySearchParams) ([]GetAllSignsByCountrySearchRow, error)
	GetAllSignsByCountySearch(ctx context.Context, arg GetAllSignsByCountySearchParams) ([]GetAllSignsByCountySearchRow, error)
	GetAllSignsByStateSearch(ctx context.Context, arg GetAllSignsByStateSearchParams) ([]GetAllSignsByStateSearchRow, error)
	GetAllStaging(ctx context.Context) ([]SignHighwaysignStaging, error)
	GetAllStates(ctx context.Context) ([]SignAdminAreaState, error)
	GetCategories(ctx context.Context) ([]SignTag, error)
	GetCountry(ctx context.Context, id int32) (SignAdminAreaCountry, error)
	GetFeature(ctx context.Context, id int32) (GetFeatureRow, error)
	GetFeatureConnectedCount(ctx context.Context, fromFeature pgtype.Int4) (int64, error)
	GetFeatureLink(ctx context.Context, id int32) (GetFeatureLinkRow, error)
	GetFeatureLinkHighway(ctx context.Context, arg GetFeatureLinkHighwayParams) (SignFeatureLinkHighway, error)
	GetFeatureLinkHighwayDirection(ctx context.Context, arg GetFeatureLinkHighwayDirectionParams) (pgtype.Bool, error)
	GetFeatureLinkHighways(ctx context.Context, featureLinkID pgtype.Int4) ([]GetFeatureLinkHighwaysRow, error)
	GetFeatureLinksByIds(ctx context.Context, dollar_1 []int32) ([]GetFeatureLinksByIdsRow, error)
	GetFeaturesWithinBuffer(ctx context.Context, arg GetFeaturesWithinBufferParams) ([]GetFeaturesWithinBufferRow, error)
	GetHighway(ctx context.Context, id int32) (SignHighway, error)
	GetHighwayByName(ctx context.Context, highwayName pgtype.Text) (SignHighway, error)
	GetHighwayType(ctx context.Context, id int32) (SignHighwayType, error)
	GetHighwayTypeByName(ctx context.Context, highwayTypeName pgtype.Text) (SignHighwayType, error)
	GetHighwaysForStateAndCountry(ctx context.Context, arg GetHighwaysForStateAndCountryParams) ([]SignHighway, error)
	GetHighwaysOnSign(ctx context.Context, highwaysignID pgtype.Int4) ([]GetHighwaysOnSignRow, error)
	GetHighwaysStartWith(ctx context.Context, highwayName pgtype.Text) ([]SignHighway, error)
	GetHugoCounties(ctx context.Context) ([]GetHugoCountiesRow, error)
	GetHugoCountries(ctx context.Context) ([]SignVwhugocountry, error)
	GetHugoFeatureLinks(ctx context.Context) ([]GetHugoFeatureLinksRow, error)
	GetHugoFeatures(ctx context.Context) ([]SignVwhugofeature, error)
	GetHugoHighwaySigns(ctx context.Context) ([]SignVwhugohighwaysign, error)
	GetHugoHighwayTypes(ctx context.Context) ([]SignVwhugohighwaytype, error)
	GetHugoHighways(ctx context.Context) ([]SignVwhugohighway, error)
	GetHugoPlaces(ctx context.Context) ([]GetHugoPlacesRow, error)
	GetHugoStates(ctx context.Context) ([]GetHugoStatesRow, error)
	GetPendingChanges(ctx context.Context) ([]SignHighwaysignPendingChange, error)
	GetScope(ctx context.Context, id int32) (SignHighwayScope, error)
	GetScopeByName(ctx context.Context, scope pgtype.Text) (SignHighwayScope, error)
	GetScopes(ctx context.Context) ([]SignHighwayScope, error)
	GetSign(ctx context.Context, id int32) (GetSignRow, error)
	GetSignByImageId(ctx context.Context, imageid pgtype.Int8) (GetSignByImageIdRow, error)
	GetSignDetails(ctx context.Context, id int32) (GetSignDetailsRow, error)
	GetSignTags(ctx context.Context, highwaysignID pgtype.Int4) ([]pgtype.Text, error)
	GetSignsByFeatureId(ctx context.Context, featureID pgtype.Int4) ([]GetSignsByFeatureIdRow, error)
	GetSignsByIds(ctx context.Context, dollar_1 []int32) ([]GetSignsByIdsRow, error)
	GetSignsOnHighway(ctx context.Context, highwayID pgtype.Int4) ([]GetSignsOnHighwayRow, error)
	GetStaging(ctx context.Context, id pgtype.Int4) (SignHighwaysignStaging, error)
	GetState(ctx context.Context, id int32) (SignAdminAreaState, error)
	GetStateByName(ctx context.Context, name pgtype.Text) (SignAdminAreaState, error)
	GetStatesByCountry(ctx context.Context, adminareaCountryID pgtype.Int4) ([]SignAdminAreaState, error)
	GetTagById(ctx context.Context, id int32) (SignTag, error)
	GetTagByName(ctx context.Context, name pgtype.Text) (SignTag, error)
	GetTagsStartWith(ctx context.Context, name pgtype.Text) ([]SignTag, error)
	InsertAdminAreaCountry(ctx context.Context, arg InsertAdminAreaCountryParams) (SignAdminAreaCountry, error)
	InsertAdminAreaCounty(ctx context.Context, arg InsertAdminAreaCountyParams) (SignAdminAreaCounty, error)
	InsertAdminAreaPlace(ctx context.Context, arg InsertAdminAreaPlaceParams) (SignAdminAreaPlace, error)
	InsertAdminAreaState(ctx context.Context, arg InsertAdminAreaStateParams) (SignAdminAreaState, error)
	InsertAdminAreaStateWithSubdivision(ctx context.Context, arg InsertAdminAreaStateWithSubdivisionParams) (SignAdminAreaState, error)
	InsertHighwaySorting(ctx context.Context, arg InsertHighwaySortingParams) (SignHighwaysignHighway, error)
	InsertStaging(ctx context.Context, arg InsertStagingParams) (SignHighwaysignStaging, error)
	RemoveHighwayFromFeatureLink(ctx context.Context, arg RemoveHighwayFromFeatureLinkParams) error
	RemoveHighwaySorting(ctx context.Context, arg RemoveHighwaySortingParams) error
	ReverseFeatureLink(ctx context.Context, arg ReverseFeatureLinkParams) error
	ReverseFeatureLinkHighway(ctx context.Context, id int32) error
	UpdateAllHighwaySortingsOnSign(ctx context.Context, arg UpdateAllHighwaySortingsOnSignParams) error
	UpdateBeginAndEnd(ctx context.Context, arg UpdateBeginAndEndParams) error
	UpdateCategoryDetails(ctx context.Context, arg UpdateCategoryDetailsParams) error
	UpdateFeatureAdminArea(ctx context.Context, arg UpdateFeatureAdminAreaParams) error
	UpdateFeatureLinkName(ctx context.Context, arg UpdateFeatureLinkNameParams) error
	UpdateFeatureName(ctx context.Context, arg UpdateFeatureNameParams) error
	UpdateFeatured(ctx context.Context, arg UpdateFeaturedParams) error
	UpdateFeaturedSignForHighwayType(ctx context.Context, arg UpdateFeaturedSignForHighwayTypeParams) error
	UpdateFeaturedSignForState(ctx context.Context, arg UpdateFeaturedSignForStateParams) error
	UpdateHighwayType(ctx context.Context, arg UpdateHighwayTypeParams) error
	UpdateImage(ctx context.Context, arg UpdateImageParams) error
	UpdateLastUpdated(ctx context.Context, arg UpdateLastUpdatedParams) error
	UpdateSignAdminAreas(ctx context.Context, arg UpdateSignAdminAreasParams) error
	UpdateSignDescription(ctx context.Context, arg UpdateSignDescriptionParams) error
	UpdateSignFeature(ctx context.Context, arg UpdateSignFeatureParams) error
	UpdateSignLocation(ctx context.Context, arg UpdateSignLocationParams) error
	UpdateSignTitle(ctx context.Context, arg UpdateSignTitleParams) error
}

var _ Querier = (*Queries)(nil)
