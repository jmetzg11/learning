package main

import (
	"fmt"
	"log"
)

type Route struct{}

// two log outputs for the same errors
func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates") // don't log and return an error
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates") // don't log and return an error
		return Route{}, err
	}
	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat) // don't log and return an error
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng) // don't log and return an error
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// Solution, the caller of GetRoute2 should be handling the error with logs
func GetRoute2(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, err
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// Better solution with more context
func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, fmt.Errorf("Failed to validate source coordinates: %w", err)
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, fmt.Errorf("Failed to validate target coordinates: %w", err)
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}
