package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// INFO is the struct that represents the information of the permit.
type INFO struct {
	Classification int    `bson:"classification,omitempty"`
	Description    string `bson:"description,omitempty"`
}

// BSON is the struct that represents a permit as it's stored.
type BSON struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FileID   string             `bson:"fileID, omitempty"`
	SharerID string             `bson:"sharerID,omitempty"`
	UserID   string             `bson:"userID,omitempty"`
	Info     *INFO              `bson:"info,omitempty"`
}

// GetID returns the string value of the b.ID.
func (b BSON) GetID() string {
	if b.ID.IsZero() {
		return ""
	}

	return b.ID.Hex()
}

// SetID sets the b.ID ObjectID's string value to id.
func (b *BSON) SetID(id string) error {
	if b == nil {
		panic("b == nil")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	b.ID = objectID
	return nil
}

// GetFileID returns b.FileID.
func (b BSON) GetFileID() string {
	return b.FileID
}

// SetFileID sets b.FileID to fileID.
func (b *BSON) SetFileID(fileID string) error {
	if b == nil {
		panic("b == nil")
	}

	if fileID == "" {
		return fmt.Errorf("FileID is required")
	}

	b.FileID = fileID
	return nil
}

// GetSharerID returns b.SharerID.
func (b BSON) GetSharerID() string {
	return b.SharerID
}

// SetSharerID sets b.SharerID to sharerID.
func (b *BSON) SetSharerID(sharerID string) error {
	if b == nil {
		panic("b == nil")
	}

	if sharerID == "" {
		return fmt.Errorf("SharerID is required")
	}

	b.SharerID = sharerID
	return nil
}

// GetUserID returns b.UserID.
func (b BSON) GetUserID() string {
	return b.UserID
}

// SetUserID sets b.UserID to userID.
func (b *BSON) SetUserID(userID string) error {
	if b == nil {
		panic("b == nil")
	}

	if userID == "" {
		return fmt.Errorf("UserID is required")
	}

	b.UserID = userID
	return nil
}

// GetInfo returns b.UserID.
func (b BSON) GetInfo() *INFO {
	return b.Info
}

// SetInfo sets b.UserID to userID.
func (b *BSON) SetInfo(info *INFO) error {
	if b == nil {
		panic("b == nil")
	}

	if info == nil {
		return fmt.Errorf("info is required")
	}

	b.Info = info
	return nil
}
