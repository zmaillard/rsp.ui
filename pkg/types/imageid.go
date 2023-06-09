package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type ImageID string

func NewImageId(imageId string) ImageID {
	return ImageID(imageId)
}

func (i ImageID) String() string {
	return string(i)
}

func (i ImageID) Placeholder(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s_p.jpg", baseImageUrl, i, i)
}

func (i ImageID) Thumbnail(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s_t.jpg", baseImageUrl, i, i)
}

func (i ImageID) Small(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s_s.jpg", baseImageUrl, i, i)
}

func (i ImageID) Medium(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s_m.jpg", baseImageUrl, i, i)
}

func (i ImageID) Large(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s_l.jpg", baseImageUrl, i, i)
}

func (i ImageID) Original(baseImageUrl string) string {
	return fmt.Sprintf("%s%s/%s.jpg", baseImageUrl, i, i)
}

func (i ImageID) ImageS3Key() string {
	return fmt.Sprintf("%s", i)
}

func (i ImageID) ToFullImageName() string {
	return i.String() + ".jpg"
}
func (i *ImageID) Scan(value interface{}) error {
	intImageId, ok := value.(int64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal ImageID value:", value))
	}

	imageId := NewImageId(strconv.Itoa(int(intImageId)))
	*i = imageId

	return nil
}

func (i ImageID) Value() (driver.Value, error) {
	return strconv.Atoi(i.String())
}

func (i ImageID) MarshalJSON() ([]byte, error) {
	type localImageId ImageID
	jsonValue, err := json.Marshal(localImageId(i))
	if err != nil {
		return nil, err
	}

	return []byte(jsonValue), nil
}
