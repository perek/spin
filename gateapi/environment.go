/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Environment struct {

	Notifications []Notification `json:"notifications,omitempty"`

	Resources []Resource `json:"resources,omitempty"`

	Locations *interface{} `json:"locations,omitempty"`

	Name string `json:"name,omitempty"`

	Constraints []Mapstringobject `json:"constraints,omitempty"`
}
