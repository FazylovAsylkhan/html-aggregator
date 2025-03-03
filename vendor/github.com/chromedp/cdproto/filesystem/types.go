package filesystem

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/storage"
)

// File [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FileSystem#type-File
type File struct {
	Name         string              `json:"name"`
	LastModified *cdp.TimeSinceEpoch `json:"lastModified"` // Timestamp
	Size         float64             `json:"size"`         // Size in bytes
	Type         string              `json:"type"`
}

// Directory [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FileSystem#type-Directory
type Directory struct {
	Name              string   `json:"name"`
	NestedDirectories []string `json:"nestedDirectories"`
	NestedFiles       []*File  `json:"nestedFiles"` // Files that are directly nested under this directory.
}

// BucketFileSystemLocator [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FileSystem#type-BucketFileSystemLocator
type BucketFileSystemLocator struct {
	StorageKey     storage.SerializedStorageKey `json:"storageKey"`                    // Storage key
	BucketName     string                       `json:"bucketName,omitempty,omitzero"` // Bucket name. Not passing a bucketName will retrieve the default Bucket. (https://developer.mozilla.org/en-US/docs/Web/API/Storage_API#storage_buckets)
	PathComponents []string                     `json:"pathComponents"`                // Path to the directory using each path component as an array item.
}
