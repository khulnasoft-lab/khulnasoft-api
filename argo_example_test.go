package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

func ExampleAPI_ArgoSmartRouting() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	smartRoutingSettings, err := api.ArgoSmartRouting(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("smart routing is %s", smartRoutingSettings.Value)
}

func ExampleAPI_ArgoTieredCaching() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	tieredCachingSettings, err := api.ArgoTieredCaching(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tiered caching is %s", tieredCachingSettings.Value)
}

func ExampleAPI_UpdateArgoSmartRouting() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	smartRoutingSettings, err := api.UpdateArgoSmartRouting(context.Background(), "01a7362d577a6c3019a474fd6f485823", "on")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("smart routing is %s", smartRoutingSettings.Value)
}

func ExampleAPI_UpdateArgoTieredCaching() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	tieredCachingSettings, err := api.UpdateArgoTieredCaching(context.Background(), "01a7362d577a6c3019a474fd6f485823", "on")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tiered caching is %s", tieredCachingSettings.Value)
}
