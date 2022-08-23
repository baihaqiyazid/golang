package simple

type FooBarService struct{
	*FooService
	*BarService
}

func NewFooBarRepository(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{FooService: fooService, BarService: barService}
}