package types

import (
	"testing"

	"github.com/datazip-inc/olake/constants"
	"github.com/stretchr/testify/assert"
)

func TestToParquetIncludesCDCTimestampOnlyWhenConfigured(t *testing.T) {
	stream := NewStream("users", "public", nil)
	configured := stream.Wrap(0)

	parquetSchema := stream.Schema.ToParquet(true, configured)
	_, found := parquetSchema.Lookup(constants.CdcTimestamp)
	assert.False(t, found, "non-CDC schemas must not contain the CDC timestamp")

	stream.UpsertField(constants.CdcTimestamp, TimestampMicro, true, true)
	parquetSchema = stream.Schema.ToParquet(true, configured)
	_, found = parquetSchema.Lookup(constants.CdcTimestamp)
	assert.True(t, found, "CDC schemas must contain the injected CDC timestamp")
}
