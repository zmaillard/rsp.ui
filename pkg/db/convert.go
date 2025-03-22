package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func ToPg4(val int32) pgtype.Int4 {
	return pgtype.Int4{Int32: val, Valid: true}
}

func IntToPg4(val int) pgtype.Int4 {
	return ToPg4(int32(val))
}

func TextToPgText(val string) pgtype.Text {
	return pgtype.Text{String: val, Valid: true}
}

func BoolToPgBool(val bool) pgtype.Bool {
	return pgtype.Bool{Bool: val, Valid: true}
}

func IntToPg8(val int64) pgtype.Int8 {
	return pgtype.Int8{Int64: val, Valid: true}
}

func Float64ToPgFloat8(val float64) pgtype.Float8 {
	return pgtype.Float8{Float64: val, Valid: true}
}

func TimeToPgTimestamp(val time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: val, Valid: true}
}

func ToPgDate(val time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  val,
		Valid: true,
	}
}
