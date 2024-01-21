package service

type VersionService struct {
	currentVersion string
}

func NewVersionService(initialVersion string) *VersionService {
	return &VersionService{
		currentVersion: initialVersion,
	}
}

func (vs *VersionService) GetCurrentVersion() string {
	return vs.currentVersion
}

func (vs *VersionService) UpdateVersion(newVersion string) {
	vs.currentVersion = newVersion
}
