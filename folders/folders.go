package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders returns a FetchFolderReponse struct that contains a slice of Folders, wherein each Folder matches the same uuid via Organisation ID
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// err, f1 & fs variables are not used anywhere else in the code. Hence, are removed.
	/*
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	*/
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// k index not used. Hence, should be replaced with _.
	for _, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	// k1 index not used. Hence, should be replaced with _.
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	// Repetitive lines 
	// var ffr *FetchFolderResponse
	// ffr = &FetchFolderResponse{Folders: fp}
	// Rewritten to:
	ffr := &FetchFolderResponse{Folders: fp} // ffr groups FetchFolderRequest structs in a single Folders slice  
	return ffr, nil
}

// FetchAllFoldersByOrgID returns all the Folder structs found in sample.json that matches the entered organisation ID via orgId.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// GetSampleData reads and parses data from sample.json 
	folders := GetSampleData()

	// resFolder, a blank slice created to fill in structs that fulfil the desired condition
	resFolder := []*Folder{}
	// looping through the slice "folders" to find all structs that matches the desired organisation ID.
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
