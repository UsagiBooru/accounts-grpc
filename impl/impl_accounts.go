package impl

import (
	"context"
	"errors"

	"github.com/UsagiBooru/accounts-grpc/gen"
)

type AccountsService struct {
}

func (s *AccountsService) GetAccountName(ctx context.Context, message *gen.GetAccountNameRequest) (*gen.GetAccountNameResponse, error) {
	return nil, errors.New("GetAccountName is not implemented yet")
}

func (s *AccountsService) AddUploadHistory(ctx context.Context, message *gen.AddUploadHistoryRequest) (*gen.AddUploadHistoryResponse, error) {
	return nil, errors.New("AddUploadHistory is not implemented yet")
}

func (s *AccountsService) EditUploadHistory(ctx context.Context, message *gen.EditUploadHistoryRequest) (*gen.EditUploadHistoryResponse, error) {
	return nil, errors.New("EditUploadHistory is not implemented yet")
}

func (s *AccountsService) DeleteUploadHistory(ctx context.Context, message *gen.DeleteUploadHistoryRequest) (*gen.DeleteUploadHistoryResponse, error) {
	return nil, errors.New("DeleteUploadHistory is not implemented yet")
}
