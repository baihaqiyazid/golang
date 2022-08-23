package simple

type BarRepository struct{

}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct{
	*BarRepository
}

func NewBarService(barReposiory *BarRepository) *BarService {
	return &BarService{BarRepository: barReposiory}
}