package goshopify

import (
	"fmt"
	"time"
)

const customCollectionsBasePath = "admin/custom_collections"
const customCollectionsResourceName = "collections"

// CustomCollectionAPI is an interface for interacting with the custom
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/customcollection
type CustomCollectionAPI interface {
	List(interface{}) ([]CustomCollection, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*CustomCollection, error)
	Create(CustomCollection) (*CustomCollection, error)
	Update(CustomCollection) (*CustomCollection, error)
	Delete(int) error

	// MetafieldsAPI used for CustomCollection resource to communicate with Metafields resource
	MetafieldsAPI
}

// CustomCollectionAPIOp handles communication with the custom collection
// related methods of the Shopify API.
type CustomCollectionAPIOp struct {
	client *Client
}

// CustomCollection represents a Shopify custom collection.
type CustomCollection struct {
	ID             int         `json:"id"`
	Handle         string      `json:"handle"`
	Title          string      `json:"title"`
	UpdatedAt      *time.Time  `json:"updated_at"`
	BodyHTML       string      `json:"body_html"`
	SortOrder      string      `json:"sort_order"`
	TemplateSuffix string      `json:"template_suffix"`
	Image          Image       `json:"image"`
	Published      bool        `json:"published"`
	PublishedAt    *time.Time  `json:"published_at"`
	PublishedScope string      `json:"published_scope"`
	Metafields     []Metafield `json:"metafields,omitempty"`
}

// CustomCollectionResource represents the result form the custom_collections/X.json endpoint
type CustomCollectionResource struct {
	Collection *CustomCollection `json:"custom_collection"`
}

// CustomCollectionsResource represents the result from the custom_collections.json endpoint
type CustomCollectionsResource struct {
	Collections []CustomCollection `json:"custom_collections"`
}

// List custom collections
func (s *CustomCollectionAPIOp) List(options interface{}) ([]CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	resource := new(CustomCollectionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count custom collections
func (s *CustomCollectionAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual custom collection
func (s *CustomCollectionAPIOp) Get(collectionID int, options interface{}) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID)
	resource := new(CustomCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new custom collection
// See Image for the details of the Image creation for a collection.
func (s *CustomCollectionAPIOp) Create(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing custom collection
func (s *CustomCollectionAPIOp) Update(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collection.ID)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing custom collection.
func (s *CustomCollectionAPIOp) Delete(collectionID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID))
}

// ListMetafields list metafields for a custom collection
func (s *CustomCollectionAPIOp) ListMetafields(customCollectionID int, options interface{}) ([]Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.List(options)
}

// CountMetafields count metafields for a custom collection
func (s *CustomCollectionAPIOp) CountMetafields(customCollectionID int, options interface{}) (int, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.Count(options)
}

// GetMetafield get individual metafield for a custom collection
func (s *CustomCollectionAPIOp) GetMetafield(customCollectionID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.Get(metafieldID, options)
}

// CreateMetafield create a new metafield for a custom collection
func (s *CustomCollectionAPIOp) CreateMetafield(customCollectionID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.Create(metafield)
}

// UpdateMetafield update an existing metafield for a custom collection
func (s *CustomCollectionAPIOp) UpdateMetafield(customCollectionID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.Update(metafield)
}

// DeleteMetafield delete an existing metafield for a custom collection
func (s *CustomCollectionAPIOp) DeleteMetafield(customCollectionID int, metafieldID int) error {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldAPI.Delete(metafieldID)
}
