package main
import (
 "encoding/json"
 "fmt"
 "strconv"
)
type UserID int
type User struct {
 ID UserID `json:"id"`
 Name string `json:"name"`
}
// MarshalJSON implements the json.Marshaler interface for UserID.
func (id UserID) MarshalJSON() ([]byte, error) {
 return json.Marshal(strconv.Itoa(int(id)))
}
// UnmarshalJSON implements the json.Unmarshaler interface for UserID.
func (id *UserID) UnmarshalJSON(data []byte) error {
 var str string
 if err := json.Unmarshal(data, &str); err != nil {
  return err
 }
 parsedID, err := strconv.Atoi(str)
 if err != nil {
  return err
 }
 *id = UserID(parsedID)
 return nil
}
func main() {
 jsonData := []byte(`{"id": "123", "name": "Alice"}`)
 var user User
 err := json.Unmarshal(jsonData, &user)
 if err != nil {
  fmt.Println("Error unmarshaling JSON:", err)
  return
 }
 fmt.Printf("Decoded custom type JSON data: %+v\n", user)
}