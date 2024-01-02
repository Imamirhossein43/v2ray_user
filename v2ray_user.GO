package v2ray_user

import (
 "bytes"
 "encoding/json"
 "fmt"
 "net/http"
)

type User struct {
 ID       string json:"id"
 Email    string json:"email"
 Password string json:"password"
}

type CreateUserRequest struct {
 User User json:"user"
}

type DeleteUserRequest struct {
 ID string json:"id"
}

type GetUserStatusRequest struct {
 ID string json:"id"
}

type GetUserStatusResponse struct {
 Status string json:"status"
}

type V2RayUserManager struct {
 ServerURL string
 Token     string
}

func NewV2RayUserManager(serverURL, token string) *V2RayUserManager {
 return &V2RayUserManager{
  ServerURL: serverURL,
  Token:     token,
 }
}

func (mgr *V2RayUserManager) CreateUser(user User) error {
 apiURL := fmt.Sprintf("%s/api/v1/createUser", mgr.ServerURL)
 reqData := CreateUserRequest{
  User: user,
 }
 jsonData, err := json.Marshal(reqData)
 if err != nil {
  return err
 }

 req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
 if err != nil {
  return err
 }
 req.Header.Set("Authorization", "Bearer "+mgr.Token)
 req.Header.Set("Content-Type", "application/json")

 client := &http.Client{}
 resp, err := client.Do(req)
 if err != nil {
  return err
 }
 defer resp.Body.Close()

 if resp.StatusCode != http.StatusOK {
  return fmt.Errorf("failed to create user: %s", resp.Status)
 }

 return nil
}

func (mgr *V2RayUserManager) DeleteUser(userID string) error {
 apiURL := fmt.Sprintf("%s/api/v1/deleteUser", mgr.ServerURL)
 reqData := DeleteUserRequest{
  ID: userID,
 }
 jsonData, err := json.Marshal(reqData)
 if err != nil {
  return err
 }

 req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
 if err != nil {
  return err
 }
 req.Header.Set("Authorization", "Bearer "+mgr.Token)
 req.Header.Set("Content-Type", "application/json")

 client := &http.Client{}
 resp, err := client.Do(req)
 if err != nil {
  return err
 }
 defer resp.Body.Close()

 if resp.StatusCode != http.StatusOK {
  return fmt.Errorf("failed to delete user: %s", resp.Status)
 }

 return nil
}

func (mgr *V2RayUserManager) GetUserStatus(userID string) (string, error) {
 apiURL := fmt.Sprintf("%s/api/v1/getUserStatus", mgr.ServerURL)
 reqData := GetUserStatusRequest{
  ID: userID,
 }
 jsonData, err := json.Marshal(reqData)
 if err != nil {
  return "", err
 }

 req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
 if err != nil {
  return "", err
 }
 req.Header.Set("Authorization", "Bearer "+mgr.Token)
 req.Header.Set("Content-Type", "application/json")

 client := &http.Client{}
 resp, err := client.Do(req)
 if err != nil {
  return "", err
 }
 defer resp.Body.Close()

 if resp.StatusCode != http.StatusOK {
  return "", fmt.Errorf("failed to get user status: %s", resp.Status)
 }

 var statusResp GetUserStatusResponse
 err = json.NewDecoder(resp.Body).Decode(&statusResp)
 if err != nil {
  return "", err
 }

 return statusResp.Status, nil
}