package appwrite

import (
	"context"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/databases"
)

// Client represents an Appwrite client
type Client struct {
	*appwrite.Client
}

// NewClient returns a new Appwrite client
func NewClient(endpoint string, projectID string, apiKey string) (*Client, error) {
	client := appwrite.NewClient(
		appwrite.WithEndpoint(endpoint),
		appwrite.WithProject(projectID),
		appwrite.WithKey(apiKey),
	)

	return &Client{client}, nil
}

// CreateUser creates a new user document
func (c *Client) CreateUser(ctx context.Context, databaseID string, usersCollectionID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.CreateDocument(ctx, databaseID, usersCollectionID, "", data)
	return doc, err
}

// CreateProduct creates a new product document
func (c *Client) CreateProduct(ctx context.Context, databaseID string, productsCollectionID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.CreateDocument(ctx, databaseID, productsCollectionID, "", data)
	return doc, err
}

// CreateRating creates a new rating document
func (c *Client) CreateRating(ctx context.Context, databaseID string, ratingsCollectionID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.CreateDocument(ctx, databaseID, ratingsCollectionID, "", data)
	return doc, err
}

// CreateRecommendation creates a new recommendation document
func (c *Client) CreateRecommendation(ctx context.Context, databaseID string, recommendationsCollectionID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.CreateDocument(ctx, databaseID, recommendationsCollectionID, "", data)
	return doc, err
}

// ListUsers lists all user documents
func (c *Client) ListUsers(ctx context.Context, databaseID string, usersCollectionID string) ([]*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	docs, err := databases.ListDocuments(ctx, databaseID, usersCollectionID)
	return docs, err
}

// ListProducts lists all product documents
func (c *Client) ListProducts(ctx context.Context, databaseID string, productsCollectionID string) ([]*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	docs, err := databases.ListDocuments(ctx, databaseID, productsCollectionID)
	return docs, err
}

// ListRatings lists all rating documents
func (c *Client) ListRatings(ctx context.Context, databaseID string, ratingsCollectionID string) ([]*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	docs, err := databases.ListDocuments(ctx, databaseID, ratingsCollectionID)
	return docs, err
}

// ListRecommendations lists all recommendation documents
func (c *Client) ListRecommendations(ctx context.Context, databaseID string, recommendationsCollectionID string) ([]*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	docs, err := databases.ListDocuments(ctx, databaseID, recommendationsCollectionID)
	return docs, err
}

// UpdateUser updates a user document
func (c *Client) UpdateUser(ctx context.Context, databaseID string, usersCollectionID string, documentID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.UpdateDocument(ctx, databaseID, usersCollectionID, documentID, data)
	return doc, err
}

// UpdateProduct updates a product document
func (c *Client) UpdateProduct(ctx context.Context, databaseID string, productsCollectionID string, documentID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.UpdateDocument(ctx, databaseID, productsCollectionID, documentID, data)
	return doc, err
}

// UpdateRating updates a rating document
func (c *Client) UpdateRating(ctx context.Context, databaseID string, ratingsCollectionID string, documentID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.UpdateDocument(ctx, databaseID, ratingsCollectionID, documentID, data)
	return doc, err
}

// UpdateRecommendation updates a recommendation document
func (c *Client) UpdateRecommendation(ctx context.Context, databaseID string, recommendationsCollectionID string, documentID string, data map[string]interface{}) (*appwrite.Document, error) {
	databases := appwrite.NewDatabases(c.Client)
	doc, err := databases.UpdateDocument(ctx, databaseID, recommendationsCollectionID, documentID, data)
	return doc, err
}

// DeleteUser deletes a user document
func (c *Client) DeleteUser(ctx context.Context, databaseID string, usersCollectionID string, documentID string) error {
	databases := appwrite.NewDatabases(c.Client)
	return databases.DeleteDocument(ctx, databaseID, usersCollectionID, documentID)
}

// DeleteProduct deletes a product document
func (c *Client) DeleteProduct(ctx context.Context, databaseID string, productsCollectionID string, documentID string) error {
	databases := appwrite.NewDatabases(c.Client)
	return databases.DeleteDocument(ctx, databaseID, productsCollectionID, documentID)
}

// DeleteRating deletes a rating document
func (c *Client) DeleteRating(ctx context.Context, databaseID string, ratingsCollectionID string, documentID string) error {
	databases := appwrite.NewDatabases(c.Client)
	return databases.DeleteDocument(ctx, databaseID, ratingsCollectionID, documentID)
}

// DeleteRecommendation deletes a recommendation document
func (c *Client) DeleteRecommendation(ctx context.Context, databaseID string, recommendationsCollectionID string, documentID string) error {
	databases := appwrite.NewDatabases(c.Client)
	return databases.DeleteDocument(ctx, databaseID, recommendationsCollectionID, documentID)
}