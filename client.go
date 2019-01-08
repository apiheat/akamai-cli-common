package commonCLI

import (
	edgegrid "github.com/apiheat/go-edgegrid"
)

// EdgeClientInit Initializes new client based on given params
func EdgeClientInit(apiClientOpts *edgegrid.ClientOptions) (*edgegrid.Client, error) {

	apiClient, err := edgegrid.NewClient(nil, apiClientOpts)

	if err != nil {
		return nil, err
	}

	return apiClient, nil
}
