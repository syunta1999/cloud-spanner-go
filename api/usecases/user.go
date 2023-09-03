// usecases/user.go
package usecases

import (
	"context"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

type UserInteractor struct {
	Client *spanner.Client
}

func NewUserInteractor(client *spanner.Client) *UserInteractor {
	return &UserInteractor{Client: client}
}

func (u *UserInteractor) GetUsers() ([]*User, error) {
	ctx := context.Background()
	stmt := spanner.NewStatement("SELECT * FROM user")
	iter := u.Client.Single().Query(ctx, stmt)
	defer iter.Stop()

	var users []*User
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read data: %v", err)
			return nil, err
		}
		var user User
		if err := row.Columns(&user.ID, &user.Name, &user.Email); err != nil {
			log.Fatalf("Failed to parse row: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserInteractor) CreateUsers() error {
	ctx := context.Background()

	// Create a new User struct
	newUser := User{
		ID:    126,
		Name:  "John Doe",
		Email: "example@example.com",
	}

	// Create a Mutation
	mutation := spanner.Insert("user", []string{"id", "name", "email"}, []interface{}{newUser.ID, newUser.Name, newUser.Email})

	// Apply the Mutation using a ReadWriteTransaction
	_, err := u.Client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		return txn.BufferWrite([]*spanner.Mutation{mutation})
	})
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	} else {
		log.Println("Successfully inserted data.")
	}
	return nil
}
