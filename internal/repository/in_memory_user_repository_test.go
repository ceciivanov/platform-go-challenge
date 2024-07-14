package repository_test

import (
	"testing"

	"github.com/ceciivanov/platform-go-challenge/internal/models"
	"github.com/ceciivanov/platform-go-challenge/internal/repository"
	"github.com/stretchr/testify/assert"
)

func setup() *repository.InMemoryUserRepository {
	repo := repository.NewInMemoryUserRepository()
	repo.GenerateSampleUsers(1, 1) // Create 1 user with 1 asset
	return repo
}

func TestGenerateSampleUsers(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	repo.GenerateSampleUsers(2, 2)

	assert.Equal(t, 2, len(repo.Users))
	for _, user := range repo.Users {
		assert.Equal(t, 2, len(user.Favourites))
	}
}

func TestGetUserFavorites(t *testing.T) {
	repo := setup()

	// Test existing user
	favorites, err := repo.GetUserFavorites(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, favorites)

	// Test non-existing user
	_, err = repo.GetUserFavorites(999)
	assert.Error(t, err)
}

func TestAddUserFavorite(t *testing.T) {
	repo := setup()
	newAsset := models.Insight{
		ID:          2,
		Type:        models.InsightType,
		Description: "New Insight",
		Text:        "Some text",
	}

	// Test adding asset to existing user
	err := repo.AddUserFavorite(1, newAsset)
	assert.NoError(t, err)

	// Verify asset was added
	favorites, _ := repo.GetUserFavorites(1)
	assert.Equal(t, newAsset, favorites[2])

	// Test adding existing asset to user
	err = repo.AddUserFavorite(1, newAsset)
	assert.Error(t, err)

	// Test adding asset to non-existing user
	err = repo.AddUserFavorite(999, newAsset)
	assert.Error(t, err)
}

func TestDeleteUserFavorite(t *testing.T) {
	repo := setup()

	// Test deleting existing asset
	err := repo.DeleteUserFavorite(1, 1)
	assert.NoError(t, err)

	// Verify asset was deleted
	favorites, _ := repo.GetUserFavorites(1)
	assert.Empty(t, favorites)

	// Test deleting non-existing asset
	err = repo.DeleteUserFavorite(1, 999)
	assert.Error(t, err)

	// Test deleting asset from non-existing user
	err = repo.DeleteUserFavorite(999, 1)
	assert.Error(t, err)
}

func TestEditUserFavorite(t *testing.T) {
	repo := setup()
	editedAsset := models.Insight{
		ID:          1,
		Type:        models.InsightType,
		Description: "Edited Insight",
		Text:        "Edited text",
	}

	// Test editing existing asset
	err := repo.EditUserFavorite(1, 1, editedAsset)
	assert.NoError(t, err)

	// Verify asset was edited
	favorites, _ := repo.GetUserFavorites(1)
	assert.Equal(t, editedAsset, favorites[1])

	// Test editing non-existing asset
	err = repo.EditUserFavorite(1, 999, editedAsset)
	assert.Error(t, err)

	// Test editing asset for non-existing user
	err = repo.EditUserFavorite(999, 1, editedAsset)
	assert.Error(t, err)

	// Test editing asset with mismatched ID
	editedAsset.ID = 999
	err = repo.EditUserFavorite(1, 1, editedAsset)
	assert.Error(t, err)

	// Test editing asset with mismatched type
	editedAsset.ID = 1
	editedAsset.Type = models.AudienceType
	err = repo.EditUserFavorite(1, 1, editedAsset)
	assert.Error(t, err)
}
