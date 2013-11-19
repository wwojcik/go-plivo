// Public Domain (-) 2013 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type PricingService struct {
	client *Client
}

type Pricing struct {
	Pricing string
}

type PricingGetParams struct {
	CountryISO string
}

// Get fetches the pricing for a specified country
func (s *PricingService) Get(p *PricingGetParams) (*Pricing, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Pricing/", p)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Pricing{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}
