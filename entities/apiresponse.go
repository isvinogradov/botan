package entities

// The response contains a JSON object, which always has a Boolean field ‘ok’ and may have an optional
// String field ‘description’ with a human-readable description of the result.
type ApiResponse struct {
	OK          bool               `json:"ok"`                    // If ‘ok’ equals true, the request was successful and the result of the query can be found in the ‘result’ field. In case of an unsuccessful request, ‘ok’ equals false and the error is explained in the ‘description’.
	Result      interface{}        `json:"result"`                // If ‘ok’ equals true, the request was successful and the result of the query can be found in the ‘result’ field.
	Description string             `json:"description,omitempty"` // Optional. In case of an unsuccessful request, ‘ok’ equals false and the error is explained in the ‘description’.
	ErrorCode   int                `json:"error_code,omitempty"`  // Optional. An Integer ‘error_code’ field is also returned, but its contents are subject to change in the future.
	RespParams  ResponseParameters `json:"parameters,omitempty"`  // Optional. Some errors may also have an optional field ‘parameters’ of the type ResponseParameters, which can help to automatically handle the error.
}

//Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int   `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

//func (ar *ApiResponse) UnmarshalJSON(b []byte) error {
//	// if message_id in ar.Result -> this is a message
//	return nil
//}
