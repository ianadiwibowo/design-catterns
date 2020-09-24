package main

import "fmt"

// Both the proxy object and the real one use the same interface.
type CatServiceInterface interface {
	GetCats() []*Cat
}

type Cat struct {
	ID   int
	Name string
}

// This is the proxy object, aka the lite version or the real object. Why lite?
// Because in this case an expensive HTTP call to external CatCloudService
// over the internet can be avoided if there is a local data cache.
type CatLocalCacheService struct {
	ExternalService *CatCloudService
	LocalCache      []*Cat
}

// A proxy object with the same function name can add more functionalities.
// In this case checking whether local data exists, before calling the real
// object which will involve more expensive activity.
func (c *CatLocalCacheService) GetCats() []*Cat {
	if len(c.LocalCache) > 0 {
		fmt.Println("Log: Retrieve from local cache")
		return c.LocalCache
	} else {
		return c.ExternalService.GetCats()
	}
}

// This is the real object, a service in the cloud.
type CatCloudService struct {
	CloudData []*Cat
}

// Calling this function over the internet is expensive. That's why the proxy
// object tries to intercept the function first, before calling the real object.
func (c *CatCloudService) GetCats() []*Cat {
	fmt.Println("Log: Retrieve from cloud service")
	return c.CloudData
}

func main() {
	// The real object is instantiated here only for the sake of example.
	// Usually, the proxy object hides the implementation of the real object.
	// Only the proxy object has access to the real object.
	// That's why it's called a proxy, aka door.
	cloudService := &CatCloudService{
		CloudData: []*Cat{
			{1, "Lupita"},
			{2, "Bupita"},
			{3, "Lupita"},
		},
	}
	localService := &CatLocalCacheService{
		ExternalService: cloudService,
	}

	localService.GetCats() // => Log: Retrieve from cloud service

	// Now that the system has local data, it will prioritize to use it first.
	localService.LocalCache = []*Cat{
		{4, "Taeka"},
	}
	localService.GetCats() // => Log: Retrieve from local cache
}
