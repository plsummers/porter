package test

import (
	"github.com/porter-dev/porter/internal/models"
	"github.com/porter-dev/porter/internal/repository"
)

// EnvironmentRepository uses gorm.DB for querying the database
type EnvironmentRepository struct {
}

// NewEnvironmentRepository returns a DefaultEnvironmentRepository which uses
// gorm.DB for querying the database
func NewEnvironmentRepository() repository.EnvironmentRepository {
	return &EnvironmentRepository{}
}

func (repo *EnvironmentRepository) CreateEnvironment(env *models.Environment) (*models.Environment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) ReadEnvironment(projectID, clusterID, gitInstallationID uint) (*models.Environment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) DeleteEnvironment(env *models.Environment) (*models.Environment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) CreateDeployment(deployment *models.Deployment) (*models.Deployment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) UpdateDeployment(deployment *models.Deployment) (*models.Deployment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) ReadDeployment(environmentID uint, namespace string) (*models.Deployment, error) {
	panic("unimplemented")
}

func (repo *EnvironmentRepository) DeleteDeployment(deployment *models.Deployment) (*models.Deployment, error) {
	panic("unimplemented")
}