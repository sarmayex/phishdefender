package domain

type PhishDefenderService struct {
	repo DefenderRepository
}

func NewPhishDefenderService(repo DefenderRepository) *PhishDefenderService {
	return &PhishDefenderService{repo: repo}
}
