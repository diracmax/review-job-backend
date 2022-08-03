//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/usecase/repository/$GOFILE
package repository

type UserRepository interface {
	Save(name string, rawPassword string) error
}
