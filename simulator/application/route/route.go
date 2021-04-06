package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

// NewRoute creates a *Route struct
func NewRoute() *Route {
	return &Route{}
}

// LoadPositions loads from a .txt file all positions (lat and long) to the Position attribute of the struct
func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route id not informed")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ",")

		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{
			Lat:  latitude,
			Long: longitude,
		})
	}

	return nil
}

// ExportJsonPositions generates a slice of string in Json using PartialRoutePosition struct
func (route *Route) ExportJsonPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string

	total := len(route.Positions)
	for index, value := range route.Positions {
		partialRoute.ID = route.ID
		partialRoute.ClientID = route.ClientID
		partialRoute.Position = []float64{value.Lat, value.Long}
		partialRoute.Finished = false

		if total-1 == index {
			partialRoute.Finished = true
		}

		jsonRoute, err := json.Marshal(partialRoute)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
