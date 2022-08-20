/*
 * TODO List
 *
 * OPEN API for Todo List
 *
 * API version: 1.0.0
 * Contact: baihaqiyazid16@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse200 struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Priority float64 `json:"priority,omitempty"`
	Tags []string `json:"tags,omitempty"`
}