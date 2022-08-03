//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/usecase/$GOFILE
package usecase

import (
	"testing"

	mock_repository "github.com/diracmax/review-job-backend/mock/usecase/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userUsecase_Register(t *testing.T) {
	type testCaseIn struct {
		mockRepository func(r *mock_repository.MockUserRepository)
		name           string
		rawPassword    string
	}
	type testCaseOut struct {
		err error
	}
	tests := []struct {
		name        string
		testCaseIn  testCaseIn
		testCaseOut testCaseOut
	}{
		{
			name: "正常に保存できた場合",
			testCaseIn: testCaseIn{
				mockRepository: func(r *mock_repository.MockUserRepository) {
					r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				},
				name:        "diracmax",
				rawPassword: "password",
			},
			testCaseOut: testCaseOut{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockUserRepository(mockCtrl)
			tt.testCaseIn.mockRepository(mockRepo)
			u := &userUsecase{
				repository: mockRepo,
			}
			err := u.Register(tt.testCaseIn.name, tt.testCaseIn.rawPassword)
			assert.Equal(t, tt.testCaseOut.err, err)
		})
	}
}
