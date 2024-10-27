package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/chub-es/go-link-shortener/internal/entity"
	"github.com/chub-es/go-link-shortener/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func link(t *testing.T) (*usecase.LinkUseCase, *MockLinkRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockLinkRepo(mockCtl)
	link := usecase.New(repo)

	return link, repo
}

func TestSearchLink(t *testing.T) {
	t.Parallel()

	link, repo := link(t)

	tests := []test{
		{
			name: "search link repo error",
			mock: func() {
				repo.EXPECT().FindOne(context.TODO(), "short_url = ?", "value").Return(entity.Link{}, errInternalServErr)
				repo.EXPECT().UpShowned(context.TODO(), 1).Return(errInternalServErr)
			},
			res: entity.Link{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := link.SearchLink(context.TODO(), "value")

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCreateLink(t *testing.T) {
	t.Parallel()

	link, repo := link(t)

	tests := []test{
		{
			name: "create link repo error",
			mock: func() {
				repo.EXPECT().Insert(context.TODO(), entity.Link{}).Return("", errInternalServErr)
			},
			res: "",
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := link.CreateLink(context.TODO(), entity.Link{})

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
