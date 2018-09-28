package commonCLI

import edgegrid "github.com/apiheat/go-edgegrid"

// EdgeClientInit Initializes new client based on given params
//
// common
func EdgeClientInit(config, section, debug string) (*edgegrid.Client, error) {
	var (
		apiClient     *edgegrid.Client
		apiClientOpts *edgegrid.ClientOptions
	)

	apiClientOpts = &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = config
	apiClientOpts.ConfigSection = section
	apiClientOpts.DebugLevel = debug

	// create new Akamai API client
	apiClient, err := edgegrid.NewClient(nil, apiClientOpts)

	if err != nil {
		return nil, err
	}

	return apiClient, nil
}
