package main

type CreateDisk struct {
	RegionID     string //required for creating CreateDisk struct
	ZoneID       string //required for creating CreateDisk struct
	DiskName     string
	Description  string
	Encrypted    bool
	DiskCategory string
	Size         int
	SnapshotID   string
	ClientToken  string
}

type Hello struct {
	HelloMsg	string
	From	string
	To	string
}