package camunda_client_go

import "fmt"

// Message a client for Message
type Message struct {
	client *Client
}

// MessageResponse ...
type MessageResponse struct {
	ResultType      string                 `json:"resultType"`
	ProcessInstance map[string]interface{} `json:"processInstance"`
	Execution       map[string]interface{} `json:"execution"`
	Variables       string                 `json:"variables"`
}

//QueryCreateMessage ...
type QueryCreateMessage struct {
	MessageName          string                 `json:"messageName"`
	BusinessKey          string                 `json:"businessKey"`
	TenantID             string                 `json:"tenantId"`
	WithoutTenantID      string                 `json:"withoutTenantId"`
	ProcessInstanceID    string                 `json:"processInstanceId"`
	CorrelationKeys      map[string]interface{} `json:"correlationKeys"`
	LocalCorrelationKeys map[string]interface{} `json:"localCorrelationKeys"`
	ProcessVariables     map[string]interface{} `json:"processVariables"`
}

// PostMessage post message
func (t *Message) PostMessage(query *QueryCreateMessage) (*MessageResponse, error) {
	if query == nil {
		query = &QueryCreateMessage{}
	}

	queryParams := map[string]string{}

	res, err := t.client.doPostJson("/message", queryParams, query)
	if err != nil {
		return nil, err
	}

	var resp *MessageResponse
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	return resp, nil
}
