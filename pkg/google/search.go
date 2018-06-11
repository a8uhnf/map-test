package google

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/a8uhnf/map-test/api"
	"github.com/a8uhnf/map-test/config"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func init() {
	fmt.Println("Hello google.com")
}

type SearchPlaces struct {
	Config config.MapConfig
}

func (s SearchPlaces) Search(in *api.SearchPlacesRequest) (*api.SearchPlacesResponse, error) {
	ret := s.searchPlaces(in)
	return ret, nil
}
func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func (s SearchPlaces) searchPlaces(in *api.SearchPlacesRequest) *api.SearchPlacesResponse {
	flag.Parse()

	cfg := s.Config
	var err error
	var client *maps.Client
	if cfg.APIKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(cfg.APIKey))
	} else if cfg.ClientID != "" || cfg.ClientSignature != "" {
		client, err = maps.NewClient(maps.WithClientIDAndSignature(cfg.ClientID, cfg.ClientSignature))
	} else {
		usageAndExit("Please specify an API Key, or Client ID and Signature.")
	}
	check(err)

	r := &maps.TextSearchRequest{
		Query:    in.Query,
		Language: in.Language,
		Radius:   uint(in.Radius),
		OpenNow:  in.OpenNow,
		MaxPrice: maps.PriceLevel(in.MinPrice),
		MinPrice: maps.PriceLevel(in.MinPrice),
		Type:     maps.PlaceType(in.PlaceType),
	}

	if in.Location != nil {
		r.Location = &maps.LatLng{
			Lat: in.Location.Lat,
			Lng: in.Location.Lng,
		}
	}

	resp, err := client.TextSearch(context.Background(), r)
	if err != nil {
		log.Println(err)
	}

	ret, err := parseResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Returns Search Results")
	log.Printf("Total hits: %d\n", len(ret.Results))
	for i, v := range ret.Results {
		log.Printf("Places #%d:\nName: %s\nAddress: %s", i+1, v.Name, v.FormattedAddress)
	}

	return ret
}

func parseResponse(in maps.PlacesSearchResponse) (*api.SearchPlacesResponse, error) {
	resp := &api.SearchPlacesResponse{}

	res := []*api.SearchPlacesResponse_PlacesSearchResult{}

	for _, v := range in.Results {
		tmp := &api.SearchPlacesResponse_PlacesSearchResult{
			Name:             v.Name,
			PlaceID:          v.PlaceID,
			Icon:             v.Icon,
			Rating:           v.Rating,
			Scope:            v.Scope,
			Types:            v.Types,
			FormattedAddress: v.FormattedAddress,
			Geometry: &api.SearchPlacesResponse_PlacesSearchResult_AddressGeometry{
				Location: &api.LatLng{
					Lat: v.Geometry.Location.Lat,
					Lng: v.Geometry.Location.Lng,
				},
				LocationType: v.Geometry.LocationType,
				Bounds: &api.SearchPlacesResponse_PlacesSearchResult_LatLngBounds{
					Northeast: &api.LatLng{
						Lat: v.Geometry.Bounds.NorthEast.Lat,
						Lng: v.Geometry.Bounds.NorthEast.Lng,
					},
					Southwest: &api.LatLng{
						Lat: v.Geometry.Bounds.NorthEast.Lat,
						Lng: v.Geometry.Bounds.NorthEast.Lng,
					},
				},
				Viewport: &api.SearchPlacesResponse_PlacesSearchResult_LatLngBounds{
					Northeast: &api.LatLng{
						Lat: v.Geometry.Bounds.NorthEast.Lat,
						Lng: v.Geometry.Bounds.NorthEast.Lng,
					},
					Southwest: &api.LatLng{
						Lat: v.Geometry.Bounds.NorthEast.Lat,
						Lng: v.Geometry.Bounds.NorthEast.Lng,
					},
				},
			},
		}
		res = append(res, tmp)
	}
	resp.Results = res

	return resp, nil
}

func parseLocation(location string, r *maps.TextSearchRequest) {
	if location != "" {
		l, err := maps.ParseLatLng(location)
		check(err)
		r.Location = &l
	}
}

func parsePriceLevel(priceLevel string) maps.PriceLevel {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree
	case "1":
		return maps.PriceLevelInexpensive
	case "2":
		return maps.PriceLevelModerate
	case "3":
		return maps.PriceLevelExpensive
	case "4":
		return maps.PriceLevelVeryExpensive
	default:
		usageAndExit(fmt.Sprintf("Unknown price level: '%s'", priceLevel))
	}
	return maps.PriceLevelFree
}

func parsePriceLevels(minprice string, maxprice string, r *maps.TextSearchRequest) {
	if minprice != "" {
		r.MinPrice = parsePriceLevel(minprice)
	}

	if maxprice != "" {
		r.MaxPrice = parsePriceLevel(minprice)
	}
}

func parsePlaceType(placeType string, r *maps.TextSearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		if err != nil {
			usageAndExit(fmt.Sprintf("Unknown place type \"%v\"", placeType))
		}

		r.Type = t
	}
}
