package appwrite

import (
	"context"
	"os"

	"github.com/appwrite/sdk-for-go/appwrite"
)

// InitAppwrite initializes the Appwrite client
func InitAppwrite() (*appwrite.Client, error) {
	endpoint := os.Getenv("APPWRITE_ENDPOINT")
	projectID := os.Getenv("APPWRITE_PROJECT_ID")
	apiKey := os.Getenv("APPWRITE_API_KEY")

	client := appwrite.NewClient(
		appwrite.WithEndpoint(endpoint),
		appwrite.WithProject(projectID),
		appwrite.WithKey(apiKey),
	)

	return client, nil
}

// CreateDocument creates a new document in the specified collection
func (c *appwrite.Client) CreateDocument(ctx context.Context, dataBaseID string, collectionID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c)
	doc, err := databases.CreateDocument(ctx, databaseID, collectionID, appwrite.ID.unique(), data)
	return doc, err
}

// ListDocuments retrieves all documents in the specifies collection
func (c *appwrite.Client) ListDocuments(ctx context.Context, databaseID string, collectionID string) ([]*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c)
	docs, err := databases.ListDocuments(ctx, databaseID, collectionID)
	return docs, err
}

// UpdateDocument updates a document in the specified collection
func (c *appwrite.Client) UpdateDocument(ctx context.Context, databaseID string, collectionID string, documentID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c)
	doc, err := databases.UpdateDocument(ctx, databaseID, collectionID, documentID, data)
	return doc, err
}

// DeleteDocument deletes a document in the specified collection
func (c *appwrite.Client) DeleteDocument(ctx context.Context, databaseID string, collectionID string, documentID string) error {
	databases := appwrite.NewDatabases(c)
	return databases.DeleteDocument(ctx, databaseID, collectionID, documentID)
}