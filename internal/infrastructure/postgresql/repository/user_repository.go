package repository

import (
	"context"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	shared_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/shared"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/postgresql/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserPostgresqlRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) repository.UserRepository {
	return UserPostgresqlRepository{queries}
}

func (ur UserPostgresqlRepository) Find(ctx context.Context, id user_vo.UserId) (models.User, error) {
	user, err := ur.queries.GetUser(ctx, uuid.MustParse(string(id)))
	if err != nil {
		return models.User{}, err
	}

	followers, _ := user.Followers.(user_vo.UserFollowers)

	return models.User{
		Id:          user_vo.UserId(user.ID.String()),
		Name:        user_vo.UserName(user.Name),
		Picture:     user_vo.UserPicture(user.Picture),
		Description: user_vo.UserDescription(user.Description.String),
		Followers:   user_vo.UserFollowers(followers),
		Mail:        user_vo.UserMail(user.Mail),
		CreatedAt:   shared_vo.CreatedAt(user.Createdat.Time),
	}, nil
}

func (ur UserPostgresqlRepository) Save(ctx context.Context, user models.User) error {
	id, err := uuid.Parse(string(user.Id))
	if err != nil {
		return err
	}

	saveUserParams := sqlc.SaveUserParams{
		ID:          id,
		Name:        string(user.Name),
		Description: pgtype.Text{String: string(user.Id)},
		Picture:     string(user.Picture),
		Mail:        string(user.Mail),
	}

	err = ur.queries.SaveUser(ctx, saveUserParams)

	return err
}
