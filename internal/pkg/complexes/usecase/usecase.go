package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/complexes"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// ComplexUsecase represents the usecase for complex using.
type ComplexUsecase struct {
	repo   complexes.ComplexRepo
	logger *zap.Logger
}

// NewComplexUsecase creates a new instance of ComplexUsecase.
func NewComplexUsecase(repo complexes.ComplexRepo, logger *zap.Logger) *ComplexUsecase {
	return &ComplexUsecase{repo: repo, logger: logger}
}

// CreateComplex handles the complex registration process.
func (u *ComplexUsecase) CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error) {
	newComplex := &models.Complex{
		ID:                     uuid.NewV4(),
		CompanyId:              data.CompanyId,
		Name:                   data.Name,
		Description:            data.Description,
		Address:                data.Address,
		DateBeginBuild:         data.DateBeginBuild,
		DateEndBuild:           data.DateEndBuild,
		WithoutFinishingOption: data.WithoutFinishingOption,
		FinishingOption:        data.FinishingOption,
		PreFinishingOption:     data.PreFinishingOption,
		ClassHousing:           data.ClassHousing,
		Parking:                data.Parking,
		Security:               data.Security,
	}

	complex, err := u.repo.CreateComplex(ctx, newComplex)
	if err != nil {
		return nil, err
	}

	return complex, nil
}

// CreateBuilding handles the building registration process.
func (u *ComplexUsecase) CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error) {
	newBuilding := &models.Building{
		ID:           uuid.NewV4(),
		ComplexID:    data.ComplexID,
		Floor:        data.Floor,
		Material:     data.Material,
		Address:      data.Address,
		AddressPoint: data.AddressPoint,
		YearCreation: data.YearCreation,
	}

	building, err := u.repo.CreateBuilding(ctx, newBuilding)
	if err != nil {
		return nil, err
	}

	return building, nil
}

func (u *ComplexUsecase) UpdateComplexPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error) {
	newId := uuid.NewV4()
	newFileName := newId.String() + fileType
	subDirectory := "complexes"
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}
	destination, err := os.Create(directory + "/" + newFileName)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		return "", err
	}
	fileName, err := u.repo.UpdateComplexPhoto(id, subDirectory+"/"+newFileName)
	if err != nil {
		return "", nil
	}
	return fileName, nil
}

// GetComplexById handles the getting complex advert process.
func (u *ComplexUsecase) GetComplexById(ctx context.Context, id uuid.UUID) (foundComplexData *models.ComplexData, err error) {

	if foundComplexData, err = u.repo.GetComplexById(ctx, id); err != nil {
		return nil, err
	}

	return foundComplexData, nil
}

// CreateFlatAdvert handles the creation advert process.
func (u *ComplexUsecase) CreateFlatAdvert(ctx context.Context, data *models.ComplexAdvertFlatCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
		}
	}()
	newAdvertType := &models.AdvertType{
		ID:         uuid.NewV4(),
		AdvertType: data.AdvertTypePlacement,
	}

	newAdvert := &models.Advert{
		ID:             uuid.NewV4(),
		UserID:         data.UserID,
		AdvertTypeID:   newAdvertType.ID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	buildingId := data.BuildingID

	newFlat := &models.Flat{
		ID:                uuid.NewV4(),
		BuildingID:        buildingId,
		AdvertTypeID:      newAdvertType.ID,
		RoomCount:         data.RoomCount,
		Floor:             data.Floor,
		CeilingHeight:     data.CeilingHeight,
		SquareGeneral:     data.SquareGeneral,
		SquareResidential: data.SquareResidential,
		Apartment:         data.Apartment,
	}

	newPriceChange := &models.PriceChange{
		ID:       uuid.NewV4(),
		AdvertID: newAdvert.ID,
		Price:    data.Price,
	}

	if err := u.repo.CreateAdvertType(ctx, tx, newAdvertType); err != nil {
		return nil, err
	}

	if err := u.repo.CreateFlat(ctx, tx, newFlat); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return newAdvert, nil
}

// CreateFlatAdvert handles the creation advert process.
func (u *ComplexUsecase) CreateHouseAdvert(ctx context.Context, data *models.ComplexAdvertHouseCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
		}
	}()
	newAdvertType := &models.AdvertType{
		ID:         uuid.NewV4(),
		AdvertType: data.AdvertTypePlacement,
	}

	newAdvert := &models.Advert{
		ID:             uuid.NewV4(),
		UserID:         data.UserID,
		AdvertTypeID:   newAdvertType.ID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	buildingId := data.BuildingID

	newHouse := &models.House{
		ID:            uuid.NewV4(),
		BuildingID:    buildingId,
		AdvertTypeID:  newAdvertType.ID,
		CeilingHeight: data.CeilingHeight,
		SquareArea:    data.SquareArea,
		SquareHouse:   data.SquareHouse,
		BedroomCount:  data.BedroomCount,
		StatusArea:    data.StatusArea,
		Cottage:       data.Cottage,
		StatusHome:    data.StatusHome,
	}

	newPriceChange := &models.PriceChange{
		ID:       uuid.NewV4(),
		AdvertID: newAdvert.ID,
		Price:    data.Price,
	}

	if err := u.repo.CreateAdvertType(ctx, tx, newAdvertType); err != nil {
		return nil, err
	}

	if err := u.repo.CreateHouse(ctx, tx, newHouse); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return newAdvert, nil
}
