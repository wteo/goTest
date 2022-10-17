package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// err, f1 and fs variables are not used anywhere else in the code. Hence, they can be removed.
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// k index not used. Hence, should be replaced with _.
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	// k1 index not used. Hence, should be replaced with _.
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrgID, when called, this function returns all the Folder objects found in the sample.json file matching the entered organisation via orgId.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// GetSampleData is a function created to read and parse data from sample.json file
	// These data are saved and dereferenced in this "folders" variable as a Folder type slice.
	folders := GetSampleData()

	// resFolder, a blank slice is created to fill in objects that fulfil the desired condition
	resFolder := []*Folder{}
	// looping through the slice "folders" to find all objects that matches the desired organisation.
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
