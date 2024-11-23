package client

import (
	"github.com/a-finocchiaro/adsb_flightradar_top10/internal"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/geography"
)

func GetZones(requester common.Requester) (geography.Fr24Zones, error) {
	var zones geography.Fr24Zones

	if err := internal.SendRequest(requester, common.FR24_ENDPOINTS["zones"], &zones); err != nil {
		return zones, common.NewFr24Error(err)
	}

	return zones, nil
}
