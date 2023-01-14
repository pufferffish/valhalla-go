package valhalla

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

const (
	tilesPath   = "./test/data/utrecht_tiles"
	extractPath = "./test/data/utrecht_tiles/tiles.tar"
)

func testActor(t *testing.T) *Actor {
	config := DefaultConfig()
	config.SetTileDirPath(tilesPath)
	config.SetTileExtractPath(extractPath)
	actor, err := NewActorFromConfig(config)
	if err != nil {
		t.Fatal(err.Error())
		return nil
	}
	return actor
}

func TestConfig(t *testing.T) {
	config := DefaultConfig()
	config.SetTileDirPath(tilesPath)
	config.SetTileExtractPath(extractPath)

	mjolnir := config.Json["mjolnir"].(map[string]interface{})
	if mjolnir["tile_extract"] != extractPath {
		t.Fatal("tile_extract does not match")
	}
	if mjolnir["tile_dir"] != tilesPath {
		t.Fatal("tile_dir does not match")
	}
}

func TestConfigNoTiles(t *testing.T) {
	_, err := NewActorFromFile("non_existent_path")
	if err == nil {
		t.Fatal("File not found error expected")
	} else {
		t.Logf("Expected error: %s\n", err.Error())
	}
}

func TestConfigActor(t *testing.T) {
	actor := testActor(t)
	response, err := actor.Status("")
	if err != nil {
		t.Fatal(err.Error())
	}
	var status map[string]interface{}
	err = json.Unmarshal([]byte(response), &status)
	if err != nil {
		t.Fatal(err.Error())
	}
	if status["tileset_last_modified"] == nil {
		t.Fatal("tileset_last_modified expected in status response")
	}
}

func TestRoute(t *testing.T) {
	query := `{
      "locations": [{"lat": 52.08813, "lon": 5.03231}, {"lat": 52.09987, "lon": 5.14913}],
      "costing": "bicycle",
      "directions_options": {"language": "ru-RU"}
    }`
	actor := testActor(t)
	response, err := actor.Route(query)
	if err != nil {
		t.Fatal(err.Error())
	}

	var route map[string]interface{}
	err = json.Unmarshal([]byte(response), &route)
	if err != nil {
		t.Fatal(err.Error())
	}

	trip := route["trip"].(map[string]interface{})
	if trip == nil {
		t.Fatal("trip expected in route response")
	}

	units := trip["units"]
	if units == nil {
		t.Fatal("units expected in route response")
	}

	if units.(string) != "kilometers" {
		t.Fatal("units is expected to be kilometers in route response")
	}

	summary := trip["summary"].(map[string]interface{})
	if summary == nil {
		t.Fatal("summary expected in route response")
	}

	length := summary["length"]
	if length == nil {
		t.Fatal("length expected in route response")
	}
	if length.(float64) <= 0.7 {
		t.Fatal("length expected to be greater than 0.7 in route response")
	}

	legs := trip["legs"]
	if legs == nil {
		t.Fatal("legs expected in route response")
	}
	if len(legs.([]interface{})) <= 0 {
		t.Fatal("legs expected to be greater than 0 in route response")
	}

	maneuvers := legs.([]interface{})[0].(map[string]interface{})["maneuvers"]
	if maneuvers == nil {
		t.Fatal("maneuvers expected in route response")
	}
	if len(maneuvers.([]interface{})) <= 0 {
		t.Fatal("maneuvers expected to be greater than 0 in route response")
	}

	instruction := maneuvers.([]interface{})[0].(map[string]interface{})["instruction"]
	if instruction == nil {
		t.Fatal("maneuvers expected in route response")
	}
	match, err := regexp.Match("[\u0400-\u04FF]", []byte(instruction.(string)))
	if err != nil {
		t.Fatal(err.Error())
	}
	if !match {
		t.Fatal("Cyrillic not found in instruction")
	}
}
func TestIsochrone(t *testing.T) {
	query := `{
        "locations": [{"lat": 52.08813, "lon": 5.03231}],
        "costing": "pedestrian",
        "contours": [{"time": 1}, {"time": 5}, {"distance": 1}, {"distance": 5}],
        "show_locations": true
    }`
	actor := testActor(t)
	response, err := actor.Isochrone(query)
	if err != nil {
		t.Fatal(err.Error())
	}

	var isochrone map[string]interface{}
	err = json.Unmarshal([]byte(response), &isochrone)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len((isochrone["features"]).([]interface{})) != 6 {
		t.Fatal("Expected 4 isochrones and 2 point layers in response")
	}
}

func TestChangeConfig(t *testing.T) {
	query := `{
      "locations": [
          {"lat": 52.08813, "lon": 5.03231},
          {"lat": 52.09987, "lon": 5.14913}
      ],
      "costing": "bicycle",
      "directions_options": {"language": "ru-RU"}
    }`
	config := DefaultConfig()
	config.SetTileDirPath(tilesPath)
	config.SetTileExtractPath(extractPath)
	config.Json["service_limits"].(map[string]interface{})["bicycle"].(map[string]interface{})["max_distance"] = 1
	actor, err := NewActorFromConfig(config)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = actor.Route(query)
	if err == nil {
		t.Fatal("Error expected but not found")
	}
	if !strings.Contains(err.Error(), "exceeds the max distance limit") {
		t.Fatal(err.Error())
	}
}
