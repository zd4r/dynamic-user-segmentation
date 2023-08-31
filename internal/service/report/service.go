package report

type Service struct {
	reportRepository reportRepository
}

func NewService(reportRepository reportRepository) *Service {
	return &Service{
		reportRepository: reportRepository,
	}
}
