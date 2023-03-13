package mongo

import (
	"context"

	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	"github.com/CyanAsterisk/FreeCar/server/shared/id"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/profile"
	mgutil "github.com/CyanAsterisk/FreeCar/server/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Manager struct {
	col *mongo.Collection
}

const (
	accountIDField      = "accountid"
	profileField        = "profile"
	identityStatusField = profileField + ".identitystatus"
	photoBlobIDField    = "photoblobid"
)

// ProfileRecord defines the profile record in db.
type ProfileRecord struct {
	AccountID   int64            `bson:"accountid"`
	Profile     *profile.Profile `bson:"profile"`
	PhotoBlobID int64            `bson:"photoblobid"`
}

// NewManager creates a mongo manager.
func NewManager(db *mongo.Database) *Manager {
	return &Manager{col: db.Collection(consts.ProfileCollection)}
}

// GetProfile gets profile for an account.
func (m *Manager) GetProfile(c context.Context, aid id.AccountID) (*ProfileRecord, error) {
	res := m.col.FindOne(c, byAccountID(aid))
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errno.RecordNotFound
		}
		return nil, err
	}
	var pr ProfileRecord
	err := res.Decode(&pr)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

// UpdateProfile updates profile for an account.
func (m *Manager) UpdateProfile(c context.Context, aid id.AccountID, prevState profile.IdentityStatus, p *profile.Profile) error {
	filter := bson.M{
		identityStatusField: prevState,
	}
	if prevState == profile.IdentityStatus_UNSUBMITTED {
		filter = mgutil.ZeroOrDoesNotExist(identityStatusField, prevState)
	}
	filter[accountIDField] = aid.Int64()
	_, err := m.col.UpdateOne(c, filter, mgutil.Set(bson.M{
		accountIDField: aid.Int64(),
		profileField:   p,
	}), options.Update().SetUpsert(true))
	return err
}

// UpdateProfilePhoto updates profile photo blob id.
func (m *Manager) UpdateProfilePhoto(c context.Context, aid id.AccountID, bid id.BlobID) error {
	_, err := m.col.UpdateOne(c, bson.M{
		accountIDField: aid.Int64(),
	}, mgutil.Set(bson.M{
		accountIDField:   aid.Int64(),
		photoBlobIDField: bid.Int64(),
	}), options.Update().SetUpsert(true))
	return err
}

func byAccountID(aid id.AccountID) bson.M {
	return bson.M{
		accountIDField: aid.Int64(),
	}
}